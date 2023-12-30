package appUser

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	com "gofly/app/user/utils"
	"gofly/global"
	"gofly/model"
	"gofly/utils"
	"gofly/utils/results"
	"io"
	"reflect"
)

func init() {
	utils.Register(&Order{}, reflect.TypeOf(Order{}).PkgPath())
}

// 用于自动注册路由
type Order struct {
}

// 解锁聊天
func (api *Order) UnlockChat(c *gin.Context) {
	var parameter map[string]interface{}
	body, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &parameter)
	if parameter["man_uid"] == nil || parameter["women_uid"] == nil || parameter["unlock_uid"] == nil ||
		(parameter["man_uid"].(string) == parameter["women_uid"].(string)) {
		results.Failed(c, "入参错误！", paramError)
		return
	}
	unlockUid := parameter["unlock_type"].(string)

	manUid := parameter["man_uid"].(string)
	womenUid := parameter["women_uid"].(string)

	//查询用户解锁纪录--有解锁纪录直接返回
	record, err := model.DB().Table("app_user_order_record").Where("ManUid", manUid).Where("WomenUid", womenUid).First()
	if record != nil {
		results.Success(c, "已有解锁纪录！", apiSuccess, true)
		return
	}

	userInfo, err := model.DB().Table("app_user").Fields("Sexual", "Account", "ContactUnlockGift").
		Where("Uid", manUid).OrWhere("Uid", womenUid).
		Order("Sexual ASC").Get()
	if userInfo == nil || err != nil || len(userInfo) != 2 {
		results.Failed(c, "入参错误！", userNotExist)
		return
	}
	if unlockUid == womenUid { //女方解锁男方
		defaultPayment := global.App.Config.Userconf.UnlockChatDefaultTackOff
		womenAccount := userInfo[0]["Account"].(string)
		uInfoAccount := com.Account2Float(womenAccount)
		if uInfoAccount < defaultPayment {
			results.Failed(c, "创建聊天失败,解锁聊天余额不足！", unlockChatNotEnough)
			return
		}
		//余额足够,扣款后,生成支付记录
		newAccount := com.FloatDifference(uInfoAccount, -defaultPayment)
		_, err = model.DB().Table("app_user").Where("Uid", womenUid).Data(map[string]interface{}{
			"Account": com.Account2String(newAccount),
		}).Update()
		_, err = model.DB().Table("app_user_order_record").Data(map[string]interface{}{
			"OrderId":      com.GenerateOrderID(),
			"ManUid":       manUid,
			"WomenUid":     womenUid,
			"PaymentType":  1, //解锁聊天扣款
			"UsrPayAmount": defaultPayment,
			"SysInAmount":  defaultPayment,
			"UpdateTime":   com.DBTimeStamp(),
		}).Insert()
		if err != nil {
			//results.Failed(c, "创建聊天失败！", accountNotEnough) Add log
		}
	} else if unlockUid == manUid { //男方解锁女方
		contactUnlockGift := uint(userInfo[0]["ContactUnlockGift"].(int64))
		manAccount := userInfo[1]["Account"].(string)
		//查询解锁礼物价格
		manPayment := com.GetGiftConfig(contactUnlockGift)
		uInfoAccount := com.Account2Float(manAccount)
		//余额校验--Account不足以支付 礼物价格，提醒充值
		if uInfoAccount < manPayment {
			results.Failed(c, "创建聊天失败,解锁聊天余额不足！", unlockChatNotEnough)
			return
		}
		//余额足够,扣款后,生成支付记录
		manNewAccount := com.FloatDifference(uInfoAccount, -manPayment)
		_, err = model.DB().Table("app_user").Where("Uid", manUid).Data(map[string]interface{}{
			"Account": com.Account2String(manNewAccount),
		}).Update()

		//解锁金额存入女方账户
		womenAccount := userInfo[0]["Account"].(string)
		sysInAmount := float64(1)

		womenNewAccount := com.Account2Float(womenAccount) + (manPayment - sysInAmount)
		_, err = model.DB().Table("app_user").Where("Uid", womenUid).Data(map[string]interface{}{
			"Account": com.Account2String(womenNewAccount),
		}).Update()
		// TODO IM通知女性用户收入

		//生成交易记录
		_, err = model.DB().Table("app_user_order_record").Data(map[string]interface{}{
			"OrderId":      com.GenerateOrderID(),
			"ManUid":       manUid,
			"WomenUid":     womenUid,
			"PaymentType":  1, //解锁聊天扣款
			"UsrPayAmount": manPayment,
			"SysInAmount":  sysInAmount,
			"UsrInAmount":  manPayment - sysInAmount,
			"UpdateTime":   com.DBTimeStamp(),
		}).Insert()
		if err != nil {
			//results.Failed(c, "创建聊天失败！", accountNotEnough) Add log
		}
	}
	results.Success(c, "解锁成功", apiSuccess, true)
}

/*
	提示：解锁见面按照对方设置的见面礼物金额进行扣款,
		 如在对方接受邀请后申请取消见面,会扣除解锁礼物金额的百分之5%;
		 如对方拒绝或在对方接受邀请前申请取消不扣除任何费用;
*/
// 发起见面--进行预扣款
func (api *Order) ManUnlockMeet(c *gin.Context) {
	var parameter map[string]interface{}
	body, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &parameter)
	if parameter["man_uid"] == nil || parameter["women_uid"] == nil ||
		(parameter["man_uid"].(string) == parameter["women_uid"].(string)) {
		results.Failed(c, "入参错误！", paramError)
		return
	}
	manUid := parameter["man_uid"].(string)
	womenUid := parameter["women_uid"].(string)
	//查询用户预见面纪录--有预约见面直接返回
	prepayment, err := model.DB().Table("app_user_prepayment_record").
		Where("ManUid", manUid).Where("WomenUid", womenUid).First()
	if prepayment != nil {
		results.Success(c, "已邀请她,请等待！", apiSuccess, haveMeetProgress)
		return
	}
	userInfo, err := model.DB().Table("app_user").Fields("Sexual", "Account", "MeetUnlockGift").
		Where("Uid", manUid).OrWhere("Uid", womenUid).
		Order("Sexual ASC").Get()
	if userInfo == nil || err != nil || len(userInfo) != 2 {
		results.Failed(c, "入参错误！", userNotExist)
		return
	}
	meetUnlockGift := uint(userInfo[0]["MeetUnlockGift"].(int64))
	manAccount := userInfo[1]["Account"].(string)
	//查询见面解锁礼物价格
	manPayment := com.GetGiftConfig(meetUnlockGift)
	manCurrentAccount := com.Account2Float(manAccount)
	//余额校验--Account不足以支付 礼物价格，提醒充值
	if manCurrentAccount < manPayment {
		results.Failed(c, "解锁见面余额不足,请充值", unlockMeetNotEnough)
		return
	}
	manNewAccount := com.FloatDifference(manCurrentAccount, -manPayment)
	_, err = model.DB().Table("app_user").Where("Uid", manUid).Data(map[string]interface{}{
		"Account": com.Account2String(manNewAccount),
	}).Update()
	if err != nil {
		results.Failed(c, "更新用户余额", serverInternalErr)
		return
	}
	//插入
	_, err = model.DB().Table("app_user_prepayment_record").Data(map[string]interface{}{
		"OrderId":    com.GenerateOrderID(),
		"ManUid":     manUid,
		"WomenUid":   womenUid,
		"Amount":     manPayment,
		"State":      1,
		"UpdateTime": com.DBTimeStamp(),
	}).Insert()
	if err != nil {
		//results.Failed(c, "创建聊天失败！", accountNotEnough) Add log
	}
	//TODO: opemIM系统通知 1.通知ManUid 已经邀请 WomenUid 2.通知WomenUid 被ManUid,是否接受？

	results.Success(c, "已经发送见面邀请,待对方确认", apiSuccess, nil)
}

// 用户取消见面
func (api *Order) CancelMeet(c *gin.Context) {
	var parameter map[string]interface{}
	body, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &parameter)
	cancelFrom := parameter["cancel_uid"].(string) //男方或女方取消 传入取消方的UID
	if parameter["man_uid"] == nil || parameter["women_uid"] == nil || parameter["cancel_uid"] == nil ||
		(parameter["man_uid"].(string) == parameter["women_uid"].(string)) {
		results.Failed(c, "入参错误！", paramError)
		return
	}
	manUid := parameter["man_uid"].(string)
	womenUid := parameter["women_uid"].(string)
	//更新预扣款纪录
	amountInfo, err := model.DB().Table("app_user_prepayment_record").Fields("Amount", "State").
		Where("ManUid", manUid).Where("WomenUid", womenUid).First()
	if amountInfo == nil || err != nil {
		results.Failed(c, "查询数据库失败！", serverInternalErr)
		return
	}
	manInfo, err := model.DB().Table("app_user").Fields("Account").Where("Uid", manUid).First()
	if manInfo == nil || err != nil {
		results.Failed(c, "查询当前用户余额失败", serverInternalErr)
		return
	}
	currentMeetState := amountInfo["State"].(int64)
	//预付款余额返回给用户
	prePayAmount := com.Account2Float(amountInfo["Amount"].(string))
	var paymentSys float64 = 0
	if cancelFrom == manUid && currentMeetState == 2 { //男方取消见面
		paymentSys = prePayAmount * global.App.Config.Userconf.UserCancelTackOffRate
	}
	prePayAmount = com.FloatDifference(prePayAmount, -paymentSys)
	currentAccount := com.Account2Float(manInfo["Account"].(string))
	manNewAccount := com.FloatDifference(currentAccount, prePayAmount)
	_, err = model.DB().Table("app_user").Where("Uid", manUid).Data(map[string]interface{}{
		"Account": com.Account2String(manNewAccount),
	}).Update()
	if err != nil {
		//TODO add log
	}

	//删除该预付款纪录
	_, err = model.DB().Table("app_user_prepayment_record").
		Where("ManUid", manUid).Where("WomenUid", womenUid).Delete()
	if err != nil {
		//TODO add log
	}

	if cancelFrom == manUid { //男方取消见面
		//TODO 1openIM通知女方取消消息 2 男方方强制取消 信誉值减少
		if currentMeetState == 2 { //男方强制取消,系统扣款
			_, err = model.DB().Table("app_user_order_record").Data(map[string]interface{}{
				"OrderId":      com.GenerateOrderID(),
				"ManUid":       manUid,
				"WomenUid":     womenUid,
				"PaymentType":  5, //取消见面系统扣款
				"UsrPayAmount": prePayAmount * global.App.Config.Userconf.UserCancelTackOffRate,
				"SysInAmount":  prePayAmount * global.App.Config.Userconf.UserCancelTackOffRate,
				"UpdateTime":   com.DBTimeStamp(),
			}).Insert()
			UpdateUserCreditScore(manUid, manCancelMeetTakeoff)
		}
	}
	if cancelFrom == womenUid { //女方拒绝见面
		//TODO 1 openIM通知女方取消消息   2女方强制取消 信誉值减少 -2
		if currentMeetState == 2 {
			UpdateUserCreditScore(womenUid, womenCancelMeetTakeoff)
		}
	}
	results.Success(c, "取消成功", apiSuccess, nil)
}

// 接受邀请
func (api *Order) WomenAcceptMeet(c *gin.Context) {
	var parameter map[string]interface{}
	body, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &parameter)
	if parameter["man_uid"] == nil || parameter["women_uid"] == nil ||
		(parameter["man_uid"].(string) == parameter["women_uid"].(string)) {
		results.Failed(c, "入参错误！", paramError)
		return
	}
	manUid := parameter["man_uid"].(string)
	womenUid := parameter["women_uid"].(string)
	//更新预扣款订单状态
	_, err := model.DB().Table("app_user_prepayment_record").Where("ManUid", manUid).Where("WomenUid", womenUid).
		Data(map[string]interface{}{"State": 2, "WomenAcceptTime": com.DBTimeStamp()}).Update()
	if err != nil {
		results.Failed(c, "取消失败请重试！", cancelMeetError)
		return
	}
	//TODO  openIM 通知对男方,女方接受邀请
	results.Success(c, "更新成功", apiSuccess, nil)
}

// 女方请求搭车
func (api *Order) WomenRequestLift(c *gin.Context) {
	//var parameter map[string]interface{}
	//body, _ := io.ReadAll(c.Request.Body)
	//_ = json.Unmarshal(body, &parameter)
	//if parameter["man_uid"] == nil || parameter["women_uid"] == nil ||
	//	(parameter["man_uid"].(string) == parameter["women_uid"].(string)) {
	//	results.Failed(c, "入参错误！", paramError)
	//	return
	//}
	//
	//manUid := parameter["man_uid"].(string)
	//womenUid := parameter["women_uid"].(string)
	////查询用户预见面纪录--有预约见面直接返回
	//manInfo, _ := model.DB().Table("app_user_prepayment_record").
	//	Where("ManUid", manUid).Where("WomenUid", womenUid).First()
	//if manInfo != nil {
	//	results.Failed(c, "对方已同意,", haveMeetProgress)
	//	return
	//}
	//
	////TODO  openIM 通知对男方,女方请求搭车
	//results.Success(c, "更新成功", apiSuccess, nil)nil
}

// 男方同意见面请求
func (api *Order) ManConfirmLift(c *gin.Context) {
	var parameter map[string]interface{}
	body, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &parameter)
	if parameter["man_uid"] == nil || parameter["women_uid"] == nil ||
		(parameter["man_uid"].(string) == parameter["women_uid"].(string)) {
		results.Failed(c, "入参错误！", paramError)
		return
	}
	manUid := parameter["man_uid"].(string)
	womenUid := parameter["women_uid"].(string)

	manInfo, err := model.DB().Table("app_user").Fields("Account", "MeetGift", "CustomGift").
		Where("Uid", manUid).First()
	if manInfo == nil || err != nil {
		results.Failed(c, "入参错误！", serverInternalErr)
		return
	}
	manAccount := manInfo["Account"].(string)
	//查询见面解锁礼物价格
	var manPayment float64 = 0
	if manInfo["MeetGift"].(int64) == 0 && manInfo["CustomGift"].(string) != "" {
		manPayment = com.Account2Float(manInfo["CustomGift"].(string))
	} else if manInfo["MeetGift"].(int64) > 0 && manInfo["CustomGift"].(string) == "" {
		manPayment = com.GetGiftConfig(uint(manInfo["MeetGift"].(int64)))
	}
	manCurrentAccount := com.Account2Float(manAccount)
	//余额校验--Account不足以支付 礼物价格，提醒充值
	if manCurrentAccount < manPayment || manCurrentAccount == 0 {
		results.Failed(c, "余额不足,请充值", unlockMeetNotEnough)
		return
	}
	manNewAccount := com.FloatDifference(manCurrentAccount, -manPayment)
	_, err = model.DB().Table("app_user").Where("Uid", manUid).Data(map[string]interface{}{
		"Account": com.Account2String(manNewAccount),
	}).Update()
	if err != nil {
		results.Failed(c, "更新用户余额", serverInternalErr)
		return
	}
	//更新预扣款纪录
	_, err = model.DB().Table("app_user_prepayment_record").Data(map[string]interface{}{
		"OrderId":       com.GenerateOrderID(),
		"ManUid":        manUid,
		"WomenUid":      womenUid,
		"Amount":        manPayment,
		"State":         2,
		"ManAcceptTime": com.DBTimeStamp(),
		"UpdateTime":    com.DBTimeStamp(),
	}).Insert()
	if err != nil {
		//results.Failed(c, "创建聊天失败！", accountNotEnough) Add log
	}
	//TODO: opemIM系统通知 1.通知ManUid 已经邀请 WomenUid 2.通知WomenUid 被ManUid,是否接受？

	results.Success(c, "已接受对方邀请", apiSuccess, nil)
}

// 完成见面--送出礼物
func (api *Order) ManSendOut(c *gin.Context) {
	var parameter map[string]interface{}
	body, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &parameter)
	if parameter["man_uid"] == nil || parameter["women_uid"] == nil ||
		(parameter["man_uid"].(string) == parameter["women_uid"].(string)) {
		results.Failed(c, "入参错误！", paramError)
		return
	}
	manUid := parameter["man_uid"].(string)
	womenUid := parameter["women_uid"].(string)
	//更新预订单
	amountInfo, err := model.DB().Table("app_user_prepayment_record").Fields("Amount").
		Where("ManUid", manUid).Where("WomenUid", womenUid).Where("State", "=", 2).First()
	if err != nil {
		results.Failed(c, "更新失败！", serverInternalErr)
		return
	}
	manPayAmount := com.Account2Float(amountInfo["Amount"].(string))
	sysInAmount := manPayAmount * com.GetAmountPayCoefficient(manPayAmount) //系统扣除
	womenInAmount := com.FloatDifference(manPayAmount, -sysInAmount)

	womenAccountInfo, err := model.DB().Table("app_user").Fields("Account").Where("Uid", womenUid).First()
	if err != nil {
		results.Failed(c, "取消失败请重试！", serverInternalErr)
		return
	}
	womenCurrentAccount := com.Account2Float(womenAccountInfo["Account"].(string))
	womenNewAccount := com.FloatDifference(womenCurrentAccount, womenInAmount)
	_, err = model.DB().Table("app_user").Where("Uid", womenUid).Data(map[string]interface{}{
		"Account": com.Account2String(womenNewAccount),
	}).Update()
	if err != nil {
		results.Failed(c, "更新女方账户余额失败！", serverInternalErr)
		return
	}

	//TODO  openIM 系统通知对女方收入信息
	//插入order纪录
	model.DB().Table("app_user_order_record").Data(map[string]interface{}{
		"OrderId":      com.GenerateOrderID(),
		"ManUid":       manUid,
		"WomenUid":     womenUid,
		"PaymentType":  2, //男方确认
		"UsrPayAmount": manPayAmount,
		"SysInAmount":  sysInAmount,
		"UsrInAmount":  womenInAmount,
		"UpdateTime":   com.DBTimeStamp(),
	}).Insert()

	//删除预付款纪录
	_, err = model.DB().Table("app_user_prepayment_record").
		Where("ManUid", manUid).Where("WomenUid", womenUid).Delete()
	if err != nil {
		results.Failed(c, "更新失败！", serverInternalErr)
		return
	}
	UpdateUserCreditScore(manUid, manFinishMeetAddScore)
	UpdateUserCreditScore(womenUid, womenFinishMeetAddScore)
	results.Success(c, "确认成功", apiSuccess, nil)
}

// 男方拒绝女方
func (api *Order) ManCancelLift(c *gin.Context) {
	//var parameter map[string]interface{}
	//body, _ := io.ReadAll(c.Request.Body)
	//_ = json.Unmarshal(body, &parameter)
	//if parameter["man_uid"] == nil || parameter["women_uid"] == nil ||
	//	(parameter["man_uid"].(string) == parameter["women_uid"].(string)) {
	//	results.Failed(c, "入参错误！", paramError)
	//	return
	//}
	//TODO  openIM 通知女方
}

// 查询进行中的见面
func (api *Order) GetMeeting(c *gin.Context) {
	manUid := c.Query("man_uid")
	womenUid := c.Query("women_uid")
	if manUid == "" && womenUid == "" {
		results.Failed(c, "入参错误！", paramError)
		return
	}
	queryStr := fmt.Sprintf("ManUid,WomenUid,Amount,ManAcceptTime,WomenAcceptTime")
	//更新预扣款订单状态
	res, err := model.DB().Table("app_user_prepayment_record").Fields(queryStr).Where("ManUid", manUid).OrWhere("WomenUid", womenUid).
		Where("State", "=", 2).Get()
	if res == nil || err != nil {
		results.Failed(c, "取消失败请重试！", getMeetingInfoError)
		return
	}
	results.Success(c, "查询", apiSuccess, res)
}

// 获取已完成纪录
func (api *Order) GetMeetRecord(c *gin.Context) {
	manUid := c.Query("man_uid")
	womenUid := c.Query("women_uid")
	if manUid == "" && womenUid == "" {
		results.Failed(c, "入参错误！", paramError)
		return
	}

	var fields string
	if manUid != "" {
		fields = fmt.Sprintf("WomenUid,PaymentType,UsrPayAmount,UpdateTime")
	} else if womenUid != "" {
		fields = fmt.Sprintf("ManUid,PaymentType,UsrInAmount,UpdateTime")
	}
	res, err := model.DB().Table("app_user_order_record").Fields(fields).
		Where("ManUid", manUid).OrWhere("WomenUid", womenUid).Get()
	if err != nil {
		results.Failed(c, "查询失败！", serverInternalErr)
		return
	}

	results.Success(c, "查询成功", apiSuccess, res)
	//TODO  openIM 通知女方
}
func UpdateUserCreditScore(Uid string, creditScore int) {
	currentScoreInfo, err := model.DB().Table("app_user").Fields("CreditScore").Where("Uid", Uid).First()
	if currentScoreInfo == nil || err != nil {
		//TODO add log,
		return
	}
	currentScore := int(currentScoreInfo["CreditScore"].(int64))

	newScore := currentScore + creditScore
	if newScore <= userWarningCreditScore {
		//TODO 系统通知 信用分危险
	}
	if currentScore < newScore {
		//TODO 系统通知 信用分增加
	}
	if newScore > userMaxCreditScore {
		newScore = userMaxCreditScore
	}
	_, err = model.DB().Table("app_user").Where("Uid", Uid).
		Data(map[string]interface{}{"CreditScore": newScore}).Update()
	if err != nil {
		//TODO add log,
		return
	}
}
