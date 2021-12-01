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

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretCustomerManagedEncryption) DeepCopyInto(out *SecretCustomerManagedEncryption) {
	*out = *in
	out.KmsKeyRef = in.KmsKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretCustomerManagedEncryption.
func (in *SecretCustomerManagedEncryption) DeepCopy() *SecretCustomerManagedEncryption {
	if in == nil {
		return nil
	}
	out := new(SecretCustomerManagedEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretManagerSecret) DeepCopyInto(out *SecretManagerSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretManagerSecret.
func (in *SecretManagerSecret) DeepCopy() *SecretManagerSecret {
	if in == nil {
		return nil
	}
	out := new(SecretManagerSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretManagerSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretManagerSecretList) DeepCopyInto(out *SecretManagerSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretManagerSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretManagerSecretList.
func (in *SecretManagerSecretList) DeepCopy() *SecretManagerSecretList {
	if in == nil {
		return nil
	}
	out := new(SecretManagerSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretManagerSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretManagerSecretSpec) DeepCopyInto(out *SecretManagerSecretSpec) {
	*out = *in
	if in.ExpireTime != nil {
		in, out := &in.ExpireTime, &out.ExpireTime
		*out = new(string)
		**out = **in
	}
	in.Replication.DeepCopyInto(&out.Replication)
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Rotation != nil {
		in, out := &in.Rotation, &out.Rotation
		*out = new(SecretRotation)
		(*in).DeepCopyInto(*out)
	}
	if in.Topics != nil {
		in, out := &in.Topics, &out.Topics
		*out = make([]v1alpha1.ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.Ttl != nil {
		in, out := &in.Ttl, &out.Ttl
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretManagerSecretSpec.
func (in *SecretManagerSecretSpec) DeepCopy() *SecretManagerSecretSpec {
	if in == nil {
		return nil
	}
	out := new(SecretManagerSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretManagerSecretStatus) DeepCopyInto(out *SecretManagerSecretStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretManagerSecretStatus.
func (in *SecretManagerSecretStatus) DeepCopy() *SecretManagerSecretStatus {
	if in == nil {
		return nil
	}
	out := new(SecretManagerSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretManagerSecretVersion) DeepCopyInto(out *SecretManagerSecretVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretManagerSecretVersion.
func (in *SecretManagerSecretVersion) DeepCopy() *SecretManagerSecretVersion {
	if in == nil {
		return nil
	}
	out := new(SecretManagerSecretVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretManagerSecretVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretManagerSecretVersionList) DeepCopyInto(out *SecretManagerSecretVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretManagerSecretVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretManagerSecretVersionList.
func (in *SecretManagerSecretVersionList) DeepCopy() *SecretManagerSecretVersionList {
	if in == nil {
		return nil
	}
	out := new(SecretManagerSecretVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretManagerSecretVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretManagerSecretVersionSpec) DeepCopyInto(out *SecretManagerSecretVersionSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	in.SecretData.DeepCopyInto(&out.SecretData)
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretManagerSecretVersionSpec.
func (in *SecretManagerSecretVersionSpec) DeepCopy() *SecretManagerSecretVersionSpec {
	if in == nil {
		return nil
	}
	out := new(SecretManagerSecretVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretManagerSecretVersionStatus) DeepCopyInto(out *SecretManagerSecretVersionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretManagerSecretVersionStatus.
func (in *SecretManagerSecretVersionStatus) DeepCopy() *SecretManagerSecretVersionStatus {
	if in == nil {
		return nil
	}
	out := new(SecretManagerSecretVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretReplicas) DeepCopyInto(out *SecretReplicas) {
	*out = *in
	if in.CustomerManagedEncryption != nil {
		in, out := &in.CustomerManagedEncryption, &out.CustomerManagedEncryption
		*out = new(SecretCustomerManagedEncryption)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretReplicas.
func (in *SecretReplicas) DeepCopy() *SecretReplicas {
	if in == nil {
		return nil
	}
	out := new(SecretReplicas)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretReplication) DeepCopyInto(out *SecretReplication) {
	*out = *in
	if in.Automatic != nil {
		in, out := &in.Automatic, &out.Automatic
		*out = new(bool)
		**out = **in
	}
	if in.UserManaged != nil {
		in, out := &in.UserManaged, &out.UserManaged
		*out = new(SecretUserManaged)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretReplication.
func (in *SecretReplication) DeepCopy() *SecretReplication {
	if in == nil {
		return nil
	}
	out := new(SecretReplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRotation) DeepCopyInto(out *SecretRotation) {
	*out = *in
	if in.NextRotationTime != nil {
		in, out := &in.NextRotationTime, &out.NextRotationTime
		*out = new(string)
		**out = **in
	}
	if in.RotationPeriod != nil {
		in, out := &in.RotationPeriod, &out.RotationPeriod
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRotation.
func (in *SecretRotation) DeepCopy() *SecretRotation {
	if in == nil {
		return nil
	}
	out := new(SecretRotation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretUserManaged) DeepCopyInto(out *SecretUserManaged) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]SecretReplicas, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretUserManaged.
func (in *SecretUserManaged) DeepCopy() *SecretUserManaged {
	if in == nil {
		return nil
	}
	out := new(SecretUserManaged)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretversionSecretData) DeepCopyInto(out *SecretversionSecretData) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(SecretversionValueFrom)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretversionSecretData.
func (in *SecretversionSecretData) DeepCopy() *SecretversionSecretData {
	if in == nil {
		return nil
	}
	out := new(SecretversionSecretData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretversionValueFrom) DeepCopyInto(out *SecretversionValueFrom) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretversionValueFrom.
func (in *SecretversionValueFrom) DeepCopy() *SecretversionValueFrom {
	if in == nil {
		return nil
	}
	out := new(SecretversionValueFrom)
	in.DeepCopyInto(out)
	return out
}
