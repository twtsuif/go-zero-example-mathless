package project

import (
	"context"
	"mathless-backend/app/project/cmd/rpc/projectservice"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/project/cmd/api/internal/svc"
	"mathless-backend/app/project/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProjectLogic {
	return &CreateProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProjectLogic) CreateProject(req *types.CreateProjectReq) (resp *types.CreateProjectResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := projectservice.CreateProjectRequest{
		Uid:          uid,
		ProjectName:  req.ProjectName,
		FilePathJson: req.FilePathJson,
	}
	rpcResp, err := l.svcCtx.ProjectRpc.CreateProject(l.ctx, &rpcReq)
	if err != nil {
		return nil, err
	}
	resp.Id = rpcResp.ProjectId
	return
}
