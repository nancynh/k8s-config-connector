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

type DeidentifytemplateBucketingConfig struct {
	/* Set of buckets. Ranges must be non-overlapping. */
	// +optional
	Buckets []DeidentifytemplateBuckets `json:"buckets,omitempty"`
}

type DeidentifytemplateBuckets struct {
	/* Upper bound of the range, exclusive; type must match min. */
	// +optional
	Max *DeidentifytemplateMax `json:"max,omitempty"`

	/* Lower bound of the range, inclusive. Type should be the same as max if used. */
	// +optional
	Min *DeidentifytemplateMin `json:"min,omitempty"`

	/* Required. Replacement value for this bucket. */
	ReplacementValue DeidentifytemplateReplacementValue `json:"replacementValue"`
}

type DeidentifytemplateCharacterMaskConfig struct {
	/* When masking a string, items in this list will be skipped when replacing characters. For example, if the input string is `555-555-5555` and you instruct Cloud DLP to skip `-` and mask 5 characters with `*`, Cloud DLP returns `***-**5-5555`. */
	// +optional
	CharactersToIgnore []DeidentifytemplateCharactersToIgnore `json:"charactersToIgnore,omitempty"`

	/* Character to use to mask the sensitive values—for example, `*` for an alphabetic string such as a name, or `0` for a numeric string such as ZIP code or credit card number. This string must have a length of 1. If not supplied, this value defaults to `*` for strings, and `0` for digits. */
	// +optional
	MaskingCharacter *string `json:"maskingCharacter,omitempty"`

	/* Number of characters to mask. If not set, all matching chars will be masked. Skipped characters do not count towards this tally. */
	// +optional
	NumberToMask *int `json:"numberToMask,omitempty"`

	/* Mask characters in reverse order. For example, if `masking_character` is `0`, `number_to_mask` is `14`, and `reverse_order` is `false`, then the input string `1234-5678-9012-3456` is masked as `00000000000000-3456`. If `masking_character` is `*`, `number_to_mask` is `3`, and `reverse_order` is `true`, then the string `12345` is masked as `12***`. */
	// +optional
	ReverseOrder *bool `json:"reverseOrder,omitempty"`
}

type DeidentifytemplateCharactersToIgnore struct {
	/* Characters to not transform when masking. */
	// +optional
	CharactersToSkip *string `json:"charactersToSkip,omitempty"`

	/* Common characters to not transform when masking. Useful to avoid removing punctuation. Possible values: COMMON_CHARS_TO_IGNORE_UNSPECIFIED, NUMERIC, ALPHA_UPPER_CASE, ALPHA_LOWER_CASE, PUNCTUATION, WHITESPACE */
	// +optional
	CommonCharactersToIgnore *string `json:"commonCharactersToIgnore,omitempty"`
}

type DeidentifytemplateCondition struct {
	/* An expression. */
	// +optional
	Expressions *DeidentifytemplateExpressions `json:"expressions,omitempty"`
}

type DeidentifytemplateConditions struct {
	/* Required. Field within the record this condition is evaluated against. */
	Field DeidentifytemplateField `json:"field"`

	/* Required. Operator used to compare the field or infoType to the value. Possible values: LOGICAL_OPERATOR_UNSPECIFIED, AND */
	Operator string `json:"operator"`

	/* Value to compare against. [Mandatory, except for `EXISTS` tests.] */
	// +optional
	Value *DeidentifytemplateValue `json:"value,omitempty"`
}

type DeidentifytemplateContext struct {
	/* Name describing the field. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type DeidentifytemplateCryptoDeterministicConfig struct {
	/* A context may be used for higher security and maintaining referential integrity such that the same identifier in two different contexts will be given a distinct surrogate. The context is appended to plaintext value being encrypted. On decryption the provided context is validated against the value used during encryption. If a context was provided during encryption, same context must be provided during decryption as well. If the context is not set, plaintext would be used as is for encryption. If the context is set but: 1. there is no record present when transforming a given value or 2. the field is not present when transforming a given value, plaintext would be used as is for encryption. Note that case (1) is expected when an `InfoTypeTransformation` is applied to both structured and non-structured `ContentItem`s. */
	// +optional
	Context *DeidentifytemplateContext `json:"context,omitempty"`

	/* The key used by the encryption function. For deterministic encryption using AES-SIV, the provided key is internally expanded to 64 bytes prior to use. */
	// +optional
	CryptoKey *DeidentifytemplateCryptoKey `json:"cryptoKey,omitempty"`

	/* The custom info type to annotate the surrogate with. This annotation will be applied to the surrogate by prefixing it with the name of the custom info type followed by the number of characters comprising the surrogate. The following scheme defines the format: {info type name}({surrogate character count}):{surrogate} For example, if the name of custom info type is 'MY_TOKEN_INFO_TYPE' and the surrogate is 'abc', the full replacement value will be: 'MY_TOKEN_INFO_TYPE(3):abc' This annotation identifies the surrogate when inspecting content using the custom info type 'Surrogate'. This facilitates reversal of the surrogate when it occurs in free text. Note: For record transformations where the entire cell in a table is being transformed, surrogates are not mandatory. Surrogates are used to denote the location of the token and are necessary for re-identification in free form text. In order for inspection to work properly, the name of this info type must not occur naturally anywhere in your data; otherwise, inspection may either - reverse a surrogate that does not correspond to an actual identifier - be unable to parse the surrogate and result in an error Therefore, choose your custom info type name carefully after considering what your data looks like. One way to select a name that has a high chance of yielding reliable detection is to include one or more unicode characters that are highly improbable to exist in your data. For example, assuming your data is entered from a regular ASCII keyboard, the symbol with the hex code point 29DD might be used like so: ⧝MY_TOKEN_TYPE. */
	// +optional
	SurrogateInfoType *DeidentifytemplateSurrogateInfoType `json:"surrogateInfoType,omitempty"`
}

type DeidentifytemplateCryptoHashConfig struct {
	/* The key used by the hash function. */
	// +optional
	CryptoKey *DeidentifytemplateCryptoKey `json:"cryptoKey,omitempty"`
}

type DeidentifytemplateCryptoKey struct {
	/* Key wrapped using Cloud KMS */
	// +optional
	KmsWrapped *DeidentifytemplateKmsWrapped `json:"kmsWrapped,omitempty"`

	/* Transient crypto key */
	// +optional
	Transient *DeidentifytemplateTransient `json:"transient,omitempty"`

	/* Unwrapped crypto key */
	// +optional
	Unwrapped *DeidentifytemplateUnwrapped `json:"unwrapped,omitempty"`
}

type DeidentifytemplateCryptoReplaceFfxFpeConfig struct {
	/* Common alphabets. Possible values: FFX_COMMON_NATIVE_ALPHABET_UNSPECIFIED, NUMERIC, HEXADECIMAL, UPPER_CASE_ALPHA_NUMERIC, ALPHA_NUMERIC */
	// +optional
	CommonAlphabet *string `json:"commonAlphabet,omitempty"`

	/* The 'tweak', a context may be used for higher security since the same identifier in two different contexts won't be given the same surrogate. If the context is not set, a default tweak will be used. If the context is set but: 1. there is no record present when transforming a given value or 1. the field is not present when transforming a given value, a default tweak will be used. Note that case (1) is expected when an `InfoTypeTransformation` is applied to both structured and non-structured `ContentItem`s. Currently, the referenced field may be of value type integer or string. The tweak is constructed as a sequence of bytes in big endian byte order such that: - a 64 bit integer is encoded followed by a single byte of value 1 - a string is encoded in UTF-8 format followed by a single byte of value 2 */
	// +optional
	Context *DeidentifytemplateContext `json:"context,omitempty"`

	/* Required. The key used by the encryption algorithm. */
	CryptoKey DeidentifytemplateCryptoKey `json:"cryptoKey"`

	/* This is supported by mapping these to the alphanumeric characters that the FFX mode natively supports. This happens before/after encryption/decryption. Each character listed must appear only once. Number of characters must be in the range [2, 95]. This must be encoded as ASCII. The order of characters does not matter. The full list of allowed characters is: ``0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ~`!@#$%^&*()_-+={[}]|:;"'<,>.?/`` */
	// +optional
	CustomAlphabet *string `json:"customAlphabet,omitempty"`

	/* The native way to select the alphabet. Must be in the range [2, 95]. */
	// +optional
	Radix *int `json:"radix,omitempty"`

	/* The custom infoType to annotate the surrogate with. This annotation will be applied to the surrogate by prefixing it with the name of the custom infoType followed by the number of characters comprising the surrogate. The following scheme defines the format: info_type_name(surrogate_character_count):surrogate For example, if the name of custom infoType is 'MY_TOKEN_INFO_TYPE' and the surrogate is 'abc', the full replacement value will be: 'MY_TOKEN_INFO_TYPE(3):abc' This annotation identifies the surrogate when inspecting content using the custom infoType [`SurrogateType`](https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#surrogatetype). This facilitates reversal of the surrogate when it occurs in free text. In order for inspection to work properly, the name of this infoType must not occur naturally anywhere in your data; otherwise, inspection may find a surrogate that does not correspond to an actual identifier. Therefore, choose your custom infoType name carefully after considering what your data looks like. One way to select a name that has a high chance of yielding reliable detection is to include one or more unicode characters that are highly improbable to exist in your data. For example, assuming your data is entered from a regular ASCII keyboard, the symbol with the hex code point 29DD might be used like so: ⧝MY_TOKEN_TYPE */
	// +optional
	SurrogateInfoType *DeidentifytemplateSurrogateInfoType `json:"surrogateInfoType,omitempty"`
}

type DeidentifytemplateDateShiftConfig struct {
	/* Points to the field that contains the context, for example, an entity id. If set, must also set cryptoKey. If set, shift will be consistent for the given context. */
	// +optional
	Context *DeidentifytemplateContext `json:"context,omitempty"`

	/* Causes the shift to be computed based on this key and the context. This results in the same shift for the same context and crypto_key. If set, must also set context. Can only be applied to table items. */
	// +optional
	CryptoKey *DeidentifytemplateCryptoKey `json:"cryptoKey,omitempty"`

	/* Required. For example, -5 means shift date to at most 5 days back in the past. */
	LowerBoundDays int `json:"lowerBoundDays"`

	/* Required. Range of shift in days. Actual shift will be selected at random within this range (inclusive ends). Negative means shift to earlier in time. Must not be more than 365250 days (1000 years) each direction. For example, 3 means shift date to at most 3 days into the future. */
	UpperBoundDays int `json:"upperBoundDays"`
}

type DeidentifytemplateDateValue struct {
	/* Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant. */
	// +optional
	Day *int `json:"day,omitempty"`

	/* Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day. */
	// +optional
	Month *int `json:"month,omitempty"`

	/* Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year. */
	// +optional
	Year *int `json:"year,omitempty"`
}

type DeidentifytemplateDeidentifyConfig struct {
	/* Treat the dataset as free-form text and apply the same free text transformation everywhere. */
	// +optional
	InfoTypeTransformations *DeidentifytemplateInfoTypeTransformations `json:"infoTypeTransformations,omitempty"`

	/* Treat the dataset as structured. Transformations can be applied to specific locations within structured datasets, such as transforming a column within a table. */
	// +optional
	RecordTransformations *DeidentifytemplateRecordTransformations `json:"recordTransformations,omitempty"`

	/* Mode for handling transformation errors. If left unspecified, the default mode is `TransformationErrorHandling.ThrowError`. */
	// +optional
	TransformationErrorHandling *DeidentifytemplateTransformationErrorHandling `json:"transformationErrorHandling,omitempty"`
}

type DeidentifytemplateExpressions struct {
	/* Conditions to apply to the expression. */
	// +optional
	Conditions *DeidentifytemplateConditions `json:"conditions,omitempty"`

	/* The operator to apply to the result of conditions. Default and currently only supported value is `AND`. Possible values: LOGICAL_OPERATOR_UNSPECIFIED, AND */
	// +optional
	LogicalOperator *string `json:"logicalOperator,omitempty"`
}

type DeidentifytemplateField struct {
	/* Name describing the field. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type DeidentifytemplateFieldTransformations struct {
	/* Only apply the transformation if the condition evaluates to true for the given `RecordCondition`. The conditions are allowed to reference fields that are not used in the actual transformation. Example Use Cases: - Apply a different bucket transformation to an age column if the zip code column for the same record is within a specific range. - Redact a field if the date of birth field is greater than 85. */
	// +optional
	Condition *DeidentifytemplateCondition `json:"condition,omitempty"`

	/* Required. Input field(s) to apply the transformation to. When you have columns that reference their position within a list, omit the index from the FieldId. FieldId name matching ignores the index. For example, instead of "contact.nums[0].type", use "contact.nums.type". */
	Fields []DeidentifytemplateFields `json:"fields"`

	/* Treat the contents of the field as free text, and selectively transform content that matches an `InfoType`. */
	// +optional
	InfoTypeTransformations *DeidentifytemplateInfoTypeTransformations `json:"infoTypeTransformations,omitempty"`

	/* Apply the transformation to the entire field. */
	// +optional
	PrimitiveTransformation *DeidentifytemplatePrimitiveTransformation `json:"primitiveTransformation,omitempty"`
}

type DeidentifytemplateFields struct {
	/* Name describing the field. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type DeidentifytemplateFixedSizeBucketingConfig struct {
	/* Required. Size of each bucket (except for minimum and maximum buckets). So if `lower_bound` = 10, `upper_bound` = 89, and `bucket_size` = 10, then the following buckets would be used: -10, 10-20, 20-30, 30-40, 40-50, 50-60, 60-70, 70-80, 80-89, 89+. Precision up to 2 decimals works. */
	BucketSize float64 `json:"bucketSize"`

	/* Required. Lower bound value of buckets. All values less than `lower_bound` are grouped together into a single bucket; for example if `lower_bound` = 10, then all values less than 10 are replaced with the value "-10". */
	LowerBound DeidentifytemplateLowerBound `json:"lowerBound"`

	/* Required. Upper bound value of buckets. All values greater than upper_bound are grouped together into a single bucket; for example if `upper_bound` = 89, then all values greater than 89 are replaced with the value "89+". */
	UpperBound DeidentifytemplateUpperBound `json:"upperBound"`
}

type DeidentifytemplateInfoTypeTransformations struct {
	/* Required. Transformation for each infoType. Cannot specify more than one for a given infoType. */
	Transformations []DeidentifytemplateTransformations `json:"transformations"`
}

type DeidentifytemplateInfoTypes struct {
	/* Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type DeidentifytemplateKmsWrapped struct {
	CryptoKeyRef v1alpha1.ResourceRef `json:"cryptoKeyRef"`

	/* Required. The wrapped data crypto key. */
	WrappedKey string `json:"wrappedKey"`
}

type DeidentifytemplateLeaveUntransformed struct {
}

type DeidentifytemplateLowerBound struct {
	/* boolean */
	// +optional
	BooleanValue *bool `json:"booleanValue,omitempty"`

	/* date */
	// +optional
	DateValue *DeidentifytemplateDateValue `json:"dateValue,omitempty"`

	/* day of week Possible values: DAY_OF_WEEK_UNSPECIFIED, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY */
	// +optional
	DayOfWeekValue *string `json:"dayOfWeekValue,omitempty"`

	/* float */
	// +optional
	FloatValue *float64 `json:"floatValue,omitempty"`

	/* integer */
	// +optional
	IntegerValue *int `json:"integerValue,omitempty"`

	/* string */
	// +optional
	StringValue *string `json:"stringValue,omitempty"`

	/* time of day */
	// +optional
	TimeValue *DeidentifytemplateTimeValue `json:"timeValue,omitempty"`

	/* timestamp */
	// +optional
	TimestampValue *string `json:"timestampValue,omitempty"`
}

type DeidentifytemplateMax struct {
	/* boolean */
	// +optional
	BooleanValue *bool `json:"booleanValue,omitempty"`

	/* date */
	// +optional
	DateValue *DeidentifytemplateDateValue `json:"dateValue,omitempty"`

	/* day of week Possible values: DAY_OF_WEEK_UNSPECIFIED, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY */
	// +optional
	DayOfWeekValue *string `json:"dayOfWeekValue,omitempty"`

	/* float */
	// +optional
	FloatValue *float64 `json:"floatValue,omitempty"`

	/* integer */
	// +optional
	IntegerValue *int `json:"integerValue,omitempty"`

	/* string */
	// +optional
	StringValue *string `json:"stringValue,omitempty"`

	/* time of day */
	// +optional
	TimeValue *DeidentifytemplateTimeValue `json:"timeValue,omitempty"`

	/* timestamp */
	// +optional
	TimestampValue *string `json:"timestampValue,omitempty"`
}

type DeidentifytemplateMin struct {
	/* boolean */
	// +optional
	BooleanValue *bool `json:"booleanValue,omitempty"`

	/* date */
	// +optional
	DateValue *DeidentifytemplateDateValue `json:"dateValue,omitempty"`

	/* day of week Possible values: DAY_OF_WEEK_UNSPECIFIED, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY */
	// +optional
	DayOfWeekValue *string `json:"dayOfWeekValue,omitempty"`

	/* float */
	// +optional
	FloatValue *float64 `json:"floatValue,omitempty"`

	/* integer */
	// +optional
	IntegerValue *int `json:"integerValue,omitempty"`

	/* string */
	// +optional
	StringValue *string `json:"stringValue,omitempty"`

	/* time of day */
	// +optional
	TimeValue *DeidentifytemplateTimeValue `json:"timeValue,omitempty"`

	/* timestamp */
	// +optional
	TimestampValue *string `json:"timestampValue,omitempty"`
}

type DeidentifytemplateNewValue struct {
	/* boolean */
	// +optional
	BooleanValue *bool `json:"booleanValue,omitempty"`

	/* date */
	// +optional
	DateValue *DeidentifytemplateDateValue `json:"dateValue,omitempty"`

	/* day of week Possible values: DAY_OF_WEEK_UNSPECIFIED, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY */
	// +optional
	DayOfWeekValue *string `json:"dayOfWeekValue,omitempty"`

	/* float */
	// +optional
	FloatValue *float64 `json:"floatValue,omitempty"`

	/* integer */
	// +optional
	IntegerValue *int `json:"integerValue,omitempty"`

	/* string */
	// +optional
	StringValue *string `json:"stringValue,omitempty"`

	/* time of day */
	// +optional
	TimeValue *DeidentifytemplateTimeValue `json:"timeValue,omitempty"`

	/* timestamp */
	// +optional
	TimestampValue *string `json:"timestampValue,omitempty"`
}

type DeidentifytemplatePrimitiveTransformation struct {
	/* Bucketing */
	// +optional
	BucketingConfig *DeidentifytemplateBucketingConfig `json:"bucketingConfig,omitempty"`

	/* Mask */
	// +optional
	CharacterMaskConfig *DeidentifytemplateCharacterMaskConfig `json:"characterMaskConfig,omitempty"`

	/* Deterministic Crypto */
	// +optional
	CryptoDeterministicConfig *DeidentifytemplateCryptoDeterministicConfig `json:"cryptoDeterministicConfig,omitempty"`

	/* Crypto */
	// +optional
	CryptoHashConfig *DeidentifytemplateCryptoHashConfig `json:"cryptoHashConfig,omitempty"`

	/* Ffx-Fpe */
	// +optional
	CryptoReplaceFfxFpeConfig *DeidentifytemplateCryptoReplaceFfxFpeConfig `json:"cryptoReplaceFfxFpeConfig,omitempty"`

	/* Date Shift */
	// +optional
	DateShiftConfig *DeidentifytemplateDateShiftConfig `json:"dateShiftConfig,omitempty"`

	/* Fixed size bucketing */
	// +optional
	FixedSizeBucketingConfig *DeidentifytemplateFixedSizeBucketingConfig `json:"fixedSizeBucketingConfig,omitempty"`

	/* Redact */
	// +optional
	RedactConfig *DeidentifytemplateRedactConfig `json:"redactConfig,omitempty"`

	/* Replace with a specified value. */
	// +optional
	ReplaceConfig *DeidentifytemplateReplaceConfig `json:"replaceConfig,omitempty"`

	/* Replace with infotype */
	// +optional
	ReplaceWithInfoTypeConfig *DeidentifytemplateReplaceWithInfoTypeConfig `json:"replaceWithInfoTypeConfig,omitempty"`

	/* Time extraction */
	// +optional
	TimePartConfig *DeidentifytemplateTimePartConfig `json:"timePartConfig,omitempty"`
}

type DeidentifytemplateRecordSuppressions struct {
	/* A condition that when it evaluates to true will result in the record being evaluated to be suppressed from the transformed content. */
	// +optional
	Condition *DeidentifytemplateCondition `json:"condition,omitempty"`
}

type DeidentifytemplateRecordTransformations struct {
	/* Transform the record by applying various field transformations. */
	// +optional
	FieldTransformations []DeidentifytemplateFieldTransformations `json:"fieldTransformations,omitempty"`

	/* Configuration defining which records get suppressed entirely. Records that match any suppression rule are omitted from the output. */
	// +optional
	RecordSuppressions []DeidentifytemplateRecordSuppressions `json:"recordSuppressions,omitempty"`
}

type DeidentifytemplateRedactConfig struct {
}

type DeidentifytemplateReplaceConfig struct {
	/* Value to replace it with. */
	// +optional
	NewValue *DeidentifytemplateNewValue `json:"newValue,omitempty"`
}

type DeidentifytemplateReplaceWithInfoTypeConfig struct {
}

type DeidentifytemplateReplacementValue struct {
	/* boolean */
	// +optional
	BooleanValue *bool `json:"booleanValue,omitempty"`

	/* date */
	// +optional
	DateValue *DeidentifytemplateDateValue `json:"dateValue,omitempty"`

	/* day of week Possible values: DAY_OF_WEEK_UNSPECIFIED, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY */
	// +optional
	DayOfWeekValue *string `json:"dayOfWeekValue,omitempty"`

	/* float */
	// +optional
	FloatValue *float64 `json:"floatValue,omitempty"`

	/* integer */
	// +optional
	IntegerValue *int `json:"integerValue,omitempty"`

	/* string */
	// +optional
	StringValue *string `json:"stringValue,omitempty"`

	/* time of day */
	// +optional
	TimeValue *DeidentifytemplateTimeValue `json:"timeValue,omitempty"`

	/* timestamp */
	// +optional
	TimestampValue *string `json:"timestampValue,omitempty"`
}

type DeidentifytemplateSurrogateInfoType struct {
	/* Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type DeidentifytemplateThrowError struct {
}

type DeidentifytemplateTimePartConfig struct {
	/* The part of the time to keep. Possible values: TIME_PART_UNSPECIFIED, YEAR, MONTH, DAY_OF_MONTH, DAY_OF_WEEK, WEEK_OF_YEAR, HOUR_OF_DAY */
	// +optional
	PartToExtract *string `json:"partToExtract,omitempty"`
}

type DeidentifytemplateTimeValue struct {
	/* Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time. */
	// +optional
	Hours *int `json:"hours,omitempty"`

	/* Minutes of hour of day. Must be from 0 to 59. */
	// +optional
	Minutes *int `json:"minutes,omitempty"`

	/* Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999. */
	// +optional
	Nanos *int `json:"nanos,omitempty"`

	/* Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds. */
	// +optional
	Seconds *int `json:"seconds,omitempty"`
}

type DeidentifytemplateTransformationErrorHandling struct {
	/* Ignore errors */
	// +optional
	LeaveUntransformed *DeidentifytemplateLeaveUntransformed `json:"leaveUntransformed,omitempty"`

	/* Throw an error */
	// +optional
	ThrowError *DeidentifytemplateThrowError `json:"throwError,omitempty"`
}

type DeidentifytemplateTransformations struct {
	/* InfoTypes to apply the transformation to. An empty list will cause this transformation to apply to all findings that correspond to infoTypes that were requested in `InspectConfig`. */
	// +optional
	InfoTypes []DeidentifytemplateInfoTypes `json:"infoTypes,omitempty"`

	/* Required. Primitive transformation to apply to the infoType. */
	PrimitiveTransformation DeidentifytemplatePrimitiveTransformation `json:"primitiveTransformation"`
}

type DeidentifytemplateTransient struct {
	/* Required. Name of the key. This is an arbitrary string used to differentiate different keys. A unique key is generated per name: two separate `TransientCryptoKey` protos share the same generated key if their names are the same. When the data crypto key is generated, this name is not used in any way (repeating the api call will result in a different key being generated). */
	Name string `json:"name"`
}

type DeidentifytemplateUnwrapped struct {
	/* Required. A 128/192/256 bit key. */
	Key string `json:"key"`
}

type DeidentifytemplateUpperBound struct {
	/* boolean */
	// +optional
	BooleanValue *bool `json:"booleanValue,omitempty"`

	/* date */
	// +optional
	DateValue *DeidentifytemplateDateValue `json:"dateValue,omitempty"`

	/* day of week Possible values: DAY_OF_WEEK_UNSPECIFIED, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY */
	// +optional
	DayOfWeekValue *string `json:"dayOfWeekValue,omitempty"`

	/* float */
	// +optional
	FloatValue *float64 `json:"floatValue,omitempty"`

	/* integer */
	// +optional
	IntegerValue *int `json:"integerValue,omitempty"`

	/* string */
	// +optional
	StringValue *string `json:"stringValue,omitempty"`

	/* time of day */
	// +optional
	TimeValue *DeidentifytemplateTimeValue `json:"timeValue,omitempty"`

	/* timestamp */
	// +optional
	TimestampValue *string `json:"timestampValue,omitempty"`
}

type DeidentifytemplateValue struct {
	/* boolean */
	// +optional
	BooleanValue *bool `json:"booleanValue,omitempty"`

	/* date */
	// +optional
	DateValue *DeidentifytemplateDateValue `json:"dateValue,omitempty"`

	/* day of week Possible values: DAY_OF_WEEK_UNSPECIFIED, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY */
	// +optional
	DayOfWeekValue *string `json:"dayOfWeekValue,omitempty"`

	/* float */
	// +optional
	FloatValue *float64 `json:"floatValue,omitempty"`

	/* integer */
	// +optional
	IntegerValue *int `json:"integerValue,omitempty"`

	/* string */
	// +optional
	StringValue *string `json:"stringValue,omitempty"`

	/* time of day */
	// +optional
	TimeValue *DeidentifytemplateTimeValue `json:"timeValue,omitempty"`

	/* timestamp */
	// +optional
	TimestampValue *string `json:"timestampValue,omitempty"`
}

type DLPDeidentifyTemplateSpec struct {
	/* The core content of the template. */
	// +optional
	DeidentifyConfig *DeidentifytemplateDeidentifyConfig `json:"deidentifyConfig,omitempty"`

	/* Short description (max 256 chars). */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Display name (max 256 chars). */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

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

type DLPDeidentifyTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   DLPDeidentifyTemplate's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The creation timestamp of an inspectTemplate. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. The geographic location where this resource is stored. */
	// +optional
	LocationId *string `json:"locationId,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Output only. The last update timestamp of an inspectTemplate. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdlpdeidentifytemplate;gcpdlpdeidentifytemplates
// +kubebuilder:subresource:status

// DLPDeidentifyTemplate is the Schema for the dlp API
// +k8s:openapi-gen=true
type DLPDeidentifyTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DLPDeidentifyTemplateSpec   `json:"spec,omitempty"`
	Status DLPDeidentifyTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DLPDeidentifyTemplateList contains a list of DLPDeidentifyTemplate
type DLPDeidentifyTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DLPDeidentifyTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DLPDeidentifyTemplate{}, &DLPDeidentifyTemplateList{})
}
