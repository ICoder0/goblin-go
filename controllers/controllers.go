package controllers

// Response 返回格式
type Response struct {
	Status int         `json:"status"`         // 状态码
	Msg    string      `json:"msg,omitempty"`  // 描述
	Data   interface{} `json:"data,omitempty"` // 数据
}
