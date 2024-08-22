# GoFly快速开发后台管理系统介绍
## 一、框架介绍
框架采用前后端分离，将Go与Vue结合开发中后台系统，Go作为一种高效、安全的编程语言，可以帮助开发者快速构建高效、可靠、安全的应用，Vue作为前端优秀框架，可以快速搭建漂亮，高体验，稳定前端页面。能让开发者开发时顺手，客户使用时满意，性能与颜值并存，让开每一个项目交付都能让您和您的客户双方都满意。Go 开发业务接口，vue 开发前端界面。后台管理系统从业务上分为：
总管理系统（admin 端简称 A 端）和业务端管理系统（专门编写业务的，方便系统做出 saas 形系统，
减少后期需要多个应用重构成本，遇到买系统时不要单独重新部署直接再 A 端开一个账号就可以，
业务端 business 简称 B 端）。天生自带SAAS多账号数据分离，可以做到不用重新部署，即可单独拉出新的一套。

GoFly快速开发框架来自我们的医疗项目，从2019年开始用于医疗系统开发，医疗项目已经运行多年，框架的安全性、并发性能、稳定性已经得到验证。特别是在疫情期间的疫苗预约接种留观等并发和反应速度都表现良好,你可放心使用于你的项目中去。

框架采用前后端分离，Go开发使用热编译，开发目录不建议太多文件，影响编译扫描效率，前段建议放在另外位置，
git上传了全部的Go源码，前端代码和数据库请到[GoFly社区下载](https://goflys.cn/prdetail?id=6)，使用时有问题请加技术客服咨询，我们社区的初心是为不让开发者为难，打造一个让大家都舒服的社区，与大家共建一个伟大的社区。

**如果框架能帮助到你，无需捐赠，给我们⭐️star就好**，让更多人使用，开发者们都找到好工作或快速开发自己项目，企业可以降本增效。

需要框架其他扩展功能如：定时任务、cms、websocket、mqtt、工作流OA审批等插件请移步到企业版使用。
## 二、优势简介
1. 基于优秀成熟框架集成，保证系统可靠性。集成的主要有 Gin、Arco Design 、Mysql 等主流框架技术《我们不生产框架，我们是优秀框架的搬运工》。

2. 系统已集成开发常用基础功能，开箱即用，快速开始您业务开发，快人一步，比同行节省成本，降本增效首选。

3. 框架根据app目录下文件成交自动生成路由，无需手动添加，这种生成方式会避免路由重复，也减少手动添加麻烦。

4. 框架提供其他开发者开发的插件，可快速安装或卸载，让开个资源共享，同意功能无需重复造车，一键安装即可使用。 框架搭建了一键 CRUD 生成前后端代码，建数据库一键生成，节省您的复制粘贴时间，进一步为您节省时间。

5. 框架自带 API 接口文档管理，接口带有请求 token 等配置，添加接口只需配置路径和数据库或者备注，其部分信息如数据字段，系统自动根据数据库字段补齐，开发配套接口文档尽可能的为您节省一点时间。不需要其他接口文档工具复制粘贴，登录注册等时间。还有一个重点！接口文档可以一键生成接口 CRUD 的代码和通用的操作数据的 CRUD 接口，根据您的业务选择自己写接口代码、一键生成接口代码、不用写和生成代码调用通用接口。让写接口工作节省更多时间。

6. 前后端分离解耦业务，让前段人员与后端人协调开发，提高项目交付，并且可以开发出功能复杂度高的项目。

7. 前端用 Vue3+TypeScript 的 UI 框架 [Arco Design](https://arco.design/vue/component/button)，好用的 UI 框架前端可以设计出优秀且交互不错的界面，完善的大厂 UI 支持，前端开发效率也很高！ 以上只是框架比较明显优势点，还有很多优势等你自己体验，我们从各个开发环节，努力为您节省每一分时间。
8. 集成操作简单的 ORM 框架，操作数据非常简单，就像使用php的Laravel一样，您可以去文档看看 [框架的ROM数据库操作文档](https://doc.goflys.cn/docview?id=25&fid=289)
   例如下面语句就可以查找一条数据：
 ```
  db.Table("users").Fields("uid,name,age").First()
```
9. 框架以“大道至简，唯快不破”为思想，在每个细节处理都坚持让“开发”变得简单，即使你是新手也可以跟着开发文档快手上手并能开发出企业级产品。
10. 我们开源的框架不是阉割版、不留后门、没有任何开发和使用限制、没有任何收费项，框架直接从我们以往开发项目整理出来的，是个纯粹开源项目，不存在使用问题。
## 三、目录结构

```
├── app                     # 应用目录
│   ├── admin               # 后台管理应用模块
│   ├── business            # 业务端应用模块
│   ├── common              # 公共应用模块
│   ├── home                # 可以编写平台对应网站
│   ├── wxapp               # 微信小程序模块
│   ├── wxoffi              # 微信公众号模块
│   └── controller.go       # 应用控制器
├── bootstrap               # 工具方法
├── global                  # 全局变量
├── model                   # 数据模型
├── resource                # 静态资源和config配置文件
├── route                   # 路由
├── runtime                 # 运行日志文件
├── tmp                     # 开发是使用fresh热编译 产生临时文件
├── utils                   # 工具包
├── go.mod                  # 依赖包管理工具
├── go.sum         
├── main.go                 # main函数        
└── README.md               # 项目介绍
```
开发时仅需在app目录下添加你新的需求，app外部文件建议不要改动，除了config配置需要改，其他不要修改，
框架已经为您封装好，你只需在app应用目录书写你的业务，路由、访问权限、跨域、限流、Token验证、ORM等
框架已集成好，开发只需添加新的方法或者新增一个文件即可。
## 四、快速安装
1. 首先在GOPATH路径下的src目录下现在放代码的文件夹下载代码解压到项目文件夹中（或者直接git clone 代码到src目录下）。
2. 再运行服务 go run main.go 或者 编译 fresh (go install github.com/pilu/fresh@latest 安装fresh热编译工具)，启动成功如下：
![运行启动命令](https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230912/00ab0aa6dbbaea7135421d9d58fc7d53.png)
在浏览器打开安装界面进行安装：
![安装界面](https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240219/30b533c7a6d3bf711498089dd0f1337f.png)

注意：前端代码安装设置是安装时同时把前端vue代码安装到开发前端代码目录下，为了防止热编译效率框架不建议把前端代码放到go目录下。
## 五、在线预览
 [1.GoFly全栈开发社区了解更多](https://goflys.cn/home)  

 [2.Go快速后台系统开发框架完整代码包下载](https://goflys.cn/prdetail?id=6)

 [3.Go快速后台系统开发文档](https://doc.goflys.cn/docview?id=25)

 [4.A端在线预览](https://sg.goflys.cn/webadmin)

 [5.B端在线预览](https://sg.goflys.cn/webbusiness)
 
 [6.企业版admin端在线体验](https://bs.goflys.cn/webadmin/)

 [7.企业版business端在线体验](https://bs.goflys.cn/webbusiness/)

## 六、效果图片预览
### 开源版界面
<table>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230507/f7c95d545b8c6b2efcdc67411717dff9.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230808/b8304ca001cda4a94b86dad216ca5219.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230802/c4de74ba182c5037a4fd0390fb7a6ecf.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230802/f894c904f617b32a8da0bb5310ed95e0.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230802/c54d9d9141bad3aaa5a4923e7abcc32e.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230802/23ec93d6787bfcbca2e6c930213671bd.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230802/fba8e679546d1f3fe450b94e7f239a51.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230802/79770e6d1fb7e4155c67f6637a4a33df.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230802/10132ac752b08efd8b2b2c56c6492775.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230802/595c0301371762910ea3c20c1ce737ca.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230808/0708c3ad360324d3af90ebebbf47db67.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240321/8eb6bf0697fed4a40cfba0e12cba0e9e.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230808/23d844127703ba85731097a305571b89.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230802/2622a5071f8f512e8f0a31e23990da3c.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230802/85c36eef5e37779858f2e912885f71c5.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240219/f9d6376f4ed31719be29838db64542ec.png"/></td>
    </tr>
</table>

### 企业版更好用、更便捷、更耐看！
<table>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240518/07abb7d17e4da3380ef1dbc0a685f734.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240522/80aedfd2cbac9f9fdb8fc4ec62f5efa3.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240521/7f40ff3a811f81a54699dc1c0dbba20c.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240521/12ff1119fbbc300a264c6f0703fb2eea.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240522/cbe90ba418e857d49876dcff32d38476.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240522/f48f74368ee62530828f457858000fcb.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240522/1cd500bc3ad29fb7c27a16920b3e790e.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240522/efc52bb38024b08691ba757fdaa9f86e.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240522/22a34fac7575b1e075f92ef6b7499398.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240522/ed4ad0ff3e402e12acd4edc5cddcfdf0.png"/></td>
    </tr>
    <tr >
       <td colspan="2" align="center">定时任务</td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240606/c24c4400a774a77b8808ecfb866acdde.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240530/8dbe2abe66f893e5af957cc33418575a.png"/></td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240530/65342de0e5e4d8f53d59fe7660d8c054.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240530/30698a566883c400f5614272382f1878.png"/></td>
    </tr>
    <tr >
       <td colspan="2" align="center">附件管理</td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240725/0531a941abdc5a199ee5ff5dc24b0d0d.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240725/da794b591f4cfe88d58334ec94200315.png"/></td>
    </tr>
   <tr >
       <td colspan="2" align="center">工作流、审批流</td>
    </tr>
    <tr>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240804/f239b37a48885c73c38866879ef596fd.png"/></td>
        <td><img src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20240807/39930d5ef8b901b48594d31f31690c17.png?_t=1723369499"/></td>
    </tr>
</table>

## 七、安装及部署打包说明
### 1. 后端代码
#### 安装fresh 热更新-边开发边编译
```
go install github.com/pilu/fresh@latest
```
#### 初始化mod
```
go mod tidy
```
#### 热编译运行
```
bee run 或 fresh
```
#### 打包
```
go build main.go
```
#### 打包（此时会打包成Linux上可运行的二进制文件，不带后缀名的文件）
```
SET GOOS=linux
SET GOARCH=amd64
go build
```
#### widows
```
// 配置环境变量
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64

go build main.go

// 编译命令
```
#### 编译成Linux环境可执行文件
```

// 配置参数
SET CGO_ENABLED=0 
SET GOOS=linux 
SET GOARCH=amd64 

go build main.go

// 编译命令
```
#### 服务器部署
部署是把打包生成的二进制文件(Linux:gofly，windows:gofly.exe)和资源文件resource复制过去即可。
### 2. 前端端代码
#### 初始化依赖
 ```
 npm install 或者 yarn install
 ```
如果第一次使用Arco Design Pro install初始化可以报错，如果保存请运行下面命令（安装项目模版的工具）：
```
npm i -g arco-cli
```
#### 运行
```
npm run serve 或者  yarn serve
```
#### 打包
```
npm run build 或者 yarn build
```

## 八、前端代码安装及源码位置
由于框架是前端后端分离，且在Go本地开发使用fresh热编译，Go目录不能用太多文件影响编译时间，
所以我们开发是建议前端代码放在其他位置。在安装界面填写你前端代码放置位置或者手动在Go项目config/settings.yml配置文件中vueobjroot手动配置前端业务端开发路径：
```
vueobjroot: D:/Project/develop/vue/gofly_base/gofly_business
```
如果你想要手动安装前端代码，源码在代码包的resource/staticfile/template/vuecode目录下文件夹中，自己复制到开发文件夹下即可。

如果以需要了解更多关于gofl社区框架，也可以移步到[GoFly全栈开发社区](https://goflys.cn/prdetail?id=6)了解或者下载最新版本。
## 联系我们
如果使用过程有问题，可以添加GoFly技术客服咨询（微信：goflycn）,社区不建群，唯一技术交流在：https://goflys.cn/knowhow
<div align="center">
 <img src=https://goflys.cn/assets/itservice-6334c7e6.jpg width=220 />
</div>
