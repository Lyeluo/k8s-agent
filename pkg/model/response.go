package model

type Response struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type MoreResponse struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

func NewResponse(success bool, objects interface{}, msg string) *Response {
	return &Response{Success: success, Data: objects, Msg: msg}
}

type ErrMsg struct {
	Code int64
	Msg  string
}

var (
	NoErr         = ErrMsg{200, "请求成功"}
	AuthErr       = ErrMsg{401, "认证错误"}
	AuthExpireErr = ErrMsg{402, "token 过期，请刷新token"}
	AuthActionErr = ErrMsg{403, "权限错误"}
	SystemErr     = ErrMsg{500, "系统错误，请联系管理员"}
	DataEmptyErr  = ErrMsg{501, "数据为空"}
	TokenCacheErr = ErrMsg{502, "TOKEN CACHE 错误"}
)
