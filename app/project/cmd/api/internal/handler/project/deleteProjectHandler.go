package project

import (
	"mathless-backend/common/interceptor"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mathless-backend/app/project/cmd/api/internal/logic/project"
	"mathless-backend/app/project/cmd/api/internal/svc"
	"mathless-backend/app/project/cmd/api/internal/types"
)

func DeleteProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteProjectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := project.NewDeleteProjectLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProject(&req)
		interceptor.HttpResultInterceptor(r, w, resp, err)
	}
}
