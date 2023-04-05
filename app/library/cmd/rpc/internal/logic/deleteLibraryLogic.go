package logic

import (
	"context"
	"mathless-backend/common/xerr"

	"mathless-backend/app/library/cmd/rpc/internal/svc"
	"mathless-backend/app/library/cmd/rpc/library"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLibraryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLibraryLogic {
	return &DeleteLibraryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLibraryLogic) DeleteLibrary(in *library.DeleteLibraryRequest) (*library.DeleteLibraryResponse, error) {
	// 校验函数库属于该用户
	libraryModel, err := l.svcCtx.LibraryModel.FindOne(l.ctx, in.LibraryId)
	if err != nil {
		return nil, err
	}
	if libraryModel.Uid != in.Uid {
		return nil, xerr.NewCustomErrorByStatus(xerr.USER_INVALID_ACCESS)
	}

	// 删除
	err = l.svcCtx.LibraryModel.Delete(l.ctx, in.LibraryId)
	if err != nil {
		return nil, err
	}
	return &library.DeleteLibraryResponse{}, nil
}
