package logic

import (
	"context"
	"mathless-backend/common/xerr"

	"mathless-backend/app/project/cmd/rpc/internal/svc"
	"mathless-backend/app/project/cmd/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProjectLogic {
	return &DeleteProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProject 删除工程
func (l *DeleteProjectLogic) DeleteProject(in *project.DeleteProjectRequest) (*project.DeleteProjectResponse, error) {
	// 判断工程是否属于该用户
	projectModel, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.ProjectId)
	if projectModel.Uid != in.Uid {
		return nil, xerr.NewCustomErrorByStatus(xerr.USER_INVALID_ACCESS)
	}

	// 删除
	err = l.svcCtx.ProjectModel.Delete(l.ctx, in.ProjectId)
	return &project.DeleteProjectResponse{}, err
}
