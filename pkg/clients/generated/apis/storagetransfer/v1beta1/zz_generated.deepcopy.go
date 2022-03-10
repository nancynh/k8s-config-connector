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
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobAccessKeyId) DeepCopyInto(out *JobAccessKeyId) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(JobValueFrom)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobAccessKeyId.
func (in *JobAccessKeyId) DeepCopy() *JobAccessKeyId {
	if in == nil {
		return nil
	}
	out := new(JobAccessKeyId)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobAwsAccessKey) DeepCopyInto(out *JobAwsAccessKey) {
	*out = *in
	in.AccessKeyId.DeepCopyInto(&out.AccessKeyId)
	in.SecretAccessKey.DeepCopyInto(&out.SecretAccessKey)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobAwsAccessKey.
func (in *JobAwsAccessKey) DeepCopy() *JobAwsAccessKey {
	if in == nil {
		return nil
	}
	out := new(JobAwsAccessKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobAwsS3DataSource) DeepCopyInto(out *JobAwsS3DataSource) {
	*out = *in
	if in.AwsAccessKey != nil {
		in, out := &in.AwsAccessKey, &out.AwsAccessKey
		*out = new(JobAwsAccessKey)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobAwsS3DataSource.
func (in *JobAwsS3DataSource) DeepCopy() *JobAwsS3DataSource {
	if in == nil {
		return nil
	}
	out := new(JobAwsS3DataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobAzureBlobStorageDataSource) DeepCopyInto(out *JobAzureBlobStorageDataSource) {
	*out = *in
	in.AzureCredentials.DeepCopyInto(&out.AzureCredentials)
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobAzureBlobStorageDataSource.
func (in *JobAzureBlobStorageDataSource) DeepCopy() *JobAzureBlobStorageDataSource {
	if in == nil {
		return nil
	}
	out := new(JobAzureBlobStorageDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobAzureCredentials) DeepCopyInto(out *JobAzureCredentials) {
	*out = *in
	in.SasToken.DeepCopyInto(&out.SasToken)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobAzureCredentials.
func (in *JobAzureCredentials) DeepCopy() *JobAzureCredentials {
	if in == nil {
		return nil
	}
	out := new(JobAzureCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobGcsDataSink) DeepCopyInto(out *JobGcsDataSink) {
	*out = *in
	out.BucketRef = in.BucketRef
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobGcsDataSink.
func (in *JobGcsDataSink) DeepCopy() *JobGcsDataSink {
	if in == nil {
		return nil
	}
	out := new(JobGcsDataSink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobGcsDataSource) DeepCopyInto(out *JobGcsDataSource) {
	*out = *in
	out.BucketRef = in.BucketRef
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobGcsDataSource.
func (in *JobGcsDataSource) DeepCopy() *JobGcsDataSource {
	if in == nil {
		return nil
	}
	out := new(JobGcsDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobHttpDataSource) DeepCopyInto(out *JobHttpDataSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobHttpDataSource.
func (in *JobHttpDataSource) DeepCopy() *JobHttpDataSource {
	if in == nil {
		return nil
	}
	out := new(JobHttpDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobObjectConditions) DeepCopyInto(out *JobObjectConditions) {
	*out = *in
	if in.ExcludePrefixes != nil {
		in, out := &in.ExcludePrefixes, &out.ExcludePrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludePrefixes != nil {
		in, out := &in.IncludePrefixes, &out.IncludePrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MaxTimeElapsedSinceLastModification != nil {
		in, out := &in.MaxTimeElapsedSinceLastModification, &out.MaxTimeElapsedSinceLastModification
		*out = new(string)
		**out = **in
	}
	if in.MinTimeElapsedSinceLastModification != nil {
		in, out := &in.MinTimeElapsedSinceLastModification, &out.MinTimeElapsedSinceLastModification
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobObjectConditions.
func (in *JobObjectConditions) DeepCopy() *JobObjectConditions {
	if in == nil {
		return nil
	}
	out := new(JobObjectConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobPosixDataSink) DeepCopyInto(out *JobPosixDataSink) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobPosixDataSink.
func (in *JobPosixDataSink) DeepCopy() *JobPosixDataSink {
	if in == nil {
		return nil
	}
	out := new(JobPosixDataSink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobPosixDataSource) DeepCopyInto(out *JobPosixDataSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobPosixDataSource.
func (in *JobPosixDataSource) DeepCopy() *JobPosixDataSource {
	if in == nil {
		return nil
	}
	out := new(JobPosixDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSasToken) DeepCopyInto(out *JobSasToken) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(JobValueFrom)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSasToken.
func (in *JobSasToken) DeepCopy() *JobSasToken {
	if in == nil {
		return nil
	}
	out := new(JobSasToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSchedule) DeepCopyInto(out *JobSchedule) {
	*out = *in
	if in.ScheduleEndDate != nil {
		in, out := &in.ScheduleEndDate, &out.ScheduleEndDate
		*out = new(JobScheduleEndDate)
		**out = **in
	}
	out.ScheduleStartDate = in.ScheduleStartDate
	if in.StartTimeOfDay != nil {
		in, out := &in.StartTimeOfDay, &out.StartTimeOfDay
		*out = new(JobStartTimeOfDay)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSchedule.
func (in *JobSchedule) DeepCopy() *JobSchedule {
	if in == nil {
		return nil
	}
	out := new(JobSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobScheduleEndDate) DeepCopyInto(out *JobScheduleEndDate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobScheduleEndDate.
func (in *JobScheduleEndDate) DeepCopy() *JobScheduleEndDate {
	if in == nil {
		return nil
	}
	out := new(JobScheduleEndDate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobScheduleStartDate) DeepCopyInto(out *JobScheduleStartDate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobScheduleStartDate.
func (in *JobScheduleStartDate) DeepCopy() *JobScheduleStartDate {
	if in == nil {
		return nil
	}
	out := new(JobScheduleStartDate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSecretAccessKey) DeepCopyInto(out *JobSecretAccessKey) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(JobValueFrom)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSecretAccessKey.
func (in *JobSecretAccessKey) DeepCopy() *JobSecretAccessKey {
	if in == nil {
		return nil
	}
	out := new(JobSecretAccessKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobStartTimeOfDay) DeepCopyInto(out *JobStartTimeOfDay) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobStartTimeOfDay.
func (in *JobStartTimeOfDay) DeepCopy() *JobStartTimeOfDay {
	if in == nil {
		return nil
	}
	out := new(JobStartTimeOfDay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTransferOptions) DeepCopyInto(out *JobTransferOptions) {
	*out = *in
	if in.DeleteObjectsFromSourceAfterTransfer != nil {
		in, out := &in.DeleteObjectsFromSourceAfterTransfer, &out.DeleteObjectsFromSourceAfterTransfer
		*out = new(bool)
		**out = **in
	}
	if in.DeleteObjectsUniqueInSink != nil {
		in, out := &in.DeleteObjectsUniqueInSink, &out.DeleteObjectsUniqueInSink
		*out = new(bool)
		**out = **in
	}
	if in.OverwriteObjectsAlreadyExistingInSink != nil {
		in, out := &in.OverwriteObjectsAlreadyExistingInSink, &out.OverwriteObjectsAlreadyExistingInSink
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTransferOptions.
func (in *JobTransferOptions) DeepCopy() *JobTransferOptions {
	if in == nil {
		return nil
	}
	out := new(JobTransferOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTransferSpec) DeepCopyInto(out *JobTransferSpec) {
	*out = *in
	if in.AwsS3DataSource != nil {
		in, out := &in.AwsS3DataSource, &out.AwsS3DataSource
		*out = new(JobAwsS3DataSource)
		(*in).DeepCopyInto(*out)
	}
	if in.AzureBlobStorageDataSource != nil {
		in, out := &in.AzureBlobStorageDataSource, &out.AzureBlobStorageDataSource
		*out = new(JobAzureBlobStorageDataSource)
		(*in).DeepCopyInto(*out)
	}
	if in.GcsDataSink != nil {
		in, out := &in.GcsDataSink, &out.GcsDataSink
		*out = new(JobGcsDataSink)
		(*in).DeepCopyInto(*out)
	}
	if in.GcsDataSource != nil {
		in, out := &in.GcsDataSource, &out.GcsDataSource
		*out = new(JobGcsDataSource)
		(*in).DeepCopyInto(*out)
	}
	if in.HttpDataSource != nil {
		in, out := &in.HttpDataSource, &out.HttpDataSource
		*out = new(JobHttpDataSource)
		**out = **in
	}
	if in.ObjectConditions != nil {
		in, out := &in.ObjectConditions, &out.ObjectConditions
		*out = new(JobObjectConditions)
		(*in).DeepCopyInto(*out)
	}
	if in.PosixDataSink != nil {
		in, out := &in.PosixDataSink, &out.PosixDataSink
		*out = new(JobPosixDataSink)
		**out = **in
	}
	if in.PosixDataSource != nil {
		in, out := &in.PosixDataSource, &out.PosixDataSource
		*out = new(JobPosixDataSource)
		**out = **in
	}
	if in.TransferOptions != nil {
		in, out := &in.TransferOptions, &out.TransferOptions
		*out = new(JobTransferOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTransferSpec.
func (in *JobTransferSpec) DeepCopy() *JobTransferSpec {
	if in == nil {
		return nil
	}
	out := new(JobTransferSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobValueFrom) DeepCopyInto(out *JobValueFrom) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobValueFrom.
func (in *JobValueFrom) DeepCopy() *JobValueFrom {
	if in == nil {
		return nil
	}
	out := new(JobValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageTransferJob) DeepCopyInto(out *StorageTransferJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageTransferJob.
func (in *StorageTransferJob) DeepCopy() *StorageTransferJob {
	if in == nil {
		return nil
	}
	out := new(StorageTransferJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageTransferJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageTransferJobList) DeepCopyInto(out *StorageTransferJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageTransferJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageTransferJobList.
func (in *StorageTransferJobList) DeepCopy() *StorageTransferJobList {
	if in == nil {
		return nil
	}
	out := new(StorageTransferJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageTransferJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageTransferJobSpec) DeepCopyInto(out *StorageTransferJobSpec) {
	*out = *in
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(JobSchedule)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	in.TransferSpec.DeepCopyInto(&out.TransferSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageTransferJobSpec.
func (in *StorageTransferJobSpec) DeepCopy() *StorageTransferJobSpec {
	if in == nil {
		return nil
	}
	out := new(StorageTransferJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageTransferJobStatus) DeepCopyInto(out *StorageTransferJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageTransferJobStatus.
func (in *StorageTransferJobStatus) DeepCopy() *StorageTransferJobStatus {
	if in == nil {
		return nil
	}
	out := new(StorageTransferJobStatus)
	in.DeepCopyInto(out)
	return out
}