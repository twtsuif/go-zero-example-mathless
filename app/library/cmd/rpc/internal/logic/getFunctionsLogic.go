package logic

import (
	"context"

	"mathless-backend/app/library/cmd/rpc/internal/svc"
	"mathless-backend/app/library/cmd/rpc/library"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFunctionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFunctionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFunctionsLogic {
	return &GetFunctionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFunctionsLogic) GetFunctions(in *library.GetFunctionsRequest) (*library.GetFunctionsResponse, error) {
	functions, err := l.svcCtx.FunctionModel.FindFunctionsByLibraryId(l.ctx, in.LibraryId)
	if err != nil {
		return nil, err
	}

	var resp []*library.Function
	for _, functionModel := range functions {
		var functionResp library.Function
		functionResp.Id = functionModel.Id
		functionResp.Name = functionModel.Name
		functionResp.FilePathJson = functionModel.FilePathJson
		functionResp.OssBucket = functionModel.OssBucket
		functionResp.OssObject = functionModel.OssPath
		resp = append(resp, &functionResp)
	}

	return &library.GetFunctionsResponse{
		Functions: resp,
	}, nil
}
