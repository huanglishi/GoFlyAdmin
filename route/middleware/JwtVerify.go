package middleware

import (
	"gofly/global"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 用户信息类，作为生成token的参数
type UserClaims struct {
	ID         int64  `json:"id"`
	Accountid  int64  `json:"accountid"`  //A端主账号id
	BusinessID int64  `json:"businessID"` //B端主账号id
	Openid     string `json:"openid"`     //微信openid
	Name       string `json:"name"`
	Username   string `json:"username"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	//自定义的token秘钥
	secret = []byte("16849841325189456f489")
	// effectTime = 2 * time.Minute //两分钟
)

// 加载配置
// token有效时间（纳秒）
var Expirre = "180"
var effectTime = time.Duration(getiInt()) * time.Minute //分钟单位
// var effectTime = time.Duration(getiInt(Expirre)) * time.Hour //2小时的时间
// 写个返回int64-默认2个小时
func getiInt() int64 {
	//加载配置

	num := global.App.Config.App.TokenOutTime
	intnum, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 2 * 60 //默认2个小时
	} else {
		return intnum
	}
}

// 返回超时时间
func TokenOutTime(claims *UserClaims) int64 {
	return time.Now().Add(effectTime).Unix()
}

// 生成token
func GenerateToken(claims *UserClaims) interface{} {
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
		panic(err)
	}
	return sign
	// return map[string]interface{}{"sign": sign, "expiresat": claims.ExpiresAt}
}

// 验证token
func JwtVerify(c *gin.Context) {
	//根路径
	var NoVerifyTokenRoot_arr []string
	if global.App.Config.App.NoVerifyTokenRoot != "" {
		NoVerifyTokenRoot_arr = strings.Split(global.App.Config.App.NoVerifyTokenRoot, `,`)
	} else {
		NoVerifyTokenRoot_arr = make([]string, 0)
	}
	//具体路径
	var NoVerifyToken_arr []string
	if global.App.Config.App.NoVerifyToken != "" {
		NoVerifyToken_arr = strings.Split(global.App.Config.App.NoVerifyToken, `,`)
	} else {
		NoVerifyToken_arr = make([]string, 0)
	}
	rootPath := strings.Split(c.Request.URL.Path, "/")
	if len(rootPath) > 2 && IsContain(NoVerifyTokenRoot_arr, rootPath[1]) { //不需要token验证-根路径
		return
	} else if IsContain(NoVerifyToken_arr, c.Request.URL.Path) { //不需要token验证-具体路径
		return
	}
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.GetHeader("authorization")
	}
	if token == "" {
		panic("token 不存在")
	}
	//验证token，并存储在请求中
	c.Set("user", ParseToken(token))
}

// 解析Token
func ParseToken(tokenString string) *UserClaims {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("The token is invalid")
	}
	return claims
}

// 更新token
func Refresh(tokenString string) interface{} {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("The token is invalid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(effectTime).Unix()
	return GenerateToken(claims)
}
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
