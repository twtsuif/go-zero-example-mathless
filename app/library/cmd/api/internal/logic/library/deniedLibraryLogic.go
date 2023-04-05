package library

import (
	"context"
	"mathless-backend/app/library/cmd/rpc/libraryservice"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeniedLibraryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeniedLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeniedLibraryLogic {
	return &DeniedLibraryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeniedLibraryLogic) DeniedLibrary(req *types.DeniedLibraryReq) (resp *types.DeniedLibraryResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := libraryservice.PublishLibraryDeniedRequest{
		Uid:        uid,
		LibraryIds: req.LibraryIds,
	}
	_, err = l.svcCtx.LibraryRpc.PublishLibraryDenied(l.ctx, &rpcReq)
	if err != nil {
		return nil, err
	}
	return
}
