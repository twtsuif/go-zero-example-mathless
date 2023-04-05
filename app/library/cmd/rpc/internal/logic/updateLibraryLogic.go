package logic

import (
	"context"
	"database/sql"
	"mathless-backend/common/xerr"

	"mathless-backend/app/library/cmd/rpc/internal/svc"
	"mathless-backend/app/library/cmd/rpc/library"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLibraryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLibraryLogic {
	return &UpdateLibraryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLibraryLogic) UpdateLibrary(in *library.UpdateLibraryRequest) (*library.UpdateLibraryResponse, error) {
	// 校验函数库是否属于该用户
	libraryModel, err := l.svcCtx.LibraryModel.FindOne(l.ctx, in.LibraryId)
	if err != nil {
		return nil, err
	}
	if libraryModel.Uid != in.Uid {
		return nil, xerr.NewCustomErrorByStatus(xerr.USER_INVALID_ACCESS)
	}

	// 修改并update数据库
	libraryModel.Name = in.LibraryName
	libraryModel.Description = sql.NullString{
		String: in.LibraryDescription,
		Valid:  true,
	}
	err = l.svcCtx.LibraryModel.Update(l.ctx, libraryModel)
	if err != nil {
		return nil, err
	}
	return &library.UpdateLibraryResponse{}, nil
}
