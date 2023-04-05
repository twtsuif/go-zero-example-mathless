package project

import (
	"context"

	"mathless-backend/app/project/cmd/api/internal/svc"
	"mathless-backend/app/project/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RunProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRunProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunProjectLogic {
	return &RunProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RunProjectLogic) RunProject(req *types.RunProjectReq) (resp *types.RunProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
