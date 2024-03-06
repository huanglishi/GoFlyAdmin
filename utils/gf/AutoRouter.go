package gf

/**
* 自动路由工具
 */
import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

// 路由结构体
type Route struct {
	path       string         //url路径
	httpMethod string         //http方法 get post
	Method     reflect.Value  //方法路由
	Args       []reflect.Type //参数类型
}

// 路由集合
var Routes = []Route{}

// 注册控制器
func Register(controller interface{}, PkgPathstr string) bool {
	// fmt.Printf("日志：%v\n", PkgPathstr)
	vbf := reflect.ValueOf(controller)
	//非控制器或无方法则直接返回
	if vbf.NumMethod() == 0 {
		return false
	}
	rootPkg := ""
	if strings.Contains(PkgPathstr, "/app") {
		PkgPath_arr := strings.Split(PkgPathstr, "/app")
		rootPkg = PkgPath_arr[len(PkgPath_arr)-1]
	}
	ctrlName := reflect.TypeOf(controller).String()
	// fmt.Println("ctrlName=", ctrlName)
	module := ctrlName
	if strings.Contains(ctrlName, ".") {
		module = ctrlName[strings.Index(ctrlName, ".")+1:]
	}
	// fmt.Println("module=", module)
	if module == "Index" { //去index
		module = "/"
	} else {
		module = "/" + strings.ToLower(module) + "/"
	}
	v := reflect.ValueOf(controller)
	// fmt.Println("遍历方法:")
	// fmt.Println(ctrlName)
	//遍历方法
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name
		//拼接路由地址
		path := rootPkg + module + FirstLower(action)
		//遍历参数
		params := make([]reflect.Type, 0, v.NumMethod())
		httpMethod := "POST" //默认POST
		if (strings.HasPrefix(action, "Get") && !strings.HasPrefix(action, "GetPost")) || action == "Index" {
			httpMethod = "GET"
		} else if strings.HasPrefix(action, "Del") || action == "Del" {
			httpMethod = "DELETE"
		} else if strings.HasPrefix(action, "Put") || action == "Put" {
			httpMethod = "PUT"
		}
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
		}
		// fmt.Println("params=", params)
		// fmt.Println("action=", action)
		route := Route{path: path, Method: method, Args: params, httpMethod: httpMethod}
		Routes = append(Routes, route)
		if strings.HasPrefix(action, "GetPost") { //再增加一个get请求
			route := Route{path: path, Method: method, Args: params, httpMethod: "GET"}
			Routes = append(Routes, route)
		}
	}
	// fmt.Println("Routes=", Routes)
	return true
}

// 绑定路由 m是方法GET POST等
// 绑定基本路由
func Bind(e *gin.Engine) {
	for _, route := range Routes {
		if route.httpMethod == "GET" {
			e.GET(route.path, match(route.path, route))
		}
		if route.httpMethod == "POST" {
			e.POST(route.path, match(route.path, route))
		}
		if route.httpMethod == "DELETE" {
			e.DELETE(route.path, match(route.path, route))
		}
		if route.httpMethod == "PUT" {
			e.PUT(route.path, match(route.path, route))
		}
	}
}

// 根据path匹配对应的方法
func match(path string, route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")
		// fmt.Println("fields,len(fields)=", fields, len(fields))
		if len(fields) < 3 {
			return
		}
		if len(Routes) > 0 {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c) // *gin.Context
			route.Method.Call(arguments)
		}
	}
}
