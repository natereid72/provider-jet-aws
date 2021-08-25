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
// +groupName=codestarnotifications.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/codestarnotifications/v1alpha1"
)

type CodestarnotificationsNotificationRuleObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type CodestarnotificationsNotificationRuleParameters struct {
	DetailType string `json:"detailType" tf:"detail_type"`

	EventTypeIds []string `json:"eventTypeIds" tf:"event_type_ids"`

	Name string `json:"name" tf:"name"`

	Resource string `json:"resource" tf:"resource"`

	Status *string `json:"status,omitempty" tf:"status"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Target []TargetParameters `json:"target,omitempty" tf:"target"`
}

type TargetObservation struct {
	Status string `json:"status" tf:"status"`
}

type TargetParameters struct {
	Address string `json:"address" tf:"address"`

	Type *string `json:"type,omitempty" tf:"type"`
}

// CodestarnotificationsNotificationRuleSpec defines the desired state of CodestarnotificationsNotificationRule
type CodestarnotificationsNotificationRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodestarnotificationsNotificationRuleParameters `json:"forProvider"`
}

// CodestarnotificationsNotificationRuleStatus defines the observed state of CodestarnotificationsNotificationRule.
type CodestarnotificationsNotificationRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodestarnotificationsNotificationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodestarnotificationsNotificationRule is the Schema for the CodestarnotificationsNotificationRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CodestarnotificationsNotificationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodestarnotificationsNotificationRuleSpec   `json:"spec"`
	Status            CodestarnotificationsNotificationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodestarnotificationsNotificationRuleList contains a list of CodestarnotificationsNotificationRules
type CodestarnotificationsNotificationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodestarnotificationsNotificationRule `json:"items"`
}

// Repository type metadata.
var (
	CodestarnotificationsNotificationRuleKind             = "CodestarnotificationsNotificationRule"
	CodestarnotificationsNotificationRuleGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CodestarnotificationsNotificationRuleKind}.String()
	CodestarnotificationsNotificationRuleKindAPIVersion   = CodestarnotificationsNotificationRuleKind + "." + v1alpha1.GroupVersion.String()
	CodestarnotificationsNotificationRuleGroupVersionKind = v1alpha1.GroupVersion.WithKind(CodestarnotificationsNotificationRuleKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CodestarnotificationsNotificationRule{}, &CodestarnotificationsNotificationRuleList{})
}