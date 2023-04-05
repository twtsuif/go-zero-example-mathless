package logic

import (
	"github.com/golang-jwt/jwt/v4"
	jwtTool "mathless-backend/common/tool/jwt"
	"time"
)

func (l *RegisterUserLogic) CreateJWT(uid int64) (string, error) {
	secret := l.svcCtx.Config.JwtAuth.AccessSecret
	last := l.svcCtx.Config.JwtAuth.AccessExpire
	return generateJWT(secret, last, uid)
}

func (l *LoginUserLogic) CreateJWT(uid int64) (string, error) {
	secret := l.svcCtx.Config.JwtAuth.AccessSecret
	last := l.svcCtx.Config.JwtAuth.AccessExpire
	return generateJWT(secret, last, uid)
}

func generateJWT(secret string, last, uid int64) (string, error) {
	now := time.Now().Unix()

	claims := make(jwt.MapClaims)
	claims["start"] = now
	claims["expired"] = now + last
	claims[jwtTool.UidKey] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secret))
}
