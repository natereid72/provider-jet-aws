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
func (in *AccountsObservation) DeepCopyInto(out *AccountsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountsObservation.
func (in *AccountsObservation) DeepCopy() *AccountsObservation {
	if in == nil {
		return nil
	}
	out := new(AccountsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountsParameters) DeepCopyInto(out *AccountsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountsParameters.
func (in *AccountsParameters) DeepCopy() *AccountsParameters {
	if in == nil {
		return nil
	}
	out := new(AccountsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganizationalUnit) DeepCopyInto(out *OrganizationsOrganizationalUnit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganizationalUnit.
func (in *OrganizationsOrganizationalUnit) DeepCopy() *OrganizationsOrganizationalUnit {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganizationalUnit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationsOrganizationalUnit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganizationalUnitList) DeepCopyInto(out *OrganizationsOrganizationalUnitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OrganizationsOrganizationalUnit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganizationalUnitList.
func (in *OrganizationsOrganizationalUnitList) DeepCopy() *OrganizationsOrganizationalUnitList {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganizationalUnitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationsOrganizationalUnitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganizationalUnitObservation) DeepCopyInto(out *OrganizationsOrganizationalUnitObservation) {
	*out = *in
	if in.Accounts != nil {
		in, out := &in.Accounts, &out.Accounts
		*out = make([]AccountsObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganizationalUnitObservation.
func (in *OrganizationsOrganizationalUnitObservation) DeepCopy() *OrganizationsOrganizationalUnitObservation {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganizationalUnitObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganizationalUnitParameters) DeepCopyInto(out *OrganizationsOrganizationalUnitParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganizationalUnitParameters.
func (in *OrganizationsOrganizationalUnitParameters) DeepCopy() *OrganizationsOrganizationalUnitParameters {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganizationalUnitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganizationalUnitSpec) DeepCopyInto(out *OrganizationsOrganizationalUnitSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganizationalUnitSpec.
func (in *OrganizationsOrganizationalUnitSpec) DeepCopy() *OrganizationsOrganizationalUnitSpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganizationalUnitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganizationalUnitStatus) DeepCopyInto(out *OrganizationsOrganizationalUnitStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganizationalUnitStatus.
func (in *OrganizationsOrganizationalUnitStatus) DeepCopy() *OrganizationsOrganizationalUnitStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganizationalUnitStatus)
	in.DeepCopyInto(out)
	return out
}