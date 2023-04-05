package logic

import (
	"context"

	"mathless-backend/app/library/cmd/rpc/internal/svc"
	"mathless-backend/app/library/cmd/rpc/library"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLibraryDeniedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLibraryDeniedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLibraryDeniedLogic {
	return &PublishLibraryDeniedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLibraryDeniedLogic) PublishLibraryDenied(in *library.PublishLibraryDeniedRequest) (*library.PublishLibraryDeniedResponse, error) {
	for _, libraryId := range in.LibraryIds {
		err := l.svcCtx.LibraryModel.ChangeLibraryStatus(l.ctx, libraryId, 0)
		if err != nil {
			return nil, err
		}
	}
	return &library.PublishLibraryDeniedResponse{}, nil
}
