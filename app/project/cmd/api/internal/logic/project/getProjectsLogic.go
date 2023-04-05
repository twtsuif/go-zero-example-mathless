package project

import (
	"context"
	"mathless-backend/app/project/cmd/rpc/projectservice"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/project/cmd/api/internal/svc"
	"mathless-backend/app/project/cmd/api/internal/types"

	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjectsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectsLogic {
	return &GetProjectsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectsLogic) GetProjects(req *types.GetProjectsReq) (*types.GetProjectsResp, error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := projectservice.GetProjectsRequest{
		Uid: uid,
	}
	rpcResp, err := l.svcCtx.ProjectRpc.GetProjects(l.ctx, &rpcReq)

	var projects []types.Project
	if len(rpcResp.Projects) > 0 {
		for _, rpcProject := range rpcResp.Projects {
			var project types.Project
			_ = copier.Copy(&project, rpcProject)
			projects = append(projects, project)
		}
	}
	resp := types.GetProjectsResp{
		Projects: projects,
	}
	return &resp, err
}
