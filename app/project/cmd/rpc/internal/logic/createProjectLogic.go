package logic

import (
	"context"
	"mathless-backend/app/project/model"

	"mathless-backend/app/project/cmd/rpc/internal/svc"
	"mathless-backend/app/project/cmd/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProjectLogic {
	return &CreateProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateProject 创建工程
func (l *CreateProjectLogic) CreateProject(in *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	projectModel := model.Project{
		Uid:          in.Uid,
		Name:         in.ProjectName,
		FilePathJson: in.FilePathJson,
	}
	result, err := l.svcCtx.ProjectModel.Insert(l.ctx, &projectModel)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	return &project.CreateProjectResponse{
		ProjectId: id,
	}, nil
}
