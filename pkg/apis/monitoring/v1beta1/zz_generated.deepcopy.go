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
func (in *AlertpolicyAggregations) DeepCopyInto(out *AlertpolicyAggregations) {
	*out = *in
	if in.AlignmentPeriod != nil {
		in, out := &in.AlignmentPeriod, &out.AlignmentPeriod
		*out = new(string)
		**out = **in
	}
	if in.CrossSeriesReducer != nil {
		in, out := &in.CrossSeriesReducer, &out.CrossSeriesReducer
		*out = new(string)
		**out = **in
	}
	if in.GroupByFields != nil {
		in, out := &in.GroupByFields, &out.GroupByFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PerSeriesAligner != nil {
		in, out := &in.PerSeriesAligner, &out.PerSeriesAligner
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertpolicyAggregations.
func (in *AlertpolicyAggregations) DeepCopy() *AlertpolicyAggregations {
	if in == nil {
		return nil
	}
	out := new(AlertpolicyAggregations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertpolicyConditionAbsent) DeepCopyInto(out *AlertpolicyConditionAbsent) {
	*out = *in
	if in.Aggregations != nil {
		in, out := &in.Aggregations, &out.Aggregations
		*out = make([]AlertpolicyAggregations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.Trigger != nil {
		in, out := &in.Trigger, &out.Trigger
		*out = new(AlertpolicyTrigger)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertpolicyConditionAbsent.
func (in *AlertpolicyConditionAbsent) DeepCopy() *AlertpolicyConditionAbsent {
	if in == nil {
		return nil
	}
	out := new(AlertpolicyConditionAbsent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertpolicyConditionMonitoringQueryLanguage) DeepCopyInto(out *AlertpolicyConditionMonitoringQueryLanguage) {
	*out = *in
	if in.Trigger != nil {
		in, out := &in.Trigger, &out.Trigger
		*out = new(AlertpolicyTrigger)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertpolicyConditionMonitoringQueryLanguage.
func (in *AlertpolicyConditionMonitoringQueryLanguage) DeepCopy() *AlertpolicyConditionMonitoringQueryLanguage {
	if in == nil {
		return nil
	}
	out := new(AlertpolicyConditionMonitoringQueryLanguage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertpolicyConditionThreshold) DeepCopyInto(out *AlertpolicyConditionThreshold) {
	*out = *in
	if in.Aggregations != nil {
		in, out := &in.Aggregations, &out.Aggregations
		*out = make([]AlertpolicyAggregations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DenominatorAggregations != nil {
		in, out := &in.DenominatorAggregations, &out.DenominatorAggregations
		*out = make([]AlertpolicyDenominatorAggregations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DenominatorFilter != nil {
		in, out := &in.DenominatorFilter, &out.DenominatorFilter
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.ThresholdValue != nil {
		in, out := &in.ThresholdValue, &out.ThresholdValue
		*out = new(float64)
		**out = **in
	}
	if in.Trigger != nil {
		in, out := &in.Trigger, &out.Trigger
		*out = new(AlertpolicyTrigger)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertpolicyConditionThreshold.
func (in *AlertpolicyConditionThreshold) DeepCopy() *AlertpolicyConditionThreshold {
	if in == nil {
		return nil
	}
	out := new(AlertpolicyConditionThreshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertpolicyConditions) DeepCopyInto(out *AlertpolicyConditions) {
	*out = *in
	if in.ConditionAbsent != nil {
		in, out := &in.ConditionAbsent, &out.ConditionAbsent
		*out = new(AlertpolicyConditionAbsent)
		(*in).DeepCopyInto(*out)
	}
	if in.ConditionMonitoringQueryLanguage != nil {
		in, out := &in.ConditionMonitoringQueryLanguage, &out.ConditionMonitoringQueryLanguage
		*out = new(AlertpolicyConditionMonitoringQueryLanguage)
		(*in).DeepCopyInto(*out)
	}
	if in.ConditionThreshold != nil {
		in, out := &in.ConditionThreshold, &out.ConditionThreshold
		*out = new(AlertpolicyConditionThreshold)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertpolicyConditions.
func (in *AlertpolicyConditions) DeepCopy() *AlertpolicyConditions {
	if in == nil {
		return nil
	}
	out := new(AlertpolicyConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertpolicyCreationRecordStatus) DeepCopyInto(out *AlertpolicyCreationRecordStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertpolicyCreationRecordStatus.
func (in *AlertpolicyCreationRecordStatus) DeepCopy() *AlertpolicyCreationRecordStatus {
	if in == nil {
		return nil
	}
	out := new(AlertpolicyCreationRecordStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertpolicyDenominatorAggregations) DeepCopyInto(out *AlertpolicyDenominatorAggregations) {
	*out = *in
	if in.AlignmentPeriod != nil {
		in, out := &in.AlignmentPeriod, &out.AlignmentPeriod
		*out = new(string)
		**out = **in
	}
	if in.CrossSeriesReducer != nil {
		in, out := &in.CrossSeriesReducer, &out.CrossSeriesReducer
		*out = new(string)
		**out = **in
	}
	if in.GroupByFields != nil {
		in, out := &in.GroupByFields, &out.GroupByFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PerSeriesAligner != nil {
		in, out := &in.PerSeriesAligner, &out.PerSeriesAligner
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertpolicyDenominatorAggregations.
func (in *AlertpolicyDenominatorAggregations) DeepCopy() *AlertpolicyDenominatorAggregations {
	if in == nil {
		return nil
	}
	out := new(AlertpolicyDenominatorAggregations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertpolicyDocumentation) DeepCopyInto(out *AlertpolicyDocumentation) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.MimeType != nil {
		in, out := &in.MimeType, &out.MimeType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertpolicyDocumentation.
func (in *AlertpolicyDocumentation) DeepCopy() *AlertpolicyDocumentation {
	if in == nil {
		return nil
	}
	out := new(AlertpolicyDocumentation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertpolicyTrigger) DeepCopyInto(out *AlertpolicyTrigger) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int)
		**out = **in
	}
	if in.Percent != nil {
		in, out := &in.Percent, &out.Percent
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertpolicyTrigger.
func (in *AlertpolicyTrigger) DeepCopy() *AlertpolicyTrigger {
	if in == nil {
		return nil
	}
	out := new(AlertpolicyTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringAlertPolicy) DeepCopyInto(out *MonitoringAlertPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringAlertPolicy.
func (in *MonitoringAlertPolicy) DeepCopy() *MonitoringAlertPolicy {
	if in == nil {
		return nil
	}
	out := new(MonitoringAlertPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringAlertPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringAlertPolicyList) DeepCopyInto(out *MonitoringAlertPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitoringAlertPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringAlertPolicyList.
func (in *MonitoringAlertPolicyList) DeepCopy() *MonitoringAlertPolicyList {
	if in == nil {
		return nil
	}
	out := new(MonitoringAlertPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringAlertPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringAlertPolicySpec) DeepCopyInto(out *MonitoringAlertPolicySpec) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AlertpolicyConditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Documentation != nil {
		in, out := &in.Documentation, &out.Documentation
		*out = new(AlertpolicyDocumentation)
		(*in).DeepCopyInto(*out)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.NotificationChannels != nil {
		in, out := &in.NotificationChannels, &out.NotificationChannels
		*out = make([]v1alpha1.ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringAlertPolicySpec.
func (in *MonitoringAlertPolicySpec) DeepCopy() *MonitoringAlertPolicySpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringAlertPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringAlertPolicyStatus) DeepCopyInto(out *MonitoringAlertPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CreationRecord != nil {
		in, out := &in.CreationRecord, &out.CreationRecord
		*out = make([]AlertpolicyCreationRecordStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringAlertPolicyStatus.
func (in *MonitoringAlertPolicyStatus) DeepCopy() *MonitoringAlertPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(MonitoringAlertPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringGroup) DeepCopyInto(out *MonitoringGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringGroup.
func (in *MonitoringGroup) DeepCopy() *MonitoringGroup {
	if in == nil {
		return nil
	}
	out := new(MonitoringGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringGroupList) DeepCopyInto(out *MonitoringGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitoringGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringGroupList.
func (in *MonitoringGroupList) DeepCopy() *MonitoringGroupList {
	if in == nil {
		return nil
	}
	out := new(MonitoringGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringGroupSpec) DeepCopyInto(out *MonitoringGroupSpec) {
	*out = *in
	if in.IsCluster != nil {
		in, out := &in.IsCluster, &out.IsCluster
		*out = new(bool)
		**out = **in
	}
	if in.ParentRef != nil {
		in, out := &in.ParentRef, &out.ParentRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringGroupSpec.
func (in *MonitoringGroupSpec) DeepCopy() *MonitoringGroupSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringGroupStatus) DeepCopyInto(out *MonitoringGroupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringGroupStatus.
func (in *MonitoringGroupStatus) DeepCopy() *MonitoringGroupStatus {
	if in == nil {
		return nil
	}
	out := new(MonitoringGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringNotificationChannel) DeepCopyInto(out *MonitoringNotificationChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringNotificationChannel.
func (in *MonitoringNotificationChannel) DeepCopy() *MonitoringNotificationChannel {
	if in == nil {
		return nil
	}
	out := new(MonitoringNotificationChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringNotificationChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringNotificationChannelList) DeepCopyInto(out *MonitoringNotificationChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitoringNotificationChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringNotificationChannelList.
func (in *MonitoringNotificationChannelList) DeepCopy() *MonitoringNotificationChannelList {
	if in == nil {
		return nil
	}
	out := new(MonitoringNotificationChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringNotificationChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringNotificationChannelSpec) DeepCopyInto(out *MonitoringNotificationChannelSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.SensitiveLabels != nil {
		in, out := &in.SensitiveLabels, &out.SensitiveLabels
		*out = new(NotificationchannelSensitiveLabels)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringNotificationChannelSpec.
func (in *MonitoringNotificationChannelSpec) DeepCopy() *MonitoringNotificationChannelSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringNotificationChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringNotificationChannelStatus) DeepCopyInto(out *MonitoringNotificationChannelStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringNotificationChannelStatus.
func (in *MonitoringNotificationChannelStatus) DeepCopy() *MonitoringNotificationChannelStatus {
	if in == nil {
		return nil
	}
	out := new(MonitoringNotificationChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationchannelAuthToken) DeepCopyInto(out *NotificationchannelAuthToken) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(NotificationchannelValueFrom)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationchannelAuthToken.
func (in *NotificationchannelAuthToken) DeepCopy() *NotificationchannelAuthToken {
	if in == nil {
		return nil
	}
	out := new(NotificationchannelAuthToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationchannelPassword) DeepCopyInto(out *NotificationchannelPassword) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(NotificationchannelValueFrom)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationchannelPassword.
func (in *NotificationchannelPassword) DeepCopy() *NotificationchannelPassword {
	if in == nil {
		return nil
	}
	out := new(NotificationchannelPassword)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationchannelSensitiveLabels) DeepCopyInto(out *NotificationchannelSensitiveLabels) {
	*out = *in
	if in.AuthToken != nil {
		in, out := &in.AuthToken, &out.AuthToken
		*out = new(NotificationchannelAuthToken)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(NotificationchannelPassword)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceKey != nil {
		in, out := &in.ServiceKey, &out.ServiceKey
		*out = new(NotificationchannelServiceKey)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationchannelSensitiveLabels.
func (in *NotificationchannelSensitiveLabels) DeepCopy() *NotificationchannelSensitiveLabels {
	if in == nil {
		return nil
	}
	out := new(NotificationchannelSensitiveLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationchannelServiceKey) DeepCopyInto(out *NotificationchannelServiceKey) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(NotificationchannelValueFrom)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationchannelServiceKey.
func (in *NotificationchannelServiceKey) DeepCopy() *NotificationchannelServiceKey {
	if in == nil {
		return nil
	}
	out := new(NotificationchannelServiceKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationchannelValueFrom) DeepCopyInto(out *NotificationchannelValueFrom) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationchannelValueFrom.
func (in *NotificationchannelValueFrom) DeepCopy() *NotificationchannelValueFrom {
	if in == nil {
		return nil
	}
	out := new(NotificationchannelValueFrom)
	in.DeepCopyInto(out)
	return out
}
