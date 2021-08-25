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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this CloudwatchCompositeAlarm.
func (mg *CloudwatchCompositeAlarm) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this CloudwatchCompositeAlarm.
func (mg *CloudwatchCompositeAlarm) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this CloudwatchCompositeAlarm.
func (mg *CloudwatchCompositeAlarm) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this CloudwatchCompositeAlarm.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *CloudwatchCompositeAlarm) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this CloudwatchCompositeAlarm.
func (mg *CloudwatchCompositeAlarm) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this CloudwatchCompositeAlarm.
func (mg *CloudwatchCompositeAlarm) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this CloudwatchCompositeAlarm.
func (mg *CloudwatchCompositeAlarm) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this CloudwatchCompositeAlarm.
func (mg *CloudwatchCompositeAlarm) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this CloudwatchCompositeAlarm.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *CloudwatchCompositeAlarm) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this CloudwatchCompositeAlarm.
func (mg *CloudwatchCompositeAlarm) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}