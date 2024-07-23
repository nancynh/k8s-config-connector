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

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryReservationCapacityCommitment) DeepCopyInto(out *BigQueryReservationCapacityCommitment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryReservationCapacityCommitment.
func (in *BigQueryReservationCapacityCommitment) DeepCopy() *BigQueryReservationCapacityCommitment {
	if in == nil {
		return nil
	}
	out := new(BigQueryReservationCapacityCommitment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryReservationCapacityCommitment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryReservationCapacityCommitmentList) DeepCopyInto(out *BigQueryReservationCapacityCommitmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryReservationCapacityCommitment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryReservationCapacityCommitmentList.
func (in *BigQueryReservationCapacityCommitmentList) DeepCopy() *BigQueryReservationCapacityCommitmentList {
	if in == nil {
		return nil
	}
	out := new(BigQueryReservationCapacityCommitmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryReservationCapacityCommitmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryReservationCapacityCommitmentSpec) DeepCopyInto(out *BigQueryReservationCapacityCommitmentSpec) {
	*out = *in
	if in.Edition != nil {
		in, out := &in.Edition, &out.Edition
		*out = new(string)
		**out = **in
	}
	if in.EnforceSingleAdminProjectPerOrg != nil {
		in, out := &in.EnforceSingleAdminProjectPerOrg, &out.EnforceSingleAdminProjectPerOrg
		*out = new(string)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.RenewalPlan != nil {
		in, out := &in.RenewalPlan, &out.RenewalPlan
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryReservationCapacityCommitmentSpec.
func (in *BigQueryReservationCapacityCommitmentSpec) DeepCopy() *BigQueryReservationCapacityCommitmentSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryReservationCapacityCommitmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryReservationCapacityCommitmentStatus) DeepCopyInto(out *BigQueryReservationCapacityCommitmentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CommitmentEndTime != nil {
		in, out := &in.CommitmentEndTime, &out.CommitmentEndTime
		*out = new(string)
		**out = **in
	}
	if in.CommitmentStartTime != nil {
		in, out := &in.CommitmentStartTime, &out.CommitmentStartTime
		*out = new(string)
		**out = **in
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
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryReservationCapacityCommitmentStatus.
func (in *BigQueryReservationCapacityCommitmentStatus) DeepCopy() *BigQueryReservationCapacityCommitmentStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryReservationCapacityCommitmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryReservationReservation) DeepCopyInto(out *BigQueryReservationReservation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryReservationReservation.
func (in *BigQueryReservationReservation) DeepCopy() *BigQueryReservationReservation {
	if in == nil {
		return nil
	}
	out := new(BigQueryReservationReservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryReservationReservation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryReservationReservationList) DeepCopyInto(out *BigQueryReservationReservationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryReservationReservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryReservationReservationList.
func (in *BigQueryReservationReservationList) DeepCopy() *BigQueryReservationReservationList {
	if in == nil {
		return nil
	}
	out := new(BigQueryReservationReservationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryReservationReservationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryReservationReservationSpec) DeepCopyInto(out *BigQueryReservationReservationSpec) {
	*out = *in
	if in.Autoscale != nil {
		in, out := &in.Autoscale, &out.Autoscale
		*out = new(ReservationAutoscale)
		(*in).DeepCopyInto(*out)
	}
	if in.Concurrency != nil {
		in, out := &in.Concurrency, &out.Concurrency
		*out = new(int64)
		**out = **in
	}
	if in.Edition != nil {
		in, out := &in.Edition, &out.Edition
		*out = new(string)
		**out = **in
	}
	if in.IgnoreIdleSlots != nil {
		in, out := &in.IgnoreIdleSlots, &out.IgnoreIdleSlots
		*out = new(bool)
		**out = **in
	}
	if in.MultiRegionAuxiliary != nil {
		in, out := &in.MultiRegionAuxiliary, &out.MultiRegionAuxiliary
		*out = new(bool)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryReservationReservationSpec.
func (in *BigQueryReservationReservationSpec) DeepCopy() *BigQueryReservationReservationSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryReservationReservationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryReservationReservationStatus) DeepCopyInto(out *BigQueryReservationReservationStatus) {
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryReservationReservationStatus.
func (in *BigQueryReservationReservationStatus) DeepCopy() *BigQueryReservationReservationStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryReservationReservationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservationAutoscale) DeepCopyInto(out *ReservationAutoscale) {
	*out = *in
	if in.CurrentSlots != nil {
		in, out := &in.CurrentSlots, &out.CurrentSlots
		*out = new(int64)
		**out = **in
	}
	if in.MaxSlots != nil {
		in, out := &in.MaxSlots, &out.MaxSlots
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservationAutoscale.
func (in *ReservationAutoscale) DeepCopy() *ReservationAutoscale {
	if in == nil {
		return nil
	}
	out := new(ReservationAutoscale)
	in.DeepCopyInto(out)
	return out
}
