package function

import (
	"mathless-backend/common/interceptor"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mathless-backend/app/library/cmd/api/internal/logic/function"
	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"
)

func GetFunctionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFunctionsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := function.NewGetFunctionsLogic(r.Context(), svcCtx)
		resp, err := l.GetFunctions(&req)
		interceptor.HttpResultInterceptor(r, w, resp, err)
	}
}
