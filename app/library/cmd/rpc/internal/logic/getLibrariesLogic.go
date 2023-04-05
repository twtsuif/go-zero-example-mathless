package logic

import (
	"context"

	"mathless-backend/app/library/cmd/rpc/internal/svc"
	"mathless-backend/app/library/cmd/rpc/library"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLibrariesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLibrariesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLibrariesLogic {
	return &GetLibrariesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLibrariesLogic) GetLibraries(in *library.GetLibrariesRequest) (*library.GetLibrariesResponse, error) {
	libraries, err := l.svcCtx.LibraryModel.FindLibrariesByUid(l.ctx, in.Uid)
	if err != nil {
		return nil, err
	}

	var resp []*library.Library
	for _, libraryModel := range libraries {
		var libraryResp library.Library
		libraryResp.Id = libraryModel.Id
		libraryResp.Name = libraryModel.Name
		libraryResp.Description = libraryModel.Description.String
		libraryResp.Status = libraryModel.Status
		resp = append(resp, &libraryResp)
	}
	return &library.GetLibrariesResponse{
		Libraries: resp,
	}, nil
}
