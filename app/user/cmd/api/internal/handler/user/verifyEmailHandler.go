package user

import (
	"mathless-backend/app/user/cmd/api/internal/logic/user"
	"mathless-backend/app/user/cmd/api/internal/svc"
	"mathless-backend/app/user/cmd/api/internal/types"
	"mathless-backend/common/interceptor"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func VerifyEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewVerifyEmailLogic(r.Context(), svcCtx)
		resp, err := l.VerifyEmail(&req)
		interceptor.HttpResultInterceptor(r, w, resp, err)
	}
}
