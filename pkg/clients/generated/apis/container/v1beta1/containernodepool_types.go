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

type NodepoolAdditionalNodeNetworkConfigs struct {
	/* Immutable. Name of the VPC where the additional interface belongs. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Immutable. Name of the subnetwork where the additional interface belongs. */
	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
}

type NodepoolAdditionalPodNetworkConfigs struct {
	/* Immutable. The maximum number of pods per node which use this pod network. */
	// +optional
	MaxPodsPerNode *int `json:"maxPodsPerNode,omitempty"`

	/* Immutable. The name of the secondary range on the subnet which provides IP address for this pod range. */
	// +optional
	SecondaryPodRange *string `json:"secondaryPodRange,omitempty"`

	/* Immutable. Name of the subnetwork where the additional pod network belongs. */
	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
}

type NodepoolAdvancedMachineFeatures struct {
	/* Immutable. The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed. */
	ThreadsPerCore int `json:"threadsPerCore"`
}

type NodepoolAutoscaling struct {
	/* Location policy specifies the algorithm used when scaling-up the node pool. "BALANCED" - Is a best effort policy that aims to balance the sizes of available zones. "ANY" - Instructs the cluster autoscaler to prioritize utilization of unused reservations, and reduces preemption risk for Spot VMs. */
	// +optional
	LocationPolicy *string `json:"locationPolicy,omitempty"`

	/* Maximum number of nodes per zone in the node pool. Must be >= min_node_count. Cannot be used with total limits. */
	// +optional
	MaxNodeCount *int `json:"maxNodeCount,omitempty"`

	/* Minimum number of nodes per zone in the node pool. Must be >=0 and <= max_node_count. Cannot be used with total limits. */
	// +optional
	MinNodeCount *int `json:"minNodeCount,omitempty"`

	/* Maximum number of all nodes in the node pool. Must be >= total_min_node_count. Cannot be used with per zone limits. */
	// +optional
	TotalMaxNodeCount *int `json:"totalMaxNodeCount,omitempty"`

	/* Minimum number of all nodes in the node pool. Must be >=0 and <= total_max_node_count. Cannot be used with per zone limits. */
	// +optional
	TotalMinNodeCount *int `json:"totalMinNodeCount,omitempty"`
}

type NodepoolBlueGreenSettings struct {
	/* Time needed after draining entire blue pool. After this period, blue pool will be cleaned up. */
	// +optional
	NodePoolSoakDuration *string `json:"nodePoolSoakDuration,omitempty"`

	/* Standard rollout policy is the default policy for blue-green. */
	StandardRolloutPolicy NodepoolStandardRolloutPolicy `json:"standardRolloutPolicy"`
}

type NodepoolConfidentialNodes struct {
	/* Immutable. Whether Confidential Nodes feature is enabled for all nodes in this pool. */
	Enabled bool `json:"enabled"`
}

type NodepoolEphemeralStorageConfig struct {
	/* Immutable. Number of local SSDs to use to back ephemeral storage. Uses NVMe interfaces. Each local SSD must be 375 or 3000 GB in size, and all local SSDs must share the same size. */
	LocalSsdCount int `json:"localSsdCount"`
}

type NodepoolEphemeralStorageLocalSsdConfig struct {
	/* Immutable. Number of local SSDs to use to back ephemeral storage. Uses NVMe interfaces. Each local SSD must be 375 or 3000 GB in size, and all local SSDs must share the same size. */
	LocalSsdCount int `json:"localSsdCount"`
}

type NodepoolFastSocket struct {
	/* Whether or not NCCL Fast Socket is enabled. */
	Enabled bool `json:"enabled"`
}

type NodepoolGcfsConfig struct {
	/* Immutable. Whether or not GCFS is enabled. */
	Enabled bool `json:"enabled"`
}

type NodepoolGpuDriverInstallationConfig struct {
	/* Immutable. Mode for how the GPU driver is installed. */
	GpuDriverVersion string `json:"gpuDriverVersion"`
}

type NodepoolGpuSharingConfig struct {
	/* Immutable. The type of GPU sharing strategy to enable on the GPU node. Possible values are described in the API package (https://pkg.go.dev/google.golang.org/api/container/v1#GPUSharingConfig). */
	GpuSharingStrategy string `json:"gpuSharingStrategy"`

	/* Immutable. The maximum number of containers that can share a GPU. */
	MaxSharedClientsPerGpu int `json:"maxSharedClientsPerGpu"`
}

type NodepoolGuestAccelerator struct {
	/* Immutable. The number of the accelerator cards exposed to an instance. */
	Count int `json:"count"`

	/* Immutable. Configuration for auto installation of GPU driver. */
	// +optional
	GpuDriverInstallationConfig *NodepoolGpuDriverInstallationConfig `json:"gpuDriverInstallationConfig,omitempty"`

	/* Immutable. Size of partitions to create on the GPU. Valid values are described in the NVIDIA mig user guide (https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning). */
	// +optional
	GpuPartitionSize *string `json:"gpuPartitionSize,omitempty"`

	/* Immutable. Configuration for GPU sharing. */
	// +optional
	GpuSharingConfig *NodepoolGpuSharingConfig `json:"gpuSharingConfig,omitempty"`

	/* Immutable. The accelerator type resource name. */
	Type string `json:"type"`
}

type NodepoolGvnic struct {
	/* Immutable. Whether or not gvnic is enabled. */
	Enabled bool `json:"enabled"`
}

type NodepoolHostMaintenancePolicy struct {
	/* Immutable. . */
	MaintenanceInterval string `json:"maintenanceInterval"`
}

type NodepoolKubeletConfig struct {
	/* Enable CPU CFS quota enforcement for containers that specify CPU limits. */
	// +optional
	CpuCfsQuota *bool `json:"cpuCfsQuota,omitempty"`

	/* Set the CPU CFS quota period value 'cpu.cfs_period_us'. */
	// +optional
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty"`

	/* Control the CPU management policy on the node. */
	CpuManagerPolicy string `json:"cpuManagerPolicy"`

	/* Controls the maximum number of processes allowed to run in a pod. */
	// +optional
	PodPidsLimit *int `json:"podPidsLimit,omitempty"`
}

type NodepoolLinuxNodeConfig struct {
	/* cgroupMode specifies the cgroup mode to be used on the node. */
	// +optional
	CgroupMode *string `json:"cgroupMode,omitempty"`

	/* The Linux kernel parameters to be applied to the nodes and all pods running on the nodes. */
	// +optional
	Sysctls map[string]string `json:"sysctls,omitempty"`
}

type NodepoolLocalNvmeSsdBlockConfig struct {
	/* Immutable. Number of raw-block local NVMe SSD disks to be attached to the node. Each local SSD is 375 GB in size. */
	LocalSsdCount int `json:"localSsdCount"`
}

type NodepoolManagement struct {
	/* Whether the nodes will be automatically repaired. */
	// +optional
	AutoRepair *bool `json:"autoRepair,omitempty"`

	/* Whether the nodes will be automatically upgraded. */
	// +optional
	AutoUpgrade *bool `json:"autoUpgrade,omitempty"`
}

type NodepoolNetworkConfig struct {
	/* Immutable. We specify the additional node networks for this node pool using this list. Each node network corresponds to an additional interface. */
	// +optional
	AdditionalNodeNetworkConfigs []NodepoolAdditionalNodeNetworkConfigs `json:"additionalNodeNetworkConfigs,omitempty"`

	/* Immutable. We specify the additional pod networks for this node pool using this list. Each pod network corresponds to an additional alias IP range for the node. */
	// +optional
	AdditionalPodNetworkConfigs []NodepoolAdditionalPodNetworkConfigs `json:"additionalPodNetworkConfigs,omitempty"`

	/* Immutable. Whether to create a new range for pod IPs in this node pool. Defaults are provided for pod_range and pod_ipv4_cidr_block if they are not specified. */
	// +optional
	CreatePodRange *bool `json:"createPodRange,omitempty"`

	/* Whether nodes have internal IP addresses only. */
	// +optional
	EnablePrivateNodes *bool `json:"enablePrivateNodes,omitempty"`

	/* Immutable. Configuration for node-pool level pod cidr overprovision. If not set, the cluster level setting will be inherited. */
	// +optional
	PodCidrOverprovisionConfig *NodepoolPodCidrOverprovisionConfig `json:"podCidrOverprovisionConfig,omitempty"`

	/* Immutable. The IP address range for pod IPs in this node pool. Only applicable if create_pod_range is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) to pick a specific range to use. */
	// +optional
	PodIpv4CidrBlock *string `json:"podIpv4CidrBlock,omitempty"`

	/* Immutable. The ID of the secondary range for pod IPs. If create_pod_range is true, this ID is used for the new range. If create_pod_range is false, uses an existing secondary range with this ID. */
	// +optional
	PodRange *string `json:"podRange,omitempty"`
}

type NodepoolNodeAffinity struct {
	/* Immutable. . */
	Key string `json:"key"`

	/* Immutable. . */
	Operator string `json:"operator"`

	/* Immutable. . */
	Values []string `json:"values"`
}

type NodepoolNodeConfig struct {
	/* Immutable. Specifies options for controlling advanced machine features. */
	// +optional
	AdvancedMachineFeatures *NodepoolAdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty"`

	// +optional
	BootDiskKMSCryptoKeyRef *v1alpha1.ResourceRef `json:"bootDiskKMSCryptoKeyRef,omitempty"`

	/* Immutable. Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after pool creation without deleting and recreating the entire pool. */
	// +optional
	ConfidentialNodes *NodepoolConfidentialNodes `json:"confidentialNodes,omitempty"`

	/* Immutable. Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB. */
	// +optional
	DiskSizeGb *int `json:"diskSizeGb,omitempty"`

	/* Immutable. Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd. */
	// +optional
	DiskType *string `json:"diskType,omitempty"`

	/* Immutable. Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk. */
	// +optional
	EphemeralStorageConfig *NodepoolEphemeralStorageConfig `json:"ephemeralStorageConfig,omitempty"`

	/* Immutable. Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk. */
	// +optional
	EphemeralStorageLocalSsdConfig *NodepoolEphemeralStorageLocalSsdConfig `json:"ephemeralStorageLocalSsdConfig,omitempty"`

	/* Enable or disable NCCL Fast Socket in the node pool. */
	// +optional
	FastSocket *NodepoolFastSocket `json:"fastSocket,omitempty"`

	/* Immutable. GCFS configuration for this node. */
	// +optional
	GcfsConfig *NodepoolGcfsConfig `json:"gcfsConfig,omitempty"`

	/* Immutable. List of the type and count of accelerator cards attached to the instance. */
	// +optional
	GuestAccelerator []NodepoolGuestAccelerator `json:"guestAccelerator,omitempty"`

	/* Immutable. Enable or disable gvnic in the node pool. */
	// +optional
	Gvnic *NodepoolGvnic `json:"gvnic,omitempty"`

	/* Immutable. The maintenance policy for the hosts on which the GKE VMs run on. */
	// +optional
	HostMaintenancePolicy *NodepoolHostMaintenancePolicy `json:"hostMaintenancePolicy,omitempty"`

	/* The image type to use for this node. Note that for a given image type, the latest version of it will be used. */
	// +optional
	ImageType *string `json:"imageType,omitempty"`

	/* Node kubelet configs. */
	// +optional
	KubeletConfig *NodepoolKubeletConfig `json:"kubeletConfig,omitempty"`

	/* The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	/* Parameters that can be configured on Linux nodes. */
	// +optional
	LinuxNodeConfig *NodepoolLinuxNodeConfig `json:"linuxNodeConfig,omitempty"`

	/* Immutable. Parameters for raw-block local NVMe SSDs. */
	// +optional
	LocalNvmeSsdBlockConfig *NodepoolLocalNvmeSsdBlockConfig `json:"localNvmeSsdBlockConfig,omitempty"`

	/* Immutable. The number of local SSD disks to be attached to the node. */
	// +optional
	LocalSsdCount *int `json:"localSsdCount,omitempty"`

	/* Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT. */
	// +optional
	LoggingVariant *string `json:"loggingVariant,omitempty"`

	/* Immutable. The name of a Google Compute Engine machine type. */
	// +optional
	MachineType *string `json:"machineType,omitempty"`

	/* Immutable. The metadata key/value pairs assigned to instances in the cluster. */
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`

	/* Immutable. Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform. */
	// +optional
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. Setting this field will assign instances
	of this pool to run on the specified node group. This is useful
	for running workloads on sole tenant nodes. */
	// +optional
	NodeGroupRef *v1alpha1.ResourceRef `json:"nodeGroupRef,omitempty"`

	/* Immutable. The set of Google API scopes to be made available on all of the node VMs. */
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty"`

	/* Immutable. Whether the nodes are created as preemptible VM instances. */
	// +optional
	Preemptible *bool `json:"preemptible,omitempty"`

	/* Immutable. The reservation affinity configuration for the node pool. */
	// +optional
	ReservationAffinity *NodepoolReservationAffinity `json:"reservationAffinity,omitempty"`

	/* The GCE resource labels (a map of key/value pairs) to be applied to the node pool. */
	// +optional
	ResourceLabels map[string]string `json:"resourceLabels,omitempty"`

	/* Immutable. Sandbox configuration for this node. */
	// +optional
	SandboxConfig *NodepoolSandboxConfig `json:"sandboxConfig,omitempty"`

	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	/* Immutable. Shielded Instance options. */
	// +optional
	ShieldedInstanceConfig *NodepoolShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	/* Immutable. Node affinity options for sole tenant node pools. */
	// +optional
	SoleTenantConfig *NodepoolSoleTenantConfig `json:"soleTenantConfig,omitempty"`

	/* Immutable. Whether the nodes are created as spot VM instances. */
	// +optional
	Spot *bool `json:"spot,omitempty"`

	/* The list of instance tags applied to all nodes. */
	// +optional
	Tags []string `json:"tags,omitempty"`

	/* Immutable. List of Kubernetes taints to be applied to each node. */
	// +optional
	Taint []NodepoolTaint `json:"taint,omitempty"`

	/* The workload metadata configuration for this node. */
	// +optional
	WorkloadMetadataConfig *NodepoolWorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty"`
}

type NodepoolPlacementPolicy struct {
	/* Immutable. If set, refers to the name of a custom resource policy supplied by the user. The resource policy must be in the same project and region as the node pool. If not found, InvalidArgument error is returned. */
	// +optional
	PolicyNameRef *v1alpha1.ResourceRef `json:"policyNameRef,omitempty"`

	/* TPU placement topology for pod slice node pool. https://cloud.google.com/tpu/docs/types-topologies#tpu_topologies. */
	// +optional
	TpuTopology *string `json:"tpuTopology,omitempty"`

	/* Type defines the type of placement policy. */
	Type string `json:"type"`
}

type NodepoolPodCidrOverprovisionConfig struct {
	Disabled bool `json:"disabled"`
}

type NodepoolReservationAffinity struct {
	/* Immutable. Corresponds to the type of reservation consumption. */
	ConsumeReservationType string `json:"consumeReservationType"`

	/* Immutable. The label key of a reservation resource. */
	// +optional
	Key *string `json:"key,omitempty"`

	/* Immutable. The label values of the reservation resource. */
	// +optional
	Values []string `json:"values,omitempty"`
}

type NodepoolSandboxConfig struct {
	/* Type of the sandbox to use for the node (e.g. 'gvisor'). */
	SandboxType string `json:"sandboxType"`
}

type NodepoolShieldedInstanceConfig struct {
	/* Immutable. Defines whether the instance has integrity monitoring enabled. */
	// +optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`

	/* Immutable. Defines whether the instance has Secure Boot enabled. */
	// +optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`
}

type NodepoolSoleTenantConfig struct {
	/* Immutable. . */
	NodeAffinity []NodepoolNodeAffinity `json:"nodeAffinity"`
}

type NodepoolStandardRolloutPolicy struct {
	/* Number of blue nodes to drain in a batch. */
	// +optional
	BatchNodeCount *int `json:"batchNodeCount,omitempty"`

	/* Percentage of the blue pool nodes to drain in a batch. */
	// +optional
	BatchPercentage *float64 `json:"batchPercentage,omitempty"`

	/* Soak time after each batch gets drained. */
	// +optional
	BatchSoakDuration *string `json:"batchSoakDuration,omitempty"`
}

type NodepoolTaint struct {
	/* Immutable. Effect for taint. */
	Effect string `json:"effect"`

	/* Immutable. Key for taint. */
	Key string `json:"key"`

	/* Immutable. Value for taint. */
	Value string `json:"value"`
}

type NodepoolUpgradeSettings struct {
	/* Settings for BlueGreen node pool upgrade. */
	// +optional
	BlueGreenSettings *NodepoolBlueGreenSettings `json:"blueGreenSettings,omitempty"`

	/* The number of additional nodes that can be added to the node pool during an upgrade. Increasing max_surge raises the number of nodes that can be upgraded simultaneously. Can be set to 0 or greater. */
	// +optional
	MaxSurge *int `json:"maxSurge,omitempty"`

	/* The number of nodes that can be simultaneously unavailable during an upgrade. Increasing max_unavailable raises the number of nodes that can be upgraded in parallel. Can be set to 0 or greater. */
	// +optional
	MaxUnavailable *int `json:"maxUnavailable,omitempty"`

	/* Update strategy for the given nodepool. */
	// +optional
	Strategy *string `json:"strategy,omitempty"`
}

type NodepoolWorkloadMetadataConfig struct {
	/* Mode is the configuration for how to expose metadata to workloads running on the node. */
	// +optional
	Mode *string `json:"mode,omitempty"`

	/* DEPRECATED. Deprecated in favor of mode. NodeMetadata is the configuration for how to expose metadata to the workloads running on the node. */
	// +optional
	NodeMetadata *string `json:"nodeMetadata,omitempty"`
}

type ContainerNodePoolSpec struct {
	/* Configuration required by cluster autoscaler to adjust the size of the node pool to the current cluster usage. To disable autoscaling, set minNodeCount and maxNodeCount to 0. */
	// +optional
	Autoscaling *NodepoolAutoscaling `json:"autoscaling,omitempty"`

	ClusterRef v1alpha1.ResourceRef `json:"clusterRef"`

	/* Immutable. The initial number of nodes for the pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Changing this will force recreation of the resource. */
	// +optional
	InitialNodeCount *int `json:"initialNodeCount,omitempty"`

	/* Immutable. The location (region or zone) of the cluster. */
	Location string `json:"location"`

	/* Node management configuration, wherein auto-repair and auto-upgrade is configured. */
	// +optional
	Management *NodepoolManagement `json:"management,omitempty"`

	/* Immutable. The maximum number of pods per node in this node pool. Note that this does not work on node pools which are "route-based" - that is, node pools belonging to clusters that do not have IP Aliasing enabled. */
	// +optional
	MaxPodsPerNode *int `json:"maxPodsPerNode,omitempty"`

	/* Immutable. Creates a unique name for the node pool beginning with the specified prefix. Conflicts with name. */
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty"`

	/* Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults. */
	// +optional
	NetworkConfig *NodepoolNetworkConfig `json:"networkConfig,omitempty"`

	/* Immutable. The configuration of the nodepool. */
	// +optional
	NodeConfig *NodepoolNodeConfig `json:"nodeConfig,omitempty"`

	/* The number of nodes per instance group. This field can be used to update the number of nodes per instance group but should not be used alongside autoscaling. */
	// +optional
	NodeCount *int `json:"nodeCount,omitempty"`

	/* The list of zones in which the node pool's nodes should be located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If unspecified, the cluster-level node_locations will be used. */
	// +optional
	NodeLocations []string `json:"nodeLocations,omitempty"`

	/* Immutable. Specifies the node placement policy. */
	// +optional
	PlacementPolicy *NodepoolPlacementPolicy `json:"placementPolicy,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Specify node upgrade settings to change how many nodes GKE attempts to upgrade at once. The number of nodes upgraded simultaneously is the sum of max_surge and max_unavailable. The maximum number of nodes upgraded simultaneously is limited to 20. */
	// +optional
	UpgradeSettings *NodepoolUpgradeSettings `json:"upgradeSettings,omitempty"`

	// +optional
	Version *string `json:"version,omitempty"`
}

type NodepoolObservedStateStatus struct {
	// +optional
	Version *string `json:"version,omitempty"`
}

type ContainerNodePoolStatus struct {
	/* Conditions represent the latest available observations of the
	   ContainerNodePool's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The resource URLs of the managed instance groups associated with this node pool. */
	// +optional
	InstanceGroupUrls []string `json:"instanceGroupUrls,omitempty"`

	/* List of instance group URLs which have been assigned to this node pool. */
	// +optional
	ManagedInstanceGroupUrls []string `json:"managedInstanceGroupUrls,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The observed state of the underlying GCP resource. */
	// +optional
	ObservedState *NodepoolObservedStateStatus `json:"observedState,omitempty"`

	// +optional
	Operation *string `json:"operation,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontainernodepool;gcpcontainernodepools
// +kubebuilder:subresource:status

// ContainerNodePool is the Schema for the container API
// +k8s:openapi-gen=true
type ContainerNodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerNodePoolSpec   `json:"spec,omitempty"`
	Status ContainerNodePoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerNodePoolList contains a list of ContainerNodePool
type ContainerNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerNodePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerNodePool{}, &ContainerNodePoolList{})
}
