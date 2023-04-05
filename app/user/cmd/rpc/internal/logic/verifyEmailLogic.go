package logic

import (
	"context"
	"github.com/pkg/errors"
	"math/rand"
	"mathless-backend/app/user/cmd/rpc/internal/svc"
	"mathless-backend/app/user/cmd/rpc/user"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/gomail.v2"
)

const VerifyEmailRedisKey = "VerifyEmail_"

type VerifyEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyEmailLogic {
	return &VerifyEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// VerifyEmail 发送邮箱验证码
func (l *VerifyEmailLogic) VerifyEmail(in *user.VerifyEmailRequest) (*user.VerifyEmailResponse, error) {
	// 发送验证码
	code, err := l.sendEmail(in.Email)
	if err != nil {
		return nil, errors.Wrapf(err, "SendEmail err email:%s , err:%v", in.Email, err)
	}

	// 放入缓存
	if err := l.svcCtx.RedisClient.Setex(VerifyEmailRedisKey+in.Email, code, 5*60); err != nil {
		return nil, errors.Wrapf(err, "Redis set verify code err , err:%v", err)
	}

	return &user.VerifyEmailResponse{Code: code}, nil
}

func (l *VerifyEmailLogic) sendEmail(mailTo string) (string, error) {
	message := gomail.NewMessage()
	// 邮件头
	message.SetHeader("From", "mathless@yeah.net")
	message.SetHeader("To", mailTo)
	message.SetHeader("Subject", "【Mathless】验证码")

	// 邮件内容
	code := generateCode(6)
	message.SetBody("text/html", "您的验证码是: <b>"+code+"</b> 五分钟内有效，非本人操作请忽略")

	// 发送
	dialer := gomail.NewDialer(l.svcCtx.Config.SMTPConf.Host,
		l.svcCtx.Config.SMTPConf.Port,
		l.svcCtx.Config.SMTPConf.Username,
		l.svcCtx.Config.SMTPConf.Password)

	return code, dialer.DialAndSend(message)
}

// 生成验证码
func generateCode(length int) string {
	var code string
	for i := 0; i < length; i++ {
		randInt := rand.Int() % 10
		code = code + strconv.Itoa(randInt)
	}
	return code
}
