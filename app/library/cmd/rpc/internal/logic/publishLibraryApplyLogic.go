package logic

import (
	"context"

	"mathless-backend/app/library/cmd/rpc/internal/svc"
	"mathless-backend/app/library/cmd/rpc/library"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLibraryApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLibraryApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLibraryApplyLogic {
	return &PublishLibraryApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLibraryApplyLogic) PublishLibraryApply(in *library.PublishLibraryApplyRequest) (*library.PublishLibraryApplyResponse, error) {
	err := l.svcCtx.LibraryModel.ChangeLibraryStatus(l.ctx, in.LibraryId, 1)
	if err != nil {
		return nil, err
	}
	return &library.PublishLibraryApplyResponse{}, nil
}
