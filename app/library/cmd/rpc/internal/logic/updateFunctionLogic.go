package logic

import (
	"context"
	"mathless-backend/app/library/cmd/rpc/internal/svc"
	"mathless-backend/app/library/cmd/rpc/library"
	"mathless-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFunctionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFunctionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFunctionLogic {
	return &UpdateFunctionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFunctionLogic) UpdateFunction(in *library.UpdateFunctionRequest) (*library.UpdateFunctionResponse, error) {
	// 根据id查询
	functionModel, err := l.svcCtx.FunctionModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	// 校验函数是否属于该用户
	libraryModel, err := l.svcCtx.LibraryModel.FindOne(l.ctx, functionModel.LibraryId)
	if err != nil {
		return nil, err
	}
	if libraryModel.Uid != in.Uid {
		return nil, xerr.NewCustomErrorByStatus(xerr.USER_INVALID_ACCESS)
	}

	// 修改字段并update
	functionModel.Name = in.Name
	functionModel.FilePathJson = in.FilePathJson
	err = l.svcCtx.FunctionModel.Update(l.ctx, functionModel)
	if err != nil {
		return nil, err
	}
	return &library.UpdateFunctionResponse{}, nil
}
