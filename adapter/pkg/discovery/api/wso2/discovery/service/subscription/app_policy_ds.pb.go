// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/service/subscription/app_policy_ds.proto

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

var File_wso2_discovery_service_subscription_app_policy_ds_proto protoreflect.FileDescriptor

var file_wso2_discovery_service_subscription_app_policy_ds_proto_rawDesc = []byte{
	0x0a, 0x37, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa3, 0x01, 0x0a, 0x21, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a, 0x19, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x99, 0x01, 0x0a, 0x34,
	0x6f, 0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x61, 0x70, 0x6b, 0x2e, 0x65, 0x6e, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x10, 0x41, 0x70, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44,
	0x53, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_wso2_discovery_service_subscription_app_policy_ds_proto_goTypes = []interface{}{
	(*v3.DiscoveryRequest)(nil),  // 0: envoy.service.discovery.v3.DiscoveryRequest
	(*v3.DiscoveryResponse)(nil), // 1: envoy.service.discovery.v3.DiscoveryResponse
}
var file_wso2_discovery_service_subscription_app_policy_ds_proto_depIdxs = []int32{
	0, // 0: discovery.service.subscription.ApplicationPolicyDiscoveryService.StreamApplicationPolicies:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	1, // 1: discovery.service.subscription.ApplicationPolicyDiscoveryService.StreamApplicationPolicies:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wso2_discovery_service_subscription_app_policy_ds_proto_init() }
func file_wso2_discovery_service_subscription_app_policy_ds_proto_init() {
	if File_wso2_discovery_service_subscription_app_policy_ds_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wso2_discovery_service_subscription_app_policy_ds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wso2_discovery_service_subscription_app_policy_ds_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_service_subscription_app_policy_ds_proto_depIdxs,
	}.Build()
	File_wso2_discovery_service_subscription_app_policy_ds_proto = out.File
	file_wso2_discovery_service_subscription_app_policy_ds_proto_rawDesc = nil
	file_wso2_discovery_service_subscription_app_policy_ds_proto_goTypes = nil
	file_wso2_discovery_service_subscription_app_policy_ds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApplicationPolicyDiscoveryServiceClient is the client API for ApplicationPolicyDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationPolicyDiscoveryServiceClient interface {
	StreamApplicationPolicies(ctx context.Context, opts ...grpc.CallOption) (ApplicationPolicyDiscoveryService_StreamApplicationPoliciesClient, error)
}

type applicationPolicyDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationPolicyDiscoveryServiceClient(cc grpc.ClientConnInterface) ApplicationPolicyDiscoveryServiceClient {
	return &applicationPolicyDiscoveryServiceClient{cc}
}

func (c *applicationPolicyDiscoveryServiceClient) StreamApplicationPolicies(ctx context.Context, opts ...grpc.CallOption) (ApplicationPolicyDiscoveryService_StreamApplicationPoliciesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ApplicationPolicyDiscoveryService_serviceDesc.Streams[0], "/discovery.service.subscription.ApplicationPolicyDiscoveryService/StreamApplicationPolicies", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationPolicyDiscoveryServiceStreamApplicationPoliciesClient{stream}
	return x, nil
}

type ApplicationPolicyDiscoveryService_StreamApplicationPoliciesClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type applicationPolicyDiscoveryServiceStreamApplicationPoliciesClient struct {
	grpc.ClientStream
}

func (x *applicationPolicyDiscoveryServiceStreamApplicationPoliciesClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *applicationPolicyDiscoveryServiceStreamApplicationPoliciesClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ApplicationPolicyDiscoveryServiceServer is the server API for ApplicationPolicyDiscoveryService service.
type ApplicationPolicyDiscoveryServiceServer interface {
	StreamApplicationPolicies(ApplicationPolicyDiscoveryService_StreamApplicationPoliciesServer) error
}

// UnimplementedApplicationPolicyDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationPolicyDiscoveryServiceServer struct {
}

func (*UnimplementedApplicationPolicyDiscoveryServiceServer) StreamApplicationPolicies(ApplicationPolicyDiscoveryService_StreamApplicationPoliciesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamApplicationPolicies not implemented")
}

func RegisterApplicationPolicyDiscoveryServiceServer(s *grpc.Server, srv ApplicationPolicyDiscoveryServiceServer) {
	s.RegisterService(&_ApplicationPolicyDiscoveryService_serviceDesc, srv)
}

func _ApplicationPolicyDiscoveryService_StreamApplicationPolicies_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ApplicationPolicyDiscoveryServiceServer).StreamApplicationPolicies(&applicationPolicyDiscoveryServiceStreamApplicationPoliciesServer{stream})
}

type ApplicationPolicyDiscoveryService_StreamApplicationPoliciesServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type applicationPolicyDiscoveryServiceStreamApplicationPoliciesServer struct {
	grpc.ServerStream
}

func (x *applicationPolicyDiscoveryServiceStreamApplicationPoliciesServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *applicationPolicyDiscoveryServiceStreamApplicationPoliciesServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ApplicationPolicyDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.service.subscription.ApplicationPolicyDiscoveryService",
	HandlerType: (*ApplicationPolicyDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamApplicationPolicies",
			Handler:       _ApplicationPolicyDiscoveryService_StreamApplicationPolicies_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "wso2/discovery/service/subscription/app_policy_ds.proto",
}
