package project

import (
	"context"
	"mathless-backend/app/project/cmd/rpc/projectservice"
	jwtTool "mathless-backend/common/tool/jwt"

	"mathless-backend/app/project/cmd/api/internal/svc"
	"mathless-backend/app/project/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectLogic {
	return &UpdateProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProjectLogic) UpdateProject(req *types.UpdateProjectReq) (resp *types.UpdateProjectResp, err error) {
	uid := jwtTool.GetUidFromCtx(l.ctx)
	rpcReq := projectservice.UpdateProjectRequest{
		Uid:          uid,
		ProjectId:    req.ProjectId,
		ProjectName:  req.ProjectName,
		FilePathJson: req.FilePathJson,
	}
	_, err = l.svcCtx.ProjectRpc.UpdateProject(l.ctx, &rpcReq)
	return
}
