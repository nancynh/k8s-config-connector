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

type ApigeeEnvironmentSpec struct {
	/* Immutable. */
	ApigeeOrganizationRef v1alpha1.ResourceRef `json:"apigeeOrganizationRef"`

	/* Optional. Description of the environment. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Optional. Display name for this environment. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Optional. Key-value pairs that may be used for customizing the environment. */
	// +optional
	Properties map[string]string `json:"properties,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ApigeeEnvironmentStatus struct {
	/* Conditions represent the latest available observations of the
	   ApigeeEnvironment's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. Creation time of this environment as milliseconds since epoch. */
	// +optional
	CreatedAt *int `json:"createdAt,omitempty"`

	/* Output only. Last modification time of this environment as milliseconds since epoch. */
	// +optional
	LastModifiedAt *int `json:"lastModifiedAt,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Output only. State of the environment. Values other than ACTIVE means the resource is not ready to use. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING */
	// +optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeenvironment;gcpapigeeenvironments
// +kubebuilder:subresource:status

// ApigeeEnvironment is the Schema for the apigee API
// +k8s:openapi-gen=true
type ApigeeEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeEnvironmentSpec   `json:"spec,omitempty"`
	Status ApigeeEnvironmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApigeeEnvironmentList contains a list of ApigeeEnvironment
type ApigeeEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeEnvironment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeEnvironment{}, &ApigeeEnvironmentList{})
}
