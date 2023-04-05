package logic

import (
	"bytes"
	"context"
	"encoding/json"

	"mathless-backend/app/dependency/cmd/rpc/dependency"
	"mathless-backend/app/dependency/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDependencyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDependencyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDependencyLogic {
	return &GetDependencyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetDependency 获取依赖内容
func (l *GetDependencyLogic) GetDependency(in *dependency.GetDependencyRequest) (*dependency.GetDependencyResponse, error) {
	data := newInvokeGetLayerDependencyRequestRequest(in.DependencyFileLocation)
	result, err := l.svcCtx.FcTool.InvokeFunction("mathless-plugin", "getLayerRequirements", data)
	return &dependency.GetDependencyResponse{PackagesJson: result}, err
}

type InvokeGetLayerDependencyRequest struct {
	ObjectName string `json:"objectName"`
}

func newInvokeGetLayerDependencyRequestRequest(objectName string) []byte {
	request := InvokeGetLayerDependencyRequest{
		ObjectName: objectName,
	}
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	_ = jsonEncoder.Encode(request)
	return bf.Bytes()
}
