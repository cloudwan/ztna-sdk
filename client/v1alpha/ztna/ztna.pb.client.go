// Code generated by protoc-gen-goten-client
// Service: ZTNA
// DO NOT EDIT!!!

package ztna_client

import (
	"google.golang.org/grpc"
)

// proto imports
import (
	access_point_client "github.com/cloudwan/ztna-sdk/client/v1alpha/access_point"
	broker_client "github.com/cloudwan/ztna-sdk/client/v1alpha/broker"
	port_forwarding_service_client "github.com/cloudwan/ztna-sdk/client/v1alpha/port_forwarding_service"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &access_point_client.GetAccessPointRequest{}
	_ = &port_forwarding_service_client.GetPortForwardingServiceRequest{}
)

type ZTNAClient interface {
	access_point_client.AccessPointServiceClient
	broker_client.BrokerServiceClient
	port_forwarding_service_client.PortForwardingServiceServiceClient
}

type ztnaClient struct {
	access_point_client.AccessPointServiceClient
	broker_client.BrokerServiceClient
	port_forwarding_service_client.PortForwardingServiceServiceClient
}

func NewZTNAClient(cc grpc.ClientConnInterface) ZTNAClient {
	return &ztnaClient{
		AccessPointServiceClient:           access_point_client.NewAccessPointServiceClient(cc),
		BrokerServiceClient:                broker_client.NewBrokerServiceClient(cc),
		PortForwardingServiceServiceClient: port_forwarding_service_client.NewPortForwardingServiceServiceClient(cc),
	}
}
