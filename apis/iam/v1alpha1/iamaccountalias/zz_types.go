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

type IamAccountAliasObservation struct {
}

type IamAccountAliasParameters struct {
	AccountAlias string `json:"accountAlias" tf:"account_alias"`
}

// IamAccountAliasSpec defines the desired state of IamAccountAlias
type IamAccountAliasSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamAccountAliasParameters `json:"forProvider"`
}

// IamAccountAliasStatus defines the observed state of IamAccountAlias.
type IamAccountAliasStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamAccountAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamAccountAlias is the Schema for the IamAccountAliass API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IamAccountAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamAccountAliasSpec   `json:"spec"`
	Status            IamAccountAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamAccountAliasList contains a list of IamAccountAliass
type IamAccountAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamAccountAlias `json:"items"`
}

// Repository type metadata.
var (
	IamAccountAliasKind             = "IamAccountAlias"
	IamAccountAliasGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: IamAccountAliasKind}.String()
	IamAccountAliasKindAPIVersion   = IamAccountAliasKind + "." + v1alpha1.GroupVersion.String()
	IamAccountAliasGroupVersionKind = v1alpha1.GroupVersion.WithKind(IamAccountAliasKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&IamAccountAlias{}, &IamAccountAliasList{})
}