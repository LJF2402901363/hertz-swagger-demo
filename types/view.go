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
