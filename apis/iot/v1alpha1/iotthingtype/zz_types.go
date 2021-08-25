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
// +groupName=iot.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/iot/v1alpha1"
)

type IotThingTypeObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type IotThingTypeParameters struct {
	Deprecated *bool `json:"deprecated,omitempty" tf:"deprecated"`

	Name string `json:"name" tf:"name"`

	Properties []PropertiesParameters `json:"properties,omitempty" tf:"properties"`
}

type PropertiesObservation struct {
}

type PropertiesParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	SearchableAttributes []string `json:"searchableAttributes,omitempty" tf:"searchable_attributes"`
}

// IotThingTypeSpec defines the desired state of IotThingType
type IotThingTypeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IotThingTypeParameters `json:"forProvider"`
}

// IotThingTypeStatus defines the observed state of IotThingType.
type IotThingTypeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IotThingTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IotThingType is the Schema for the IotThingTypes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IotThingType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotThingTypeSpec   `json:"spec"`
	Status            IotThingTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotThingTypeList contains a list of IotThingTypes
type IotThingTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotThingType `json:"items"`
}

// Repository type metadata.
var (
	IotThingTypeKind             = "IotThingType"
	IotThingTypeGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: IotThingTypeKind}.String()
	IotThingTypeKindAPIVersion   = IotThingTypeKind + "." + v1alpha1.GroupVersion.String()
	IotThingTypeGroupVersionKind = v1alpha1.GroupVersion.WithKind(IotThingTypeKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&IotThingType{}, &IotThingTypeList{})
}