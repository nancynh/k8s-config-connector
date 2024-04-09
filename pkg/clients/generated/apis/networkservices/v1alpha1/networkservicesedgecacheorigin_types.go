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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EdgecacheoriginAwsV4Authentication struct {
	/* The access key ID your origin uses to identify the key. */
	AccessKeyId string `json:"accessKeyId"`

	/* The name of the AWS region that your origin is in. */
	OriginRegion string `json:"originRegion"`

	/* The Secret Manager secret version of the secret access key used by your origin.

	This is the resource name of the secret version in the format 'projects/* /secrets/* /versions/*' where the '*' values are replaced by the project, secret, and version you require. */
	SecretAccessKeyVersion string `json:"secretAccessKeyVersion"`
}

type EdgecacheoriginHeaderAction struct {
	/* Describes a header to add.

	You may add a maximum of 25 request headers. */
	// +optional
	RequestHeadersToAdd []EdgecacheoriginRequestHeadersToAdd `json:"requestHeadersToAdd,omitempty"`
}

type EdgecacheoriginOriginOverrideAction struct {
	/* The header actions, including adding and removing
	headers, for request handled by this origin. */
	// +optional
	HeaderAction *EdgecacheoriginHeaderAction `json:"headerAction,omitempty"`

	/* The URL rewrite configuration for request that are
	handled by this origin. */
	// +optional
	UrlRewrite *EdgecacheoriginUrlRewrite `json:"urlRewrite,omitempty"`
}

type EdgecacheoriginOriginRedirect struct {
	/* The set of redirect response codes that the CDN
	follows. Values of
	[RedirectConditions](https://cloud.google.com/media-cdn/docs/reference/rest/v1/projects.locations.edgeCacheOrigins#redirectconditions)
	are accepted. */
	// +optional
	RedirectConditions []string `json:"redirectConditions,omitempty"`
}

type EdgecacheoriginRequestHeadersToAdd struct {
	/* The name of the header to add. */
	HeaderName string `json:"headerName"`

	/* The value of the header to add. */
	HeaderValue string `json:"headerValue"`

	/* Whether to replace all existing headers with the same name.

	By default, added header values are appended
	to the response or request headers with the
	same field names. The added values are
	separated by commas.

	To overwrite existing values, set 'replace' to 'true'. */
	// +optional
	Replace *bool `json:"replace,omitempty"`
}

type EdgecacheoriginTimeout struct {
	/* The maximum duration to wait for a single origin connection to be established, including DNS lookup, TLS handshake and TCP/QUIC connection establishment.

	Defaults to 5 seconds. The timeout must be a value between 1s and 15s.

	The connectTimeout capped by the deadline set by the request's maxAttemptsTimeout.  The last connection attempt may have a smaller connectTimeout in order to adhere to the overall maxAttemptsTimeout. */
	// +optional
	ConnectTimeout *string `json:"connectTimeout,omitempty"`

	/* The maximum time across all connection attempts to the origin, including failover origins, before returning an error to the client. A HTTP 504 will be returned if the timeout is reached before a response is returned.

	Defaults to 15 seconds. The timeout must be a value between 1s and 30s.

	If a failoverOrigin is specified, the maxAttemptsTimeout of the first configured origin sets the deadline for all connection attempts across all failoverOrigins. */
	// +optional
	MaxAttemptsTimeout *string `json:"maxAttemptsTimeout,omitempty"`

	/* The maximum duration to wait between reads of a single HTTP connection/stream.

	Defaults to 15 seconds.  The timeout must be a value between 1s and 30s.

	The readTimeout is capped by the responseTimeout.  All reads of the HTTP connection/stream must be completed by the deadline set by the responseTimeout.

	If the response headers have already been written to the connection, the response will be truncated and logged. */
	// +optional
	ReadTimeout *string `json:"readTimeout,omitempty"`

	/* The maximum duration to wait for the last byte of a response to arrive when reading from the HTTP connection/stream.

	Defaults to 30 seconds. The timeout must be a value between 1s and 120s.

	The responseTimeout starts after the connection has been established.

	This also applies to HTTP Chunked Transfer Encoding responses, and/or when an open-ended Range request is made to the origin. Origins that take longer to write additional bytes to the response than the configured responseTimeout will result in an error being returned to the client.

	If the response headers have already been written to the connection, the response will be truncated and logged. */
	// +optional
	ResponseTimeout *string `json:"responseTimeout,omitempty"`
}

type EdgecacheoriginUrlRewrite struct {
	/* Prior to forwarding the request to the selected
	origin, the request's host header is replaced with
	contents of the hostRewrite.

	This value must be between 1 and 255 characters. */
	// +optional
	HostRewrite *string `json:"hostRewrite,omitempty"`
}

type NetworkServicesEdgeCacheOriginSpec struct {
	/* Enable AWS Signature Version 4 origin authentication. */
	// +optional
	AwsV4Authentication *EdgecacheoriginAwsV4Authentication `json:"awsV4Authentication,omitempty"`

	/* A human-readable description of the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The Origin resource to try when the current origin cannot be reached.
	After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.

	The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
	A reference to a Topic resource. */
	// +optional
	FailoverOrigin *string `json:"failoverOrigin,omitempty"`

	/* The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.

	Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
	retryConditions and failoverOrigin to control its own cache fill failures.

	The total number of allowed attempts to cache fill across this and failover origins is limited to four.
	The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.

	The last valid, non-retried response from all origins will be returned to the client.
	If no origin returns a valid response, an HTTP 502 will be returned to the client.

	Defaults to 1. Must be a value greater than 0 and less than 4. */
	// +optional
	MaxAttempts *int `json:"maxAttempts,omitempty"`

	/* A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.

	This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com, IPv4: 35.218.1.1, IPv6: 2607:f8b0:4012:809::200e, Cloud Storage: gs://bucketname

	When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.  It must not contain a protocol (e.g., https://) and it must not contain any slashes.
	If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected. */
	OriginAddress string `json:"originAddress"`

	/* The override actions, including url rewrites and header
	additions, for requests that use this origin. */
	// +optional
	OriginOverrideAction *EdgecacheoriginOriginOverrideAction `json:"originOverrideAction,omitempty"`

	/* Follow redirects from this origin. */
	// +optional
	OriginRedirect *EdgecacheoriginOriginRedirect `json:"originRedirect,omitempty"`

	/* The port to connect to the origin on.
	Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP. */
	// +optional
	Port *int `json:"port,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.

	When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server. Possible values: ["HTTP2", "HTTPS", "HTTP"]. */
	// +optional
	Protocol *string `json:"protocol,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Specifies one or more retry conditions for the configured origin.

	If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
	the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.

	The default retryCondition is "CONNECT_FAILURE".

	retryConditions apply to this origin, and not subsequent failoverOrigin(s),
	which may specify their own retryConditions and maxAttempts.

	Valid values are:

	- CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
	- HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
	- GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
	- RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
	- NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet.
	- FORBIDDEN: Retry if the origin returns a HTTP 403 (Forbidden). Possible values: ["CONNECT_FAILURE", "HTTP_5XX", "GATEWAY_ERROR", "RETRIABLE_4XX", "NOT_FOUND", "FORBIDDEN"]. */
	// +optional
	RetryConditions []string `json:"retryConditions,omitempty"`

	/* The connection and HTTP timeout configuration for this origin. */
	// +optional
	Timeout *EdgecacheoriginTimeout `json:"timeout,omitempty"`
}

type NetworkServicesEdgeCacheOriginStatus struct {
	/* Conditions represent the latest available observations of the
	   NetworkServicesEdgeCacheOrigin's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkservicesedgecacheorigin;gcpnetworkservicesedgecacheorigins
// +kubebuilder:subresource:status

// NetworkServicesEdgeCacheOrigin is the Schema for the networkservices API
// +k8s:openapi-gen=true
type NetworkServicesEdgeCacheOrigin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesEdgeCacheOriginSpec   `json:"spec,omitempty"`
	Status NetworkServicesEdgeCacheOriginStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkServicesEdgeCacheOriginList contains a list of NetworkServicesEdgeCacheOrigin
type NetworkServicesEdgeCacheOriginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesEdgeCacheOrigin `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesEdgeCacheOrigin{}, &NetworkServicesEdgeCacheOriginList{})
}
