package logic

import (
	"context"

	"mathless-backend/app/dependency/cmd/rpc/dependency"
	"mathless-backend/app/dependency/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchDependenciesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchDependenciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchDependenciesLogic {
	return &SearchDependenciesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 搜索依赖
func (l *SearchDependenciesLogic) SearchDependencies(in *dependency.SearchDependenciesRequest) (*dependency.SearchDependenciesResponse, error) {
	requirements, err := l.svcCtx.DependencyModel.FindRequirementPageByName(l.ctx, in.Name, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.DependencyModel.FindRequirementCountByName(l.ctx, in.Name)
	if err != nil {
		return nil, err
	}

	var requirementsString []string
	for _, requirement := range requirements {
		requirementsString = append(requirementsString, requirement.Name)
	}

	return &dependency.SearchDependenciesResponse{
		Packages: requirementsString,
		Total:    total,
	}, nil
}
