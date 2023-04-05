package dependency

import (
	"context"
	"mathless-backend/app/dependency/cmd/rpc/dependencyservice"
	jwtTool "mathless-backend/common/tool/jwt"
	"strconv"

	"mathless-backend/app/dependency/cmd/api/internal/svc"
	"mathless-backend/app/dependency/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadDependencyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadDependencyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadDependencyLogic {
	return &DownloadDependencyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadDependencyLogic) DownloadDependency(req *types.DownloadDependencyReq) (resp *types.DownloadDependencyResp, err error) {
	// 上下文获取uid
	uid := jwtTool.GetUidFromCtx(l.ctx)
	bucket := l.svcCtx.Config.AliyunConf.OssConf.Bucket

	var rpcReq dependencyservice.DownloadDependencyRequest

	// file和layer的前缀
	downloadLocation := StoreLocationDir + "/" + strconv.FormatInt(uid, 10) + "/"
	layerName := strconv.FormatInt(uid, 10) + "_"

	// 根据类型拼接
	if req.Typ == ProjectDependencyType {
		downloadLocation += ProjectPrefix + "/" + req.ProjectId + ".zip"
		layerName += ProjectPrefix + "_" + req.ProjectId
	} else if req.Typ == LibraryDependencyType {
		downloadLocation += FunctionPrefix + "/" + req.FunctionId + ".zip"
		layerName += FunctionPrefix + "_" + req.FunctionId
	}

	// 若依赖为空 则删除oss和layer
	if req.Requirements == "" {
		err = l.svcCtx.OssTool.DeleteOss(downloadLocation)
		if err != nil {
			return nil, err
		}
		err = l.svcCtx.FcTool.DeleteLayer(layerName)
		if err != nil {
			return nil, err
		}
	}

	// 构造并发送rpc请求
	rpcReq.FileLocation = downloadLocation
	rpcReq.Requirements = req.Requirements
	_, err = l.svcCtx.DependencyRpc.DownloadDependency(l.ctx, &rpcReq)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.FcTool.CreateLayer(layerName, bucket, downloadLocation)

	return
}
