// Copyright 2021 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mockgrafeas/v1/slsa_provenance_zero_two.proto

package grafeas

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type SlsaProvenanceZeroTwo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Builder     *SlsaProvenanceZeroTwo_SlsaBuilder    `protobuf:"bytes,1,opt,name=builder,proto3" json:"builder,omitempty"`
	BuildType   string                                `protobuf:"bytes,2,opt,name=build_type,json=buildType,proto3" json:"build_type,omitempty"`
	Invocation  *SlsaProvenanceZeroTwo_SlsaInvocation `protobuf:"bytes,3,opt,name=invocation,proto3" json:"invocation,omitempty"`
	BuildConfig *_struct.Struct                       `protobuf:"bytes,4,opt,name=build_config,json=buildConfig,proto3" json:"build_config,omitempty"`
	Metadata    *SlsaProvenanceZeroTwo_SlsaMetadata   `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Materials   []*SlsaProvenanceZeroTwo_SlsaMaterial `protobuf:"bytes,6,rep,name=materials,proto3" json:"materials,omitempty"`
}

func (x *SlsaProvenanceZeroTwo) Reset() {
	*x = SlsaProvenanceZeroTwo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenanceZeroTwo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenanceZeroTwo) ProtoMessage() {}

func (x *SlsaProvenanceZeroTwo) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenanceZeroTwo.ProtoReflect.Descriptor instead.
func (*SlsaProvenanceZeroTwo) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescGZIP(), []int{0}
}

func (x *SlsaProvenanceZeroTwo) GetBuilder() *SlsaProvenanceZeroTwo_SlsaBuilder {
	if x != nil {
		return x.Builder
	}
	return nil
}

func (x *SlsaProvenanceZeroTwo) GetBuildType() string {
	if x != nil {
		return x.BuildType
	}
	return ""
}

func (x *SlsaProvenanceZeroTwo) GetInvocation() *SlsaProvenanceZeroTwo_SlsaInvocation {
	if x != nil {
		return x.Invocation
	}
	return nil
}

func (x *SlsaProvenanceZeroTwo) GetBuildConfig() *_struct.Struct {
	if x != nil {
		return x.BuildConfig
	}
	return nil
}

func (x *SlsaProvenanceZeroTwo) GetMetadata() *SlsaProvenanceZeroTwo_SlsaMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SlsaProvenanceZeroTwo) GetMaterials() []*SlsaProvenanceZeroTwo_SlsaMaterial {
	if x != nil {
		return x.Materials
	}
	return nil
}

// Identifies the entity that executed the recipe, which is trusted to have
// correctly performed the operation and populated this provenance.
type SlsaProvenanceZeroTwo_SlsaBuilder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SlsaProvenanceZeroTwo_SlsaBuilder) Reset() {
	*x = SlsaProvenanceZeroTwo_SlsaBuilder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenanceZeroTwo_SlsaBuilder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenanceZeroTwo_SlsaBuilder) ProtoMessage() {}

func (x *SlsaProvenanceZeroTwo_SlsaBuilder) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenanceZeroTwo_SlsaBuilder.ProtoReflect.Descriptor instead.
func (*SlsaProvenanceZeroTwo_SlsaBuilder) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SlsaProvenanceZeroTwo_SlsaBuilder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// The collection of artifacts that influenced the build including sources,
// dependencies, build tools, base images, and so on.
type SlsaProvenanceZeroTwo_SlsaMaterial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri    string            `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Digest map[string]string `protobuf:"bytes,2,rep,name=digest,proto3" json:"digest,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SlsaProvenanceZeroTwo_SlsaMaterial) Reset() {
	*x = SlsaProvenanceZeroTwo_SlsaMaterial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenanceZeroTwo_SlsaMaterial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenanceZeroTwo_SlsaMaterial) ProtoMessage() {}

func (x *SlsaProvenanceZeroTwo_SlsaMaterial) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenanceZeroTwo_SlsaMaterial.ProtoReflect.Descriptor instead.
func (*SlsaProvenanceZeroTwo_SlsaMaterial) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SlsaProvenanceZeroTwo_SlsaMaterial) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *SlsaProvenanceZeroTwo_SlsaMaterial) GetDigest() map[string]string {
	if x != nil {
		return x.Digest
	}
	return nil
}

// Identifies the event that kicked off the build.
type SlsaProvenanceZeroTwo_SlsaInvocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigSource *SlsaProvenanceZeroTwo_SlsaConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	Parameters   *_struct.Struct                         `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"`
	Environment  *_struct.Struct                         `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *SlsaProvenanceZeroTwo_SlsaInvocation) Reset() {
	*x = SlsaProvenanceZeroTwo_SlsaInvocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenanceZeroTwo_SlsaInvocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenanceZeroTwo_SlsaInvocation) ProtoMessage() {}

func (x *SlsaProvenanceZeroTwo_SlsaInvocation) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenanceZeroTwo_SlsaInvocation.ProtoReflect.Descriptor instead.
func (*SlsaProvenanceZeroTwo_SlsaInvocation) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescGZIP(), []int{0, 2}
}

func (x *SlsaProvenanceZeroTwo_SlsaInvocation) GetConfigSource() *SlsaProvenanceZeroTwo_SlsaConfigSource {
	if x != nil {
		return x.ConfigSource
	}
	return nil
}

func (x *SlsaProvenanceZeroTwo_SlsaInvocation) GetParameters() *_struct.Struct {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *SlsaProvenanceZeroTwo_SlsaInvocation) GetEnvironment() *_struct.Struct {
	if x != nil {
		return x.Environment
	}
	return nil
}

// Describes where the config file that kicked off the build came from.
// This is effectively a pointer to the source where buildConfig came from.
type SlsaProvenanceZeroTwo_SlsaConfigSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri        string            `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Digest     map[string]string `protobuf:"bytes,2,rep,name=digest,proto3" json:"digest,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EntryPoint string            `protobuf:"bytes,3,opt,name=entry_point,json=entryPoint,proto3" json:"entry_point,omitempty"`
}

func (x *SlsaProvenanceZeroTwo_SlsaConfigSource) Reset() {
	*x = SlsaProvenanceZeroTwo_SlsaConfigSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenanceZeroTwo_SlsaConfigSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenanceZeroTwo_SlsaConfigSource) ProtoMessage() {}

func (x *SlsaProvenanceZeroTwo_SlsaConfigSource) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenanceZeroTwo_SlsaConfigSource.ProtoReflect.Descriptor instead.
func (*SlsaProvenanceZeroTwo_SlsaConfigSource) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescGZIP(), []int{0, 3}
}

func (x *SlsaProvenanceZeroTwo_SlsaConfigSource) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *SlsaProvenanceZeroTwo_SlsaConfigSource) GetDigest() map[string]string {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *SlsaProvenanceZeroTwo_SlsaConfigSource) GetEntryPoint() string {
	if x != nil {
		return x.EntryPoint
	}
	return ""
}

// Other properties of the build.
type SlsaProvenanceZeroTwo_SlsaMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildInvocationId string                                  `protobuf:"bytes,1,opt,name=build_invocation_id,json=buildInvocationId,proto3" json:"build_invocation_id,omitempty"`
	BuildStartedOn    *timestamp.Timestamp                    `protobuf:"bytes,2,opt,name=build_started_on,json=buildStartedOn,proto3" json:"build_started_on,omitempty"`
	BuildFinishedOn   *timestamp.Timestamp                    `protobuf:"bytes,3,opt,name=build_finished_on,json=buildFinishedOn,proto3" json:"build_finished_on,omitempty"`
	Completeness      *SlsaProvenanceZeroTwo_SlsaCompleteness `protobuf:"bytes,4,opt,name=completeness,proto3" json:"completeness,omitempty"`
	Reproducible      bool                                    `protobuf:"varint,5,opt,name=reproducible,proto3" json:"reproducible,omitempty"`
}

func (x *SlsaProvenanceZeroTwo_SlsaMetadata) Reset() {
	*x = SlsaProvenanceZeroTwo_SlsaMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenanceZeroTwo_SlsaMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenanceZeroTwo_SlsaMetadata) ProtoMessage() {}

func (x *SlsaProvenanceZeroTwo_SlsaMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenanceZeroTwo_SlsaMetadata.ProtoReflect.Descriptor instead.
func (*SlsaProvenanceZeroTwo_SlsaMetadata) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescGZIP(), []int{0, 4}
}

func (x *SlsaProvenanceZeroTwo_SlsaMetadata) GetBuildInvocationId() string {
	if x != nil {
		return x.BuildInvocationId
	}
	return ""
}

func (x *SlsaProvenanceZeroTwo_SlsaMetadata) GetBuildStartedOn() *timestamp.Timestamp {
	if x != nil {
		return x.BuildStartedOn
	}
	return nil
}

func (x *SlsaProvenanceZeroTwo_SlsaMetadata) GetBuildFinishedOn() *timestamp.Timestamp {
	if x != nil {
		return x.BuildFinishedOn
	}
	return nil
}

func (x *SlsaProvenanceZeroTwo_SlsaMetadata) GetCompleteness() *SlsaProvenanceZeroTwo_SlsaCompleteness {
	if x != nil {
		return x.Completeness
	}
	return nil
}

func (x *SlsaProvenanceZeroTwo_SlsaMetadata) GetReproducible() bool {
	if x != nil {
		return x.Reproducible
	}
	return false
}

// Indicates that the builder claims certain fields in this message to be
// complete.
type SlsaProvenanceZeroTwo_SlsaCompleteness struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters  bool `protobuf:"varint,1,opt,name=parameters,proto3" json:"parameters,omitempty"`
	Environment bool `protobuf:"varint,2,opt,name=environment,proto3" json:"environment,omitempty"`
	Materials   bool `protobuf:"varint,3,opt,name=materials,proto3" json:"materials,omitempty"`
}

func (x *SlsaProvenanceZeroTwo_SlsaCompleteness) Reset() {
	*x = SlsaProvenanceZeroTwo_SlsaCompleteness{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenanceZeroTwo_SlsaCompleteness) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenanceZeroTwo_SlsaCompleteness) ProtoMessage() {}

func (x *SlsaProvenanceZeroTwo_SlsaCompleteness) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenanceZeroTwo_SlsaCompleteness.ProtoReflect.Descriptor instead.
func (*SlsaProvenanceZeroTwo_SlsaCompleteness) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescGZIP(), []int{0, 5}
}

func (x *SlsaProvenanceZeroTwo_SlsaCompleteness) GetParameters() bool {
	if x != nil {
		return x.Parameters
	}
	return false
}

func (x *SlsaProvenanceZeroTwo_SlsaCompleteness) GetEnvironment() bool {
	if x != nil {
		return x.Environment
	}
	return false
}

func (x *SlsaProvenanceZeroTwo_SlsaCompleteness) GetMaterials() bool {
	if x != nil {
		return x.Materials
	}
	return false
}

var File_mockgrafeas_v1_slsa_provenance_zero_two_proto protoreflect.FileDescriptor

var file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x6c, 0x73, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x74, 0x77, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92,
	0x0c, 0x0a, 0x15, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x5a, 0x65, 0x72, 0x6f, 0x54, 0x77, 0x6f, 0x12, 0x4b, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6d, 0x6f, 0x63, 0x6b,
	0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x50,
	0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5a, 0x65, 0x72, 0x6f, 0x54, 0x77, 0x6f,
	0x2e, 0x53, 0x6c, 0x73, 0x61, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x07, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x54, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72,
	0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5a, 0x65, 0x72, 0x6f, 0x54, 0x77, 0x6f, 0x2e,
	0x53, 0x6c, 0x73, 0x61, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0c, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72,
	0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5a, 0x65, 0x72, 0x6f, 0x54, 0x77, 0x6f, 0x2e,
	0x53, 0x6c, 0x73, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x50, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6d, 0x6f, 0x63, 0x6b,
	0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x50,
	0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5a, 0x65, 0x72, 0x6f, 0x54, 0x77, 0x6f,
	0x2e, 0x53, 0x6c, 0x73, 0x61, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x09, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x1d, 0x0a, 0x0b, 0x53, 0x6c, 0x73, 0x61,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x1a, 0xb3, 0x01, 0x0a, 0x0c, 0x53, 0x6c, 0x73, 0x61,
	0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x56, 0x0a, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x6d, 0x6f, 0x63,
	0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61,
	0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5a, 0x65, 0x72, 0x6f, 0x54, 0x77,
	0x6f, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x44,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xe1, 0x01,
	0x0a, 0x0e, 0x53, 0x6c, 0x73, 0x61, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x5b, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72,
	0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72, 0x6f,
	0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5a, 0x65, 0x72, 0x6f, 0x54, 0x77, 0x6f, 0x2e, 0x53,
	0x6c, 0x73, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x37, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x1a, 0xdc, 0x01, 0x0a, 0x10, 0x53, 0x6c, 0x73, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x5a, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72,
	0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5a, 0x65, 0x72, 0x6f, 0x54, 0x77, 0x6f, 0x2e,
	0x53, 0x6c, 0x73, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0xcc, 0x02, 0x0a, 0x0c, 0x53, 0x6c, 0x73, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x2e, 0x0a, 0x13, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x44, 0x0a, 0x10, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x46, 0x0a, 0x11, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4f, 0x6e, 0x12,
	0x5a, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5a, 0x65, 0x72, 0x6f, 0x54, 0x77, 0x6f, 0x2e, 0x53, 0x6c, 0x73,
	0x61, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x0c, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x69, 0x62, 0x6c, 0x65, 0x1a,
	0x72, 0x0a, 0x10, 0x53, 0x6c, 0x73, 0x61, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e,
	0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x73, 0x42, 0x75, 0x0a, 0x11, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72,
	0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x1a, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72,
	0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5a, 0x65, 0x72, 0x6f, 0x54, 0x77, 0x6f, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f,
	0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x61,
	0x66, 0x65, 0x61, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescOnce sync.Once
	file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescData = file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDesc
)

func file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescGZIP() []byte {
	file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescOnce.Do(func() {
		file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescData)
	})
	return file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDescData
}

var file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mockgrafeas_v1_slsa_provenance_zero_two_proto_goTypes = []interface{}{
	(*SlsaProvenanceZeroTwo)(nil),                  // 0: mockgrafeas.v1.SlsaProvenanceZeroTwo
	(*SlsaProvenanceZeroTwo_SlsaBuilder)(nil),      // 1: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaBuilder
	(*SlsaProvenanceZeroTwo_SlsaMaterial)(nil),     // 2: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaMaterial
	(*SlsaProvenanceZeroTwo_SlsaInvocation)(nil),   // 3: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaInvocation
	(*SlsaProvenanceZeroTwo_SlsaConfigSource)(nil), // 4: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaConfigSource
	(*SlsaProvenanceZeroTwo_SlsaMetadata)(nil),     // 5: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaMetadata
	(*SlsaProvenanceZeroTwo_SlsaCompleteness)(nil), // 6: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaCompleteness
	nil,                         // 7: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaMaterial.DigestEntry
	nil,                         // 8: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaConfigSource.DigestEntry
	(*_struct.Struct)(nil),      // 9: google.protobuf.Struct
	(*timestamp.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_mockgrafeas_v1_slsa_provenance_zero_two_proto_depIdxs = []int32{
	1,  // 0: mockgrafeas.v1.SlsaProvenanceZeroTwo.builder:type_name -> mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaBuilder
	3,  // 1: mockgrafeas.v1.SlsaProvenanceZeroTwo.invocation:type_name -> mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaInvocation
	9,  // 2: mockgrafeas.v1.SlsaProvenanceZeroTwo.build_config:type_name -> google.protobuf.Struct
	5,  // 3: mockgrafeas.v1.SlsaProvenanceZeroTwo.metadata:type_name -> mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaMetadata
	2,  // 4: mockgrafeas.v1.SlsaProvenanceZeroTwo.materials:type_name -> mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaMaterial
	7,  // 5: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaMaterial.digest:type_name -> mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaMaterial.DigestEntry
	4,  // 6: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaInvocation.config_source:type_name -> mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaConfigSource
	9,  // 7: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaInvocation.parameters:type_name -> google.protobuf.Struct
	9,  // 8: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaInvocation.environment:type_name -> google.protobuf.Struct
	8,  // 9: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaConfigSource.digest:type_name -> mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaConfigSource.DigestEntry
	10, // 10: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaMetadata.build_started_on:type_name -> google.protobuf.Timestamp
	10, // 11: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaMetadata.build_finished_on:type_name -> google.protobuf.Timestamp
	6,  // 12: mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaMetadata.completeness:type_name -> mockgrafeas.v1.SlsaProvenanceZeroTwo.SlsaCompleteness
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_mockgrafeas_v1_slsa_provenance_zero_two_proto_init() }
func file_mockgrafeas_v1_slsa_provenance_zero_two_proto_init() {
	if File_mockgrafeas_v1_slsa_provenance_zero_two_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenanceZeroTwo); i {
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
		file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenanceZeroTwo_SlsaBuilder); i {
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
		file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenanceZeroTwo_SlsaMaterial); i {
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
		file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenanceZeroTwo_SlsaInvocation); i {
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
		file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenanceZeroTwo_SlsaConfigSource); i {
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
		file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenanceZeroTwo_SlsaMetadata); i {
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
		file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenanceZeroTwo_SlsaCompleteness); i {
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
			RawDescriptor: file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgrafeas_v1_slsa_provenance_zero_two_proto_goTypes,
		DependencyIndexes: file_mockgrafeas_v1_slsa_provenance_zero_two_proto_depIdxs,
		MessageInfos:      file_mockgrafeas_v1_slsa_provenance_zero_two_proto_msgTypes,
	}.Build()
	File_mockgrafeas_v1_slsa_provenance_zero_two_proto = out.File
	file_mockgrafeas_v1_slsa_provenance_zero_two_proto_rawDesc = nil
	file_mockgrafeas_v1_slsa_provenance_zero_two_proto_goTypes = nil
	file_mockgrafeas_v1_slsa_provenance_zero_two_proto_depIdxs = nil
}
