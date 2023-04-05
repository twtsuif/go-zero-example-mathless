package library

import (
	"mathless-backend/common/interceptor"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mathless-backend/app/library/cmd/api/internal/logic/library"
	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"
)

func ApplyLibraryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApplyLibraryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := library.NewApplyLibraryLogic(r.Context(), svcCtx)
		resp, err := l.ApplyLibrary(&req)
		interceptor.HttpResultInterceptor(r, w, resp, err)
	}
}
