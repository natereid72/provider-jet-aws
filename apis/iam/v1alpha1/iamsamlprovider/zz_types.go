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
// +groupName=iam.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1"
)

type IamSamlProviderObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ValidUntil string `json:"validUntil" tf:"valid_until"`
}

type IamSamlProviderParameters struct {
	Name string `json:"name" tf:"name"`

	SamlMetadataDocument string `json:"samlMetadataDocument" tf:"saml_metadata_document"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// IamSamlProviderSpec defines the desired state of IamSamlProvider
type IamSamlProviderSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamSamlProviderParameters `json:"forProvider"`
}

// IamSamlProviderStatus defines the observed state of IamSamlProvider.
type IamSamlProviderStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamSamlProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamSamlProvider is the Schema for the IamSamlProviders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IamSamlProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamSamlProviderSpec   `json:"spec"`
	Status            IamSamlProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamSamlProviderList contains a list of IamSamlProviders
type IamSamlProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamSamlProvider `json:"items"`
}

// Repository type metadata.
var (
	IamSamlProviderKind             = "IamSamlProvider"
	IamSamlProviderGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: IamSamlProviderKind}.String()
	IamSamlProviderKindAPIVersion   = IamSamlProviderKind + "." + v1alpha1.GroupVersion.String()
	IamSamlProviderGroupVersionKind = v1alpha1.GroupVersion.WithKind(IamSamlProviderKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&IamSamlProvider{}, &IamSamlProviderList{})
}