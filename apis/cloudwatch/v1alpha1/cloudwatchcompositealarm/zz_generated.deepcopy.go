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
func (in *CloudwatchCompositeAlarm) DeepCopyInto(out *CloudwatchCompositeAlarm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchCompositeAlarm.
func (in *CloudwatchCompositeAlarm) DeepCopy() *CloudwatchCompositeAlarm {
	if in == nil {
		return nil
	}
	out := new(CloudwatchCompositeAlarm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudwatchCompositeAlarm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchCompositeAlarmList) DeepCopyInto(out *CloudwatchCompositeAlarmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudwatchCompositeAlarm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchCompositeAlarmList.
func (in *CloudwatchCompositeAlarmList) DeepCopy() *CloudwatchCompositeAlarmList {
	if in == nil {
		return nil
	}
	out := new(CloudwatchCompositeAlarmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudwatchCompositeAlarmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchCompositeAlarmObservation) DeepCopyInto(out *CloudwatchCompositeAlarmObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchCompositeAlarmObservation.
func (in *CloudwatchCompositeAlarmObservation) DeepCopy() *CloudwatchCompositeAlarmObservation {
	if in == nil {
		return nil
	}
	out := new(CloudwatchCompositeAlarmObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchCompositeAlarmParameters) DeepCopyInto(out *CloudwatchCompositeAlarmParameters) {
	*out = *in
	if in.ActionsEnabled != nil {
		in, out := &in.ActionsEnabled, &out.ActionsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AlarmActions != nil {
		in, out := &in.AlarmActions, &out.AlarmActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AlarmDescription != nil {
		in, out := &in.AlarmDescription, &out.AlarmDescription
		*out = new(string)
		**out = **in
	}
	if in.InsufficientDataActions != nil {
		in, out := &in.InsufficientDataActions, &out.InsufficientDataActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OkActions != nil {
		in, out := &in.OkActions, &out.OkActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchCompositeAlarmParameters.
func (in *CloudwatchCompositeAlarmParameters) DeepCopy() *CloudwatchCompositeAlarmParameters {
	if in == nil {
		return nil
	}
	out := new(CloudwatchCompositeAlarmParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchCompositeAlarmSpec) DeepCopyInto(out *CloudwatchCompositeAlarmSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchCompositeAlarmSpec.
func (in *CloudwatchCompositeAlarmSpec) DeepCopy() *CloudwatchCompositeAlarmSpec {
	if in == nil {
		return nil
	}
	out := new(CloudwatchCompositeAlarmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchCompositeAlarmStatus) DeepCopyInto(out *CloudwatchCompositeAlarmStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchCompositeAlarmStatus.
func (in *CloudwatchCompositeAlarmStatus) DeepCopy() *CloudwatchCompositeAlarmStatus {
	if in == nil {
		return nil
	}
	out := new(CloudwatchCompositeAlarmStatus)
	in.DeepCopyInto(out)
	return out
}