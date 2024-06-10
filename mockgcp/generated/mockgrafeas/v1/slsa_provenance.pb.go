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
// source: mockgrafeas/v1/slsa_provenance.proto

package grafeas

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

type SlsaProvenance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Builder *SlsaProvenance_SlsaBuilder `protobuf:"bytes,1,opt,name=builder,proto3" json:"builder,omitempty"` // required
	// Identifies the configuration used for the build.
	// When combined with materials, this SHOULD fully describe the build,
	// such that re-running this recipe results in bit-for-bit identical output
	// (if the build is reproducible).
	Recipe   *SlsaProvenance_SlsaRecipe   `protobuf:"bytes,2,opt,name=recipe,proto3" json:"recipe,omitempty"` // required
	Metadata *SlsaProvenance_SlsaMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The collection of artifacts that influenced the build including sources,
	// dependencies, build tools, base images, and so on. This is considered to be
	// incomplete unless metadata.completeness.materials is true. Unset or null is
	// equivalent to empty.
	Materials []*SlsaProvenance_Material `protobuf:"bytes,4,rep,name=materials,proto3" json:"materials,omitempty"`
}

func (x *SlsaProvenance) Reset() {
	*x = SlsaProvenance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenance) ProtoMessage() {}

func (x *SlsaProvenance) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenance.ProtoReflect.Descriptor instead.
func (*SlsaProvenance) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_proto_rawDescGZIP(), []int{0}
}

func (x *SlsaProvenance) GetBuilder() *SlsaProvenance_SlsaBuilder {
	if x != nil {
		return x.Builder
	}
	return nil
}

func (x *SlsaProvenance) GetRecipe() *SlsaProvenance_SlsaRecipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

func (x *SlsaProvenance) GetMetadata() *SlsaProvenance_SlsaMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SlsaProvenance) GetMaterials() []*SlsaProvenance_Material {
	if x != nil {
		return x.Materials
	}
	return nil
}

// Steps taken to build the artifact.
// For a TaskRun, typically each container corresponds to one step in the
// recipe.
type SlsaProvenance_SlsaRecipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URI indicating what type of recipe was performed. It determines the
	// meaning of recipe.entryPoint, recipe.arguments, recipe.environment, and
	// materials.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Index in materials containing the recipe steps that are not implied by
	// recipe.type. For example, if the recipe type were "make", then this would
	// point to the source containing the Makefile, not the make program itself.
	// Set to -1 if the recipe doesn't come from a material, as zero is default
	// unset value for int64.
	DefinedInMaterial int64 `protobuf:"varint,2,opt,name=defined_in_material,json=definedInMaterial,proto3" json:"defined_in_material,omitempty"`
	// String identifying the entry point into the build.
	// This is often a path to a configuration file and/or a target label within
	// that file. The syntax and meaning are defined by recipe.type. For
	// example, if the recipe type were "make", then this would reference the
	// directory in which to run make as well as which target to use.
	EntryPoint string `protobuf:"bytes,3,opt,name=entry_point,json=entryPoint,proto3" json:"entry_point,omitempty"`
	// Collection of all external inputs that influenced the build on top of
	// recipe.definedInMaterial and recipe.entryPoint. For example, if the
	// recipe type were "make", then this might be the flags passed to make
	// aside from the target, which is captured in recipe.entryPoint. Depending
	// on the recipe Type, the structure may be different.
	Arguments *any1.Any `protobuf:"bytes,4,opt,name=arguments,proto3" json:"arguments,omitempty"`
	// Any other builder-controlled inputs necessary for correctly evaluating
	// the recipe. Usually only needed for reproducing the build but not
	// evaluated as part of policy. Depending on the recipe Type, the structure
	// may be different.
	Environment *any1.Any `protobuf:"bytes,5,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *SlsaProvenance_SlsaRecipe) Reset() {
	*x = SlsaProvenance_SlsaRecipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenance_SlsaRecipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenance_SlsaRecipe) ProtoMessage() {}

func (x *SlsaProvenance_SlsaRecipe) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenance_SlsaRecipe.ProtoReflect.Descriptor instead.
func (*SlsaProvenance_SlsaRecipe) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SlsaProvenance_SlsaRecipe) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SlsaProvenance_SlsaRecipe) GetDefinedInMaterial() int64 {
	if x != nil {
		return x.DefinedInMaterial
	}
	return 0
}

func (x *SlsaProvenance_SlsaRecipe) GetEntryPoint() string {
	if x != nil {
		return x.EntryPoint
	}
	return ""
}

func (x *SlsaProvenance_SlsaRecipe) GetArguments() *any1.Any {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *SlsaProvenance_SlsaRecipe) GetEnvironment() *any1.Any {
	if x != nil {
		return x.Environment
	}
	return nil
}

// Indicates that the builder claims certain fields in this message to be
// complete.
type SlsaProvenance_SlsaCompleteness struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, the builder claims that recipe.arguments is complete, meaning
	// that all external inputs are properly captured in the recipe.
	Arguments bool `protobuf:"varint,1,opt,name=arguments,proto3" json:"arguments,omitempty"`
	// If true, the builder claims that recipe.environment is claimed to be
	// complete.
	Environment bool `protobuf:"varint,2,opt,name=environment,proto3" json:"environment,omitempty"`
	// If true, the builder claims that materials are complete, usually through
	// some controls to prevent network access. Sometimes called "hermetic".
	Materials bool `protobuf:"varint,3,opt,name=materials,proto3" json:"materials,omitempty"`
}

func (x *SlsaProvenance_SlsaCompleteness) Reset() {
	*x = SlsaProvenance_SlsaCompleteness{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenance_SlsaCompleteness) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenance_SlsaCompleteness) ProtoMessage() {}

func (x *SlsaProvenance_SlsaCompleteness) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenance_SlsaCompleteness.ProtoReflect.Descriptor instead.
func (*SlsaProvenance_SlsaCompleteness) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SlsaProvenance_SlsaCompleteness) GetArguments() bool {
	if x != nil {
		return x.Arguments
	}
	return false
}

func (x *SlsaProvenance_SlsaCompleteness) GetEnvironment() bool {
	if x != nil {
		return x.Environment
	}
	return false
}

func (x *SlsaProvenance_SlsaCompleteness) GetMaterials() bool {
	if x != nil {
		return x.Materials
	}
	return false
}

// Other properties of the build.
type SlsaProvenance_SlsaMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies the particular build invocation, which can be useful for
	// finding associated logs or other ad-hoc analysis. The value SHOULD be
	// globally unique, per in-toto Provenance spec.
	BuildInvocationId string `protobuf:"bytes,1,opt,name=build_invocation_id,json=buildInvocationId,proto3" json:"build_invocation_id,omitempty"`
	// The timestamp of when the build started.
	BuildStartedOn *timestamp.Timestamp `protobuf:"bytes,2,opt,name=build_started_on,json=buildStartedOn,proto3" json:"build_started_on,omitempty"`
	// The timestamp of when the build completed.
	BuildFinishedOn *timestamp.Timestamp `protobuf:"bytes,3,opt,name=build_finished_on,json=buildFinishedOn,proto3" json:"build_finished_on,omitempty"`
	// Indicates that the builder claims certain fields in this message to be
	// complete.
	Completeness *SlsaProvenance_SlsaCompleteness `protobuf:"bytes,4,opt,name=completeness,proto3" json:"completeness,omitempty"`
	// If true, the builder claims that running the recipe on materials will
	// produce bit-for-bit identical output.
	Reproducible bool `protobuf:"varint,5,opt,name=reproducible,proto3" json:"reproducible,omitempty"`
}

func (x *SlsaProvenance_SlsaMetadata) Reset() {
	*x = SlsaProvenance_SlsaMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenance_SlsaMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenance_SlsaMetadata) ProtoMessage() {}

func (x *SlsaProvenance_SlsaMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenance_SlsaMetadata.ProtoReflect.Descriptor instead.
func (*SlsaProvenance_SlsaMetadata) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_proto_rawDescGZIP(), []int{0, 2}
}

func (x *SlsaProvenance_SlsaMetadata) GetBuildInvocationId() string {
	if x != nil {
		return x.BuildInvocationId
	}
	return ""
}

func (x *SlsaProvenance_SlsaMetadata) GetBuildStartedOn() *timestamp.Timestamp {
	if x != nil {
		return x.BuildStartedOn
	}
	return nil
}

func (x *SlsaProvenance_SlsaMetadata) GetBuildFinishedOn() *timestamp.Timestamp {
	if x != nil {
		return x.BuildFinishedOn
	}
	return nil
}

func (x *SlsaProvenance_SlsaMetadata) GetCompleteness() *SlsaProvenance_SlsaCompleteness {
	if x != nil {
		return x.Completeness
	}
	return nil
}

func (x *SlsaProvenance_SlsaMetadata) GetReproducible() bool {
	if x != nil {
		return x.Reproducible
	}
	return false
}

type SlsaProvenance_SlsaBuilder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SlsaProvenance_SlsaBuilder) Reset() {
	*x = SlsaProvenance_SlsaBuilder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenance_SlsaBuilder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenance_SlsaBuilder) ProtoMessage() {}

func (x *SlsaProvenance_SlsaBuilder) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenance_SlsaBuilder.ProtoReflect.Descriptor instead.
func (*SlsaProvenance_SlsaBuilder) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_proto_rawDescGZIP(), []int{0, 3}
}

func (x *SlsaProvenance_SlsaBuilder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SlsaProvenance_Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri    string            `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Digest map[string]string `protobuf:"bytes,2,rep,name=digest,proto3" json:"digest,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SlsaProvenance_Material) Reset() {
	*x = SlsaProvenance_Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlsaProvenance_Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlsaProvenance_Material) ProtoMessage() {}

func (x *SlsaProvenance_Material) ProtoReflect() protoreflect.Message {
	mi := &file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlsaProvenance_Material.ProtoReflect.Descriptor instead.
func (*SlsaProvenance_Material) Descriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_slsa_provenance_proto_rawDescGZIP(), []int{0, 4}
}

func (x *SlsaProvenance_Material) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *SlsaProvenance_Material) GetDigest() map[string]string {
	if x != nil {
		return x.Digest
	}
	return nil
}

var File_mockgrafeas_v1_slsa_provenance_proto protoreflect.FileDescriptor

var file_mockgrafeas_v1_slsa_provenance_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x6c, 0x73, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x89, 0x09, 0x0a, 0x0e, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61,
	0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72, 0x6f, 0x76,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x06, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x6f,
	0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73,
	0x61, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x53, 0x6c, 0x73, 0x61,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x47,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x53, 0x6c, 0x73, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x45, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x6f, 0x63,
	0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61,
	0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0xdd,
	0x01, 0x0a, 0x0a, 0x53, 0x6c, 0x73, 0x61, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x5f,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x61, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x70,
	0x0a, 0x10, 0x53, 0x6c, 0x73, 0x61, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65,
	0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73,
	0x1a, 0xc5, 0x02, 0x0a, 0x0c, 0x53, 0x6c, 0x73, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
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
	0x53, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x69, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x69, 0x62, 0x6c, 0x65, 0x1a, 0x1d, 0x0a, 0x0b, 0x53, 0x6c, 0x73, 0x61,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x1a, 0xa4, 0x01, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x4b, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61,
	0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72, 0x6f, 0x76,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e,
	0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x6e,
	0x0a, 0x11, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x13, 0x53, 0x6c, 0x73, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockgrafeas_v1_slsa_provenance_proto_rawDescOnce sync.Once
	file_mockgrafeas_v1_slsa_provenance_proto_rawDescData = file_mockgrafeas_v1_slsa_provenance_proto_rawDesc
)

func file_mockgrafeas_v1_slsa_provenance_proto_rawDescGZIP() []byte {
	file_mockgrafeas_v1_slsa_provenance_proto_rawDescOnce.Do(func() {
		file_mockgrafeas_v1_slsa_provenance_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgrafeas_v1_slsa_provenance_proto_rawDescData)
	})
	return file_mockgrafeas_v1_slsa_provenance_proto_rawDescData
}

var file_mockgrafeas_v1_slsa_provenance_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mockgrafeas_v1_slsa_provenance_proto_goTypes = []interface{}{
	(*SlsaProvenance)(nil),                  // 0: mockgrafeas.v1.SlsaProvenance
	(*SlsaProvenance_SlsaRecipe)(nil),       // 1: mockgrafeas.v1.SlsaProvenance.SlsaRecipe
	(*SlsaProvenance_SlsaCompleteness)(nil), // 2: mockgrafeas.v1.SlsaProvenance.SlsaCompleteness
	(*SlsaProvenance_SlsaMetadata)(nil),     // 3: mockgrafeas.v1.SlsaProvenance.SlsaMetadata
	(*SlsaProvenance_SlsaBuilder)(nil),      // 4: mockgrafeas.v1.SlsaProvenance.SlsaBuilder
	(*SlsaProvenance_Material)(nil),         // 5: mockgrafeas.v1.SlsaProvenance.Material
	nil,                                     // 6: mockgrafeas.v1.SlsaProvenance.Material.DigestEntry
	(*any1.Any)(nil),                        // 7: google.protobuf.Any
	(*timestamp.Timestamp)(nil),             // 8: google.protobuf.Timestamp
}
var file_mockgrafeas_v1_slsa_provenance_proto_depIdxs = []int32{
	4,  // 0: mockgrafeas.v1.SlsaProvenance.builder:type_name -> mockgrafeas.v1.SlsaProvenance.SlsaBuilder
	1,  // 1: mockgrafeas.v1.SlsaProvenance.recipe:type_name -> mockgrafeas.v1.SlsaProvenance.SlsaRecipe
	3,  // 2: mockgrafeas.v1.SlsaProvenance.metadata:type_name -> mockgrafeas.v1.SlsaProvenance.SlsaMetadata
	5,  // 3: mockgrafeas.v1.SlsaProvenance.materials:type_name -> mockgrafeas.v1.SlsaProvenance.Material
	7,  // 4: mockgrafeas.v1.SlsaProvenance.SlsaRecipe.arguments:type_name -> google.protobuf.Any
	7,  // 5: mockgrafeas.v1.SlsaProvenance.SlsaRecipe.environment:type_name -> google.protobuf.Any
	8,  // 6: mockgrafeas.v1.SlsaProvenance.SlsaMetadata.build_started_on:type_name -> google.protobuf.Timestamp
	8,  // 7: mockgrafeas.v1.SlsaProvenance.SlsaMetadata.build_finished_on:type_name -> google.protobuf.Timestamp
	2,  // 8: mockgrafeas.v1.SlsaProvenance.SlsaMetadata.completeness:type_name -> mockgrafeas.v1.SlsaProvenance.SlsaCompleteness
	6,  // 9: mockgrafeas.v1.SlsaProvenance.Material.digest:type_name -> mockgrafeas.v1.SlsaProvenance.Material.DigestEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_mockgrafeas_v1_slsa_provenance_proto_init() }
func file_mockgrafeas_v1_slsa_provenance_proto_init() {
	if File_mockgrafeas_v1_slsa_provenance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenance); i {
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
		file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenance_SlsaRecipe); i {
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
		file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenance_SlsaCompleteness); i {
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
		file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenance_SlsaMetadata); i {
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
		file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenance_SlsaBuilder); i {
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
		file_mockgrafeas_v1_slsa_provenance_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlsaProvenance_Material); i {
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
			RawDescriptor: file_mockgrafeas_v1_slsa_provenance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgrafeas_v1_slsa_provenance_proto_goTypes,
		DependencyIndexes: file_mockgrafeas_v1_slsa_provenance_proto_depIdxs,
		MessageInfos:      file_mockgrafeas_v1_slsa_provenance_proto_msgTypes,
	}.Build()
	File_mockgrafeas_v1_slsa_provenance_proto = out.File
	file_mockgrafeas_v1_slsa_provenance_proto_rawDesc = nil
	file_mockgrafeas_v1_slsa_provenance_proto_goTypes = nil
	file_mockgrafeas_v1_slsa_provenance_proto_depIdxs = nil
}
