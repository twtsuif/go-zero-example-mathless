package logic

import (
	"context"
	"database/sql"
	"mathless-backend/app/library/model"

	"mathless-backend/app/library/cmd/rpc/internal/svc"
	"mathless-backend/app/library/cmd/rpc/library"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLibraryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLibraryLogic {
	return &CreateLibraryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLibraryLogic) CreateLibrary(in *library.CreateLibraryRequest) (*library.CreateLibraryResponse, error) {
	libraryModel := model.Library{
		Uid:  in.Uid,
		Name: in.Name,
		Description: sql.NullString{
			String: in.Description,
			Valid:  true,
		},
	}
	result, err := l.svcCtx.LibraryModel.Insert(l.ctx, &libraryModel)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &library.CreateLibraryResponse{
		Id: id,
	}, nil
}
