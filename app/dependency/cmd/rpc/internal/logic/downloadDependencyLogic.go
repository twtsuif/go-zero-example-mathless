package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"

	"mathless-backend/app/dependency/cmd/rpc/dependency"
	"mathless-backend/app/dependency/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadDependencyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDownloadDependencyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadDependencyLogic {
	return &DownloadDependencyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DownloadDependency 下载依赖
func (l *DownloadDependencyLogic) DownloadDependency(in *dependency.DownloadDependencyRequest) (*dependency.DownloadDependencyResponse, error) {
	data := newInvokeCreateLayerOssRequest(in.FileLocation, in.Requirements)
	result, err := l.svcCtx.FcTool.InvokeFunction("mathless-plugin", "createLayerOSS", data)
	if err != nil {
		return nil, err
	}
	if result != "success" {
		return nil, errors.New("下载依赖失败 请核对依赖是否正确")
	}

	return &dependency.DownloadDependencyResponse{}, nil
}

type InvokeCreateLayerOssRequest struct {
	ObjectName   string `json:"objectName"`
	Requirements string `json:"requirements"`
}

func newInvokeCreateLayerOssRequest(objectName, requirements string) []byte {
	request := InvokeCreateLayerOssRequest{
		ObjectName:   objectName,
		Requirements: requirements,
	}
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	_ = jsonEncoder.Encode(request)
	return bf.Bytes()
}
