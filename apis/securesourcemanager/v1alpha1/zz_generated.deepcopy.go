//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance_HostConfig) DeepCopyInto(out *Instance_HostConfig) {
	*out = *in
	if in.HTML != nil {
		in, out := &in.HTML, &out.HTML
		*out = new(string)
		**out = **in
	}
	if in.Api != nil {
		in, out := &in.Api, &out.Api
		*out = new(string)
		**out = **in
	}
	if in.GitHTTP != nil {
		in, out := &in.GitHTTP, &out.GitHTTP
		*out = new(string)
		**out = **in
	}
	if in.GitSSH != nil {
		in, out := &in.GitSSH, &out.GitSSH
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance_HostConfig.
func (in *Instance_HostConfig) DeepCopy() *Instance_HostConfig {
	if in == nil {
		return nil
	}
	out := new(Instance_HostConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance_PrivateConfig) DeepCopyInto(out *Instance_PrivateConfig) {
	*out = *in
	if in.IsPrivate != nil {
		in, out := &in.IsPrivate, &out.IsPrivate
		*out = new(bool)
		**out = **in
	}
	if in.CaPool != nil {
		in, out := &in.CaPool, &out.CaPool
		*out = new(string)
		**out = **in
	}
	if in.HTTPServiceAttachment != nil {
		in, out := &in.HTTPServiceAttachment, &out.HTTPServiceAttachment
		*out = new(string)
		**out = **in
	}
	if in.SSHServiceAttachment != nil {
		in, out := &in.SSHServiceAttachment, &out.SSHServiceAttachment
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance_PrivateConfig.
func (in *Instance_PrivateConfig) DeepCopy() *Instance_PrivateConfig {
	if in == nil {
		return nil
	}
	out := new(Instance_PrivateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		*out = new(string)
		**out = **in
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Uris != nil {
		in, out := &in.Uris, &out.Uris
		*out = new(Repository_URIs)
		(*in).DeepCopyInto(*out)
	}
	if in.InitialConfig != nil {
		in, out := &in.InitialConfig, &out.InitialConfig
		*out = new(Repository_InitialConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository_InitialConfig) DeepCopyInto(out *Repository_InitialConfig) {
	*out = *in
	if in.DefaultBranch != nil {
		in, out := &in.DefaultBranch, &out.DefaultBranch
		*out = new(string)
		**out = **in
	}
	if in.Gitignores != nil {
		in, out := &in.Gitignores, &out.Gitignores
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.License != nil {
		in, out := &in.License, &out.License
		*out = new(string)
		**out = **in
	}
	if in.Readme != nil {
		in, out := &in.Readme, &out.Readme
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository_InitialConfig.
func (in *Repository_InitialConfig) DeepCopy() *Repository_InitialConfig {
	if in == nil {
		return nil
	}
	out := new(Repository_InitialConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository_URIs) DeepCopyInto(out *Repository_URIs) {
	*out = *in
	if in.HTML != nil {
		in, out := &in.HTML, &out.HTML
		*out = new(string)
		**out = **in
	}
	if in.GitHTTPS != nil {
		in, out := &in.GitHTTPS, &out.GitHTTPS
		*out = new(string)
		**out = **in
	}
	if in.Api != nil {
		in, out := &in.Api, &out.Api
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository_URIs.
func (in *Repository_URIs) DeepCopy() *Repository_URIs {
	if in == nil {
		return nil
	}
	out := new(Repository_URIs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureSourceManagerInstance) DeepCopyInto(out *SecureSourceManagerInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureSourceManagerInstance.
func (in *SecureSourceManagerInstance) DeepCopy() *SecureSourceManagerInstance {
	if in == nil {
		return nil
	}
	out := new(SecureSourceManagerInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecureSourceManagerInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureSourceManagerInstanceList) DeepCopyInto(out *SecureSourceManagerInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecureSourceManagerInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureSourceManagerInstanceList.
func (in *SecureSourceManagerInstanceList) DeepCopy() *SecureSourceManagerInstanceList {
	if in == nil {
		return nil
	}
	out := new(SecureSourceManagerInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecureSourceManagerInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureSourceManagerInstanceObservedState) DeepCopyInto(out *SecureSourceManagerInstanceObservedState) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateNote != nil {
		in, out := &in.StateNote, &out.StateNote
		*out = new(string)
		**out = **in
	}
	if in.HostConfig != nil {
		in, out := &in.HostConfig, &out.HostConfig
		*out = new(Instance_HostConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureSourceManagerInstanceObservedState.
func (in *SecureSourceManagerInstanceObservedState) DeepCopy() *SecureSourceManagerInstanceObservedState {
	if in == nil {
		return nil
	}
	out := new(SecureSourceManagerInstanceObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureSourceManagerInstanceSpec) DeepCopyInto(out *SecureSourceManagerInstanceSpec) {
	*out = *in
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.KmsKey != nil {
		in, out := &in.KmsKey, &out.KmsKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureSourceManagerInstanceSpec.
func (in *SecureSourceManagerInstanceSpec) DeepCopy() *SecureSourceManagerInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(SecureSourceManagerInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureSourceManagerInstanceStatus) DeepCopyInto(out *SecureSourceManagerInstanceStatus) {
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
		*out = new(SecureSourceManagerInstanceObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureSourceManagerInstanceStatus.
func (in *SecureSourceManagerInstanceStatus) DeepCopy() *SecureSourceManagerInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(SecureSourceManagerInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
