package interceptor

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"mathless-backend/common/result"
	"mathless-backend/common/xerr"
	"net/http"
)

// HttpResultInterceptor 拦截http响应 封装返回类型
func HttpResultInterceptor(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		r := result.Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//// TODO 验证是否可以强转
		//customError := err.(*xerr.CustomError)
		//// 只返回自定义错误 内部错误不返回
		//if xerr.IsCustomError(customError.GetStatus()) {
		//	httpx.WriteJson(w, http.StatusBadRequest, result.FailByError(customError))
		//} else {
		//	httpx.WriteJson(w, http.StatusBadRequest, result.ServerFail())
		//}
		//
		//// 记录日志
		//logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "服务器异常"

		causeErr := errors.Cause(err)                  // err类型
		if e, ok := causeErr.(*xerr.CustomError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetStatus()
			errmsg = e.GetMessage()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCustomError(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusBadRequest, result.Fail(errcode, errmsg))
	}
}
