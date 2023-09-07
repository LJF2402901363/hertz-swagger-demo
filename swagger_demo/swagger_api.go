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
