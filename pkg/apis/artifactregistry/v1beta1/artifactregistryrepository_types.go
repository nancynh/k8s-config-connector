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

type ArtifactRegistryRepositorySpec struct {
	/* The user-provided description of the repository. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The format of packages that are stored in the repository. */
	Format string `json:"format"`

	/* The customer managed encryption key that’s used to encrypt the
	contents of the Repository. */
	// +optional
	KmsKeyRef *v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`

	/* Immutable. The name of the location this repository is located in. */
	Location string `json:"location"`

	/* Immutable. Optional. The repositoryId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ArtifactRegistryRepositoryStatus struct {
	/* Conditions represent the latest available observations of the
	   ArtifactRegistryRepository's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The time when the repository was created. */
	CreateTime string `json:"createTime,omitempty"`
	/* The name of the repository, for example:
	"projects/p1/locations/us-central1/repositories/repo1" */
	Name string `json:"name,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* The time when the repository was last updated. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArtifactRegistryRepository is the Schema for the artifactregistry API
// +k8s:openapi-gen=true
type ArtifactRegistryRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArtifactRegistryRepositorySpec   `json:"spec,omitempty"`
	Status ArtifactRegistryRepositoryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArtifactRegistryRepositoryList contains a list of ArtifactRegistryRepository
type ArtifactRegistryRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArtifactRegistryRepository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArtifactRegistryRepository{}, &ArtifactRegistryRepositoryList{})
}
