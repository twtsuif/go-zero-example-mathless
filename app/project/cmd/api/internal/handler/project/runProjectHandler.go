package project

import (
	"mathless-backend/common/interceptor"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mathless-backend/app/project/cmd/api/internal/logic/project"
	"mathless-backend/app/project/cmd/api/internal/svc"
	"mathless-backend/app/project/cmd/api/internal/types"
)

func RunProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RunProjectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := project.NewRunProjectLogic(r.Context(), svcCtx)
		resp, err := l.RunProject(&req)
		interceptor.HttpResultInterceptor(r, w, resp, err)
	}
}
