package logic

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"mathless-backend/app/user/cmd/rpc/internal/svc"
	"mathless-backend/app/user/cmd/rpc/user"
	"mathless-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginUser 登录
func (l *LoginUserLogic) LoginUser(in *user.LoginUserRequest) (*user.LoginUserResponse, error) {
	// 查询用户
	userModel, err := l.svcCtx.UserModel.FindOneByEmail(context.Background(), in.Email)
	if err != nil {
		return nil, xerr.NewCustomErrorByStatus(xerr.USER_LOGIN_ERROR)
	}

	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(in.Password))
	if err != nil {
		return nil, xerr.NewCustomErrorByStatus(xerr.USER_LOGIN_ERROR)
	}

	// 生成token
	token, err := l.CreateJWT(userModel.Id)
	if err != nil {
		return nil, err
	}

	return &user.LoginUserResponse{Token: token}, err
}
