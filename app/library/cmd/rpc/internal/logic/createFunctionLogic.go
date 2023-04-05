package logic

import (
	"context"
	"mathless-backend/app/library/model"

	"mathless-backend/app/library/cmd/rpc/internal/svc"
	"mathless-backend/app/library/cmd/rpc/library"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFunctionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateFunctionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFunctionLogic {
	return &CreateFunctionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateFunctionLogic) CreateFunction(in *library.CreateFunctionRequest) (*library.CreateFunctionResponse, error) {
	functionsModel := model.Functions{
		LibraryId:    in.LibraryId,
		Name:         in.Name,
		FilePathJson: in.FilePathJson,
		OssBucket:    in.OssBucket,
		OssPath:      in.OssObject,
	}
	result, err := l.svcCtx.FunctionModel.Insert(l.ctx, &functionsModel)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &library.CreateFunctionResponse{
		Id: id,
	}, nil
}
