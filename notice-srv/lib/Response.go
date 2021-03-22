package lib

import (
	"encoding/json"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/micro/go-micro/v2/util/log"
	"notice-srv/submodules/common/protob"
	"strconv"
)

//响应设计
//参考 https://www.icode9.com/content-4-813814.html

//Any字段需要通过MarshalAny(m proto.Message) (anypb.Any, error)方法转换为anypb.Any类型，然后grpc将其传给client端，
//client端再通过ptypes.UnmarshalAny(any *anypb.Any, m proto.Message) 转换成protoc3支持的类型
//参考：https://blog.csdn.net/weixin_41431016/article/details/109699817

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

//通过code获取错误配置信息
func getErrByCode(code int64) string {
	return Err.GetString("error." + strconv.FormatInt(code, 10))
}

// 转换为proto响应
func (res *Response) ToProtoRes(protoRes *proto.Response) {
	/**
	主要是将 Response.Data 转换为 proto.Response.Data,
	从go的interface转换为proto的any,
	如果是proto定义的结构体就不用这么麻烦，直接marshal就行了，无需先转成json，再用any.Any处理
	*/
	j, err := json.Marshal(res.Data)
	if err != nil {
		log.Error("json.Marshal错误", err)
	}
	orig := &any.Any{Value: []byte(j)}
	data, mErr := ptypes.MarshalAny(orig)
	if mErr != nil {
		log.Error("ptypes.MarshalAny错误", mErr)
	}

	protoRes.Code = res.Code
	protoRes.Message = res.Message
	protoRes.Data = data
}
