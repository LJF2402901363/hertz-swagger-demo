# go swagger使用简介以及hertz框架中集成go-swagger

# 1.背景

​	最近负责的项目模块需要从一个大单体代码仓库中迁移到新的代码仓库另起一个新服务，因此需要将我负责的代码模块从老代码仓库中迁移到新的代码仓库中然后进行重构。老的单体服务是一个基于go的web框架gin进行包装后的web框架，而我负责的代码迁移到新的服务后使用公司内部的go web框架hertz进行重构。老的代码框架中没有接入类似swagger相关的生成open API文档的组件，导致在前后端联调过程中对于http API接口的定义只能是写在飞书文档中，沟通起来非常不方便。因此为了提高前后端http接口文档的对接效率，在新的迁移服务中接入go swagger生成http API接口的文档。

# 2.go swagger的安装和使用

## 2.1安装swagger cli

~~~shell
 #  go的版本是 go1.20.3 linux/amd64
 go install github.com/swaggo/swag/cmd/swag@latest
 
 # 查看swagger的版本
 swag -v
// swag version v1.16.2
~~~



## 2.2swagger cli的使用

### 2.2.1查看 swag cli支持的命令

~~~shell
swag init -h
~~~

可以看到当前swag只支持三个命令：

>1. swag init
>2. swag fmt
>3. swag help

~~~sh
NAME:
   swag - Automatically generate RESTful API documentation with Swagger 2.0 for Go.

USAGE:
   swag [global options] command [command options] [arguments...]

VERSION:
   v1.16.2

COMMANDS:
   init, i  Create docs.go
   fmt, f   format swag comments
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
~~~

### 2.2.2swag init的使用

~~~sh
swag init -h
NAME:
   swag init - Create docs.go

USAGE:
   swag init [command options] [arguments...]

OPTIONS:
   --generalInfo value, -g value          API通用信息所在的go源文件路径，如果是相对路径则基于API解析目录 (默认: "main.go")
   --dir value, -d value                  API解析目录 (默认: "./")
   --exclude value                        解析扫描时排除的目录，多个目录可用逗号分隔（默认：空）
   --propertyStrategy value, -p value     结构体字段命名规则，三种：snakecase,camelcase,pascalcase (默认: "camelcase")
   --output value, -o value               文件(swagger.json, swagger.yaml and doc.go)输出目录 (默认: "./docs")
   --parseVendor                          是否解析vendor目录里的go源文件，默认不
   --parseDependency                      是否解析依赖目录中的go源文件，默认不
   --markdownFiles value, --md value      指定API的描述信息所使用的markdown文件所在的目录
   --generatedTime                        是否输出时间到输出文件docs.go的顶部，默认是
   --codeExampleFiles value, --cef value  解析包含用于 x-codeSamples 扩展的代码示例文件的文件夹，默认禁用
   --parseInternal                        解析 internal 包中的go文件，默认禁用
   --parseDepth value                     依赖解析深度 (默认: 100)
   --instanceName value                   设置文档实例名 (默认: "swagger")
~~~

在swag init命令中，以我的经验来看主要需要关注以下几个命令参数：

>1.    --dir value, -d value                  API解析目录 (默认: "./")，多个目录可以用英文逗号","隔开。如果不指定该参数，将会默认解析当前目录 ./ 下的所有API注释。该参数可以指定多个目录，每个目录之间用。比如指定生成两个文件夹 ./test1、./test2下的所有API文档，那么可以这样指定：swag init -d ./test1,./test2。
>
>2.  --generalInfo value, -g value          API通用信息所在的go源文件路径.如果不指定该参数，默认会使用当前项目目录下 ./main.go作为API通用信息所在的go源文件路径。如果指定了 -d参数，那么则是相对于-d参数目录下的.go文件路径。比如
>
>   swag init -d ./test1,./test2 -g swagger_api.go，那么就说明API通用信息所在的go源文件 swagger_api.go是在 ./test1或者 ./test2目录之下。
>
>3.    --exclude value                        解析扫描时排除的目录，多个目录可用逗号分隔（默认：空）。指定了目录之后解析API时候不会解析该目录列表下的API。
>
>4.   --parseVendor                          是否解析vendor目录里的go源文件，默认不不解析。
>
>5.    --parseDependency                      是否解析依赖目录中的go源文件，默认不解析。如果需要解析的API中引用了依赖文件的结构体，那么需要指定该参数。

### 2.2.3 swag fmt的使用

~~~sh
swag fmt
~~~

```sh
NAME:
   swag fmt - format swag comments 用于格式化API的注释

USAGE:
   swag fmt [command options] [arguments...]

OPTIONS:
   --dir value, -d value          API解析目录 (默认: "./")
   --exclude value                解析扫描时排除的目录，多个目录可用逗号分隔（默认：空）
   --generalInfo value, -g value  API通用信息所在的go源文件路径，如果是相对路径则基于API解析目录 (默认: "main.go")
   --help, -h                     show help (default: false)
```

# 3.hertz框架中引入go-swagger

## 3.1快速生成hert服务代码

具体参考：[hz 安装](https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/install/)

~~~
# 在 GOPATH 外执行，需要指定 go mod 名
hz new -module hertz-web

# 整理 & 拉取依赖
go mod tidy

~~~

删除不必要的包目录，并新建 swagger_demo,types目录：

~~~sql
.
├── build.sh
├── go.mod
├── go.sum
├── main.go    			# main方法入口
├── router.go
├── router_gen.go
├── swagger_demo        # 存放API接口以及 swagger通用信息接口的包
│   └── swagger_api.go  # 包含了swagger通用信息接口以及我们需要测试的接口
└── types               # 存放请求和响应的结构体
    └── view.go

~~~



## 3.2引入go-swagger

### 3.2.1  view.go文件

~~~go
/*
	@author: 24029

@since: 2023/9/3 22:03:24
@desc:
*/
package types

// BaseRes 响应的基类,这里包含泛型T
type BaseRes[T any] struct {
	Code int    // 0成功，1失败
	Msg  string // 响应信息
	Data T      // 响应数据
}

type AddUserReq struct {
	Name        string `json:"name,omitempty"  validate:"true"`        // 用户名,必填
	Password    string `json:"password,omitempty" validate:"true"`     // 用户密码，必填
	Age         int    `json:"age,omitempty" validate:"false"`         // 用户年龄，选填
	Description string `json:"description,omitempty" validate:"false"` // 用户描述，选填
}
type GetUserReq struct {
	Name     string `json:"name,omitempty"  validate:"true"`    // 用户名,必填
	Password string `json:"password,omitempty" validate:"true"` // 用户密码，必填
}

type GetUserRes struct {
	ID       int64  // 用户ID
	Name     string `json:"name,omitempty"`     // 用户名
	Password string `json:"password,omitempty"` // 用户密码
	Age      int    `json:"age"`                // 用户年龄
}

type DeleteUserReq struct {
	UserID string `json:"user_id,omitempty"  validate:"true"`
}

type UserDetailsReq struct {
	Name    string `json:"name,omitempty"  validate:"true"` // 用户名，必填
	Age     string `json:"age,omitempty"  validate:"true"`  // 用户年龄,必填
	Details string `json:"details"  validate:"false" `      // 用户详细，选填
}

~~~

### 3.2.2swagger_api.go文件

~~~go
/*
	@author: 24029

@since: 2023/9/3 21:47:51
@desc:
*/
package swagger_demo

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"hertz-web/types"
	"net/http"
)

// SaggerAPI godoc
//
//	@title						测试 hertz集成swagger API组件
//	@version					1.0
//	@description				用户测试hertz框架集成go swagger组件.
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@swagger_demo.io
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@host						localhost:8888
//	@BasePath					/api/v1
//	@securityDefinitions.basic	BasicAuth
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func SaggerAPI() {

}

// GetUser godoc
//
//	@Summary		根据用户名和用户密码获取用户信息
//	@Description	1.如果用户名不存在，那么提示用户该用户不存在
//	@Description	2.如果密码不正常那么提示用户密码不正确。
//	@Accept			application/json
//	@Produce		application/json
//	@Param			name		query		string	true	"用户名"
//	@Param			password	query		string	true	"用户密码"
//	@Success		200			{object}	types.BaseRes[types.GetUserRes]
//	@Failure		400			{object}	types.BaseRes[error]
//	@Failure		401			{object}	types.BaseRes[error]
//	@Failure		404			{object}	types.BaseRes[error]
//	@Router			/user/get [get]
func GetUser(c context.Context, ctx *app.RequestContext) {
	var req types.GetUserReq
	err := ctx.BindAndValidate(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.BaseRes[error]{
			Code: 0,
			Msg:  "参数绑定失败",
			Data: err,
		})
		return
	}
	res := types.GetUserRes{
		ID:       1,
		Name:     "",
		Password: "",
		Age:      0,
	}
	result := types.BaseRes[types.GetUserRes]{
		Code: 0,
		Data: res,
		Msg:  "",
	}
	ctx.JSON(http.StatusOK, result)
}

// AddUser godoc
//
//	@Summary		新增用户信息
//	@Description	1.如果用户名不符合规范，则禁止新增该用户
//	@Description	2.如果用户名已存在，那么新增用户失败
//	@Accept			application/json
//	@Produce		application/json
//	@Param			req	body		types.AddUserReq	true	"新增用户请求参数"
//	@Success		200	{object}	types.BaseRes[any]
//	@Failure		400	{object}	types.BaseRes[error]
//	@Failure		401	{object}	types.BaseRes[error]
//	@Failure		404	{object}	types.BaseRes[error]
//	@Failure		500	{object}	types.BaseRes[error]
//	@Router			/user/add [post]
func AddUser(c context.Context, ctx *app.RequestContext) {
	var req types.AddUserReq
	err := ctx.BindAndValidate(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.BaseRes[error]{
			Code: 0,
			Msg:  "参数绑定失败",
			Data: err,
		})
		return
	}
	result := types.BaseRes[any]{
		Code: 0,
		Msg:  "",
	}
	ctx.JSON(http.StatusOK, result)
}

// DeleteUser godoc
//
//	@Summary		根据用户ID删除用户信息
//	@Description	1.如果用户名不符合规范，则删除失败
//	@Description	2.如果用户名已存在，那么可以删除该用户
//	@Accept			application/json
//	@Produce		application/json
//	@Param			req	body		types.DeleteUserReq	true	"删除用户请求参数"
//	@Success		200	{object}	types.BaseRes[any]
//	@Failure		400	{object}	types.BaseRes[error]
//	@Failure		401	{object}	types.BaseRes[error]
//	@Failure		404	{object}	types.BaseRes[error]
//	@Failure		500	{object}	types.BaseRes[error]
//	@Router			/user/delete [delete]
func DeleteUser(c context.Context, ctx *app.RequestContext) {
	var req types.DeleteUserReq
	err := ctx.BindAndValidate(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.BaseRes[error]{
			Code: 0,
			Msg:  "参数绑定失败",
			Data: err,
		})
		return
	}
	result := types.BaseRes[any]{
		Code: 0,
		Msg:  "",
	}
	ctx.JSON(http.StatusOK, result)
}

// GetUserByID godoc
//
//	@Summary		根据用户ID获取用户信息
//	@Description	如果用户名不存在，那么返回空信息；如果用户存在那么返回具体信息
//	@Accept			application/json
//	@Produce		application/json
//	@Param			user_id		query		string	true	"用户ID"
//	@Param			password	query		string	true	"用户密码"
//	@Success		200			{object}	types.UserDetailsReq
//	@Failure		400			{object}	types.BaseRes[error]
//	@Failure		401			{object}	types.BaseRes[error]
//	@Failure		404			{object}	types.BaseRes[error]
//	@Failure		500			{object}	types.BaseRes[error]
//	@Router			/user/get/:id [get]
func GetUserByID(c context.Context, ctx *app.RequestContext) {
	var req types.GetUserReq
	err := ctx.BindAndValidate(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.BaseRes[error]{
			Code: 0,
			Msg:  "参数绑定失败",
			Data: err,
		})
		return
	}
	res := types.GetUserRes{
		ID:       1,
		Name:     "",
		Password: "",
		Age:      0,
	}
	result := types.BaseRes[types.GetUserRes]{
		Code: 0,
		Data: res,
		Msg:  "",
	}
	ctx.JSON(http.StatusOK, result)
}

~~~

>注意事项：
>
>1. 在API注释中，用户自定义结构体如果有包含泛型，那么在使用类型时候需要将泛型指定为具体类型，否则会报错。
    >
    >   比如  types.BaseRes[T any]的自定义BaseRes是一个包含泛型的结构体，因此在使用时候需要指定其具体类型：
    >     types.BaseRes[types.GetUserRes],这样才能正常解析成功。

### 3.2.3生成 swag doc文档代码

~~~sh
# 格式化并解析 ./swagger_demo,./types包下的API注释生成 swag API文档
swag fmt & swag init -d ./swagger_demo,./types -g swagger_api.go 
~~~



~~~
[1] 20301
2023/09/03 23:23:03 Generate swagger docs....
2023/09/03 23:23:03 Generate general API Info, search dir:./swagger_demo
[1]  + 20301 done       swag fmt
2023/09/03 23:23:06 Generate general API Info, search dir:./types
2023/09/03 23:23:07 Generating types.BaseRes-types_GetUserRes
2023/09/03 23:23:07 Generating types.GetUserRes
2023/09/03 23:23:07 Generating types.BaseRes-error 
2023/09/03 23:23:07 Generating types.AddUserReq
2023/09/03 23:23:07 Generating types.BaseRes-any
2023/09/03 23:23:07 Generating types.DeleteUserReq
2023/09/03 23:23:07 Generating types.UserDetailsReq
2023/09/03 23:23:07 create docs.go at docs/docs.go
2023/09/03 23:23:07 create swagger.json at docs/swagger.json
2023/09/03 23:23:07 create swagger.yaml at docs/swagger.yaml
~~~

~~~sh
./
├── build.sh
├── docs                      #生成的swagger API文档目录
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── main.go
├── router.go
├── router_gen.go
├── swagger_demo
│   └── swagger_api.go
└── types
    └── view.go

~~~

![image-20230903234145966](https://raw.githubusercontent.com/LJF2402901363/typora-images/main/202309032341291.png)

使用goland查看

### 3.2.4在main.go文件中引入适配hertz框架的中间件 hertz-contrib/swagger

~~~sh
# 安装 swaggo/files
go  install github.com/swaggo/files
# 安装hertz框架适配swagger的中间件
go install github.com/hertz-contrib/swagger
~~~

~~~
// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "hertz-web/docs"          // 必须要引入生成swagger API代码的docs包进行初始化 init()方法
	"hertz-web/swagger_demo"
)


func main() {
	h := server.Default()

	h.GET("/user/get", swagger_demo.GetUser)
	h.POST("/user/add", swagger_demo.AddUser)
	h.DELETE("/user/delete", swagger_demo.DeleteUser)
	h.GET("/user/get/:id", swagger_demo.GetUserByID)

	url := swagger.URL("http://localhost:8888/swagger/doc.json") // The url pointing to API definition
    // swagger API首页
    h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	h.Spin()
}

~~~

## 3.3启动hertz服务

编译项目：

~~~sh
go build ./
~~~

启动项目：

~~~sh
./hert-web
~~~



访问swagger API首页：
http://localhost:8888/swagger/index.html

![image-20230904001515518](https://raw.githubusercontent.com/LJF2402901363/typora-images/main/202309040015672.png)



# 4.声明式注释详解

中文版参考：[swag中文使用文档](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)

英文版参考：[swag中文使用文档](https://github.com/swaggo/swag/blob/master/README.md)

**（建议直接看英文版文档，因为英文版文档会实时更新到最新，但是中文版不一定会更新到最新）。**

5.测试用例代码仓库

[hertz-swagger-demo](https://github.com/LJF2402901363/hertz-swagger-demo)

