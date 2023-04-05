package plugin

import (
	"context"
	"mathless-backend/app/plugin/internal/svc"
	"mathless-backend/app/plugin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RunFcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRunFcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunFcLogic {
	return &RunFcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RunFcLogic) RunFc(req *types.InvokeFcReq) (*types.InvokeFcResp, error) {
	result, err := l.svcCtx.FcTool.InvokeFunction(req.ServiceName, req.FunctionName, []byte(req.Data))
	if err != nil {
		return nil, err
	}
	resp := &types.InvokeFcResp{
		Result: result,
	}
	return resp, nil
}
