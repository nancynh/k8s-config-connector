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

type IntentParameters struct {
	/* The entity type of the parameter.
	Format: projects/-/locations/-/agents/-/entityTypes/<System Entity Type ID> for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<Entity Type ID> for developer entity types. */
	EntityType string `json:"entityType"`

	/* The unique identifier of the parameter. This field is used by training phrases to annotate their parts. */
	Id string `json:"id"`

	/* Indicates whether the parameter represents a list of values. */
	// +optional
	IsList *bool `json:"isList,omitempty"`

	/* Indicates whether the parameter content should be redacted in log. If redaction is enabled, the parameter content will be replaced by parameter name during logging.
	Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled. */
	// +optional
	Redact *bool `json:"redact,omitempty"`
}

type IntentParts struct {
	/* The parameter used to annotate this part of the training phrase. This field is required for annotated parts of the training phrase. */
	// +optional
	ParameterId *string `json:"parameterId,omitempty"`

	/* The text for this part. */
	Text string `json:"text"`
}

type IntentTrainingPhrases struct {
	/* The unique identifier of the training phrase. */
	// +optional
	Id *string `json:"id,omitempty"`

	/* The ordered list of training phrase parts. The parts are concatenated in order to form the training phrase.
	Note: The API does not automatically annotate training phrases like the Dialogflow Console does.
	Note: Do not forget to include whitespace at part boundaries, so the training phrase is well formatted when the parts are concatenated.
	If the training phrase does not need to be annotated with parameters, you just need a single part with only the Part.text field set.
	If you want to annotate the training phrase, you must create multiple parts, where the fields of each part are populated in one of two ways:
	Part.text is set to a part of the phrase that has no parameters.
	Part.text is set to a part of the phrase that you want to annotate, and the parameterId field is set. */
	Parts []IntentParts `json:"parts"`

	/* Indicates how many times this example was added to the intent. */
	// +optional
	RepeatCount *int `json:"repeatCount,omitempty"`
}

type DialogflowCXIntentSpec struct {
	/* Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The human-readable name of the intent, unique within the agent. */
	DisplayName string `json:"displayName"`

	/* Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
	Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event. */
	// +optional
	IsFallback *bool `json:"isFallback,omitempty"`

	/* Immutable. The language of the following fields in intent:
	Intent.training_phrases.parts.text
	If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used. */
	// +optional
	LanguageCode *string `json:"languageCode,omitempty"`

	/* The collection of parameters associated with the intent. */
	// +optional
	Parameters []IntentParameters `json:"parameters,omitempty"`

	/* Immutable. The agent to create an intent for.
	Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>. */
	// +optional
	Parent *string `json:"parent,omitempty"`

	/* The priority of this intent. Higher numbers represent higher priorities.
	If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
	If the supplied value is negative, the intent is ignored in runtime detect intent requests. */
	// +optional
	Priority *int `json:"priority,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The collection of training phrases the agent is trained on to identify the intent. */
	// +optional
	TrainingPhrases []IntentTrainingPhrases `json:"trainingPhrases,omitempty"`
}

type DialogflowCXIntentStatus struct {
	/* Conditions represent the latest available observations of the
	   DialogflowCXIntent's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The unique identifier of the intent.
	Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowcxintent;gcpdialogflowcxintents
// +kubebuilder:subresource:status

// DialogflowCXIntent is the Schema for the dialogflowcx API
// +k8s:openapi-gen=true
type DialogflowCXIntent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowCXIntentSpec   `json:"spec,omitempty"`
	Status DialogflowCXIntentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DialogflowCXIntentList contains a list of DialogflowCXIntent
type DialogflowCXIntentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowCXIntent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowCXIntent{}, &DialogflowCXIntentList{})
}
