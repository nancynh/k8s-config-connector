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

type InspecttemplateCloudStoragePath struct {
	/* A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt */
	// +optional
	Path *string `json:"path,omitempty"`
}

type InspecttemplateCustomInfoTypes struct {
	/* A list of phrases to detect as a CustomInfoType. */
	// +optional
	Dictionary *InspecttemplateDictionary `json:"dictionary,omitempty"`

	/* If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned. It still can be used for rules matching. Possible values: EXCLUSION_TYPE_UNSPECIFIED, EXCLUSION_TYPE_EXCLUDE */
	// +optional
	ExclusionType *string `json:"exclusionType,omitempty"`

	/* CustomInfoType can either be a new infoType, or an extension of built-in infoType, when the name matches one of existing infoTypes and that infoType is specified in `InspectContent.info_types` field. Specifying the latter adds findings to the one detected by the system. If built-in info type is not specified in `InspectContent.info_types` list then the name is treated as a custom info type. */
	// +optional
	InfoType *InspecttemplateInfoType `json:"infoType,omitempty"`

	/* Likelihood to return for this CustomInfoType. This base value can be altered by a detection rule if the finding meets the criteria specified by the rule. Defaults to `VERY_LIKELY` if not specified. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY */
	// +optional
	Likelihood *string `json:"likelihood,omitempty"`

	/* Regular expression based CustomInfoType. */
	// +optional
	Regex *InspecttemplateRegex `json:"regex,omitempty"`

	/* Load an existing `StoredInfoType` resource for use in `InspectDataSource`. Not currently supported in `InspectContent`. */
	// +optional
	StoredType *InspecttemplateStoredType `json:"storedType,omitempty"`

	/* Message for detecting output from deidentification transformations that support reversing. */
	// +optional
	SurrogateType *InspecttemplateSurrogateType `json:"surrogateType,omitempty"`
}

type InspecttemplateDictionary struct {
	/* Newline-delimited file of words in Cloud Storage. Only a single file is accepted. */
	// +optional
	CloudStoragePath *InspecttemplateCloudStoragePath `json:"cloudStoragePath,omitempty"`

	/* List of words or phrases to search for. */
	// +optional
	WordList *InspecttemplateWordList `json:"wordList,omitempty"`
}

type InspecttemplateExcludeInfoTypes struct {
	/* InfoType list in ExclusionRule rule drops a finding when it overlaps or contained within with a finding of an infoType from this list. For example, for `InspectionRuleSet.info_types` containing "PHONE_NUMBER"` and `exclusion_rule` containing `exclude_info_types.info_types` with "EMAIL_ADDRESS" the phone number findings are dropped if they overlap with EMAIL_ADDRESS finding. That leads to "555-222-2222@example.org" to generate only a single finding, namely email address. */
	// +optional
	InfoTypes []InspecttemplateInfoTypes `json:"infoTypes,omitempty"`
}

type InspecttemplateExclusionRule struct {
	/* Dictionary which defines the rule. */
	// +optional
	Dictionary *InspecttemplateDictionary `json:"dictionary,omitempty"`

	/* Set of infoTypes for which findings would affect this rule. */
	// +optional
	ExcludeInfoTypes *InspecttemplateExcludeInfoTypes `json:"excludeInfoTypes,omitempty"`

	/* How the rule is applied, see MatchingType documentation for details. Possible values: MATCHING_TYPE_UNSPECIFIED, MATCHING_TYPE_FULL_MATCH, MATCHING_TYPE_PARTIAL_MATCH, MATCHING_TYPE_INVERSE_MATCH */
	// +optional
	MatchingType *string `json:"matchingType,omitempty"`

	/* Regular expression which defines the rule. */
	// +optional
	Regex *InspecttemplateRegex `json:"regex,omitempty"`
}

type InspecttemplateHotwordRegex struct {
	/* The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included. */
	// +optional
	GroupIndexes []int `json:"groupIndexes,omitempty"`

	/* Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub. */
	// +optional
	Pattern *string `json:"pattern,omitempty"`
}

type InspecttemplateHotwordRule struct {
	/* Regular expression pattern defining what qualifies as a hotword. */
	// +optional
	HotwordRegex *InspecttemplateHotwordRegex `json:"hotwordRegex,omitempty"`

	/* Likelihood adjustment to apply to all matching findings. */
	// +optional
	LikelihoodAdjustment *InspecttemplateLikelihoodAdjustment `json:"likelihoodAdjustment,omitempty"`

	/* Proximity of the finding within which the entire hotword must reside. The total length of the window cannot exceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be used to match substrings of the finding itself. For example, the certainty of a phone number regex "(d{3}) d{3}-d{4}" could be adjusted upwards if the area code is known to be the local area code of a company office using the hotword regex "(xxx)", where "xxx" is the area code in question. */
	// +optional
	Proximity *InspecttemplateProximity `json:"proximity,omitempty"`
}

type InspecttemplateInfoType struct {
	/* Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type InspecttemplateInfoTypes struct {
	/* Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type InspecttemplateInspectConfig struct {
	/* List of options defining data content to scan. If empty, text, images, and other content will be included. */
	// +optional
	ContentOptions []string `json:"contentOptions,omitempty"`

	/* CustomInfoTypes provided by the user. See https://cloud.google.com/dlp/docs/creating-custom-infotypes to learn more. */
	// +optional
	CustomInfoTypes []InspecttemplateCustomInfoTypes `json:"customInfoTypes,omitempty"`

	/* When true, excludes type information of the findings. */
	// +optional
	ExcludeInfoTypes *bool `json:"excludeInfoTypes,omitempty"`

	/* When true, a contextual quote from the data that triggered a finding is included in the response; see Finding.quote. */
	// +optional
	IncludeQuote *bool `json:"includeQuote,omitempty"`

	/* Restricts what info_types to look for. The values must correspond to InfoType values returned by ListInfoTypes or listed at https://cloud.google.com/dlp/docs/infotypes-reference. When no InfoTypes or CustomInfoTypes are specified in a request, the system may automatically choose what detectors to run. By default this may be all types, but may change over time as detectors are updated. If you need precise control and predictability as to what detectors are run you should specify specific InfoTypes listed in the reference, otherwise a default list will be used, which may change over time. */
	// +optional
	InfoTypes []InspecttemplateInfoTypes `json:"infoTypes,omitempty"`

	/* Configuration to control the number of findings returned. */
	// +optional
	Limits *InspecttemplateLimits `json:"limits,omitempty"`

	/* Only returns findings equal or above this threshold. The default is POSSIBLE. See https://cloud.google.com/dlp/docs/likelihood to learn more. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY */
	// +optional
	MinLikelihood *string `json:"minLikelihood,omitempty"`

	/* Set of rules to apply to the findings for this InspectConfig. Exclusion rules, contained in the set are executed in the end, other rules are executed in the order they are specified for each info type. */
	// +optional
	RuleSet []InspecttemplateRuleSet `json:"ruleSet,omitempty"`
}

type InspecttemplateLikelihoodAdjustment struct {
	/* Set the likelihood of a finding to a fixed value. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY */
	// +optional
	FixedLikelihood *string `json:"fixedLikelihood,omitempty"`

	/* Increase or decrease the likelihood by the specified number of levels. For example, if a finding would be `POSSIBLE` without the detection rule and `relative_likelihood` is 1, then it is upgraded to `LIKELY`, while a value of -1 would downgrade it to `UNLIKELY`. Likelihood may never drop below `VERY_UNLIKELY` or exceed `VERY_LIKELY`, so applying an adjustment of 1 followed by an adjustment of -1 when base likelihood is `VERY_LIKELY` will result in a final likelihood of `LIKELY`. */
	// +optional
	RelativeLikelihood *int `json:"relativeLikelihood,omitempty"`
}

type InspecttemplateLimits struct {
	/* Configuration of findings limit given for specified infoTypes. */
	// +optional
	MaxFindingsPerInfoType []InspecttemplateMaxFindingsPerInfoType `json:"maxFindingsPerInfoType,omitempty"`

	/* Max number of findings that will be returned for each item scanned. When set within `InspectJobConfig`, the maximum returned is 2000 regardless if this is set higher. When set within `InspectContentRequest`, this field is ignored. */
	// +optional
	MaxFindingsPerItem *int `json:"maxFindingsPerItem,omitempty"`

	/* Max number of findings that will be returned per request/job. When set within `InspectContentRequest`, the maximum returned is 2000 regardless if this is set higher. */
	// +optional
	MaxFindingsPerRequest *int `json:"maxFindingsPerRequest,omitempty"`
}

type InspecttemplateMaxFindingsPerInfoType struct {
	/* Type of information the findings limit applies to. Only one limit per info_type should be provided. If InfoTypeLimit does not have an info_type, the DLP API applies the limit against all info_types that are found but not specified in another InfoTypeLimit. */
	// +optional
	InfoType *InspecttemplateInfoType `json:"infoType,omitempty"`

	/* Max findings limit for the given infoType. */
	// +optional
	MaxFindings *int `json:"maxFindings,omitempty"`
}

type InspecttemplateProximity struct {
	/* Number of characters after the finding to consider. */
	// +optional
	WindowAfter *int `json:"windowAfter,omitempty"`

	/* Number of characters before the finding to consider. */
	// +optional
	WindowBefore *int `json:"windowBefore,omitempty"`
}

type InspecttemplateRegex struct {
	/* The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included. */
	// +optional
	GroupIndexes []int `json:"groupIndexes,omitempty"`

	/* Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub. */
	// +optional
	Pattern *string `json:"pattern,omitempty"`
}

type InspecttemplateRuleSet struct {
	/* List of infoTypes this rule set is applied to. */
	// +optional
	InfoTypes []InspecttemplateInfoTypes `json:"infoTypes,omitempty"`

	/* Set of rules to be applied to infoTypes. The rules are applied in order. */
	// +optional
	Rules []InspecttemplateRules `json:"rules,omitempty"`
}

type InspecttemplateRules struct {
	/* Exclusion rule. */
	// +optional
	ExclusionRule *InspecttemplateExclusionRule `json:"exclusionRule,omitempty"`

	/*  */
	// +optional
	HotwordRule *InspecttemplateHotwordRule `json:"hotwordRule,omitempty"`
}

type InspecttemplateStoredType struct {
	/* Timestamp indicating when the version of the `StoredInfoType` used for inspection was created. Output-only field, populated by the system. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/*  */
	// +optional
	NameRef *v1alpha1.ResourceRef `json:"nameRef,omitempty"`
}

type InspecttemplateSurrogateType struct {
}

type InspecttemplateWordList struct {
	/* Words or phrases defining the dictionary. The dictionary must contain at least one phrase and every phrase must contain at least 2 characters that are letters or digits. [required] */
	// +optional
	Words []string `json:"words,omitempty"`
}

type DLPInspectTemplateSpec struct {
	/* Short description (max 256 chars). */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Display name (max 256 chars). */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* The core content of the template. Configuration of the scanning process. */
	// +optional
	InspectConfig *InspecttemplateInspectConfig `json:"inspectConfig,omitempty"`

	/* Immutable. The location of the resource */
	// +optional
	Location *string `json:"location,omitempty"`

	/* Immutable. The Organization that this resource belongs to. Only one of [organizationRef, projectRef] may be specified. */
	// +optional
	OrganizationRef *v1alpha1.ResourceRef `json:"organizationRef,omitempty"`

	/* Immutable. The Project that this resource belongs to. Only one of [organizationRef, projectRef] may be specified. */
	// +optional
	ProjectRef *v1alpha1.ResourceRef `json:"projectRef,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type DLPInspectTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   DLPInspectTemplate's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The creation timestamp of an inspectTemplate. */
	CreateTime string `json:"createTime,omitempty"`
	/* Output only. The geographic location where this resource is stored. */
	LocationId string `json:"locationId,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. The last update timestamp of an inspectTemplate. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DLPInspectTemplate is the Schema for the dlp API
// +k8s:openapi-gen=true
type DLPInspectTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DLPInspectTemplateSpec   `json:"spec,omitempty"`
	Status DLPInspectTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DLPInspectTemplateList contains a list of DLPInspectTemplate
type DLPInspectTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DLPInspectTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DLPInspectTemplate{}, &DLPInspectTemplateList{})
}
