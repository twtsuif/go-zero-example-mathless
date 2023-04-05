package library

import (
	"context"
	"mathless-backend/app/library/cmd/rpc/libraryservice"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLibraryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLibraryLogic {
	return &DeleteLibraryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLibraryLogic) DeleteLibrary(req *types.DeleteLibraryReq) (resp *types.DeleteLibraryReq, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := libraryservice.DeleteLibraryRequest{
		Uid:       uid,
		LibraryId: req.Id,
	}
	_, err = l.svcCtx.LibraryRpc.DeleteLibrary(l.ctx, &rpcReq)
	return
}
