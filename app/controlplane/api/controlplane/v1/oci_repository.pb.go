//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: controlplane/v1/oci_repository.proto

package v1

import (
	_ "github.com/chainloop-dev/chainloop/app/controlplane/api/errors/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type OCIRepositoryErrorReason int32

const (
	// TODO: add support for PRECONDITION_FAILED
	OCIRepositoryErrorReason_OCI_REPOSITORY_ERROR_REASON_UNSPECIFIED OCIRepositoryErrorReason = 0
	OCIRepositoryErrorReason_OCI_REPOSITORY_ERROR_REASON_REQUIRED    OCIRepositoryErrorReason = 1
	// The repository does not seem to be operational
	// a previous validation has failed
	OCIRepositoryErrorReason_OCI_REPOSITORY_ERROR_REASON_INVALID OCIRepositoryErrorReason = 2
)

// Enum value maps for OCIRepositoryErrorReason.
var (
	OCIRepositoryErrorReason_name = map[int32]string{
		0: "OCI_REPOSITORY_ERROR_REASON_UNSPECIFIED",
		1: "OCI_REPOSITORY_ERROR_REASON_REQUIRED",
		2: "OCI_REPOSITORY_ERROR_REASON_INVALID",
	}
	OCIRepositoryErrorReason_value = map[string]int32{
		"OCI_REPOSITORY_ERROR_REASON_UNSPECIFIED": 0,
		"OCI_REPOSITORY_ERROR_REASON_REQUIRED":    1,
		"OCI_REPOSITORY_ERROR_REASON_INVALID":     2,
	}
)

func (x OCIRepositoryErrorReason) Enum() *OCIRepositoryErrorReason {
	p := new(OCIRepositoryErrorReason)
	*p = x
	return p
}

func (x OCIRepositoryErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OCIRepositoryErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_controlplane_v1_oci_repository_proto_enumTypes[0].Descriptor()
}

func (OCIRepositoryErrorReason) Type() protoreflect.EnumType {
	return &file_controlplane_v1_oci_repository_proto_enumTypes[0]
}

func (x OCIRepositoryErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OCIRepositoryErrorReason.Descriptor instead.
func (OCIRepositoryErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_controlplane_v1_oci_repository_proto_rawDescGZIP(), []int{0}
}

type OCIRepositoryServiceSaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FQDN of the OCI repository, including paths
	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// Types that are assignable to Credentials:
	//
	//	*OCIRepositoryServiceSaveRequest_KeyPair
	Credentials isOCIRepositoryServiceSaveRequest_Credentials `protobuf_oneof:"credentials"`
}

func (x *OCIRepositoryServiceSaveRequest) Reset() {
	*x = OCIRepositoryServiceSaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_oci_repository_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCIRepositoryServiceSaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCIRepositoryServiceSaveRequest) ProtoMessage() {}

func (x *OCIRepositoryServiceSaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_oci_repository_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCIRepositoryServiceSaveRequest.ProtoReflect.Descriptor instead.
func (*OCIRepositoryServiceSaveRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_oci_repository_proto_rawDescGZIP(), []int{0}
}

func (x *OCIRepositoryServiceSaveRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (m *OCIRepositoryServiceSaveRequest) GetCredentials() isOCIRepositoryServiceSaveRequest_Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (x *OCIRepositoryServiceSaveRequest) GetKeyPair() *OCIRepositoryServiceSaveRequest_Keypair {
	if x, ok := x.GetCredentials().(*OCIRepositoryServiceSaveRequest_KeyPair); ok {
		return x.KeyPair
	}
	return nil
}

type isOCIRepositoryServiceSaveRequest_Credentials interface {
	isOCIRepositoryServiceSaveRequest_Credentials()
}

type OCIRepositoryServiceSaveRequest_KeyPair struct {
	KeyPair *OCIRepositoryServiceSaveRequest_Keypair `protobuf:"bytes,2,opt,name=key_pair,json=keyPair,proto3,oneof"`
}

func (*OCIRepositoryServiceSaveRequest_KeyPair) isOCIRepositoryServiceSaveRequest_Credentials() {}

type OCIRepositoryServiceSaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OCIRepositoryServiceSaveResponse) Reset() {
	*x = OCIRepositoryServiceSaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_oci_repository_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCIRepositoryServiceSaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCIRepositoryServiceSaveResponse) ProtoMessage() {}

func (x *OCIRepositoryServiceSaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_oci_repository_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCIRepositoryServiceSaveResponse.ProtoReflect.Descriptor instead.
func (*OCIRepositoryServiceSaveResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_oci_repository_proto_rawDescGZIP(), []int{1}
}

type OCIRepositoryServiceSaveRequest_Keypair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *OCIRepositoryServiceSaveRequest_Keypair) Reset() {
	*x = OCIRepositoryServiceSaveRequest_Keypair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_oci_repository_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCIRepositoryServiceSaveRequest_Keypair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCIRepositoryServiceSaveRequest_Keypair) ProtoMessage() {}

func (x *OCIRepositoryServiceSaveRequest_Keypair) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_oci_repository_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCIRepositoryServiceSaveRequest_Keypair.ProtoReflect.Descriptor instead.
func (*OCIRepositoryServiceSaveRequest_Keypair) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_oci_repository_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OCIRepositoryServiceSaveRequest_Keypair) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *OCIRepositoryServiceSaveRequest_Keypair) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_controlplane_v1_oci_repository_proto protoreflect.FileDescriptor

var file_controlplane_v1_oci_repository_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x63, 0x69, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x1f, 0x4f, 0x43, 0x49,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x55, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x69,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x43, 0x49, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4b, 0x65, 0x79, 0x70, 0x61, 0x69,
	0x72, 0x48, 0x00, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x53, 0x0a, 0x07,
	0x4b, 0x65, 0x79, 0x70, 0x61, 0x69, 0x72, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x42, 0x12, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x22, 0x0a, 0x20, 0x4f, 0x43, 0x49, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0xac, 0x01, 0x0a, 0x18, 0x4f, 0x43,
	0x49, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x27, 0x4f, 0x43, 0x49, 0x5f, 0x52, 0x45,
	0x50, 0x4f, 0x53, 0x49, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x2e, 0x0a, 0x24, 0x4f, 0x43, 0x49, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x53,
	0x49, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x01, 0x1a, 0x04, 0xb0,
	0x45, 0x93, 0x03, 0x12, 0x2d, 0x0a, 0x23, 0x4f, 0x43, 0x49, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x53,
	0x49, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xb0, 0x45,
	0x93, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x32, 0x83, 0x01, 0x0a, 0x14, 0x4f, 0x43, 0x49,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6b, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x30, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x43, 0x49, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x43,
	0x49, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4c,
	0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controlplane_v1_oci_repository_proto_rawDescOnce sync.Once
	file_controlplane_v1_oci_repository_proto_rawDescData = file_controlplane_v1_oci_repository_proto_rawDesc
)

func file_controlplane_v1_oci_repository_proto_rawDescGZIP() []byte {
	file_controlplane_v1_oci_repository_proto_rawDescOnce.Do(func() {
		file_controlplane_v1_oci_repository_proto_rawDescData = protoimpl.X.CompressGZIP(file_controlplane_v1_oci_repository_proto_rawDescData)
	})
	return file_controlplane_v1_oci_repository_proto_rawDescData
}

var file_controlplane_v1_oci_repository_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_controlplane_v1_oci_repository_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_controlplane_v1_oci_repository_proto_goTypes = []interface{}{
	(OCIRepositoryErrorReason)(0),                   // 0: controlplane.v1.OCIRepositoryErrorReason
	(*OCIRepositoryServiceSaveRequest)(nil),         // 1: controlplane.v1.OCIRepositoryServiceSaveRequest
	(*OCIRepositoryServiceSaveResponse)(nil),        // 2: controlplane.v1.OCIRepositoryServiceSaveResponse
	(*OCIRepositoryServiceSaveRequest_Keypair)(nil), // 3: controlplane.v1.OCIRepositoryServiceSaveRequest.Keypair
}
var file_controlplane_v1_oci_repository_proto_depIdxs = []int32{
	3, // 0: controlplane.v1.OCIRepositoryServiceSaveRequest.key_pair:type_name -> controlplane.v1.OCIRepositoryServiceSaveRequest.Keypair
	1, // 1: controlplane.v1.OCIRepositoryService.Save:input_type -> controlplane.v1.OCIRepositoryServiceSaveRequest
	2, // 2: controlplane.v1.OCIRepositoryService.Save:output_type -> controlplane.v1.OCIRepositoryServiceSaveResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_controlplane_v1_oci_repository_proto_init() }
func file_controlplane_v1_oci_repository_proto_init() {
	if File_controlplane_v1_oci_repository_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controlplane_v1_oci_repository_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCIRepositoryServiceSaveRequest); i {
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
		file_controlplane_v1_oci_repository_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCIRepositoryServiceSaveResponse); i {
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
		file_controlplane_v1_oci_repository_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCIRepositoryServiceSaveRequest_Keypair); i {
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
	file_controlplane_v1_oci_repository_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OCIRepositoryServiceSaveRequest_KeyPair)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controlplane_v1_oci_repository_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controlplane_v1_oci_repository_proto_goTypes,
		DependencyIndexes: file_controlplane_v1_oci_repository_proto_depIdxs,
		EnumInfos:         file_controlplane_v1_oci_repository_proto_enumTypes,
		MessageInfos:      file_controlplane_v1_oci_repository_proto_msgTypes,
	}.Build()
	File_controlplane_v1_oci_repository_proto = out.File
	file_controlplane_v1_oci_repository_proto_rawDesc = nil
	file_controlplane_v1_oci_repository_proto_goTypes = nil
	file_controlplane_v1_oci_repository_proto_depIdxs = nil
}
