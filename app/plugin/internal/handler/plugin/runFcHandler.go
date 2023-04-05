package plugin

import (
	"mathless-backend/common/interceptor"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mathless-backend/app/plugin/internal/logic/plugin"
	"mathless-backend/app/plugin/internal/svc"
	"mathless-backend/app/plugin/internal/types"
)

func RunFcHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InvokeFcReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := plugin.NewRunFcLogic(r.Context(), svcCtx)
		resp, err := l.RunFc(&req)
		interceptor.HttpResultInterceptor(r, w, resp, err)
	}
}
