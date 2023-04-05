package user

import (
	"context"
	"mathless-backend/app/user/cmd/api/internal/svc"
	"mathless-backend/app/user/cmd/api/internal/types"
	"mathless-backend/app/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyEmailLogic {
	return &VerifyEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyEmailLogic) VerifyEmail(req *types.VerifyEmailReq) (*types.VerifyEmailResp, error) {
	rpcReq := &user.VerifyEmailRequest{Email: req.Email}
	rpcResp, err := l.svcCtx.UserRpc.VerifyEmail(context.Background(), rpcReq)

	var resp types.VerifyEmailResp
	resp.Code = rpcResp.Code
	return &resp, err
}
