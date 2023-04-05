package library

import (
	"context"
	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"
	"mathless-backend/app/library/cmd/rpc/library"
	jwtTool "mathless-backend/common/tool/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLibraryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLibraryLogic {
	return &UpdateLibraryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLibraryLogic) UpdateLibrary(req *types.UpdateLibraryReq) (resp *types.UpdateLibraryResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := library.UpdateLibraryRequest{
		LibraryId:          req.Id,
		Uid:                uid,
		LibraryName:        req.Name,
		LibraryDescription: req.Description,
	}
	_, err = l.svcCtx.LibraryRpc.UpdateLibrary(l.ctx, &rpcReq)
	return
}
