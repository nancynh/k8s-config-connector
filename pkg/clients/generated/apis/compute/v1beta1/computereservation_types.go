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

type ReservationGuestAccelerators struct {
	/* Immutable. The number of the guest accelerator cards exposed to
	this instance. */
	AcceleratorCount int `json:"acceleratorCount"`

	/* Immutable. The full or partial URL of the accelerator type to
	attach to this instance. For example:
	'projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100'

	If you are creating an instance template, specify only the accelerator name. */
	AcceleratorType string `json:"acceleratorType"`
}

type ReservationInstanceProperties struct {
	/* Immutable. Guest accelerator type and count. */
	// +optional
	GuestAccelerators []ReservationGuestAccelerators `json:"guestAccelerators,omitempty"`

	/* Immutable. The amount of local ssd to reserve with each instance. This
	reserves disks of type 'local-ssd'. */
	// +optional
	LocalSsds []ReservationLocalSsds `json:"localSsds,omitempty"`

	/* Immutable. The name of the machine type to reserve. */
	MachineType string `json:"machineType"`

	/* Immutable. The minimum CPU platform for the reservation. For example,
	'"Intel Skylake"'. See
	the CPU platform availability reference](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform#availablezones)
	for information on available CPU platforms. */
	// +optional
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`
}

type ReservationLocalSsds struct {
	/* Immutable. The size of the disk in base-2 GB. */
	DiskSizeGb int `json:"diskSizeGb"`

	/* Immutable. The disk interface to use for attaching this disk. Default value: "SCSI" Possible values: ["SCSI", "NVME"]. */
	// +optional
	Interface *string `json:"interface,omitempty"`
}

type ReservationSpecificReservation struct {
	/* The number of resources that are allocated. */
	Count int `json:"count"`

	/* How many instances are in use. */
	// +optional
	InUseCount *int `json:"inUseCount,omitempty"`

	/* Immutable. The instance properties for the reservation. */
	InstanceProperties ReservationInstanceProperties `json:"instanceProperties"`
}

type ComputeReservationSpec struct {
	/* Immutable. An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Reservation for instances with specific machine shapes. */
	SpecificReservation ReservationSpecificReservation `json:"specificReservation"`

	/* Immutable. When set to true, only VMs that target this reservation by name can
	consume this reservation. Otherwise, it can be consumed by VMs with
	affinity for any reservation. Defaults to false. */
	// +optional
	SpecificReservationRequired *bool `json:"specificReservationRequired,omitempty"`

	/* Immutable. The zone where the reservation is made. */
	Zone string `json:"zone"`
}

type ComputeReservationStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeReservation's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Full or partial URL to a parent commitment. This field displays for
	reservations that are tied to a commitment. */
	// +optional
	Commitment *string `json:"commitment,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* The status of the reservation. */
	// +optional
	Status *string `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputereservation;gcpcomputereservations
// +kubebuilder:subresource:status

// ComputeReservation is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeReservationSpec   `json:"spec,omitempty"`
	Status ComputeReservationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeReservationList contains a list of ComputeReservation
type ComputeReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeReservation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeReservation{}, &ComputeReservationList{})
}
