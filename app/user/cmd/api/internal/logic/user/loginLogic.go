package user

import (
	"context"
	"mathless-backend/app/user/cmd/api/internal/svc"
	"mathless-backend/app/user/cmd/api/internal/types"
	"mathless-backend/app/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	rpcReq := &user.LoginUserRequest{
		Email:    req.Email,
		Password: req.Password,
	}
	loginUser, err := l.svcCtx.UserRpc.LoginUser(context.Background(), rpcReq)
	if err != nil {
		return nil, err
	}

	var resp types.LoginResp
	resp.Token = loginUser.Token
	return &resp, err
}
