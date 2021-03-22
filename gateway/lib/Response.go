package lib

import (
	"encoding/json"
	"strconv"
)

type Response struct {
	Code    int64       `json:"code"`    // 错误码
	Message string      `json:"message"` // 错误描述
	Data    interface{} `json:"data"`    // 返回数据
}

// 构造函数
func NewResponse(code int64) *Response {
	return &Response{
		Code:    code,
		Message: getErrByCode(code),
		Data:    nil,
	}
}

//成功响应
func Success() *Response {
	code := Err.GetInt64("error.ok")
	return &Response{
		Code:    code,
		Message: getErrByCode(code),
		Data:    nil,
	}
}

// 自定义响应信息
func (res *Response) WithMsg(message string) *Response {
	res.Message = message
	return res
}

// 追加响应数据
func (res *Response) WithData(data interface{}) *Response {
	res.Data = data
	return res
}

//通过code获取错误配置信息
func getErrByCode(code int64) string {
	return Err.GetString("error." + strconv.FormatInt(code, 10))
}

// ToString 返回 JSON 格式的错误详情
func (res *Response) ToString() string {
	err := &struct {
		Code    int64       `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    res.Code,
		Message: res.Message,
		Data:    res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}
