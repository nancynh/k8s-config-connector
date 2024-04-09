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

type CryptokeyVersionTemplate struct {
	/* The algorithm to use when creating a version based on this template.
	See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs. */
	Algorithm string `json:"algorithm"`

	/* Immutable. The protection level to use when creating a version based on this template. Possible values include "SOFTWARE", "HSM", "EXTERNAL", "EXTERNAL_VPC". Defaults to "SOFTWARE". */
	// +optional
	ProtectionLevel *string `json:"protectionLevel,omitempty"`
}

type KMSCryptoKeySpec struct {
	/* Immutable. The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
	If not specified at creation time, the default duration is 24 hours. */
	// +optional
	DestroyScheduledDuration *string `json:"destroyScheduledDuration,omitempty"`

	/* Immutable. Whether this key may contain imported versions only. */
	// +optional
	ImportOnly *bool `json:"importOnly,omitempty"`

	/* The KMSKeyRing that this key belongs to. */
	KeyRingRef v1alpha1.ResourceRef `json:"keyRingRef"`

	/* Immutable. The immutable purpose of this CryptoKey. See the
	[purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
	for possible inputs.
	Default value is "ENCRYPT_DECRYPT". */
	// +optional
	Purpose *string `json:"purpose,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	The first rotation will take place after the specified period. The rotation period has
	the format of a decimal number with up to 9 fractional digits, followed by the
	letter 's' (seconds). It must be greater than a day (ie, 86400). */
	// +optional
	RotationPeriod *string `json:"rotationPeriod,omitempty"`

	/* Immutable. If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	You must use the 'google_kms_key_ring_import_job' resource to import the CryptoKeyVersion. */
	// +optional
	SkipInitialVersionCreation *bool `json:"skipInitialVersionCreation,omitempty"`

	/* A template describing settings for new crypto key versions. */
	// +optional
	VersionTemplate *CryptokeyVersionTemplate `json:"versionTemplate,omitempty"`
}

type KMSCryptoKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   KMSCryptoKey's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The self link of the created key in the format projects/{project}/locations/{location}/keyRings/{keyRingName}/cryptoKeys/{name}. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpkmscryptokey;gcpkmscryptokeys
// +kubebuilder:subresource:status

// KMSCryptoKey is the Schema for the kms API
// +k8s:openapi-gen=true
type KMSCryptoKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSCryptoKeySpec   `json:"spec,omitempty"`
	Status KMSCryptoKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KMSCryptoKeyList contains a list of KMSCryptoKey
type KMSCryptoKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KMSCryptoKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KMSCryptoKey{}, &KMSCryptoKeyList{})
}
