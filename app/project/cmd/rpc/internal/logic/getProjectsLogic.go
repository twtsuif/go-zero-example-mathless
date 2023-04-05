package logic

import (
	"context"
	"mathless-backend/app/project/cmd/rpc/internal/svc"
	"mathless-backend/app/project/cmd/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectsLogic {
	return &GetProjectsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProjects 获取用户的工程
func (l *GetProjectsLogic) GetProjects(in *project.GetProjectsRequest) (*project.GetProjectsResponse, error) {
	projects, err := l.svcCtx.ProjectModel.FindProjectsByUid(l.ctx, in.Uid)
	if err != nil {
		return nil, err
	}

	var resp []*project.Project
	for _, projectModel := range projects {
		var projectResp project.Project
		projectResp.Id = projectModel.Id
		projectResp.Name = projectModel.Name
		projectResp.FilePathJson = projectModel.FilePathJson
		resp = append(resp, &projectResp)
	}

	return &project.GetProjectsResponse{
		Projects: resp,
	}, nil
}
