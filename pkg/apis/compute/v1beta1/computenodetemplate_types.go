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

type NodetemplateNodeTypeFlexibility struct {
	/* Immutable. Number of virtual CPUs to use. */
	// +optional
	Cpus *string `json:"cpus,omitempty"`

	/* Use local SSD */
	// +optional
	LocalSsd *string `json:"localSsd,omitempty"`

	/* Immutable. Physical memory available to the node, defined in MB. */
	// +optional
	Memory *string `json:"memory,omitempty"`
}

type NodetemplateServerBinding struct {
	/* Immutable. Type of server binding policy. If 'RESTART_NODE_ON_ANY_SERVER',
	nodes using this template will restart on any physical server
	following a maintenance event.

	If 'RESTART_NODE_ON_MINIMAL_SERVER', nodes using this template
	will restart on the same physical server following a maintenance
	event, instead of being live migrated to or restarted on a new
	physical server. This option may be useful if you are using
	software licenses tied to the underlying server characteristics
	such as physical sockets or cores, to avoid the need for
	additional licenses when maintenance occurs. However, VMs on such
	nodes will experience outages while maintenance is applied. Possible values: ["RESTART_NODE_ON_ANY_SERVER", "RESTART_NODE_ON_MINIMAL_SERVERS"] */
	Type string `json:"type"`
}

type ComputeNodeTemplateSpec struct {
	/* Immutable. CPU overcommit. Default value: "NONE" Possible values: ["ENABLED", "NONE"] */
	// +optional
	CpuOvercommitType *string `json:"cpuOvercommitType,omitempty"`

	/* Immutable. An optional textual description of the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Node type to use for nodes group that are created from this template.
	Only one of nodeTypeFlexibility and nodeType can be specified. */
	// +optional
	NodeType *string `json:"nodeType,omitempty"`

	/* Immutable. Flexible properties for the desired node type. Node groups that
	use this node template will create nodes of a type that matches
	these properties. Only one of nodeTypeFlexibility and nodeType can
	be specified. */
	// +optional
	NodeTypeFlexibility *NodetemplateNodeTypeFlexibility `json:"nodeTypeFlexibility,omitempty"`

	/* Immutable. Region where nodes using the node template will be created.
	If it is not provided, the provider region is used. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The server binding policy for nodes using this template. Determines
	where the nodes should restart following a maintenance event. */
	// +optional
	ServerBinding *NodetemplateServerBinding `json:"serverBinding,omitempty"`
}

type ComputeNodeTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeNodeTemplate's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/*  */
	SelfLink string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeNodeTemplate is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeNodeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNodeTemplateSpec   `json:"spec,omitempty"`
	Status ComputeNodeTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeNodeTemplateList contains a list of ComputeNodeTemplate
type ComputeNodeTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNodeTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNodeTemplate{}, &ComputeNodeTemplateList{})
}
