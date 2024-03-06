package gf

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gofly/global"
	"gofly/model"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 发送GET请求
func HttpGet(url_text string, data map[string]interface{}) (map[string]interface{}, error) {
	u, err := url.Parse(url_text)
	if err != nil {
		log.Fatal(err)
	}
	paras := &url.Values{}
	//设置请求参数
	for k, v := range data {
		paras.Set(k, fmt.Sprintf("%v", v))
	}
	u.RawQuery = paras.Encode()
	resp, err := http.Get(u.String())
	//关闭资源
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, errors.New("request token err :" + err.Error())
	}
	jMap := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&jMap)
	if err != nil {
		return nil, errors.New("request token response json parse err :" + err.Error())
	} else {
		return jMap, nil
	}

}

// 发送POST请求
func HttpPost(url_text string, urldata map[string]interface{}, postdata map[string]interface{}, contentType string) (map[string]interface{}, error) {
	u, err := url.Parse(url_text)
	if err != nil {
		log.Fatal(err)
	}
	paras := &url.Values{}
	//设置请求参数
	for k, v := range urldata {
		paras.Set(k, v.(string))
	}
	u.RawQuery = paras.Encode()
	//json序列化
	jsonData := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(jsonData)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(postdata); err != nil {
		return nil, errors.New("请求错误 :" + err.Error())
	}
	body := bytes.NewBufferString(string(jsonData.Bytes()))
	resp, erro := http.Post(u.String(), contentType, body)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if erro != nil {
		return nil, errors.New("请求错误 :" + erro.Error())
	}
	jMap := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&jMap)
	if err != nil {
		return nil, errors.New(" 返回结果解析错误 :" + err.Error())
	} else {
		return jMap, nil
	}

}

// 发送POST请求-备用
func HttpPost_c(url_text string, urldata map[string]interface{}, postdata map[string]interface{}, contentType string) (map[string]interface{}, error) {
	u, err := url.Parse(url_text)
	if err != nil {
		log.Fatal(err)
	}
	paras := &url.Values{}
	//设置请求参数
	for k, v := range urldata {
		paras.Set(k, v.(string))
	}
	u.RawQuery = paras.Encode()
	jsonStr, _ := json.Marshal(postdata)
	body := bytes.NewBuffer([]byte(jsonStr))
	resp, erro := http.Post(u.String(), contentType, body)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if erro != nil {
		return nil, errors.New("请求错误 :" + erro.Error())
	}
	jMap := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&jMap)
	if err != nil {
		return nil, errors.New(" 返回结果解析错误 :" + err.Error())
	} else {
		return jMap, nil
	}

}

// 请求失败返回
type Response struct {
	Code      int         `json:"code"`
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

func ServerError(c *gin.Context, err interface{}) {
	conf := global.App.Config
	msg := "内部服务器错误"
	if os.Getenv(gin.EnvGinMode) != gin.ReleaseMode && reflect.TypeOf(err).Name() == "string" {
		msg = err.(string)
	} else {
		if conf.App.Env != "pro" && os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
			if _, ok := err.(error); ok {
				msg = err.(error).Error()
			}
		} else {
			str := fmt.Sprintf("内部服务器错误： %s\n", err.(error).Error()) //拼接字符串
			global.App.Log.Error(str)
		}
	}
	//判断错误类型
	if res := strings.Contains(msg, "token is expired by"); res { //token超时
		c.JSON(200, Response{
			401,
			http.StatusInternalServerError,
			nil,
			msg,
		})
	} else if res := strings.Contains(msg, "invalid memory address or nil pointer dereference"); res { //数据库链接失败
		model.MyInit(3) //重连数据库-初始化数据
		c.JSON(http.StatusInternalServerError, Response{1,
			http.StatusInternalServerError,
			"可能是数据库链接失败，请查看数据库链接是否正常",
			msg + "，可能是数据库链接失败，请查看数据库配置及是否启动，再刷新试试！",
		})
	} else {
		c.JSON(http.StatusInternalServerError, Response{1,
			http.StatusInternalServerError,
			nil,
			msg,
		})
	}
	c.Abort()
}

// 返回错误
func Get_x(url string) (string, error) {
	// 超时时间：2秒
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}
	return result.String(), nil
}

/**
*  post 请求
*  data string
 */
func Post_strdata(url string, data string, contentType string) (string, error) {
	if contentType == "" {
		contentType = "application/json"
	}
	payload := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", contentType)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		return "", error
	}
	defer resp.Body.Close()
	result, _ := io.ReadAll(resp.Body)
	return string(result), nil
}

// tool
// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}

/**
* 发送POST请求
* url：         请求地址
* data：        POST请求提交的数据 interface{}
* contentType： 请求体格式，如：application/json
* content：     请求放回的内容
 */
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := io.ReadAll(resp.Body)
	return string(result)
}
