// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/subscription/subscription_policy.proto

package subscription

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SubscriptionPolicy data model
type SubscriptionPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TenantId             int32  `protobuf:"varint,2,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	Name                 string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	QuotaType            string `protobuf:"bytes,4,opt,name=quotaType,proto3" json:"quotaType,omitempty"`
	GraphQLMaxComplexity int32  `protobuf:"varint,5,opt,name=graphQLMaxComplexity,proto3" json:"graphQLMaxComplexity,omitempty"`
	GraphQLMaxDepth      int32  `protobuf:"varint,6,opt,name=graphQLMaxDepth,proto3" json:"graphQLMaxDepth,omitempty"`
	RateLimitCount       int32  `protobuf:"varint,7,opt,name=rateLimitCount,proto3" json:"rateLimitCount,omitempty"`
	RateLimitTimeUnit    string `protobuf:"bytes,8,opt,name=rateLimitTimeUnit,proto3" json:"rateLimitTimeUnit,omitempty"`
	StopOnQuotaReach     bool   `protobuf:"varint,9,opt,name=stopOnQuotaReach,proto3" json:"stopOnQuotaReach,omitempty"`
	TenantDomain         string `protobuf:"bytes,10,opt,name=tenantDomain,proto3" json:"tenantDomain,omitempty"`
	Timestamp            int64  `protobuf:"varint,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SubscriptionPolicy) Reset() {
	*x = SubscriptionPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_subscription_subscription_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionPolicy) ProtoMessage() {}

func (x *SubscriptionPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_subscription_subscription_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionPolicy.ProtoReflect.Descriptor instead.
func (*SubscriptionPolicy) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_subscription_subscription_policy_proto_rawDescGZIP(), []int{0}
}

func (x *SubscriptionPolicy) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SubscriptionPolicy) GetTenantId() int32 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *SubscriptionPolicy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubscriptionPolicy) GetQuotaType() string {
	if x != nil {
		return x.QuotaType
	}
	return ""
}

func (x *SubscriptionPolicy) GetGraphQLMaxComplexity() int32 {
	if x != nil {
		return x.GraphQLMaxComplexity
	}
	return 0
}

func (x *SubscriptionPolicy) GetGraphQLMaxDepth() int32 {
	if x != nil {
		return x.GraphQLMaxDepth
	}
	return 0
}

func (x *SubscriptionPolicy) GetRateLimitCount() int32 {
	if x != nil {
		return x.RateLimitCount
	}
	return 0
}

func (x *SubscriptionPolicy) GetRateLimitTimeUnit() string {
	if x != nil {
		return x.RateLimitTimeUnit
	}
	return ""
}

func (x *SubscriptionPolicy) GetStopOnQuotaReach() bool {
	if x != nil {
		return x.StopOnQuotaReach
	}
	return false
}

func (x *SubscriptionPolicy) GetTenantDomain() string {
	if x != nil {
		return x.TenantDomain
	}
	return ""
}

func (x *SubscriptionPolicy) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_wso2_discovery_subscription_subscription_policy_proto protoreflect.FileDescriptor

var file_wso2_discovery_subscription_subscription_policy_proto_rawDesc = []byte{
	0x0a, 0x35, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x94, 0x03, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x51, 0x4c, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x74,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x67, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c,
	0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a,
	0x0f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x4d, 0x61, 0x78, 0x44, 0x65, 0x70, 0x74, 0x68,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x4d,
	0x61, 0x78, 0x44, 0x65, 0x70, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x11, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x55, 0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x2a, 0x0a,
	0x10, 0x73, 0x74, 0x6f, 0x70, 0x4f, 0x6e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x61, 0x63,
	0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x74, 0x6f, 0x70, 0x4f, 0x6e, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x61, 0x63, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x9a, 0x01, 0x0a, 0x2c,
	0x6f, 0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x61, 0x70, 0x6b, 0x2e, 0x65, 0x6e, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x17, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_subscription_subscription_policy_proto_rawDescOnce sync.Once
	file_wso2_discovery_subscription_subscription_policy_proto_rawDescData = file_wso2_discovery_subscription_subscription_policy_proto_rawDesc
)

func file_wso2_discovery_subscription_subscription_policy_proto_rawDescGZIP() []byte {
	file_wso2_discovery_subscription_subscription_policy_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_subscription_subscription_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_subscription_subscription_policy_proto_rawDescData)
	})
	return file_wso2_discovery_subscription_subscription_policy_proto_rawDescData
}

var file_wso2_discovery_subscription_subscription_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wso2_discovery_subscription_subscription_policy_proto_goTypes = []interface{}{
	(*SubscriptionPolicy)(nil), // 0: wso2.discovery.subscription.SubscriptionPolicy
}
var file_wso2_discovery_subscription_subscription_policy_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wso2_discovery_subscription_subscription_policy_proto_init() }
func file_wso2_discovery_subscription_subscription_policy_proto_init() {
	if File_wso2_discovery_subscription_subscription_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_subscription_subscription_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wso2_discovery_subscription_subscription_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_subscription_subscription_policy_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_subscription_subscription_policy_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_subscription_subscription_policy_proto_msgTypes,
	}.Build()
	File_wso2_discovery_subscription_subscription_policy_proto = out.File
	file_wso2_discovery_subscription_subscription_policy_proto_rawDesc = nil
	file_wso2_discovery_subscription_subscription_policy_proto_goTypes = nil
	file_wso2_discovery_subscription_subscription_policy_proto_depIdxs = nil
}
