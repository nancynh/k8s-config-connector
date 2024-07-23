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
func (in *InstanceMaintenancePolicy) DeepCopyInto(out *InstanceMaintenancePolicy) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceWindow != nil {
		in, out := &in.WeeklyMaintenanceWindow, &out.WeeklyMaintenanceWindow
		*out = make([]InstanceWeeklyMaintenanceWindow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMaintenancePolicy.
func (in *InstanceMaintenancePolicy) DeepCopy() *InstanceMaintenancePolicy {
	if in == nil {
		return nil
	}
	out := new(InstanceMaintenancePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceMaintenanceScheduleStatus) DeepCopyInto(out *InstanceMaintenanceScheduleStatus) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.ScheduleDeadlineTime != nil {
		in, out := &in.ScheduleDeadlineTime, &out.ScheduleDeadlineTime
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMaintenanceScheduleStatus.
func (in *InstanceMaintenanceScheduleStatus) DeepCopy() *InstanceMaintenanceScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceMaintenanceScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceMemcacheNodesStatus) DeepCopyInto(out *InstanceMemcacheNodesStatus) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.NodeId != nil {
		in, out := &in.NodeId, &out.NodeId
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMemcacheNodesStatus.
func (in *InstanceMemcacheNodesStatus) DeepCopy() *InstanceMemcacheNodesStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceMemcacheNodesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceMemcacheParameters) DeepCopyInto(out *InstanceMemcacheParameters) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMemcacheParameters.
func (in *InstanceMemcacheParameters) DeepCopy() *InstanceMemcacheParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceMemcacheParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceNodeConfig) DeepCopyInto(out *InstanceNodeConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceNodeConfig.
func (in *InstanceNodeConfig) DeepCopy() *InstanceNodeConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStartTime) DeepCopyInto(out *InstanceStartTime) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(int64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(int64)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(int64)
		**out = **in
	}
	if in.Seconds != nil {
		in, out := &in.Seconds, &out.Seconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStartTime.
func (in *InstanceStartTime) DeepCopy() *InstanceStartTime {
	if in == nil {
		return nil
	}
	out := new(InstanceStartTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceWeeklyMaintenanceWindow) DeepCopyInto(out *InstanceWeeklyMaintenanceWindow) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceWeeklyMaintenanceWindow.
func (in *InstanceWeeklyMaintenanceWindow) DeepCopy() *InstanceWeeklyMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := new(InstanceWeeklyMaintenanceWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcacheInstance) DeepCopyInto(out *MemcacheInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcacheInstance.
func (in *MemcacheInstance) DeepCopy() *MemcacheInstance {
	if in == nil {
		return nil
	}
	out := new(MemcacheInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemcacheInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcacheInstanceList) DeepCopyInto(out *MemcacheInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MemcacheInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcacheInstanceList.
func (in *MemcacheInstanceList) DeepCopy() *MemcacheInstanceList {
	if in == nil {
		return nil
	}
	out := new(MemcacheInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemcacheInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcacheInstanceSpec) DeepCopyInto(out *MemcacheInstanceSpec) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.MaintenancePolicy != nil {
		in, out := &in.MaintenancePolicy, &out.MaintenancePolicy
		*out = new(InstanceMaintenancePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.MemcacheParameters != nil {
		in, out := &in.MemcacheParameters, &out.MemcacheParameters
		*out = new(InstanceMemcacheParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.MemcacheVersion != nil {
		in, out := &in.MemcacheVersion, &out.MemcacheVersion
		*out = new(string)
		**out = **in
	}
	if in.NetworkRef != nil {
		in, out := &in.NetworkRef, &out.NetworkRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	out.NodeConfig = in.NodeConfig
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcacheInstanceSpec.
func (in *MemcacheInstanceSpec) DeepCopy() *MemcacheInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(MemcacheInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcacheInstanceStatus) DeepCopyInto(out *MemcacheInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.DiscoveryEndpoint != nil {
		in, out := &in.DiscoveryEndpoint, &out.DiscoveryEndpoint
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceSchedule != nil {
		in, out := &in.MaintenanceSchedule, &out.MaintenanceSchedule
		*out = make([]InstanceMaintenanceScheduleStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemcacheFullVersion != nil {
		in, out := &in.MemcacheFullVersion, &out.MemcacheFullVersion
		*out = new(string)
		**out = **in
	}
	if in.MemcacheNodes != nil {
		in, out := &in.MemcacheNodes, &out.MemcacheNodes
		*out = make([]InstanceMemcacheNodesStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcacheInstanceStatus.
func (in *MemcacheInstanceStatus) DeepCopy() *MemcacheInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(MemcacheInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
