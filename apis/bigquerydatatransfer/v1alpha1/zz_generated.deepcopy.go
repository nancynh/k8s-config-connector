//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDataTransferConfig) DeepCopyInto(out *BigQueryDataTransferConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDataTransferConfig.
func (in *BigQueryDataTransferConfig) DeepCopy() *BigQueryDataTransferConfig {
	if in == nil {
		return nil
	}
	out := new(BigQueryDataTransferConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryDataTransferConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDataTransferConfigList) DeepCopyInto(out *BigQueryDataTransferConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryDataTransferConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDataTransferConfigList.
func (in *BigQueryDataTransferConfigList) DeepCopy() *BigQueryDataTransferConfigList {
	if in == nil {
		return nil
	}
	out := new(BigQueryDataTransferConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryDataTransferConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDataTransferConfigObservedState) DeepCopyInto(out *BigQueryDataTransferConfigObservedState) {
	*out = *in
	if in.DatasetRegion != nil {
		in, out := &in.DatasetRegion, &out.DatasetRegion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NextRunTime != nil {
		in, out := &in.NextRunTime, &out.NextRunTime
		*out = new(string)
		**out = **in
	}
	if in.OwnerInfo != nil {
		in, out := &in.OwnerInfo, &out.OwnerInfo
		*out = new(UserInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDataTransferConfigObservedState.
func (in *BigQueryDataTransferConfigObservedState) DeepCopy() *BigQueryDataTransferConfigObservedState {
	if in == nil {
		return nil
	}
	out := new(BigQueryDataTransferConfigObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDataTransferConfigSpec) DeepCopyInto(out *BigQueryDataTransferConfigSpec) {
	*out = *in
	if in.DataRefreshWindowDays != nil {
		in, out := &in.DataRefreshWindowDays, &out.DataRefreshWindowDays
		*out = new(int32)
		**out = **in
	}
	if in.DataSourceID != nil {
		in, out := &in.DataSourceID, &out.DataSourceID
		*out = new(string)
		**out = **in
	}
	out.DatasetRef = in.DatasetRef
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EmailPreferences != nil {
		in, out := &in.EmailPreferences, &out.EmailPreferences
		*out = new(EmailPreferences)
		(*in).DeepCopyInto(*out)
	}
	if in.EncryptionConfiguration != nil {
		in, out := &in.EncryptionConfiguration, &out.EncryptionConfiguration
		*out = new(EncryptionConfiguration)
		**out = **in
	}
	out.PubSubTopicRef = in.PubSubTopicRef
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Parent.DeepCopyInto(&out.Parent)
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.ScheduleOptions != nil {
		in, out := &in.ScheduleOptions, &out.ScheduleOptions
		*out = new(ScheduleOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDataTransferConfigSpec.
func (in *BigQueryDataTransferConfigSpec) DeepCopy() *BigQueryDataTransferConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryDataTransferConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDataTransferConfigStatus) DeepCopyInto(out *BigQueryDataTransferConfigStatus) {
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
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(BigQueryDataTransferConfigObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDataTransferConfigStatus.
func (in *BigQueryDataTransferConfigStatus) DeepCopy() *BigQueryDataTransferConfigStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryDataTransferConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailPreferences) DeepCopyInto(out *EmailPreferences) {
	*out = *in
	if in.EnableFailureEmail != nil {
		in, out := &in.EnableFailureEmail, &out.EnableFailureEmail
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailPreferences.
func (in *EmailPreferences) DeepCopy() *EmailPreferences {
	if in == nil {
		return nil
	}
	out := new(EmailPreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfiguration) DeepCopyInto(out *EncryptionConfiguration) {
	*out = *in
	out.KmsKeyRef = in.KmsKeyRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfiguration.
func (in *EncryptionConfiguration) DeepCopy() *EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parent) DeepCopyInto(out *Parent) {
	*out = *in
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parent.
func (in *Parent) DeepCopy() *Parent {
	if in == nil {
		return nil
	}
	out := new(Parent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleOptions) DeepCopyInto(out *ScheduleOptions) {
	*out = *in
	if in.DisableAutoScheduling != nil {
		in, out := &in.DisableAutoScheduling, &out.DisableAutoScheduling
		*out = new(bool)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleOptions.
func (in *ScheduleOptions) DeepCopy() *ScheduleOptions {
	if in == nil {
		return nil
	}
	out := new(ScheduleOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserInfo) DeepCopyInto(out *UserInfo) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserInfo.
func (in *UserInfo) DeepCopy() *UserInfo {
	if in == nil {
		return nil
	}
	out := new(UserInfo)
	in.DeepCopyInto(out)
	return out
}
