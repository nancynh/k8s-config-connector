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

type FirewallAllow struct {
	/* An optional list of ports to which this rule applies. This field
	is only applicable for UDP or TCP protocol. Each entry must be
	either an integer or a range. If not specified, this rule
	applies to connections through any port.

	Example inputs include: ["22"], ["80","443"], and
	["12345-12349"]. */
	// +optional
	Ports []string `json:"ports,omitempty"`

	/* The IP protocol to which this rule applies. The protocol type is
	required when creating a firewall rule. This value can either be
	one of the following well known protocol strings (tcp, udp,
	icmp, esp, ah, sctp, ipip, all), or the IP protocol number. */
	Protocol string `json:"protocol"`
}

type FirewallDeny struct {
	/* An optional list of ports to which this rule applies. This field
	is only applicable for UDP or TCP protocol. Each entry must be
	either an integer or a range. If not specified, this rule
	applies to connections through any port.

	Example inputs include: ["22"], ["80","443"], and
	["12345-12349"]. */
	// +optional
	Ports []string `json:"ports,omitempty"`

	/* The IP protocol to which this rule applies. The protocol type is
	required when creating a firewall rule. This value can either be
	one of the following well known protocol strings (tcp, udp,
	icmp, esp, ah, sctp, ipip, all), or the IP protocol number. */
	Protocol string `json:"protocol"`
}

type FirewallLogConfig struct {
	/* This field denotes whether to include or exclude metadata for firewall logs. Possible values: ["EXCLUDE_ALL_METADATA", "INCLUDE_ALL_METADATA"] */
	Metadata string `json:"metadata"`
}

type ComputeFirewallSpec struct {
	/* The list of ALLOW rules specified by this firewall. Each rule
	specifies a protocol and port-range tuple that describes a permitted
	connection. */
	// +optional
	Allow []FirewallAllow `json:"allow,omitempty"`

	/* The list of DENY rules specified by this firewall. Each rule specifies
	a protocol and port-range tuple that describes a denied connection. */
	// +optional
	Deny []FirewallDeny `json:"deny,omitempty"`

	/* An optional description of this resource. Provide this property when
	you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* If destination ranges are specified, the firewall will apply only to
	traffic that has destination IP address in these ranges. These ranges
	must be expressed in CIDR format. Only IPv4 is supported. */
	// +optional
	DestinationRanges []string `json:"destinationRanges,omitempty"`

	/* Immutable. Direction of traffic to which this firewall applies; default is
	INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
	destinationRanges; For EGRESS traffic, it is NOT supported to specify
	sourceRanges OR sourceTags. Possible values: ["INGRESS", "EGRESS"] */
	// +optional
	Direction *string `json:"direction,omitempty"`

	/* Denotes whether the firewall rule is disabled, i.e not applied to the
	network it is associated with. When set to true, the firewall rule is
	not enforced and the network behaves as if it did not exist. If this
	is unspecified, the firewall rule will be enabled. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	/* DEPRECATED — Deprecated in favor of log_config. This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported to Stackdriver. */
	// +optional
	EnableLogging *bool `json:"enableLogging,omitempty"`

	/* This field denotes the logging options for a particular firewall rule.
	If defined, logging is enabled, and logs will be exported to Cloud Logging. */
	// +optional
	LogConfig *FirewallLogConfig `json:"logConfig,omitempty"`

	/* The network to attach this firewall to. */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef"`

	/* Priority for this rule. This is an integer between 0 and 65535, both
	inclusive. When not specified, the value assumed is 1000. Relative
	priorities determine precedence of conflicting rules. Lower value of
	priority implies higher precedence (eg, a rule with priority 0 has
	higher precedence than a rule with priority 1). DENY rules take
	precedence over ALLOW rules having equal priority. */
	// +optional
	Priority *int `json:"priority,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* If source ranges are specified, the firewall will apply only to
	traffic that has source IP address in these ranges. These ranges must
	be expressed in CIDR format. One or both of sourceRanges and
	sourceTags may be set. If both properties are set, the firewall will
	apply to traffic that has source IP address within sourceRanges OR the
	source IP that belongs to a tag listed in the sourceTags property. The
	connection does not need to match both properties for the firewall to
	apply. Only IPv4 is supported. */
	// +optional
	SourceRanges []string `json:"sourceRanges,omitempty"`

	/*  */
	// +optional
	SourceServiceAccounts []v1alpha1.ResourceRef `json:"sourceServiceAccounts,omitempty"`

	/* If source tags are specified, the firewall will apply only to traffic
	with source IP that belongs to a tag listed in source tags. Source
	tags cannot be used to control traffic to an instance's external IP
	address. Because tags are associated with an instance, not an IP
	address. One or both of sourceRanges and sourceTags may be set. If
	both properties are set, the firewall will apply to traffic that has
	source IP address within sourceRanges OR the source IP that belongs to
	a tag listed in the sourceTags property. The connection does not need
	to match both properties for the firewall to apply. */
	// +optional
	SourceTags []string `json:"sourceTags,omitempty"`

	/*  */
	// +optional
	TargetServiceAccounts []v1alpha1.ResourceRef `json:"targetServiceAccounts,omitempty"`

	/* A list of instance tags indicating sets of instances located in the
	network that may make network connections as specified in allowed[].
	If no targetTags are specified, the firewall rule applies to all
	instances on the specified network. */
	// +optional
	TargetTags []string `json:"targetTags,omitempty"`
}

type ComputeFirewallStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeFirewall's current state. */
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

// ComputeFirewall is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeFirewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeFirewallSpec   `json:"spec,omitempty"`
	Status ComputeFirewallStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeFirewallList contains a list of ComputeFirewall
type ComputeFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeFirewall `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeFirewall{}, &ComputeFirewallList{})
}
