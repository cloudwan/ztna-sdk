// Code generated by protoc-gen-goten-client
// Service: ZTNA
// DO NOT EDIT!!!

package ztna_client

import (
	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	access_point_client "github.com/cloudwan/ztna-sdk/client/v1alpha/access_point"
	broker_client "github.com/cloudwan/ztna-sdk/client/v1alpha/broker"
	port_forwarding_service_client "github.com/cloudwan/ztna-sdk/client/v1alpha/port_forwarding_service"
	access_point "github.com/cloudwan/ztna-sdk/resources/v1alpha/access_point"
	port_forwarding_service "github.com/cloudwan/ztna-sdk/resources/v1alpha/port_forwarding_service"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &access_point.AccessPoint{}
	_ = &access_point_client.GetAccessPointRequest{}
	_ = &port_forwarding_service.PortForwardingService{}
	_ = &port_forwarding_service_client.GetPortForwardingServiceRequest{}
)

var (
	descriptorInitialized bool
	ztnaDescriptor        *ZTNADescriptor
)

type ZTNADescriptor struct{}

func (d *ZTNADescriptor) GetServiceName() string {
	return "ztna"
}

func (d *ZTNADescriptor) GetServiceDomain() string {
	return "ztna.edgelq.com"
}

func (d *ZTNADescriptor) GetVersion() string {
	return "v1alpha"
}

func (d *ZTNADescriptor) GetNextVersion() string {

	return ""
}

func (d *ZTNADescriptor) AllResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		access_point.GetDescriptor(),
		port_forwarding_service.GetDescriptor(),
	}
}

func (d *ZTNADescriptor) AllApiDescriptors() []gotenclient.ApiDescriptor {
	return []gotenclient.ApiDescriptor{
		access_point_client.GetAccessPointServiceDescriptor(),
		broker_client.GetBrokerServiceDescriptor(),
		port_forwarding_service_client.GetPortForwardingServiceServiceDescriptor(),
	}
}

func (d *ZTNADescriptor) AllImportedServiceInfos() []gotenclient.ServiceImportInfo {
	return []gotenclient.ServiceImportInfo{
		{
			Domain:  "iam.edgelq.com",
			Version: "v1alpha",
		},
	}
}

func GetZTNADescriptor() *ZTNADescriptor {
	return ztnaDescriptor
}

func initDescriptor() {
	ztnaDescriptor = &ZTNADescriptor{}
	gotenclient.GetRegistry().RegisterSvcDescriptor(ztnaDescriptor)
}

func init() {
	if !descriptorInitialized {
		initDescriptor()
		descriptorInitialized = true
	}
}
