package user

import (
	"context"
	"mathless-backend/app/user/cmd/api/internal/svc"
	"mathless-backend/app/user/cmd/api/internal/types"
	"mathless-backend/app/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	rpcReq := &user.RegisterUserRequest{
		Email:      req.Email,
		UserName:   req.Username,
		Password:   req.Password,
		VerifyCode: req.VerifyCode,
	}
	registerUser, err := l.svcCtx.UserRpc.RegisterUser(context.Background(), rpcReq)
	if err != nil {
		return nil, err
	}
	var resp types.RegisterResp
	resp.Token = registerUser.Token
	return &resp, nil
}
