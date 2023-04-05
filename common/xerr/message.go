package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "success"
	message[SERVER_COMMON_ERROR] = "服务器内部异常，请联系管理员"
	message[REUQEST_PARAM_ERROR] = "请求参数非法"
	message[USER_LOGIN_ERROR] = "用户不存在或账号密码错误"
	message[USER_INVALID_ACCESS] = "非法访问"
}

func GetErrMessage(status uint32) string {
	if msg, ok := message[status]; ok {
		return msg
	} else {
		return message[SERVER_COMMON_ERROR]
	}
}

// IsCustomError  是否为自定义错误
func IsCustomError(status uint32) bool {
	if _, ok := message[status]; ok {
		return true
	} else {
		return false
	}
}
