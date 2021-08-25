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
func (in *ContentConfigObservation) DeepCopyInto(out *ContentConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentConfigObservation.
func (in *ContentConfigObservation) DeepCopy() *ContentConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ContentConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentConfigParameters) DeepCopyInto(out *ContentConfigParameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentConfigParameters.
func (in *ContentConfigParameters) DeepCopy() *ContentConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ContentConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentConfigPermissionsObservation) DeepCopyInto(out *ContentConfigPermissionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentConfigPermissionsObservation.
func (in *ContentConfigPermissionsObservation) DeepCopy() *ContentConfigPermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(ContentConfigPermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentConfigPermissionsParameters) DeepCopyInto(out *ContentConfigPermissionsParameters) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Grantee != nil {
		in, out := &in.Grantee, &out.Grantee
		*out = new(string)
		**out = **in
	}
	if in.GranteeType != nil {
		in, out := &in.GranteeType, &out.GranteeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentConfigPermissionsParameters.
func (in *ContentConfigPermissionsParameters) DeepCopy() *ContentConfigPermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(ContentConfigPermissionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipeline) DeepCopyInto(out *ElastictranscoderPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipeline.
func (in *ElastictranscoderPipeline) DeepCopy() *ElastictranscoderPipeline {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElastictranscoderPipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipelineList) DeepCopyInto(out *ElastictranscoderPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElastictranscoderPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipelineList.
func (in *ElastictranscoderPipelineList) DeepCopy() *ElastictranscoderPipelineList {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElastictranscoderPipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipelineObservation) DeepCopyInto(out *ElastictranscoderPipelineObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipelineObservation.
func (in *ElastictranscoderPipelineObservation) DeepCopy() *ElastictranscoderPipelineObservation {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipelineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipelineParameters) DeepCopyInto(out *ElastictranscoderPipelineParameters) {
	*out = *in
	if in.AwsKmsKeyArn != nil {
		in, out := &in.AwsKmsKeyArn, &out.AwsKmsKeyArn
		*out = new(string)
		**out = **in
	}
	if in.ContentConfig != nil {
		in, out := &in.ContentConfig, &out.ContentConfig
		*out = make([]ContentConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ContentConfigPermissions != nil {
		in, out := &in.ContentConfigPermissions, &out.ContentConfigPermissions
		*out = make([]ContentConfigPermissionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]NotificationsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OutputBucket != nil {
		in, out := &in.OutputBucket, &out.OutputBucket
		*out = new(string)
		**out = **in
	}
	if in.ThumbnailConfig != nil {
		in, out := &in.ThumbnailConfig, &out.ThumbnailConfig
		*out = make([]ThumbnailConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ThumbnailConfigPermissions != nil {
		in, out := &in.ThumbnailConfigPermissions, &out.ThumbnailConfigPermissions
		*out = make([]ThumbnailConfigPermissionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipelineParameters.
func (in *ElastictranscoderPipelineParameters) DeepCopy() *ElastictranscoderPipelineParameters {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipelineParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipelineSpec) DeepCopyInto(out *ElastictranscoderPipelineSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipelineSpec.
func (in *ElastictranscoderPipelineSpec) DeepCopy() *ElastictranscoderPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipelineStatus) DeepCopyInto(out *ElastictranscoderPipelineStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipelineStatus.
func (in *ElastictranscoderPipelineStatus) DeepCopy() *ElastictranscoderPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationsObservation) DeepCopyInto(out *NotificationsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationsObservation.
func (in *NotificationsObservation) DeepCopy() *NotificationsObservation {
	if in == nil {
		return nil
	}
	out := new(NotificationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationsParameters) DeepCopyInto(out *NotificationsParameters) {
	*out = *in
	if in.Completed != nil {
		in, out := &in.Completed, &out.Completed
		*out = new(string)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(string)
		**out = **in
	}
	if in.Progressing != nil {
		in, out := &in.Progressing, &out.Progressing
		*out = new(string)
		**out = **in
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationsParameters.
func (in *NotificationsParameters) DeepCopy() *NotificationsParameters {
	if in == nil {
		return nil
	}
	out := new(NotificationsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailConfigObservation) DeepCopyInto(out *ThumbnailConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailConfigObservation.
func (in *ThumbnailConfigObservation) DeepCopy() *ThumbnailConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ThumbnailConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailConfigParameters) DeepCopyInto(out *ThumbnailConfigParameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailConfigParameters.
func (in *ThumbnailConfigParameters) DeepCopy() *ThumbnailConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ThumbnailConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailConfigPermissionsObservation) DeepCopyInto(out *ThumbnailConfigPermissionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailConfigPermissionsObservation.
func (in *ThumbnailConfigPermissionsObservation) DeepCopy() *ThumbnailConfigPermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(ThumbnailConfigPermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailConfigPermissionsParameters) DeepCopyInto(out *ThumbnailConfigPermissionsParameters) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Grantee != nil {
		in, out := &in.Grantee, &out.Grantee
		*out = new(string)
		**out = **in
	}
	if in.GranteeType != nil {
		in, out := &in.GranteeType, &out.GranteeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailConfigPermissionsParameters.
func (in *ThumbnailConfigPermissionsParameters) DeepCopy() *ThumbnailConfigPermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(ThumbnailConfigPermissionsParameters)
	in.DeepCopyInto(out)
	return out
}