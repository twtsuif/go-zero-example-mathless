package function

import (
	"context"
	"github.com/jinzhu/copier"
	"mathless-backend/app/library/cmd/rpc/library"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFunctionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFunctionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFunctionsLogic {
	return &GetFunctionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFunctionsLogic) GetFunctions(req *types.GetFunctionsReq) (resp *types.GetFunctionsResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := library.GetFunctionsRequest{
		Uid:       uid,
		LibraryId: req.LibraryId,
	}
	rpcResp, err := l.svcCtx.LibraryRpc.GetFunctions(l.ctx, &rpcReq)
	if err != nil {
		return nil, err
	}

	var functions []types.Function
	if len(rpcResp.Functions) > 0 {
		for _, rpcLibrary := range rpcResp.Functions {
			var functionModel types.Function
			_ = copier.Copy(&functionModel, rpcLibrary)
			functions = append(functions, functionModel)
		}
	}
	resp.Functions = functions
	return
}
