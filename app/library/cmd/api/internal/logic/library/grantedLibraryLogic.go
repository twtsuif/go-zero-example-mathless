package library

import (
	"context"
	"mathless-backend/app/library/cmd/rpc/libraryservice"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrantedLibraryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrantedLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrantedLibraryLogic {
	return &GrantedLibraryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrantedLibraryLogic) GrantedLibrary(req *types.GrantedLibraryReq) (resp *types.GrantedLibraryResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := libraryservice.PublishLibraryGrantedRequest{
		Uid:        uid,
		LibraryIds: req.LibraryIds,
	}
	_, err = l.svcCtx.LibraryRpc.PublishLibraryGranted(l.ctx, &rpcReq)
	return
}
