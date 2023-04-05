package library

import (
	"context"
	"mathless-backend/app/library/cmd/rpc/libraryservice"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyLibraryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyLibraryLogic {
	return &ApplyLibraryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyLibraryLogic) ApplyLibrary(req *types.ApplyLibraryReq) (resp *types.ApplyLibraryResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := libraryservice.PublishLibraryApplyRequest{
		Uid:       uid,
		LibraryId: req.LibraryId,
	}
	_, err = l.svcCtx.LibraryRpc.PublishLibraryApply(l.ctx, &rpcReq)
	if err != nil {
		return nil, err
	}
	return
}
