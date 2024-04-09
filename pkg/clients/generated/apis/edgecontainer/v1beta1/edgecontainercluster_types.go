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

type ClusterAdminUsers struct {
	UsernameRef v1alpha1.ResourceRef `json:"usernameRef"`
}

type ClusterAuthorization struct {
	/* User that will be granted the cluster-admin role on the cluster, providing
	full access to the cluster. Currently, this is a singular field, but will
	be expanded to allow multiple admins in the future. */
	AdminUsers ClusterAdminUsers `json:"adminUsers"`
}

type ClusterControlPlane struct {
	/* Immutable. Local control plane configuration. */
	// +optional
	Local *ClusterLocal `json:"local,omitempty"`

	/* Immutable. Remote control plane configuration. */
	// +optional
	Remote *ClusterRemote `json:"remote,omitempty"`
}

type ClusterControlPlaneEncryption struct {
	/* The Cloud KMS CryptoKeyVersion currently in use for protecting control
	plane disks. Only applicable if kms_key is set. */
	// +optional
	KmsKeyActiveVersion *string `json:"kmsKeyActiveVersion,omitempty"`

	// +optional
	KmsKeyRef *v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`

	/* Availability of the Cloud KMS CryptoKey. If not 'KEY_AVAILABLE', then
	nodes may go offline as they cannot access their local data. This can be
	caused by a lack of permissions to use the key, or if the key is disabled
	or deleted. */
	// +optional
	KmsKeyState *string `json:"kmsKeyState,omitempty"`

	/* Error status returned by Cloud KMS when using this key. This field may be
	populated only if 'kms_key_state' is not 'KMS_KEY_STATE_KEY_AVAILABLE'.
	If populated, this field contains the error status reported by Cloud KMS. */
	// +optional
	KmsStatus []ClusterKmsStatus `json:"kmsStatus,omitempty"`
}

type ClusterFleet struct {
	/* The name of the managed Hub Membership resource associated to this cluster.
	Membership names are formatted as
	'projects/<project-number>/locations/global/membership/<cluster-id>'. */
	// +optional
	Membership *string `json:"membership,omitempty"`

	/* The number of the Fleet host project where this cluster will be registered. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`
}

type ClusterIngress struct {
	/* Whether Ingress is disabled. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	/* Ingress VIP. */
	// +optional
	Ipv4Vip *string `json:"ipv4Vip,omitempty"`
}

type ClusterKmsStatus struct {
	/* The status code, which should be an enum value of google.rpc.Code. */
	// +optional
	Code *int `json:"code,omitempty"`

	/* A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client. */
	// +optional
	Message *string `json:"message,omitempty"`
}

type ClusterLocal struct {
	/* Only machines matching this filter will be allowed to host control
	plane nodes. The filtering language accepts strings like "name=<name>",
	and is documented here: [AIP-160](https://google.aip.dev/160). */
	// +optional
	MachineFilter *string `json:"machineFilter,omitempty"`

	/* The number of nodes to serve as replicas of the Control Plane.
	Only 1 and 3 are supported. */
	// +optional
	NodeCount *int `json:"nodeCount,omitempty"`

	/* Immutable. Name of the Google Distributed Cloud Edge zones where this node pool
	will be created. For example: 'us-central1-edge-customer-a'. */
	// +optional
	NodeLocation *string `json:"nodeLocation,omitempty"`

	/* Policy configuration about how user applications are deployed. Possible values: ["SHARED_DEPLOYMENT_POLICY_UNSPECIFIED", "ALLOWED", "DISALLOWED"]. */
	// +optional
	SharedDeploymentPolicy *string `json:"sharedDeploymentPolicy,omitempty"`
}

type ClusterMaintenancePolicy struct {
	/* Specifies the maintenance window in which maintenance may be performed. */
	Window ClusterWindow `json:"window"`
}

type ClusterNetworking struct {
	/* Immutable. All pods in the cluster are assigned an RFC1918 IPv4 address from these
	blocks. Only a single block is supported. This field cannot be changed
	after creation. */
	ClusterIpv4CidrBlocks []string `json:"clusterIpv4CidrBlocks"`

	/* Immutable. If specified, dual stack mode is enabled and all pods in the cluster are
	assigned an IPv6 address from these blocks alongside from an IPv4
	address. Only a single block is supported. This field cannot be changed
	after creation. */
	// +optional
	ClusterIpv6CidrBlocks []string `json:"clusterIpv6CidrBlocks,omitempty"`

	/* IP addressing type of this cluster i.e. SINGLESTACK_V4 vs DUALSTACK_V4_V6. */
	// +optional
	NetworkType *string `json:"networkType,omitempty"`

	/* Immutable. All services in the cluster are assigned an RFC1918 IPv4 address from these
	blocks. Only a single block is supported. This field cannot be changed
	after creation. */
	ServicesIpv4CidrBlocks []string `json:"servicesIpv4CidrBlocks"`

	/* Immutable. If specified, dual stack mode is enabled and all services in the cluster are
	assigned an IPv6 address from these blocks alongside from an IPv4
	address. Only a single block is supported. This field cannot be changed
	after creation. */
	// +optional
	ServicesIpv6CidrBlocks []string `json:"servicesIpv6CidrBlocks,omitempty"`
}

type ClusterRecurringWindow struct {
	/* An RRULE (https://tools.ietf.org/html/rfc5545#section-3.8.5.3) for how
	this window recurs. They go on for the span of time between the start and
	end time. */
	// +optional
	Recurrence *string `json:"recurrence,omitempty"`

	/* Represents an arbitrary window of time. */
	// +optional
	Window *ClusterWindow `json:"window,omitempty"`
}

type ClusterRemote struct {
	/* Immutable. Name of the Google Distributed Cloud Edge zones where this node pool
	will be created. For example: 'us-central1-edge-customer-a'. */
	// +optional
	NodeLocation *string `json:"nodeLocation,omitempty"`
}

type ClusterSystemAddonsConfig struct {
	/* Config for the Ingress add-on which allows customers to create an Ingress
	object to manage external access to the servers in a cluster. The add-on
	consists of istiod and istio-ingress. */
	// +optional
	Ingress *ClusterIngress `json:"ingress,omitempty"`
}

type ClusterWindow struct {
	/* The time that the window ends. The end time must take place after the
	start time. */
	// +optional
	EndTime *string `json:"endTime,omitempty"`

	/* The time that the window first starts. */
	// +optional
	StartTime *string `json:"startTime,omitempty"`
}

type EdgeContainerClusterSpec struct {
	/* Immutable. RBAC policy that will be applied and managed by GEC. */
	Authorization ClusterAuthorization `json:"authorization"`

	/* The configuration of the cluster control plane. */
	// +optional
	ControlPlane *ClusterControlPlane `json:"controlPlane,omitempty"`

	/* Remote control plane disk encryption options. This field is only used when
	enabling CMEK support. */
	// +optional
	ControlPlaneEncryption *ClusterControlPlaneEncryption `json:"controlPlaneEncryption,omitempty"`

	/* The default maximum number of pods per node used if a maximum value is not
	specified explicitly for a node pool in this cluster. If unspecified, the
	Kubernetes default value will be used. */
	// +optional
	DefaultMaxPodsPerNode *int `json:"defaultMaxPodsPerNode,omitempty"`

	/* Address pools for cluster data plane external load balancing. */
	// +optional
	ExternalLoadBalancerIpv4AddressPools []string `json:"externalLoadBalancerIpv4AddressPools,omitempty"`

	/* Immutable. Fleet related configuration.
	Fleets are a Google Cloud concept for logically organizing clusters,
	letting you use and manage multi-cluster capabilities and apply
	consistent policies across your systems. */
	Fleet ClusterFleet `json:"fleet"`

	/* Immutable. The location of the resource. */
	Location string `json:"location"`

	/* Cluster-wide maintenance policy configuration. */
	// +optional
	MaintenancePolicy *ClusterMaintenancePolicy `json:"maintenancePolicy,omitempty"`

	/* Fleet related configuration.
	Fleets are a Google Cloud concept for logically organizing clusters,
	letting you use and manage multi-cluster capabilities and apply
	consistent policies across your systems. */
	Networking ClusterNetworking `json:"networking"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* The release channel a cluster is subscribed to. Possible values: ["RELEASE_CHANNEL_UNSPECIFIED", "NONE", "REGULAR"]. */
	// +optional
	ReleaseChannel *string `json:"releaseChannel,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Config that customers are allowed to define for GDCE system add-ons. */
	// +optional
	SystemAddonsConfig *ClusterSystemAddonsConfig `json:"systemAddonsConfig,omitempty"`

	/* The target cluster version. For example: "1.5.0". */
	// +optional
	TargetVersion *string `json:"targetVersion,omitempty"`
}

type ClusterMaintenanceEventsStatus struct {
	/* The time when the maintenance event request was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* The time when the maintenance event ended, either successfully or not. If
	the maintenance event is split into multiple maintenance windows,
	end_time is only updated when the whole flow ends. */
	// +optional
	EndTime *string `json:"endTime,omitempty"`

	/* The operation for running the maintenance event. Specified in the format
	projects/* /locations/* /operations/*. If the maintenance event is split
	into multiple operations (e.g. due to maintenance windows), the latest
	one is recorded. */
	// +optional
	Operation *string `json:"operation,omitempty"`

	/* The schedule of the maintenance event. */
	// +optional
	Schedule *string `json:"schedule,omitempty"`

	/* The time when the maintenance event started. */
	// +optional
	StartTime *string `json:"startTime,omitempty"`

	/* Indicates the maintenance event state. */
	// +optional
	State *string `json:"state,omitempty"`

	/* The target version of the cluster. */
	// +optional
	TargetVersion *string `json:"targetVersion,omitempty"`

	/* Indicates the maintenance event type. */
	// +optional
	Type *string `json:"type,omitempty"`

	/* The time when the maintenance event message was updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	/* UUID of the maintenance event. */
	// +optional
	Uuid *string `json:"uuid,omitempty"`
}

type EdgeContainerClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   EdgeContainerCluster's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The PEM-encoded public certificate of the cluster's CA. */
	// +optional
	ClusterCaCertificate *string `json:"clusterCaCertificate,omitempty"`

	/* The control plane release version. */
	// +optional
	ControlPlaneVersion *string `json:"controlPlaneVersion,omitempty"`

	/* The time the cluster was created, in RFC3339 text format. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* The IP address of the Kubernetes API server. */
	// +optional
	Endpoint *string `json:"endpoint,omitempty"`

	/* All the maintenance events scheduled for the cluster, including the ones
	ongoing, planned for the future and done in the past (up to 90 days). */
	// +optional
	MaintenanceEvents []ClusterMaintenanceEventsStatus `json:"maintenanceEvents,omitempty"`

	/* The lowest release version among all worker nodes. This field can be empty
	if the cluster does not have any worker nodes. */
	// +optional
	NodeVersion *string `json:"nodeVersion,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The port number of the Kubernetes API server. */
	// +optional
	Port *int `json:"port,omitempty"`

	/* Indicates the status of the cluster. */
	// +optional
	Status *string `json:"status,omitempty"`

	/* The time the cluster was last updated, in RFC3339 text format. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpedgecontainercluster;gcpedgecontainerclusters
// +kubebuilder:subresource:status

// EdgeContainerCluster is the Schema for the edgecontainer API
// +k8s:openapi-gen=true
type EdgeContainerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeContainerClusterSpec   `json:"spec,omitempty"`
	Status EdgeContainerClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EdgeContainerClusterList contains a list of EdgeContainerCluster
type EdgeContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeContainerCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EdgeContainerCluster{}, &EdgeContainerClusterList{})
}
