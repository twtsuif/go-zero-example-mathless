package logic

import (
	"context"
	"mathless-backend/app/project/cmd/rpc/internal/svc"
	"mathless-backend/app/project/cmd/rpc/project"
	"mathless-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectLogic {
	return &UpdateProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProject 更新工程
func (l *UpdateProjectLogic) UpdateProject(in *project.UpdateProjectRequest) (*project.UpdateProjectResponse, error) {
	// 根据id查询
	projectModel, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.ProjectId)
	if err != nil {
		return nil, err
	}

	// 校验工程是否始于该用户
	if projectModel.Uid != in.Uid {
		return nil, xerr.NewCustomErrorByStatus(xerr.USER_INVALID_ACCESS)
	}

	// 修改并update数据库
	projectModel.Name = in.ProjectName
	projectModel.FilePathJson = in.FilePathJson
	err = l.svcCtx.ProjectModel.Update(l.ctx, projectModel)
	if err != nil {
		return nil, err
	}
	return &project.UpdateProjectResponse{}, nil
}
