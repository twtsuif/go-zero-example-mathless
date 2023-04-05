package dependency

import (
	"context"
	"mathless-backend/app/dependency/cmd/api/internal/svc"
	"mathless-backend/app/dependency/cmd/api/internal/types"
	"mathless-backend/app/dependency/cmd/rpc/dependency"
	jwtTool "mathless-backend/common/tool/jwt"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDependenciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDependenciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDependenciesLogic {
	return &GetDependenciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDependenciesLogic) GetDependencies(req *types.GetDependenciesReq) (*types.GetDependenciesResp, error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	// file前缀
	downloadLocation := StoreLocationDir + "/" + strconv.FormatInt(uid, 10) + "/"

	// 根据类型拼接
	if req.Typ == ProjectDependencyType {
		downloadLocation += ProjectPrefix + "/" + req.ProjectId + ".zip"
	} else if req.Typ == LibraryDependencyType {
		downloadLocation += FunctionPrefix + "/" + req.FunctionId + ".zip"
	}

	// 构造并调用rpc
	rpcReq := &dependency.GetDependencyRequest{DependencyFileLocation: downloadLocation}
	rpcResp, err := l.svcCtx.DependencyRpc.GetDependency(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	var resp types.GetDependenciesResp
	resp.PackageJson = rpcResp.PackagesJson
	return &resp, err
}
