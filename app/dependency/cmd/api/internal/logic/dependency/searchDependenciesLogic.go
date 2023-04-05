package dependency

import (
	"context"
	"mathless-backend/app/dependency/cmd/api/internal/svc"
	"mathless-backend/app/dependency/cmd/api/internal/types"
	"mathless-backend/app/dependency/cmd/rpc/dependency"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchDependenciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchDependenciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchDependenciesLogic {
	return &SearchDependenciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchDependenciesLogic) SearchDependencies(req *types.SearchDependenciesReq) (*types.SearchDependenciesResp, error) {
	rpcReq := dependency.SearchDependenciesRequest{
		Name:     req.Name,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	rpcResp, err := l.svcCtx.DependencyRpc.SearchDependencies(l.ctx, &rpcReq)
	resp := types.SearchDependenciesResp{
		Packages: rpcResp.Packages,
		Total:    rpcResp.Total,
	}
	return &resp, err
}
