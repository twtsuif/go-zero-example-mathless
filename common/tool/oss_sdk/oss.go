package oss_sdk

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type OssTool struct {
	BucketClient *oss.Bucket
}

func NewOssTool(endpoint, accessKeyID, accessKeySecret, bucket string) *OssTool {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		panic(err)
	}
	bucketClient, err := client.Bucket(bucket)
	return &OssTool{bucketClient}
}

func (ossTool OssTool) DeleteOss(filePath string) error {
	return ossTool.BucketClient.DeleteObject(filePath)
}
