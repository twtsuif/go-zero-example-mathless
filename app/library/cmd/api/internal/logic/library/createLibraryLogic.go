package library

import (
	"context"
	"mathless-backend/app/library/cmd/rpc/library"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLibraryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLibraryLogic {
	return &CreateLibraryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLibraryLogic) CreateLibrary(req *types.CreateLibraryReq) (resp *types.CreateLibraryResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := library.CreateLibraryRequest{
		Uid:         uid,
		Name:        req.Name,
		Description: req.Description,
	}
	rpcResp, err := l.svcCtx.LibraryRpc.CreateLibrary(l.ctx, &rpcReq)
	if err != nil {
		return nil, err
	}
	resp.Id = rpcResp.Id
	return
}
