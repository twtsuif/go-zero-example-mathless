package xerr

import "fmt"

// CustomError 自定义Http错误类型 只需实现Error接口
type CustomError struct {
	status  uint32
	message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("status:%d，message:%s", e.status, e.message)
}

// GetStatus 返回给前端的错误码
func (e *CustomError) GetStatus() uint32 {
	return e.status
}

// GetMessage 返回给前端显示端错误信息
func (e *CustomError) GetMessage() string {
	return e.message
}

func NewCustomError(status uint32, message string) *CustomError {
	return &CustomError{status, message}
}
func NewCustomErrorByStatus(status uint32) *CustomError {
	return &CustomError{status, GetErrMessage(status)}
}
