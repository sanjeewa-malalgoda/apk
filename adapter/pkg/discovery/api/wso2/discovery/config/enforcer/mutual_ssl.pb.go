//  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
//
//  WSO2 LLC. licenses this file to you under the Apache License,
//  Version 2.0 (the "License"); you may not use this file except
//  in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an
//  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
//  KIND, either express or implied.  See the License for the
//  specific language governing permissions and limitations
//  under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/config/enforcer/mutual_ssl.proto

package enforcer

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

// Client certificate header store model
type MutualSSL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertificateHeader               string `protobuf:"bytes,1,opt,name=certificateHeader,proto3" json:"certificateHeader,omitempty"`
	EnableClientValidation          bool   `protobuf:"varint,2,opt,name=enableClientValidation,proto3" json:"enableClientValidation,omitempty"`
	ClientCertificateEncode         bool   `protobuf:"varint,3,opt,name=clientCertificateEncode,proto3" json:"clientCertificateEncode,omitempty"`
	EnableOutboundCertificateHeader bool   `protobuf:"varint,4,opt,name=enableOutboundCertificateHeader,proto3" json:"enableOutboundCertificateHeader,omitempty"`
}

func (x *MutualSSL) Reset() {
	*x = MutualSSL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_config_enforcer_mutual_ssl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutualSSL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutualSSL) ProtoMessage() {}

func (x *MutualSSL) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_config_enforcer_mutual_ssl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutualSSL.ProtoReflect.Descriptor instead.
func (*MutualSSL) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDescGZIP(), []int{0}
}

func (x *MutualSSL) GetCertificateHeader() string {
	if x != nil {
		return x.CertificateHeader
	}
	return ""
}

func (x *MutualSSL) GetEnableClientValidation() bool {
	if x != nil {
		return x.EnableClientValidation
	}
	return false
}

func (x *MutualSSL) GetClientCertificateEncode() bool {
	if x != nil {
		return x.ClientCertificateEncode
	}
	return false
}

func (x *MutualSSL) GetEnableOutboundCertificateHeader() bool {
	if x != nil {
		return x.EnableOutboundCertificateHeader
	}
	return false
}

var File_wso2_discovery_config_enforcer_mutual_ssl_proto protoreflect.FileDescriptor

var file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72,
	0x2f, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x72, 0x22, 0xf5, 0x01, 0x0a, 0x09, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x53, 0x4c, 0x12,
	0x2c, 0x0a, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a,
	0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x17, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x48, 0x0a, 0x1f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x93, 0x01, 0x0a, 0x2f, 0x6f, 0x72,
	0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x61, 0x70, 0x6b, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x42, 0x0e, 0x4d,
	0x75, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x53, 0x4c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x3b, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDescOnce sync.Once
	file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDescData = file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDesc
)

func file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDescGZIP() []byte {
	file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDescData)
	})
	return file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDescData
}

var file_wso2_discovery_config_enforcer_mutual_ssl_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wso2_discovery_config_enforcer_mutual_ssl_proto_goTypes = []interface{}{
	(*MutualSSL)(nil), // 0: wso2.discovery.config.enforcer.MutualSSL
}
var file_wso2_discovery_config_enforcer_mutual_ssl_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wso2_discovery_config_enforcer_mutual_ssl_proto_init() }
func file_wso2_discovery_config_enforcer_mutual_ssl_proto_init() {
	if File_wso2_discovery_config_enforcer_mutual_ssl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_config_enforcer_mutual_ssl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutualSSL); i {
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
			RawDescriptor: file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_config_enforcer_mutual_ssl_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_config_enforcer_mutual_ssl_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_config_enforcer_mutual_ssl_proto_msgTypes,
	}.Build()
	File_wso2_discovery_config_enforcer_mutual_ssl_proto = out.File
	file_wso2_discovery_config_enforcer_mutual_ssl_proto_rawDesc = nil
	file_wso2_discovery_config_enforcer_mutual_ssl_proto_goTypes = nil
	file_wso2_discovery_config_enforcer_mutual_ssl_proto_depIdxs = nil
}
