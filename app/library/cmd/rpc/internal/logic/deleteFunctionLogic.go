package logic

import (
	"context"
	"mathless-backend/common/xerr"

	"mathless-backend/app/library/cmd/rpc/internal/svc"
	"mathless-backend/app/library/cmd/rpc/library"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFunctionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFunctionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFunctionLogic {
	return &DeleteFunctionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFunctionLogic) DeleteFunction(in *library.DeleteFunctionRequest) (*library.DeleteFunctionResponse, error) {
	// 检查函数是否属于该用户 TODO：封装方法
	functionModel, err := l.svcCtx.FunctionModel.FindOne(l.ctx, in.FunctionId)
	if err != nil {
		return nil, err
	}
	libraryModel, err := l.svcCtx.LibraryModel.FindOne(l.ctx, functionModel.LibraryId)
	if err != nil {
		return nil, err
	}
	if libraryModel.Uid != in.Uid {
		return nil, xerr.NewCustomErrorByStatus(xerr.USER_INVALID_ACCESS)
	}

	// 删除
	err = l.svcCtx.FunctionModel.Delete(l.ctx, in.FunctionId)
	if err != nil {
		return nil, err
	}

	return &library.DeleteFunctionResponse{}, nil
}
