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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyObservedState) DeepCopyInto(out *APIKeyObservedState) {
	*out = *in
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyObservedState.
func (in *APIKeyObservedState) DeepCopy() *APIKeyObservedState {
	if in == nil {
		return nil
	}
	out := new(APIKeyObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeySpec) DeepCopyInto(out *APIKeySpec) {
	*out = *in
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Restrictions != nil {
		in, out := &in.Restrictions, &out.Restrictions
		*out = new(Restrictions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeySpec.
func (in *APIKeySpec) DeepCopy() *APIKeySpec {
	if in == nil {
		return nil
	}
	out := new(APIKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyStatus) DeepCopyInto(out *APIKeyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(APIKeyObservedState)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyStatus.
func (in *APIKeyStatus) DeepCopy() *APIKeyStatus {
	if in == nil {
		return nil
	}
	out := new(APIKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeysKey) DeepCopyInto(out *APIKeysKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeysKey.
func (in *APIKeysKey) DeepCopy() *APIKeysKey {
	if in == nil {
		return nil
	}
	out := new(APIKeysKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIKeysKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeysKeyList) DeepCopyInto(out *APIKeysKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIKeysKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeysKeyList.
func (in *APIKeysKeyList) DeepCopy() *APIKeysKeyList {
	if in == nil {
		return nil
	}
	out := new(APIKeysKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIKeysKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AndroidApplication) DeepCopyInto(out *AndroidApplication) {
	*out = *in
	if in.Sha1Fingerprint != nil {
		in, out := &in.Sha1Fingerprint, &out.Sha1Fingerprint
		*out = new(string)
		**out = **in
	}
	if in.PackageName != nil {
		in, out := &in.PackageName, &out.PackageName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AndroidApplication.
func (in *AndroidApplication) DeepCopy() *AndroidApplication {
	if in == nil {
		return nil
	}
	out := new(AndroidApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AndroidKeyRestrictions) DeepCopyInto(out *AndroidKeyRestrictions) {
	*out = *in
	if in.AllowedApplications != nil {
		in, out := &in.AllowedApplications, &out.AllowedApplications
		*out = make([]AndroidApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AndroidKeyRestrictions.
func (in *AndroidKeyRestrictions) DeepCopy() *AndroidKeyRestrictions {
	if in == nil {
		return nil
	}
	out := new(AndroidKeyRestrictions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiTarget) DeepCopyInto(out *ApiTarget) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiTarget.
func (in *ApiTarget) DeepCopy() *ApiTarget {
	if in == nil {
		return nil
	}
	out := new(ApiTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrowserKeyRestrictions) DeepCopyInto(out *BrowserKeyRestrictions) {
	*out = *in
	if in.AllowedReferrers != nil {
		in, out := &in.AllowedReferrers, &out.AllowedReferrers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrowserKeyRestrictions.
func (in *BrowserKeyRestrictions) DeepCopy() *BrowserKeyRestrictions {
	if in == nil {
		return nil
	}
	out := new(BrowserKeyRestrictions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IosKeyRestrictions) DeepCopyInto(out *IosKeyRestrictions) {
	*out = *in
	if in.AllowedBundleIds != nil {
		in, out := &in.AllowedBundleIds, &out.AllowedBundleIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IosKeyRestrictions.
func (in *IosKeyRestrictions) DeepCopy() *IosKeyRestrictions {
	if in == nil {
		return nil
	}
	out := new(IosKeyRestrictions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Key_AnnotationsEntry) DeepCopyInto(out *Key_AnnotationsEntry) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Key_AnnotationsEntry.
func (in *Key_AnnotationsEntry) DeepCopy() *Key_AnnotationsEntry {
	if in == nil {
		return nil
	}
	out := new(Key_AnnotationsEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Restrictions) DeepCopyInto(out *Restrictions) {
	*out = *in
	if in.BrowserKeyRestrictions != nil {
		in, out := &in.BrowserKeyRestrictions, &out.BrowserKeyRestrictions
		*out = new(BrowserKeyRestrictions)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerKeyRestrictions != nil {
		in, out := &in.ServerKeyRestrictions, &out.ServerKeyRestrictions
		*out = new(ServerKeyRestrictions)
		(*in).DeepCopyInto(*out)
	}
	if in.AndroidKeyRestrictions != nil {
		in, out := &in.AndroidKeyRestrictions, &out.AndroidKeyRestrictions
		*out = new(AndroidKeyRestrictions)
		(*in).DeepCopyInto(*out)
	}
	if in.IosKeyRestrictions != nil {
		in, out := &in.IosKeyRestrictions, &out.IosKeyRestrictions
		*out = new(IosKeyRestrictions)
		(*in).DeepCopyInto(*out)
	}
	if in.ApiTargets != nil {
		in, out := &in.ApiTargets, &out.ApiTargets
		*out = make([]ApiTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Restrictions.
func (in *Restrictions) DeepCopy() *Restrictions {
	if in == nil {
		return nil
	}
	out := new(Restrictions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerKeyRestrictions) DeepCopyInto(out *ServerKeyRestrictions) {
	*out = *in
	if in.AllowedIps != nil {
		in, out := &in.AllowedIps, &out.AllowedIps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerKeyRestrictions.
func (in *ServerKeyRestrictions) DeepCopy() *ServerKeyRestrictions {
	if in == nil {
		return nil
	}
	out := new(ServerKeyRestrictions)
	in.DeepCopyInto(out)
	return out
}
