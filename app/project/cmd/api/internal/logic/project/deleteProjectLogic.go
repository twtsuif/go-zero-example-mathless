package project

import (
	"context"
	"mathless-backend/app/project/cmd/rpc/projectservice"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/project/cmd/api/internal/svc"
	"mathless-backend/app/project/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProjectLogic {
	return &DeleteProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProjectLogic) DeleteProject(req *types.DeleteProjectReq) (resp *types.DeleteProjectResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := projectservice.DeleteProjectRequest{
		Uid:       uid,
		ProjectId: req.ProjectId,
	}
	_, err = l.svcCtx.ProjectRpc.DeleteProject(l.ctx, &rpcReq)
	return
}
