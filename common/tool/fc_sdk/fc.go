package fc_sdk

import "github.com/aliyun/fc-go-sdk"

type FcTool struct {
	client *fc.Client
}

func NewFcTool(endpoint, apiVersion, accessKeyID, accessKeySecret string) *FcTool {
	client, err := fc.NewClient(endpoint, apiVersion, accessKeyID, accessKeySecret)
	if err != nil {
		panic("阿里云连接失败")
	}
	return &FcTool{client: client}
}

// CreateService 创建service
// param: service名
func (fcTool FcTool) CreateService(serviceName string) error {
	_, err := fcTool.client.CreateService(fc.NewCreateServiceInput().WithServiceName(serviceName))
	return err
}

// CreateFunction 创建函数
// param: service名 function名
func (fcTool FcTool) CreateFunction(serviceName, functionName string) error {
	createFunctionInput := fc.NewCreateFunctionInput(serviceName).
		WithFunctionName(functionName).
		WithHandler("index.handler").
		WithRuntime("python3").
		WithCode(fc.NewCode().WithFiles("code/index.py")).
		WithTimeout(5).
		WithMemorySize(32768).
		WithInstanceType("c1") // 指定函数所属实例类型为性能实例。

	_, err := fcTool.client.CreateFunction(createFunctionInput)
	return err
}

// InvokeFunction 调用函数
// param: service名 function名 data
func (fcTool FcTool) InvokeFunction(serviceName, functionName string, data []byte) (string, error) {
	input := fc.NewInvokeFunctionInput(serviceName, functionName).
		WithPayload(data).
		WithLogType("None")
	output, err := fcTool.client.InvokeFunction(input)
	if err != nil {
		return "", err
	}
	return string(output.Payload), nil
}

// DeleteLayer 删除层
func (fcTool FcTool) DeleteLayer(layerName string) error {
	layerVersions, err := fcTool.client.ListLayerVersions(fc.NewListLayerVersionsInput(layerName, 1))
	if err != nil {
		return err
	}
	for _, version := range layerVersions.Layers {
		_, err := fcTool.client.DeleteLayerVersion(fc.NewDeleteLayerVersionInput(layerName, version.Version))
		if err != nil {
			return err
		}
	}
	return err
}

// CreateLayer 创建层
func (fcTool FcTool) CreateLayer(layerName string, bucketName string, objectName string) (int32, error) {
	input := fc.NewPublishLayerVersionInput().
		WithLayerName(layerName).
		WithCode(fc.NewCode().WithOSSBucketName(bucketName).WithOSSObjectName(objectName)).
		WithCompatibleRuntime([]string{"python3"})
	publishLayerVersionOutput, err := fcTool.client.PublishLayerVersion(input)
	return publishLayerVersionOutput.Version, err
}
