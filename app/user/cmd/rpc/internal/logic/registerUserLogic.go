package logic

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"mathless-backend/app/user/cmd/rpc/internal/svc"
	"mathless-backend/app/user/cmd/rpc/user"
	"mathless-backend/app/user/model"
	"mathless-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RegisterUser 注册
// 查询用户是否已经存在 -> 校验邮箱验证码 -> 数据库插入数据 -> 返回token
func (l *RegisterUserLogic) RegisterUser(in *user.RegisterUserRequest) (*user.RegisterUserResponse, error) {
	// 判断邮箱和用户名是否已经存在
	existEmail, err := l.emailExist(in.Email)
	if err != nil {
		return nil, err
	}
	existUsername, err := l.usernameExist(in.UserName)
	if err != nil {
		return nil, err
	}
	if existEmail || existUsername {
		return nil, xerr.NewCustomErrorByStatus(xerr.USER_EXIST)
	}

	// 验证码是否正确
	right, err := l.verifyCode(in.Email, in.VerifyCode)
	if err != nil {
		return nil, err
	}
	if !right {
		return nil, xerr.NewCustomErrorByStatus(xerr.USER_EXIST)
	}

	// 加密
	hash, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)

	// 数据库插入数据
	insertUser, err := l.svcCtx.UserModel.Insert(context.Background(), &model.User{
		Username: in.UserName,
		Password: string(hash),
		Email:    in.Email,
	})
	if err != nil {
		return nil, xerr.NewCustomErrorByStatus(xerr.DB_ERROR)
	}

	// 生成token
	id, err := insertUser.LastInsertId()
	token, err := l.CreateJWT(id)

	return &user.RegisterUserResponse{Token: token}, err
}

// 邮箱是否存在
func (l *RegisterUserLogic) emailExist(email string) (bool, error) {
	userModel, err := l.svcCtx.UserModel.FindOneByEmail(context.Background(), email)

	if err != nil && err != model.ErrNotFound {
		return true, errors.Wrapf(xerr.NewCustomErrorByStatus(xerr.DB_ERROR), "select email err %v", err)
	}

	if userModel != nil {
		return true, nil
	}
	return false, nil
}

// 用户名是否存在
func (l *RegisterUserLogic) usernameExist(username string) (bool, error) {
	userModel, err := l.svcCtx.UserModel.FindOneByUsername(context.Background(), username)

	if err != nil && err != model.ErrNotFound {
		return true, errors.Wrapf(xerr.NewCustomErrorByStatus(xerr.DB_ERROR), "select username err %v", err)
	}

	if userModel != nil {
		return true, nil
	}

	return false, nil
}

func (l *RegisterUserLogic) verifyCode(email, code string) (bool, error) {
	rightCode, err := l.svcCtx.RedisClient.Get(VerifyEmailRedisKey + email)
	if err != nil {
		return false, errors.Wrapf(xerr.NewCustomErrorByStatus(xerr.DB_ERROR), "select email err %v", err)
	}
	return rightCode == code, nil
}
