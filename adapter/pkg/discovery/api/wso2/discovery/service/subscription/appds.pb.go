// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/service/subscription/appds.proto

package subscription

import (
	context "context"
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_wso2_discovery_service_subscription_appds_proto protoreflect.FileDescriptor

var file_wso2_discovery_service_subscription_appds_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x96, 0x01,
	0x0a, 0x1b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a,
	0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33,
	0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x93, 0x01, 0x0a, 0x34, 0x6f, 0x72, 0x67, 0x2e, 0x77,
	0x73, 0x6f, 0x32, 0x2e, 0x61, 0x70, 0x6b, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0a, 0x41, 0x70, 0x70, 0x44, 0x53, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_wso2_discovery_service_subscription_appds_proto_goTypes = []interface{}{
	(*v3.DiscoveryRequest)(nil),  // 0: envoy.service.discovery.v3.DiscoveryRequest
	(*v3.DiscoveryResponse)(nil), // 1: envoy.service.discovery.v3.DiscoveryResponse
}
var file_wso2_discovery_service_subscription_appds_proto_depIdxs = []int32{
	0, // 0: discovery.service.subscription.ApplicationDiscoveryService.StreamApplications:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	1, // 1: discovery.service.subscription.ApplicationDiscoveryService.StreamApplications:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wso2_discovery_service_subscription_appds_proto_init() }
func file_wso2_discovery_service_subscription_appds_proto_init() {
	if File_wso2_discovery_service_subscription_appds_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wso2_discovery_service_subscription_appds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wso2_discovery_service_subscription_appds_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_service_subscription_appds_proto_depIdxs,
	}.Build()
	File_wso2_discovery_service_subscription_appds_proto = out.File
	file_wso2_discovery_service_subscription_appds_proto_rawDesc = nil
	file_wso2_discovery_service_subscription_appds_proto_goTypes = nil
	file_wso2_discovery_service_subscription_appds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApplicationDiscoveryServiceClient is the client API for ApplicationDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationDiscoveryServiceClient interface {
	StreamApplications(ctx context.Context, opts ...grpc.CallOption) (ApplicationDiscoveryService_StreamApplicationsClient, error)
}

type applicationDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationDiscoveryServiceClient(cc grpc.ClientConnInterface) ApplicationDiscoveryServiceClient {
	return &applicationDiscoveryServiceClient{cc}
}

func (c *applicationDiscoveryServiceClient) StreamApplications(ctx context.Context, opts ...grpc.CallOption) (ApplicationDiscoveryService_StreamApplicationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ApplicationDiscoveryService_serviceDesc.Streams[0], "/discovery.service.subscription.ApplicationDiscoveryService/StreamApplications", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationDiscoveryServiceStreamApplicationsClient{stream}
	return x, nil
}

type ApplicationDiscoveryService_StreamApplicationsClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type applicationDiscoveryServiceStreamApplicationsClient struct {
	grpc.ClientStream
}

func (x *applicationDiscoveryServiceStreamApplicationsClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *applicationDiscoveryServiceStreamApplicationsClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ApplicationDiscoveryServiceServer is the server API for ApplicationDiscoveryService service.
type ApplicationDiscoveryServiceServer interface {
	StreamApplications(ApplicationDiscoveryService_StreamApplicationsServer) error
}

// UnimplementedApplicationDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationDiscoveryServiceServer struct {
}

func (*UnimplementedApplicationDiscoveryServiceServer) StreamApplications(ApplicationDiscoveryService_StreamApplicationsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamApplications not implemented")
}

func RegisterApplicationDiscoveryServiceServer(s *grpc.Server, srv ApplicationDiscoveryServiceServer) {
	s.RegisterService(&_ApplicationDiscoveryService_serviceDesc, srv)
}

func _ApplicationDiscoveryService_StreamApplications_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ApplicationDiscoveryServiceServer).StreamApplications(&applicationDiscoveryServiceStreamApplicationsServer{stream})
}

type ApplicationDiscoveryService_StreamApplicationsServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type applicationDiscoveryServiceStreamApplicationsServer struct {
	grpc.ServerStream
}

func (x *applicationDiscoveryServiceStreamApplicationsServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *applicationDiscoveryServiceStreamApplicationsServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ApplicationDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.service.subscription.ApplicationDiscoveryService",
	HandlerType: (*ApplicationDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamApplications",
			Handler:       _ApplicationDiscoveryService_StreamApplications_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "wso2/discovery/service/subscription/appds.proto",
}
