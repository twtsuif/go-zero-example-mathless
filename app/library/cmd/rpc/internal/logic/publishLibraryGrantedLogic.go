package logic

import (
	"context"

	"mathless-backend/app/library/cmd/rpc/internal/svc"
	"mathless-backend/app/library/cmd/rpc/library"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLibraryGrantedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLibraryGrantedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLibraryGrantedLogic {
	return &PublishLibraryGrantedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLibraryGrantedLogic) PublishLibraryGranted(in *library.PublishLibraryGrantedRequest) (*library.PublishLibraryGrantedResponse, error) {
	for _, libraryId := range in.LibraryIds {
		err := l.svcCtx.LibraryModel.ChangeLibraryStatus(l.ctx, libraryId, 2)
		if err != nil {
			return nil, err
		}
	}
	return &library.PublishLibraryGrantedResponse{}, nil
}
