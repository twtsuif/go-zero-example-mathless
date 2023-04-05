package xerr

const OK uint32 = 200

// 系统错误
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002

// 用户错误
const USER_LOGIN_ERROR uint32 = 200001
const USER_INVALID_ACCESS uint32 = 200002
