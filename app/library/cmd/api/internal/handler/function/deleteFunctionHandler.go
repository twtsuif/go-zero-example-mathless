package function

import (
	"mathless-backend/common/interceptor"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mathless-backend/app/library/cmd/api/internal/logic/function"
	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"
)

func DeleteFunctionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteFunctionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := function.NewDeleteFunctionLogic(r.Context(), svcCtx)
		resp, err := l.DeleteFunction(&req)
		interceptor.HttpResultInterceptor(r, w, resp, err)
	}
}
