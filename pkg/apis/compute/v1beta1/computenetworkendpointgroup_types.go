// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ComputeNetworkEndpointGroupSpec struct {
	/* Immutable. The default port used if the port number is not specified in the
	network endpoint. */
	// +optional
	DefaultPort *int `json:"defaultPort,omitempty"`

	/* Immutable. An optional description of this resource. Provide this property when
	you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Location represents the geographical location of the ComputeNetworkEndpointGroup. Specify a zone name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/) */
	Location string `json:"location"`

	/* Immutable. Type of network endpoints in this network endpoint group. Default value: "GCE_VM_IP_PORT" Possible values: ["GCE_VM_IP_PORT"] */
	// +optional
	NetworkEndpointType *string `json:"networkEndpointType,omitempty"`

	/* The network to which all network endpoints in the NEG belong. Uses
	"default" project network if unspecified. */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Optional subnetwork to which all network endpoints in the NEG belong. */
	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
}

type ComputeNetworkEndpointGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeNetworkEndpointGroup's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/*  */
	SelfLink string `json:"selfLink,omitempty"`
	/* Number of network endpoints in the network endpoint group. */
	Size int `json:"size,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeNetworkEndpointGroup is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeNetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkEndpointGroupSpec   `json:"spec,omitempty"`
	Status ComputeNetworkEndpointGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeNetworkEndpointGroupList contains a list of ComputeNetworkEndpointGroup
type ComputeNetworkEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNetworkEndpointGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNetworkEndpointGroup{}, &ComputeNetworkEndpointGroupList{})
}
