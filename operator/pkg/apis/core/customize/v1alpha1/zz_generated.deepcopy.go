//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerResourceSpec) DeepCopyInto(out *ContainerResourceSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerResourceSpec.
func (in *ContainerResourceSpec) DeepCopy() *ContainerResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResource) DeepCopyInto(out *ControllerResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResource.
func (in *ControllerResource) DeepCopy() *ControllerResource {
	if in == nil {
		return nil
	}
	out := new(ControllerResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResourceList) DeepCopyInto(out *ControllerResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResourceList.
func (in *ControllerResourceList) DeepCopy() *ControllerResourceList {
	if in == nil {
		return nil
	}
	out := new(ControllerResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResourceSpec) DeepCopyInto(out *ControllerResourceSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]ContainerResourceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResourceSpec.
func (in *ControllerResourceSpec) DeepCopy() *ControllerResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResourceStatus) DeepCopyInto(out *ControllerResourceStatus) {
	*out = *in
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResourceStatus.
func (in *ControllerResourceStatus) DeepCopy() *ControllerResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerResourceStatus)
	in.DeepCopyInto(out)
	return out
}
