package dependency

import (
	"mathless-backend/common/interceptor"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mathless-backend/app/dependency/cmd/api/internal/logic/dependency"
	"mathless-backend/app/dependency/cmd/api/internal/svc"
	"mathless-backend/app/dependency/cmd/api/internal/types"
)

func GetDependenciesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDependenciesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dependency.NewGetDependenciesLogic(r.Context(), svcCtx)
		resp, err := l.GetDependencies(&req)
		interceptor.HttpResultInterceptor(r, w, resp, err)
	}
}
