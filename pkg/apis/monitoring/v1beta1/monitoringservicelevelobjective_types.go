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

type ServicelevelobjectiveAvailability struct {
}

type ServicelevelobjectiveBasicSli struct {
	/* Good service is defined to be the count of requests made to this service that return successfully. */
	// +optional
	Availability *ServicelevelobjectiveAvailability `json:"availability,omitempty"`

	/* Good service is defined to be the count of requests made to this service that are fast enough with respect to `latency.threshold`. */
	// +optional
	Latency *ServicelevelobjectiveLatency `json:"latency,omitempty"`

	/* OPTIONAL: The set of locations to which this SLI is relevant. Telemetry from other locations will not be used to calculate performance for this SLI. If omitted, this SLI applies to all locations in which the Service has activity. For service types that don't support breaking down by location, setting this field will result in an error. */
	// +optional
	Location []string `json:"location,omitempty"`

	/* OPTIONAL: The set of RPCs to which this SLI is relevant. Telemetry from other methods will not be used to calculate performance for this SLI. If omitted, this SLI applies to all the Service's methods. For service types that don't support breaking down by method, setting this field will result in an error. */
	// +optional
	Method []string `json:"method,omitempty"`

	/* Good service is defined to be the count of operations performed by this service that return successfully */
	// +optional
	OperationAvailability *ServicelevelobjectiveOperationAvailability `json:"operationAvailability,omitempty"`

	/* Good service is defined to be the count of operations performed by this service that are fast enough with respect to `operation_latency.threshold`. */
	// +optional
	OperationLatency *ServicelevelobjectiveOperationLatency `json:"operationLatency,omitempty"`

	/* OPTIONAL: The set of API versions to which this SLI is relevant. Telemetry from other API versions will not be used to calculate performance for this SLI. If omitted, this SLI applies to all API versions. For service types that don't support breaking down by version, setting this field will result in an error. */
	// +optional
	Version []string `json:"version,omitempty"`
}

type ServicelevelobjectiveBasicSliPerformance struct {
	/* Good service is defined to be the count of requests made to this service that return successfully. */
	// +optional
	Availability *ServicelevelobjectiveAvailability `json:"availability,omitempty"`

	/* Good service is defined to be the count of requests made to this service that are fast enough with respect to `latency.threshold`. */
	// +optional
	Latency *ServicelevelobjectiveLatency `json:"latency,omitempty"`

	/* OPTIONAL: The set of locations to which this SLI is relevant. Telemetry from other locations will not be used to calculate performance for this SLI. If omitted, this SLI applies to all locations in which the Service has activity. For service types that don't support breaking down by location, setting this field will result in an error. */
	// +optional
	Location []string `json:"location,omitempty"`

	/* OPTIONAL: The set of RPCs to which this SLI is relevant. Telemetry from other methods will not be used to calculate performance for this SLI. If omitted, this SLI applies to all the Service's methods. For service types that don't support breaking down by method, setting this field will result in an error. */
	// +optional
	Method []string `json:"method,omitempty"`

	/* Good service is defined to be the count of operations performed by this service that return successfully */
	// +optional
	OperationAvailability *ServicelevelobjectiveOperationAvailability `json:"operationAvailability,omitempty"`

	/* Good service is defined to be the count of operations performed by this service that are fast enough with respect to `operation_latency.threshold`. */
	// +optional
	OperationLatency *ServicelevelobjectiveOperationLatency `json:"operationLatency,omitempty"`

	/* OPTIONAL: The set of API versions to which this SLI is relevant. Telemetry from other API versions will not be used to calculate performance for this SLI. If omitted, this SLI applies to all API versions. For service types that don't support breaking down by version, setting this field will result in an error. */
	// +optional
	Version []string `json:"version,omitempty"`
}

type ServicelevelobjectiveDistributionCut struct {
	/* A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) specifying a `TimeSeries` aggregating values. Must have `ValueType = DISTRIBUTION` and `MetricKind = DELTA` or `MetricKind = CUMULATIVE`. */
	// +optional
	DistributionFilter *string `json:"distributionFilter,omitempty"`

	/* Range of values considered "good." For a one-sided range, set one bound to an infinite value. */
	// +optional
	Range *ServicelevelobjectiveRange `json:"range,omitempty"`
}

type ServicelevelobjectiveGoodTotalRatio struct {
	/* A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) specifying a `TimeSeries` quantifying bad service, either demanded service that was not provided or demanded service that was of inadequate quality. Must have `ValueType = DOUBLE` or `ValueType = INT64` and must have `MetricKind = DELTA` or `MetricKind = CUMULATIVE`. */
	// +optional
	BadServiceFilter *string `json:"badServiceFilter,omitempty"`

	/* A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) specifying a `TimeSeries` quantifying good service provided. Must have `ValueType = DOUBLE` or `ValueType = INT64` and must have `MetricKind = DELTA` or `MetricKind = CUMULATIVE`. */
	// +optional
	GoodServiceFilter *string `json:"goodServiceFilter,omitempty"`

	/* A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) specifying a `TimeSeries` quantifying total demanded service. Must have `ValueType = DOUBLE` or `ValueType = INT64` and must have `MetricKind = DELTA` or `MetricKind = CUMULATIVE`. */
	// +optional
	TotalServiceFilter *string `json:"totalServiceFilter,omitempty"`
}

type ServicelevelobjectiveGoodTotalRatioThreshold struct {
	/* `BasicSli` to evaluate to judge window quality. */
	// +optional
	BasicSliPerformance *ServicelevelobjectiveBasicSliPerformance `json:"basicSliPerformance,omitempty"`

	/* `RequestBasedSli` to evaluate to judge window quality. */
	// +optional
	Performance *ServicelevelobjectivePerformance `json:"performance,omitempty"`

	/* If window `performance >= threshold`, the window is counted as good. */
	// +optional
	Threshold *float64 `json:"threshold,omitempty"`
}

type ServicelevelobjectiveLatency struct {
	/* A description of the experience associated with failing requests. Possible values: LATENCY_EXPERIENCE_UNSPECIFIED, DELIGHTING, SATISFYING, ANNOYING */
	// +optional
	Experience *string `json:"experience,omitempty"`

	/* Good service is defined to be the count of requests made to this service that return in no more than `threshold`. */
	// +optional
	Threshold *string `json:"threshold,omitempty"`
}

type ServicelevelobjectiveMetricMeanInRange struct {
	/* Range of values considered "good." For a one-sided range, set one bound to an infinite value. */
	// +optional
	Range *ServicelevelobjectiveRange `json:"range,omitempty"`

	/* A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) specifying the `TimeSeries` to use for evaluating window quality. */
	// +optional
	TimeSeries *string `json:"timeSeries,omitempty"`
}

type ServicelevelobjectiveMetricSumInRange struct {
	/* Range of values considered "good." For a one-sided range, set one bound to an infinite value. */
	// +optional
	Range *ServicelevelobjectiveRange `json:"range,omitempty"`

	/* A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) specifying the `TimeSeries` to use for evaluating window quality. */
	// +optional
	TimeSeries *string `json:"timeSeries,omitempty"`
}

type ServicelevelobjectiveOperationAvailability struct {
}

type ServicelevelobjectiveOperationLatency struct {
	/* A description of the experience associated with failing requests. Possible values: LATENCY_EXPERIENCE_UNSPECIFIED, DELIGHTING, SATISFYING, ANNOYING */
	// +optional
	Experience *string `json:"experience,omitempty"`

	/* Good service is defined to be the count of operations that are completed in no more than `threshold`. */
	// +optional
	Threshold *string `json:"threshold,omitempty"`
}

type ServicelevelobjectivePerformance struct {
	/* `distribution_cut` is used when `good_service` is a count of values aggregated in a `Distribution` that fall into a good range. The `total_service` is the total count of all values aggregated in the `Distribution`. */
	// +optional
	DistributionCut *ServicelevelobjectiveDistributionCut `json:"distributionCut,omitempty"`

	/* `good_total_ratio` is used when the ratio of `good_service` to `total_service` is computed from two `TimeSeries`. */
	// +optional
	GoodTotalRatio *ServicelevelobjectiveGoodTotalRatio `json:"goodTotalRatio,omitempty"`
}

type ServicelevelobjectiveRange struct {
	/* Range maximum. */
	// +optional
	Max *float64 `json:"max,omitempty"`

	/* Range minimum. */
	// +optional
	Min *float64 `json:"min,omitempty"`
}

type ServicelevelobjectiveRequestBased struct {
	/* `distribution_cut` is used when `good_service` is a count of values aggregated in a `Distribution` that fall into a good range. The `total_service` is the total count of all values aggregated in the `Distribution`. */
	// +optional
	DistributionCut *ServicelevelobjectiveDistributionCut `json:"distributionCut,omitempty"`

	/* `good_total_ratio` is used when the ratio of `good_service` to `total_service` is computed from two `TimeSeries`. */
	// +optional
	GoodTotalRatio *ServicelevelobjectiveGoodTotalRatio `json:"goodTotalRatio,omitempty"`
}

type ServicelevelobjectiveServiceLevelIndicator struct {
	/* Basic SLI on a well-known service type. */
	// +optional
	BasicSli *ServicelevelobjectiveBasicSli `json:"basicSli,omitempty"`

	/* Request-based SLIs */
	// +optional
	RequestBased *ServicelevelobjectiveRequestBased `json:"requestBased,omitempty"`

	/* Windows-based SLIs */
	// +optional
	WindowsBased *ServicelevelobjectiveWindowsBased `json:"windowsBased,omitempty"`
}

type ServicelevelobjectiveWindowsBased struct {
	/* A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) specifying a `TimeSeries` with `ValueType = BOOL`. The window is good if any `true` values appear in the window. */
	// +optional
	GoodBadMetricFilter *string `json:"goodBadMetricFilter,omitempty"`

	/* A window is good if its `performance` is high enough. */
	// +optional
	GoodTotalRatioThreshold *ServicelevelobjectiveGoodTotalRatioThreshold `json:"goodTotalRatioThreshold,omitempty"`

	/* A window is good if the metric's value is in a good range, averaged across returned streams. */
	// +optional
	MetricMeanInRange *ServicelevelobjectiveMetricMeanInRange `json:"metricMeanInRange,omitempty"`

	/* A window is good if the metric's value is in a good range, summed across returned streams. */
	// +optional
	MetricSumInRange *ServicelevelobjectiveMetricSumInRange `json:"metricSumInRange,omitempty"`

	/* Duration over which window quality is evaluated. Must be an integer fraction of a day and at least `60s`. */
	// +optional
	WindowPeriod *string `json:"windowPeriod,omitempty"`
}

type MonitoringServiceLevelObjectiveSpec struct {
	/* A calendar period, semantically "since the start of the current ``". At this time, only `DAY`, `WEEK`, `FORTNIGHT`, and `MONTH` are supported. Possible values: CALENDAR_PERIOD_UNSPECIFIED, DAY, WEEK, FORTNIGHT, MONTH, QUARTER, HALF, YEAR */
	// +optional
	CalendarPeriod *string `json:"calendarPeriod,omitempty"`

	/* Name used for UI elements listing this SLO. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* The fraction of service that must be good in order for this objective to be met. `0 < goal <= 0.999`. */
	Goal float64 `json:"goal"`

	/* The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* A rolling time period, semantically "in the past ``". Must be an integer multiple of 1 day no larger than 30 days. */
	// +optional
	RollingPeriod *string `json:"rollingPeriod,omitempty"`

	/* The definition of good service, used to measure and calculate the quality of the `Service`'s performance with respect to a single aspect of service quality. */
	// +optional
	ServiceLevelIndicator *ServicelevelobjectiveServiceLevelIndicator `json:"serviceLevelIndicator,omitempty"`

	/*  */
	ServiceRef v1alpha1.ResourceRef `json:"serviceRef"`
}

type MonitoringServiceLevelObjectiveStatus struct {
	/* Conditions represent the latest available observations of the
	   MonitoringServiceLevelObjective's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Time stamp of the `Create` or most recent `Update` command on this `Slo`. */
	CreateTime string `json:"createTime,omitempty"`
	/* Time stamp of the `Update` or `Delete` command that made this no longer a current `Slo`. This field is not populated in `ServiceLevelObjective`s returned from calls to `GetServiceLevelObjective` and `ListServiceLevelObjectives`, because it is always empty in the current version. It is populated in `ServiceLevelObjective`s representing previous versions in the output of `ListServiceLevelObjectiveVersions`. Because all old configuration versions are stored, `Update` operations mark the obsoleted version as deleted. */
	DeleteTime string `json:"deleteTime,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. If set, this SLO is managed at the [Service Management](https://cloud.google.com/service-management/overview) level. Therefore the service yaml file is the source of truth for this SLO, and API `Update` and `Delete` operations are forbidden. */
	ServiceManagementOwned bool `json:"serviceManagementOwned,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringServiceLevelObjective is the Schema for the monitoring API
// +k8s:openapi-gen=true
type MonitoringServiceLevelObjective struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringServiceLevelObjectiveSpec   `json:"spec,omitempty"`
	Status MonitoringServiceLevelObjectiveStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringServiceLevelObjectiveList contains a list of MonitoringServiceLevelObjective
type MonitoringServiceLevelObjectiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringServiceLevelObjective `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringServiceLevelObjective{}, &MonitoringServiceLevelObjectiveList{})
}
