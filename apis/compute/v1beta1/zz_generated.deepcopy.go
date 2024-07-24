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

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeForwardingRule) DeepCopyInto(out *ComputeForwardingRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeForwardingRule.
func (in *ComputeForwardingRule) DeepCopy() *ComputeForwardingRule {
	if in == nil {
		return nil
	}
	out := new(ComputeForwardingRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeForwardingRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeForwardingRuleList) DeepCopyInto(out *ComputeForwardingRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComputeForwardingRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeForwardingRuleList.
func (in *ComputeForwardingRuleList) DeepCopy() *ComputeForwardingRuleList {
	if in == nil {
		return nil
	}
	out := new(ComputeForwardingRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeForwardingRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeForwardingRuleSpec) DeepCopyInto(out *ComputeForwardingRuleSpec) {
	*out = *in
	if in.AllPorts != nil {
		in, out := &in.AllPorts, &out.AllPorts
		*out = new(bool)
		**out = **in
	}
	if in.AllowGlobalAccess != nil {
		in, out := &in.AllowGlobalAccess, &out.AllowGlobalAccess
		*out = new(bool)
		**out = **in
	}
	if in.AllowPscGlobalAccess != nil {
		in, out := &in.AllowPscGlobalAccess, &out.AllowPscGlobalAccess
		*out = new(bool)
		**out = **in
	}
	if in.BackendServiceRef != nil {
		in, out := &in.BackendServiceRef, &out.BackendServiceRef
		*out = new(refsv1beta1.ComputeBackendServiceRef)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(ForwardingruleIpAddress)
		(*in).DeepCopyInto(*out)
	}
	if in.IpProtocol != nil {
		in, out := &in.IpProtocol, &out.IpProtocol
		*out = new(string)
		**out = **in
	}
	if in.IpVersion != nil {
		in, out := &in.IpVersion, &out.IpVersion
		*out = new(string)
		**out = **in
	}
	if in.IsMirroringCollector != nil {
		in, out := &in.IsMirroringCollector, &out.IsMirroringCollector
		*out = new(bool)
		**out = **in
	}
	if in.LoadBalancingScheme != nil {
		in, out := &in.LoadBalancingScheme, &out.LoadBalancingScheme
		*out = new(string)
		**out = **in
	}
	if in.MetadataFilters != nil {
		in, out := &in.MetadataFilters, &out.MetadataFilters
		*out = make([]ForwardingruleMetadataFilters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkRef != nil {
		in, out := &in.NetworkRef, &out.NetworkRef
		*out = new(refsv1beta1.ComputeNetworkRef)
		**out = **in
	}
	if in.NetworkTier != nil {
		in, out := &in.NetworkTier, &out.NetworkTier
		*out = new(string)
		**out = **in
	}
	if in.NoAutomateDnsZone != nil {
		in, out := &in.NoAutomateDnsZone, &out.NoAutomateDnsZone
		*out = new(bool)
		**out = **in
	}
	if in.PortRange != nil {
		in, out := &in.PortRange, &out.PortRange
		*out = new(string)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ServiceDirectoryRegistrations != nil {
		in, out := &in.ServiceDirectoryRegistrations, &out.ServiceDirectoryRegistrations
		*out = make([]ForwardingruleServiceDirectoryRegistrations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceLabel != nil {
		in, out := &in.ServiceLabel, &out.ServiceLabel
		*out = new(string)
		**out = **in
	}
	if in.SourceIpRanges != nil {
		in, out := &in.SourceIpRanges, &out.SourceIpRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetworkRef != nil {
		in, out := &in.SubnetworkRef, &out.SubnetworkRef
		*out = new(refsv1beta1.ComputeSubnetworkRef)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(ForwardingruleTarget)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeForwardingRuleSpec.
func (in *ComputeForwardingRuleSpec) DeepCopy() *ComputeForwardingRuleSpec {
	if in == nil {
		return nil
	}
	out := new(ComputeForwardingRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeForwardingRuleStatus) DeepCopyInto(out *ComputeForwardingRuleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.BaseForwardingRule != nil {
		in, out := &in.BaseForwardingRule, &out.BaseForwardingRule
		*out = new(string)
		**out = **in
	}
	if in.CreationTimestamp != nil {
		in, out := &in.CreationTimestamp, &out.CreationTimestamp
		*out = new(string)
		**out = **in
	}
	if in.LabelFingerprint != nil {
		in, out := &in.LabelFingerprint, &out.LabelFingerprint
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.PscConnectionId != nil {
		in, out := &in.PscConnectionId, &out.PscConnectionId
		*out = new(string)
		**out = **in
	}
	if in.PscConnectionStatus != nil {
		in, out := &in.PscConnectionStatus, &out.PscConnectionStatus
		*out = new(string)
		**out = **in
	}
	if in.SelfLink != nil {
		in, out := &in.SelfLink, &out.SelfLink
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeForwardingRuleStatus.
func (in *ComputeForwardingRuleStatus) DeepCopy() *ComputeForwardingRuleStatus {
	if in == nil {
		return nil
	}
	out := new(ComputeForwardingRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingruleFilterLabels) DeepCopyInto(out *ForwardingruleFilterLabels) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingruleFilterLabels.
func (in *ForwardingruleFilterLabels) DeepCopy() *ForwardingruleFilterLabels {
	if in == nil {
		return nil
	}
	out := new(ForwardingruleFilterLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingruleIpAddress) DeepCopyInto(out *ForwardingruleIpAddress) {
	*out = *in
	if in.AddressRef != nil {
		in, out := &in.AddressRef, &out.AddressRef
		*out = new(refsv1beta1.ComputeAddressRef)
		**out = **in
	}
	if in.Ip != nil {
		in, out := &in.Ip, &out.Ip
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingruleIpAddress.
func (in *ForwardingruleIpAddress) DeepCopy() *ForwardingruleIpAddress {
	if in == nil {
		return nil
	}
	out := new(ForwardingruleIpAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingruleMetadataFilters) DeepCopyInto(out *ForwardingruleMetadataFilters) {
	*out = *in
	if in.FilterLabels != nil {
		in, out := &in.FilterLabels, &out.FilterLabels
		*out = make([]ForwardingruleFilterLabels, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingruleMetadataFilters.
func (in *ForwardingruleMetadataFilters) DeepCopy() *ForwardingruleMetadataFilters {
	if in == nil {
		return nil
	}
	out := new(ForwardingruleMetadataFilters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingruleServiceDirectoryRegistrations) DeepCopyInto(out *ForwardingruleServiceDirectoryRegistrations) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingruleServiceDirectoryRegistrations.
func (in *ForwardingruleServiceDirectoryRegistrations) DeepCopy() *ForwardingruleServiceDirectoryRegistrations {
	if in == nil {
		return nil
	}
	out := new(ForwardingruleServiceDirectoryRegistrations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingruleTarget) DeepCopyInto(out *ForwardingruleTarget) {
	*out = *in
	if in.ServiceAttachmentRef != nil {
		in, out := &in.ServiceAttachmentRef, &out.ServiceAttachmentRef
		*out = new(refsv1beta1.ComputeServiceAttachmentRef)
		**out = **in
	}
	if in.TargetGRPCProxyRef != nil {
		in, out := &in.TargetGRPCProxyRef, &out.TargetGRPCProxyRef
		*out = new(refsv1beta1.ComputeTargetGrpcProxyRef)
		**out = **in
	}
	if in.TargetHTTPProxyRef != nil {
		in, out := &in.TargetHTTPProxyRef, &out.TargetHTTPProxyRef
		*out = new(refsv1beta1.ComputeTargetHTTPProxyRef)
		**out = **in
	}
	if in.TargetHTTPSProxyRef != nil {
		in, out := &in.TargetHTTPSProxyRef, &out.TargetHTTPSProxyRef
		*out = new(refsv1beta1.ComputeTargetHTTPSProxyRef)
		**out = **in
	}
	if in.TargetSSLProxyRef != nil {
		in, out := &in.TargetSSLProxyRef, &out.TargetSSLProxyRef
		*out = new(refsv1beta1.ComputeTargetSSLProxyRef)
		**out = **in
	}
	if in.TargetTCPProxyRef != nil {
		in, out := &in.TargetTCPProxyRef, &out.TargetTCPProxyRef
		*out = new(refsv1beta1.ComputeTargetTCPProxyRef)
		**out = **in
	}
	if in.TargetVPNGatewayRef != nil {
		in, out := &in.TargetVPNGatewayRef, &out.TargetVPNGatewayRef
		*out = new(refsv1beta1.ComputeTargetVPNGatewayRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingruleTarget.
func (in *ForwardingruleTarget) DeepCopy() *ForwardingruleTarget {
	if in == nil {
		return nil
	}
	out := new(ForwardingruleTarget)
	in.DeepCopyInto(out)
	return out
}
