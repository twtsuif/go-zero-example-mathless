package function

import (
	"context"
	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"
	"mathless-backend/app/library/cmd/rpc/library"
	jwtTool "mathless-backend/common/tool/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFunctionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFunctionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFunctionLogic {
	return &UpdateFunctionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFunctionLogic) UpdateFunction(req *types.UpdateFunctionReq) (resp *types.UpdateFunctionResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := library.UpdateFunctionRequest{
		Id:           req.Id,
		Uid:          uid,
		Name:         req.Name,
		FilePathJson: req.FilePathJson,
	}
	_, err = l.svcCtx.LibraryRpc.UpdateFunction(l.ctx, &rpcReq)
	return
}
