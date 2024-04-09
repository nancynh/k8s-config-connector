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

type ClusterAuthorization struct {
	/* Users that can perform operations as a cluster admin. A managed
	ClusterRoleBinding will be created to grant the 'cluster-admin' ClusterRole
	to the users. Up to ten admin users can be provided.

	For more info on RBAC, see
	https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles. */
	// +optional
	AdminUsers []string `json:"adminUsers,omitempty"`
}

type ClusterBinaryAuthorization struct {
	/* Configure Binary Authorization evaluation mode. Possible values: ["DISABLED", "PROJECT_SINGLETON_POLICY_ENFORCE"]. */
	// +optional
	EvaluationMode *string `json:"evaluationMode,omitempty"`
}

type ClusterComponentConfig struct {
	/* The components to be enabled. Possible values: ["SYSTEM_COMPONENTS", "WORKLOADS"]. */
	// +optional
	EnableComponents []string `json:"enableComponents,omitempty"`
}

type ClusterFleet struct {
	/* The name of the managed Hub Membership resource associated to this
	cluster. Membership names are formatted as
	projects/<project-number>/locations/global/membership/<cluster-id>. */
	// +optional
	Membership *string `json:"membership,omitempty"`

	/* The number of the Fleet host project where this cluster will be registered. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`
}

type ClusterLoggingConfig struct {
	/* The configuration of the logging components. */
	// +optional
	ComponentConfig *ClusterComponentConfig `json:"componentConfig,omitempty"`
}

type ClusterManagedPrometheusConfig struct {
	/* Enable Managed Collection. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
}

type ClusterMonitoringConfig struct {
	/* Enable Google Cloud Managed Service for Prometheus in the cluster. */
	// +optional
	ManagedPrometheusConfig *ClusterManagedPrometheusConfig `json:"managedPrometheusConfig,omitempty"`
}

type ClusterOidcConfig struct {
	/* Immutable. A JSON Web Token (JWT) issuer URI. 'issuer' must start with 'https://'. */
	IssuerUrl string `json:"issuerUrl"`

	/* Immutable. OIDC verification keys in JWKS format (RFC 7517). */
	// +optional
	Jwks *string `json:"jwks,omitempty"`
}

type ContainerAttachedClusterSpec struct {
	/* Optional. Annotations on the cluster. This field has the same
	restrictions as Kubernetes annotations. The total size of all keys and
	values combined is limited to 256k. Key can have 2 segments: prefix (optional)
	and name (required), separated by a slash (/). Prefix must be a DNS subdomain.
	Name must be 63 characters or less, begin and end with alphanumerics,
	with dashes (-), underscores (_), dots (.), and alphanumerics between. */
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	/* Configuration related to the cluster RBAC settings. */
	// +optional
	Authorization *ClusterAuthorization `json:"authorization,omitempty"`

	/* Binary Authorization configuration. */
	// +optional
	BinaryAuthorization *ClusterBinaryAuthorization `json:"binaryAuthorization,omitempty"`

	/* Policy to determine what flags to send on delete. */
	// +optional
	DeletionPolicy *string `json:"deletionPolicy,omitempty"`

	/* A human readable description of this attached cluster. Cannot be longer
	than 255 UTF-8 encoded bytes. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The Kubernetes distribution of the underlying attached cluster. Supported values:
	"eks", "aks". */
	Distribution string `json:"distribution"`

	/* Fleet configuration. */
	Fleet ClusterFleet `json:"fleet"`

	/* Immutable. The location for the resource. */
	Location string `json:"location"`

	/* Logging configuration. */
	// +optional
	LoggingConfig *ClusterLoggingConfig `json:"loggingConfig,omitempty"`

	/* Monitoring configuration. */
	// +optional
	MonitoringConfig *ClusterMonitoringConfig `json:"monitoringConfig,omitempty"`

	/* OIDC discovery information of the target cluster.

	Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
	API server. This fields indicates how GCP services
	validate KSA tokens in order to allow system workloads (such as GKE Connect
	and telemetry agents) to authenticate back to GCP.

	Both clusters with public and private issuer URLs are supported.
	Clusters with public issuers only need to specify the 'issuer_url' field
	while clusters with private issuers need to provide both
	'issuer_url' and 'jwks'. */
	OidcConfig ClusterOidcConfig `json:"oidcConfig"`

	/* The platform version for the cluster (e.g. '1.23.0-gke.1'). */
	PlatformVersion string `json:"platformVersion"`

	/* The ID of the project in which the resource belongs. If it is not provided, the provider project is used. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ClusterErrorsStatus struct {
	/* Human-friendly description of the error. */
	// +optional
	Message *string `json:"message,omitempty"`
}

type ClusterWorkloadIdentityConfigStatus struct {
	/* The ID of the OIDC Identity Provider (IdP) associated to
	the Workload Identity Pool. */
	// +optional
	IdentityProvider *string `json:"identityProvider,omitempty"`

	/* The OIDC issuer URL for this cluster. */
	// +optional
	IssuerUri *string `json:"issuerUri,omitempty"`

	/* The Workload Identity Pool associated to the cluster. */
	// +optional
	WorkloadPool *string `json:"workloadPool,omitempty"`
}

type ContainerAttachedClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   ContainerAttachedCluster's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The region where this cluster runs.

	For EKS clusters, this is an AWS region. For AKS clusters,
	this is an Azure region. */
	// +optional
	ClusterRegion *string `json:"clusterRegion,omitempty"`

	/* Output only. The time at which this cluster was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* A set of errors found in the cluster. */
	// +optional
	Errors []ClusterErrorsStatus `json:"errors,omitempty"`

	/* The Kubernetes version of the cluster. */
	// +optional
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* If set, there are currently changes in flight to the cluster. */
	// +optional
	Reconciling *bool `json:"reconciling,omitempty"`

	/* The current state of the cluster. Possible values:
	STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,
	DEGRADED. */
	// +optional
	State *string `json:"state,omitempty"`

	/* A globally unique identifier for the cluster. */
	// +optional
	Uid *string `json:"uid,omitempty"`

	/* The time at which this cluster was last updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	/* Workload Identity settings. */
	// +optional
	WorkloadIdentityConfig []ClusterWorkloadIdentityConfigStatus `json:"workloadIdentityConfig,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontainerattachedcluster;gcpcontainerattachedclusters
// +kubebuilder:subresource:status

// ContainerAttachedCluster is the Schema for the containerattached API
// +k8s:openapi-gen=true
type ContainerAttachedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerAttachedClusterSpec   `json:"spec,omitempty"`
	Status ContainerAttachedClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerAttachedClusterList contains a list of ContainerAttachedCluster
type ContainerAttachedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerAttachedCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerAttachedCluster{}, &ContainerAttachedClusterList{})
}
