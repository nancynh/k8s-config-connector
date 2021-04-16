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

type PolicyAuditConfigs struct {
	/* Required. The configuration for logging of each type of permission. */
	AuditLogConfigs []PolicyAuditLogConfigs `json:"auditLogConfigs"`

	/* Required. The service for which to enable Data Access audit logs. The special value 'allServices' covers all services. Note that if there are audit configs covering both 'allServices' and a specific service, then the union of the two audit configs is used for that service: the 'logTypes' specified in each 'auditLogConfig' are enabled, and the 'exemptedMembers' in each 'auditLogConfg' are exempted. */
	Service string `json:"service"`
}

type PolicyAuditLogConfigs struct {
	/* Identities that do not cause logging for this type of permission. The format is the same as that for 'members' in IAMPolicy/IAMPolicyMember. */
	// +optional
	ExemptedMembers []string `json:"exemptedMembers,omitempty"`

	/* Permission type for which logging is to be configured. Must be one of 'DATA_READ', 'DATA_WRITE', or 'ADMIN_READ'. */
	LogType string `json:"logType"`
}

type PolicyBindings struct {
	/* Optional. The condition under which the binding applies. */
	// +optional
	Condition *PolicyCondition `json:"condition,omitempty"`

	/* Optional. The list of IAM users to be bound to the role. */
	// +optional
	Members []string `json:"members,omitempty"`

	/* Required. The role to bind the users to. */
	Role string `json:"role"`
}

type PolicyCondition struct {
	/*  */
	// +optional
	Description *string `json:"description,omitempty"`

	/*  */
	Expression string `json:"expression"`

	/*  */
	Title string `json:"title"`
}

type IAMPolicySpec struct {
	/* Optional. The list of IAM audit configs. */
	// +optional
	AuditConfigs []PolicyAuditConfigs `json:"auditConfigs,omitempty"`

	/* Optional. The list of IAM bindings. */
	// +optional
	Bindings []PolicyBindings `json:"bindings,omitempty"`

	/* Immutable. Required. The GCP resource to set the IAM policy on. */
	ResourceRef v1alpha1.ResourceRef `json:"resourceRef"`
}

type IAMPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   IAMPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMPolicy is the Schema for the iam API
// +k8s:openapi-gen=true
type IAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMPolicySpec   `json:"spec,omitempty"`
	Status IAMPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMPolicyList contains a list of IAMPolicy
type IAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAMPolicy{}, &IAMPolicyList{})
}
