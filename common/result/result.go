package result

import "mathless-backend/common/xerr"

type Result struct {
	Status  uint32      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *Result {
	return &Result{Status: 200, Message: "success", Data: data}
}

func Fail(status uint32, message string) *Result {
	return &Result{Status: status, Message: message}
}

// FailByError 根据自定义异常生成Result
func FailByError(customError *xerr.CustomError) *Result {
	return Fail(customError.GetStatus(), customError.GetMessage())
}

// ServerFail 返回服务器异常的Result
func ServerFail() *Result {
	return FailByError(xerr.NewCustomErrorByStatus(xerr.SERVER_COMMON_ERROR))
}
