package response

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code int         `json:"code"`    // 错误码
	Msg  string      `json:"message"` // 错误描述
	Data interface{} `json:"data"`    // 返回数据
}

// 根据错误判断是否有预设响应信息，若无直接返回 err
func ResponseByErr(err error) Response {
	res, ok := ERR_RSP_MAP[err]
	if ok {
		// 如果有预设，直接返回
		return *res
	} else {
		return Err.WithMsg(fmt.Sprint(err))
	}
}

// 自定义响应信息
func (res *Response) WithMsg(message string) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// 追加响应数据
func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// ToString 返回 JSON 格式的错误详情
func (res *Response) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

// 构造函数
func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
