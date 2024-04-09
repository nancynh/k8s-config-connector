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

type GroupGroupKey struct {
	/* Immutable. The ID of the entity.

	For Google-managed entities, the id must be the email address of an existing
	group or user.

	For external-identity-mapped entities, the id must be a string conforming
	to the Identity Source's requirements.

	Must be unique within a namespace. */
	Id string `json:"id"`

	/* Immutable. The namespace in which the entity exists.

	If not specified, the EntityKey represents a Google-managed entity
	such as a Google user or a Google Group.

	If specified, the EntityKey represents an external-identity-mapped group.
	The namespace must correspond to an identity source created in Admin Console
	and must be in the form of 'identitysources/{identity_source_id}'. */
	// +optional
	Namespace *string `json:"namespace,omitempty"`
}

type CloudIdentityGroupSpec struct {
	/* An extended description to help users determine the purpose of a Group.
	Must not be longer than 4,096 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The display name of the Group. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Immutable. EntityKey of the Group. */
	GroupKey GroupGroupKey `json:"groupKey"`

	/* Immutable. The initial configuration options for creating a Group.

	See the
	[API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
	for possible values. Default value: "EMPTY" Possible values: ["INITIAL_GROUP_CONFIG_UNSPECIFIED", "WITH_INITIAL_OWNER", "EMPTY"]. */
	// +optional
	InitialGroupConfig *string `json:"initialGroupConfig,omitempty"`

	/* One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value.

	Google Groups are the default type of group and have a label with a key of cloudidentity.googleapis.com/groups.discussion_forum and an empty value.

	Existing Google Groups can have an additional label with a key of cloudidentity.googleapis.com/groups.security and an empty value added to them. This is an immutable change and the security label cannot be removed once added.

	Dynamic groups have a label with a key of cloudidentity.googleapis.com/groups.dynamic.

	Identity-mapped groups for Cloud Search have a label with a key of system/groups/external and an empty value. */
	Labels map[string]string `json:"labels"`

	/* Immutable. The resource name of the entity under which this Group resides in the
	Cloud Identity resource hierarchy.

	Must be of the form identitysources/{identity_source_id} for external-identity-mapped
	groups or customers/{customer_id} for Google Groups. */
	Parent string `json:"parent"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type CloudIdentityGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   CloudIdentityGroup's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The time when the Group was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Resource name of the Group in the format: groups/{group_id}, where group_id
	is the unique ID assigned to the Group. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The time when the Group was last updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudidentitygroup;gcpcloudidentitygroups
// +kubebuilder:subresource:status

// CloudIdentityGroup is the Schema for the cloudidentity API
// +k8s:openapi-gen=true
type CloudIdentityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudIdentityGroupSpec   `json:"spec,omitempty"`
	Status CloudIdentityGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudIdentityGroupList contains a list of CloudIdentityGroup
type CloudIdentityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudIdentityGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudIdentityGroup{}, &CloudIdentityGroupList{})
}
