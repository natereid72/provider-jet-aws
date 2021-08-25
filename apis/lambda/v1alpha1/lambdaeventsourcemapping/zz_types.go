/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

// +kubebuilder:object:generate=true
// +groupName=lambda.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/lambda/v1alpha1"
)

type DestinationConfigObservation struct {
}

type DestinationConfigParameters struct {
	OnFailure []OnFailureParameters `json:"onFailure,omitempty" tf:"on_failure"`
}

type LambdaEventSourceMappingObservation struct {
	FunctionArn string `json:"functionArn" tf:"function_arn"`

	LastModified string `json:"lastModified" tf:"last_modified"`

	LastProcessingResult string `json:"lastProcessingResult" tf:"last_processing_result"`

	State string `json:"state" tf:"state"`

	StateTransitionReason string `json:"stateTransitionReason" tf:"state_transition_reason"`

	Uuid string `json:"uuid" tf:"uuid"`
}

type LambdaEventSourceMappingParameters struct {
	BatchSize *int64 `json:"batchSize,omitempty" tf:"batch_size"`

	BisectBatchOnFunctionError *bool `json:"bisectBatchOnFunctionError,omitempty" tf:"bisect_batch_on_function_error"`

	DestinationConfig []DestinationConfigParameters `json:"destinationConfig,omitempty" tf:"destination_config"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	EventSourceArn *string `json:"eventSourceArn,omitempty" tf:"event_source_arn"`

	FunctionName string `json:"functionName" tf:"function_name"`

	FunctionResponseTypes []string `json:"functionResponseTypes,omitempty" tf:"function_response_types"`

	MaximumBatchingWindowInSeconds *int64 `json:"maximumBatchingWindowInSeconds,omitempty" tf:"maximum_batching_window_in_seconds"`

	MaximumRecordAgeInSeconds *int64 `json:"maximumRecordAgeInSeconds,omitempty" tf:"maximum_record_age_in_seconds"`

	MaximumRetryAttempts *int64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts"`

	ParallelizationFactor *int64 `json:"parallelizationFactor,omitempty" tf:"parallelization_factor"`

	Queues []string `json:"queues,omitempty" tf:"queues"`

	SelfManagedEventSource []SelfManagedEventSourceParameters `json:"selfManagedEventSource,omitempty" tf:"self_managed_event_source"`

	SourceAccessConfiguration []SourceAccessConfigurationParameters `json:"sourceAccessConfiguration,omitempty" tf:"source_access_configuration"`

	StartingPosition *string `json:"startingPosition,omitempty" tf:"starting_position"`

	StartingPositionTimestamp *string `json:"startingPositionTimestamp,omitempty" tf:"starting_position_timestamp"`

	Topics []string `json:"topics,omitempty" tf:"topics"`

	TumblingWindowInSeconds *int64 `json:"tumblingWindowInSeconds,omitempty" tf:"tumbling_window_in_seconds"`
}

type OnFailureObservation struct {
}

type OnFailureParameters struct {
	DestinationArn string `json:"destinationArn" tf:"destination_arn"`
}

type SelfManagedEventSourceObservation struct {
}

type SelfManagedEventSourceParameters struct {
	Endpoints map[string]string `json:"endpoints" tf:"endpoints"`
}

type SourceAccessConfigurationObservation struct {
}

type SourceAccessConfigurationParameters struct {
	Type string `json:"type" tf:"type"`

	Uri string `json:"uri" tf:"uri"`
}

// LambdaEventSourceMappingSpec defines the desired state of LambdaEventSourceMapping
type LambdaEventSourceMappingSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LambdaEventSourceMappingParameters `json:"forProvider"`
}

// LambdaEventSourceMappingStatus defines the observed state of LambdaEventSourceMapping.
type LambdaEventSourceMappingStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LambdaEventSourceMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaEventSourceMapping is the Schema for the LambdaEventSourceMappings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LambdaEventSourceMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaEventSourceMappingSpec   `json:"spec"`
	Status            LambdaEventSourceMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaEventSourceMappingList contains a list of LambdaEventSourceMappings
type LambdaEventSourceMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaEventSourceMapping `json:"items"`
}

// Repository type metadata.
var (
	LambdaEventSourceMappingKind             = "LambdaEventSourceMapping"
	LambdaEventSourceMappingGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: LambdaEventSourceMappingKind}.String()
	LambdaEventSourceMappingKindAPIVersion   = LambdaEventSourceMappingKind + "." + v1alpha1.GroupVersion.String()
	LambdaEventSourceMappingGroupVersionKind = v1alpha1.GroupVersion.WithKind(LambdaEventSourceMappingKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&LambdaEventSourceMapping{}, &LambdaEventSourceMappingList{})
}