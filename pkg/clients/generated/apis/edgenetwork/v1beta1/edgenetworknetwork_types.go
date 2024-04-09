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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EdgeNetworkNetworkSpec struct {
	/* Immutable. A free-text description of the resource. Max length 1024 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The Google Cloud region to which the target Distributed Cloud Edge zone belongs. */
	Location string `json:"location"`

	/* Immutable. IP (L3) MTU value of the network. Default value is '1500'. Possible values are: '1500', '9000'. */
	// +optional
	Mtu *int `json:"mtu,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The networkId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The name of the target Distributed Cloud Edge zone. */
	Zone string `json:"zone"`
}

type EdgeNetworkNetworkStatus struct {
	/* Conditions represent the latest available observations of the
	   EdgeNetworkNetwork's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The time when the subnet was created.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* The canonical name of this resource, with format
	'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The time when the subnet was last updated.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpedgenetworknetwork;gcpedgenetworknetworks
// +kubebuilder:subresource:status

// EdgeNetworkNetwork is the Schema for the edgenetwork API
// +k8s:openapi-gen=true
type EdgeNetworkNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeNetworkNetworkSpec   `json:"spec,omitempty"`
	Status EdgeNetworkNetworkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EdgeNetworkNetworkList contains a list of EdgeNetworkNetwork
type EdgeNetworkNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeNetworkNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EdgeNetworkNetwork{}, &EdgeNetworkNetworkList{})
}
