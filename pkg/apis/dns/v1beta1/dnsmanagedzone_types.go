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

type ManagedzoneDefaultKeySpecs struct {
	/* String mnemonic specifying the DNSSEC algorithm of this key Possible values: ["ecdsap256sha256", "ecdsap384sha384", "rsasha1", "rsasha256", "rsasha512"] */
	// +optional
	Algorithm *string `json:"algorithm,omitempty"`

	/* Length of the keys in bits */
	// +optional
	KeyLength *int `json:"keyLength,omitempty"`

	/* Specifies whether this is a key signing key (KSK) or a zone
	signing key (ZSK). Key signing keys have the Secure Entry
	Point flag set and, when active, will only be used to sign
	resource record sets of type DNSKEY. Zone signing keys do
	not have the Secure Entry Point flag set and will be used
	to sign all other types of resource record sets. Possible values: ["keySigning", "zoneSigning"] */
	// +optional
	KeyType *string `json:"keyType,omitempty"`

	/* Identifies what kind of resource this is */
	// +optional
	Kind *string `json:"kind,omitempty"`
}

type ManagedzoneDnssecConfig struct {
	/* Specifies parameters that will be used for generating initial DnsKeys
	for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
	you must also provide one for the other.
	default_key_specs can only be updated when the state is 'off'. */
	// +optional
	DefaultKeySpecs []ManagedzoneDefaultKeySpecs `json:"defaultKeySpecs,omitempty"`

	/* Identifies what kind of resource this is */
	// +optional
	Kind *string `json:"kind,omitempty"`

	/* Specifies the mechanism used to provide authenticated denial-of-existence responses.
	non_existence can only be updated when the state is 'off'. Possible values: ["nsec", "nsec3"] */
	// +optional
	NonExistence *string `json:"nonExistence,omitempty"`

	/* Specifies whether DNSSEC is enabled, and what mode it is in Possible values: ["off", "on", "transfer"] */
	// +optional
	State *string `json:"state,omitempty"`
}

type ManagedzoneForwardingConfig struct {
	/* List of target name servers to forward to. Cloud DNS will
	select the best available name server if more than
	one target is given. */
	TargetNameServers []ManagedzoneTargetNameServers `json:"targetNameServers"`
}

type ManagedzoneNamespace struct {
	/* The fully qualified or partial URL of the service directory namespace that should be
	associated with the zone. This should be formatted like
	'https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace_id}'
	or simply 'projects/{project}/locations/{location}/namespaces/{namespace_id}'
	Ignored for 'public' visibility zones. */
	NamespaceUrl string `json:"namespaceUrl"`
}

type ManagedzoneNetworks struct {
	/* VPC network to bind to. */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef"`
}

type ManagedzonePeeringConfig struct {
	/* The network with which to peer. */
	TargetNetwork ManagedzoneTargetNetwork `json:"targetNetwork"`
}

type ManagedzonePrivateVisibilityConfig struct {
	/*  */
	Networks []ManagedzoneNetworks `json:"networks"`
}

type ManagedzoneServiceDirectoryConfig struct {
	/* The namespace associated with the zone. */
	Namespace ManagedzoneNamespace `json:"namespace"`
}

type ManagedzoneTargetNameServers struct {
	/* Forwarding path for this TargetNameServer. If unset or 'default' Cloud DNS will make forwarding
	decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
	to the Internet. When set to 'private', Cloud DNS will always send queries through VPC for this target Possible values: ["default", "private"] */
	// +optional
	ForwardingPath *string `json:"forwardingPath,omitempty"`

	/* IPv4 address of a target name server. */
	Ipv4Address string `json:"ipv4Address"`
}

type ManagedzoneTargetNetwork struct {
	/* VPC network to forward queries to. */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef"`
}

type DNSManagedZoneSpec struct {
	/* A textual description field. Defaults to 'Managed by Config Connector'. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The DNS name of this managed zone, for instance "example.com.". */
	DnsName string `json:"dnsName"`

	/* DNSSEC configuration */
	// +optional
	DnssecConfig *ManagedzoneDnssecConfig `json:"dnssecConfig,omitempty"`

	/* The presence for this field indicates that outbound forwarding is enabled
	for this zone. The value of this field contains the set of destinations
	to forward to. */
	// +optional
	ForwardingConfig *ManagedzoneForwardingConfig `json:"forwardingConfig,omitempty"`

	/* The presence of this field indicates that DNS Peering is enabled for this
	zone. The value of this field contains the network to peer with. */
	// +optional
	PeeringConfig *ManagedzonePeeringConfig `json:"peeringConfig,omitempty"`

	/* For privately visible zones, the set of Virtual Private Cloud
	resources that the zone is visible from. */
	// +optional
	PrivateVisibilityConfig *ManagedzonePrivateVisibilityConfig `json:"privateVisibilityConfig,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	lookup queries using automatically configured records for VPC resources. This only applies
	to networks listed under 'private_visibility_config'. */
	// +optional
	ReverseLookup *bool `json:"reverseLookup,omitempty"`

	/* Immutable. The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone. */
	// +optional
	ServiceDirectoryConfig *ManagedzoneServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`

	/* Immutable. The zone's visibility: public zones are exposed to the Internet,
	while private zones are visible only to Virtual Private Cloud resources. Default value: "public" Possible values: ["private", "public"] */
	// +optional
	Visibility *string `json:"visibility,omitempty"`
}

type DNSManagedZoneStatus struct {
	/* Conditions represent the latest available observations of the
	   DNSManagedZone's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Delegate your managed_zone to these virtual name servers;
	defined by the server */
	NameServers []string `json:"nameServers,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSManagedZone is the Schema for the dns API
// +k8s:openapi-gen=true
type DNSManagedZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSManagedZoneSpec   `json:"spec,omitempty"`
	Status DNSManagedZoneStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSManagedZoneList contains a list of DNSManagedZone
type DNSManagedZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSManagedZone `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DNSManagedZone{}, &DNSManagedZoneList{})
}
