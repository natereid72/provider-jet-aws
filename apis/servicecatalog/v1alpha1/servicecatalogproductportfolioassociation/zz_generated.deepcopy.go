// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogProductPortfolioAssociation) DeepCopyInto(out *ServicecatalogProductPortfolioAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogProductPortfolioAssociation.
func (in *ServicecatalogProductPortfolioAssociation) DeepCopy() *ServicecatalogProductPortfolioAssociation {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogProductPortfolioAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicecatalogProductPortfolioAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogProductPortfolioAssociationList) DeepCopyInto(out *ServicecatalogProductPortfolioAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicecatalogProductPortfolioAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogProductPortfolioAssociationList.
func (in *ServicecatalogProductPortfolioAssociationList) DeepCopy() *ServicecatalogProductPortfolioAssociationList {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogProductPortfolioAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicecatalogProductPortfolioAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogProductPortfolioAssociationObservation) DeepCopyInto(out *ServicecatalogProductPortfolioAssociationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogProductPortfolioAssociationObservation.
func (in *ServicecatalogProductPortfolioAssociationObservation) DeepCopy() *ServicecatalogProductPortfolioAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogProductPortfolioAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogProductPortfolioAssociationParameters) DeepCopyInto(out *ServicecatalogProductPortfolioAssociationParameters) {
	*out = *in
	if in.AcceptLanguage != nil {
		in, out := &in.AcceptLanguage, &out.AcceptLanguage
		*out = new(string)
		**out = **in
	}
	if in.SourcePortfolioId != nil {
		in, out := &in.SourcePortfolioId, &out.SourcePortfolioId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogProductPortfolioAssociationParameters.
func (in *ServicecatalogProductPortfolioAssociationParameters) DeepCopy() *ServicecatalogProductPortfolioAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogProductPortfolioAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogProductPortfolioAssociationSpec) DeepCopyInto(out *ServicecatalogProductPortfolioAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogProductPortfolioAssociationSpec.
func (in *ServicecatalogProductPortfolioAssociationSpec) DeepCopy() *ServicecatalogProductPortfolioAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogProductPortfolioAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogProductPortfolioAssociationStatus) DeepCopyInto(out *ServicecatalogProductPortfolioAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogProductPortfolioAssociationStatus.
func (in *ServicecatalogProductPortfolioAssociationStatus) DeepCopy() *ServicecatalogProductPortfolioAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogProductPortfolioAssociationStatus)
	in.DeepCopyInto(out)
	return out
}