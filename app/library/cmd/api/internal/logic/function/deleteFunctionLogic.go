package function

import (
	"context"
	"mathless-backend/app/library/cmd/rpc/libraryservice"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFunctionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFunctionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFunctionLogic {
	return &DeleteFunctionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFunctionLogic) DeleteFunction(req *types.DeleteFunctionReq) (resp *types.DeleteFunctionResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := libraryservice.DeleteFunctionRequest{
		Uid:        uid,
		FunctionId: req.Id,
	}
	_, err = l.svcCtx.LibraryRpc.DeleteFunction(l.ctx, &rpcReq)
	return
}
