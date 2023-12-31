// Code generated by Kitex v0.7.2. DO NOT EDIT.

package fileservice

import (
	fileservice "fileService/kitex_gen/fileService"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return fileServiceServiceInfo
}

var fileServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "fileService"
	handlerType := (*fileservice.FileService)(nil)
	methods := map[string]kitex.MethodInfo{}
	extra := map[string]interface{}{
		"PackageName":     "fileservice",
		"ServiceFilePath": `fileService.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}
