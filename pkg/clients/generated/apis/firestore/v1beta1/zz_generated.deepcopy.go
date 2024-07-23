//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
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

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirestoreIndex) DeepCopyInto(out *FirestoreIndex) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirestoreIndex.
func (in *FirestoreIndex) DeepCopy() *FirestoreIndex {
	if in == nil {
		return nil
	}
	out := new(FirestoreIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirestoreIndex) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirestoreIndexList) DeepCopyInto(out *FirestoreIndexList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FirestoreIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirestoreIndexList.
func (in *FirestoreIndexList) DeepCopy() *FirestoreIndexList {
	if in == nil {
		return nil
	}
	out := new(FirestoreIndexList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirestoreIndexList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirestoreIndexSpec) DeepCopyInto(out *FirestoreIndexSpec) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make([]IndexFields, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.QueryScope != nil {
		in, out := &in.QueryScope, &out.QueryScope
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirestoreIndexSpec.
func (in *FirestoreIndexSpec) DeepCopy() *FirestoreIndexSpec {
	if in == nil {
		return nil
	}
	out := new(FirestoreIndexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirestoreIndexStatus) DeepCopyInto(out *FirestoreIndexStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirestoreIndexStatus.
func (in *FirestoreIndexStatus) DeepCopy() *FirestoreIndexStatus {
	if in == nil {
		return nil
	}
	out := new(FirestoreIndexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexFields) DeepCopyInto(out *IndexFields) {
	*out = *in
	if in.ArrayConfig != nil {
		in, out := &in.ArrayConfig, &out.ArrayConfig
		*out = new(string)
		**out = **in
	}
	if in.FieldPath != nil {
		in, out := &in.FieldPath, &out.FieldPath
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexFields.
func (in *IndexFields) DeepCopy() *IndexFields {
	if in == nil {
		return nil
	}
	out := new(IndexFields)
	in.DeepCopyInto(out)
	return out
}
