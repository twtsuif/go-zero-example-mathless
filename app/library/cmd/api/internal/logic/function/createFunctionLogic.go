package function

import (
	"context"
	"mathless-backend/app/library/cmd/rpc/libraryservice"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/library/cmd/api/internal/svc"
	"mathless-backend/app/library/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFunctionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFunctionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFunctionLogic {
	return &CreateFunctionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFunctionLogic) CreateFunction(req *types.CreateFunctionReq) (resp *types.CreateFunctionResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := libraryservice.CreateFunctionRequest{
		Uid:          uid,
		Name:         req.Name,
		FilePathJson: req.FilePathJson,
		OssBucket:    req.OssBucket,
		OssObject:    req.OssObject,
		LibraryId:    req.LibraryId,
	}
	rpcResp, err := l.svcCtx.LibraryRpc.CreateFunction(l.ctx, &rpcReq)
	if err != nil {
		return nil, err
	}
	resp.Id = rpcResp.Id
	return
}
