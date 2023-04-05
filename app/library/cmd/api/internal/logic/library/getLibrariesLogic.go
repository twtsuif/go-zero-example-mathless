package library

import (
	"context"
	"github.com/jinzhu/copier"
	"mathless-backend/app/library/cmd/rpc/libraryservice"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLibrariesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLibrariesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLibrariesLogic {
	return &GetLibrariesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLibrariesLogic) GetLibraries(req *types.GetLibrariesReq) (*types.GetLibrariesResp, error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := libraryservice.GetLibrariesRequest{
		Uid: uid,
	}

	rpcResp, err := l.svcCtx.LibraryRpc.GetLibraries(l.ctx, &rpcReq)
	var libraries []types.Library
	if len(rpcResp.Libraries) > 0 {
		for _, rpcLibrary := range rpcResp.Libraries {
			var library types.Library
			_ = copier.Copy(&library, rpcLibrary)
			libraries = append(libraries, library)
		}
	}

	resp := types.GetLibrariesResp{
		Libraries: libraries,
	}
	return &resp, err
}
