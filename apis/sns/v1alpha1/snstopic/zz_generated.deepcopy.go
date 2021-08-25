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
func (in *SnsTopic) DeepCopyInto(out *SnsTopic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopic.
func (in *SnsTopic) DeepCopy() *SnsTopic {
	if in == nil {
		return nil
	}
	out := new(SnsTopic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsTopic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicList) DeepCopyInto(out *SnsTopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnsTopic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicList.
func (in *SnsTopicList) DeepCopy() *SnsTopicList {
	if in == nil {
		return nil
	}
	out := new(SnsTopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsTopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicObservation) DeepCopyInto(out *SnsTopicObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicObservation.
func (in *SnsTopicObservation) DeepCopy() *SnsTopicObservation {
	if in == nil {
		return nil
	}
	out := new(SnsTopicObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicParameters) DeepCopyInto(out *SnsTopicParameters) {
	*out = *in
	if in.ApplicationFailureFeedbackRoleArn != nil {
		in, out := &in.ApplicationFailureFeedbackRoleArn, &out.ApplicationFailureFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.ApplicationSuccessFeedbackRoleArn != nil {
		in, out := &in.ApplicationSuccessFeedbackRoleArn, &out.ApplicationSuccessFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.ApplicationSuccessFeedbackSampleRate != nil {
		in, out := &in.ApplicationSuccessFeedbackSampleRate, &out.ApplicationSuccessFeedbackSampleRate
		*out = new(int64)
		**out = **in
	}
	if in.ContentBasedDeduplication != nil {
		in, out := &in.ContentBasedDeduplication, &out.ContentBasedDeduplication
		*out = new(bool)
		**out = **in
	}
	if in.DeliveryPolicy != nil {
		in, out := &in.DeliveryPolicy, &out.DeliveryPolicy
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FifoTopic != nil {
		in, out := &in.FifoTopic, &out.FifoTopic
		*out = new(bool)
		**out = **in
	}
	if in.FirehoseFailureFeedbackRoleArn != nil {
		in, out := &in.FirehoseFailureFeedbackRoleArn, &out.FirehoseFailureFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.FirehoseSuccessFeedbackRoleArn != nil {
		in, out := &in.FirehoseSuccessFeedbackRoleArn, &out.FirehoseSuccessFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.FirehoseSuccessFeedbackSampleRate != nil {
		in, out := &in.FirehoseSuccessFeedbackSampleRate, &out.FirehoseSuccessFeedbackSampleRate
		*out = new(int64)
		**out = **in
	}
	if in.HttpFailureFeedbackRoleArn != nil {
		in, out := &in.HttpFailureFeedbackRoleArn, &out.HttpFailureFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.HttpSuccessFeedbackRoleArn != nil {
		in, out := &in.HttpSuccessFeedbackRoleArn, &out.HttpSuccessFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.HttpSuccessFeedbackSampleRate != nil {
		in, out := &in.HttpSuccessFeedbackSampleRate, &out.HttpSuccessFeedbackSampleRate
		*out = new(int64)
		**out = **in
	}
	if in.KmsMasterKeyId != nil {
		in, out := &in.KmsMasterKeyId, &out.KmsMasterKeyId
		*out = new(string)
		**out = **in
	}
	if in.LambdaFailureFeedbackRoleArn != nil {
		in, out := &in.LambdaFailureFeedbackRoleArn, &out.LambdaFailureFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.LambdaSuccessFeedbackRoleArn != nil {
		in, out := &in.LambdaSuccessFeedbackRoleArn, &out.LambdaSuccessFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.LambdaSuccessFeedbackSampleRate != nil {
		in, out := &in.LambdaSuccessFeedbackSampleRate, &out.LambdaSuccessFeedbackSampleRate
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.SqsFailureFeedbackRoleArn != nil {
		in, out := &in.SqsFailureFeedbackRoleArn, &out.SqsFailureFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.SqsSuccessFeedbackRoleArn != nil {
		in, out := &in.SqsSuccessFeedbackRoleArn, &out.SqsSuccessFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.SqsSuccessFeedbackSampleRate != nil {
		in, out := &in.SqsSuccessFeedbackSampleRate, &out.SqsSuccessFeedbackSampleRate
		*out = new(int64)
		**out = **in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicParameters.
func (in *SnsTopicParameters) DeepCopy() *SnsTopicParameters {
	if in == nil {
		return nil
	}
	out := new(SnsTopicParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSpec) DeepCopyInto(out *SnsTopicSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSpec.
func (in *SnsTopicSpec) DeepCopy() *SnsTopicSpec {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicStatus) DeepCopyInto(out *SnsTopicStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicStatus.
func (in *SnsTopicStatus) DeepCopy() *SnsTopicStatus {
	if in == nil {
		return nil
	}
	out := new(SnsTopicStatus)
	in.DeepCopyInto(out)
	return out
}