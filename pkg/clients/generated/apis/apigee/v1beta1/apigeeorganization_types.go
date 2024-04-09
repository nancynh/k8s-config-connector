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

type OrganizationAddonsConfig struct {
	/* Configuration for the Advanced API Ops add-on. */
	// +optional
	AdvancedApiOpsConfig *OrganizationAdvancedApiOpsConfig `json:"advancedApiOpsConfig,omitempty"`

	/* Configuration for the Monetization add-on. */
	// +optional
	MonetizationConfig *OrganizationMonetizationConfig `json:"monetizationConfig,omitempty"`
}

type OrganizationAdvancedApiOpsConfig struct {
	/* Flag that specifies whether the Advanced API Ops add-on is enabled. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
}

type OrganizationMonetizationConfig struct {
	/* Flag that specifies whether the Monetization add-on is enabled. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
}

type ApigeeOrganizationSpec struct {
	/* Addon configurations of the Apigee organization. */
	// +optional
	AddonsConfig *OrganizationAddonsConfig `json:"addonsConfig,omitempty"`

	/* Immutable. Required. Primary GCP region for analytics data storage. For valid values, see (https://cloud.google.com/apigee/docs/api-platform/get-started/create-org). */
	AnalyticsRegion string `json:"analyticsRegion"`

	// +optional
	AuthorizedNetworkRef *v1alpha1.ResourceRef `json:"authorizedNetworkRef,omitempty"`

	/* Description of the Apigee organization. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Display name for the Apigee organization. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Properties defined in the Apigee organization profile. */
	// +optional
	Properties map[string]string `json:"properties,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// +optional
	RuntimeDatabaseEncryptionKeyRef *v1alpha1.ResourceRef `json:"runtimeDatabaseEncryptionKeyRef,omitempty"`

	/* Immutable. Required. Runtime type of the Apigee organization based on the Apigee subscription purchased. Possible values: RUNTIME_TYPE_UNSPECIFIED, CLOUD, HYBRID */
	RuntimeType string `json:"runtimeType"`
}

type ApigeeOrganizationStatus struct {
	/* Conditions represent the latest available observations of the
	   ApigeeOrganization's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. Billing type of the Apigee organization. See (https://cloud.google.com/apigee/pricing). Possible values: BILLING_TYPE_UNSPECIFIED, SUBSCRIPTION, EVALUATION */
	// +optional
	BillingType *string `json:"billingType,omitempty"`

	/* Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when (#RuntimeType) is `CLOUD`. */
	// +optional
	CaCertificate *string `json:"caCertificate,omitempty"`

	/* Output only. Time that the Apigee organization was created in milliseconds since epoch. */
	// +optional
	CreatedAt *int `json:"createdAt,omitempty"`

	/* Output only. List of environments in the Apigee organization. */
	// +optional
	Environments []string `json:"environments,omitempty"`

	/* Output only. Time that the Apigee organization is scheduled for deletion. */
	// +optional
	ExpiresAt *int `json:"expiresAt,omitempty"`

	/* Output only. Time that the Apigee organization was last modified in milliseconds since epoch. */
	// +optional
	LastModifiedAt *int `json:"lastModifiedAt,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Output only. Project ID associated with the Apigee organization. */
	// +optional
	ProjectId *string `json:"projectId,omitempty"`

	/* Output only. State of the organization. Values other than ACTIVE means the resource is not ready to use. Possible values: SNAPSHOT_STATE_UNSPECIFIED, MISSING, OK_DOCSTORE, OK_SUBMITTED, OK_EXTERNAL, DELETED */
	// +optional
	State *string `json:"state,omitempty"`

	/* Output only. DEPRECATED: This will eventually be replaced by BillingType. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation purposes only) or paid (full subscription has been purchased). See (https://cloud.google.com/apigee/pricing/). Possible values: SUBSCRIPTION_TYPE_UNSPECIFIED, PAID, TRIAL */
	// +optional
	SubscriptionType *string `json:"subscriptionType,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeorganization;gcpapigeeorganizations
// +kubebuilder:subresource:status

// ApigeeOrganization is the Schema for the apigee API
// +k8s:openapi-gen=true
type ApigeeOrganization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeOrganizationSpec   `json:"spec,omitempty"`
	Status ApigeeOrganizationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApigeeOrganizationList contains a list of ApigeeOrganization
type ApigeeOrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeOrganization `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeOrganization{}, &ApigeeOrganizationList{})
}
