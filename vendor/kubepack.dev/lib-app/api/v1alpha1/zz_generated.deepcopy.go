// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketFile) DeepCopyInto(out *BucketFile) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketFile.
func (in *BucketFile) DeepCopy() *BucketFile {
	if in == nil {
		return nil
	}
	out := new(BucketFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketFileOutput) DeepCopyInto(out *BucketFileOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketFileOutput.
func (in *BucketFileOutput) DeepCopy() *BucketFileOutput {
	if in == nil {
		return nil
	}
	out := new(BucketFileOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketJsonFile) DeepCopyInto(out *BucketJsonFile) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketJsonFile.
func (in *BucketJsonFile) DeepCopy() *BucketJsonFile {
	if in == nil {
		return nil
	}
	out := new(BucketJsonFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObject) DeepCopyInto(out *BucketObject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObject.
func (in *BucketObject) DeepCopy() *BucketObject {
	if in == nil {
		return nil
	}
	out := new(BucketObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartOrder) DeepCopyInto(out *ChartOrder) {
	*out = *in
	out.ChartRepoRef = in.ChartRepoRef
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartOrder.
func (in *ChartOrder) DeepCopy() *ChartOrder {
	if in == nil {
		return nil
	}
	out := new(ChartOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartTemplate) DeepCopyInto(out *ChartTemplate) {
	*out = *in
	out.ChartRef = in.ChartRef
	if in.CRDs != nil {
		in, out := &in.CRDs, &out.CRDs
		*out = make([]BucketJsonFile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Manifest != nil {
		in, out := &in.Manifest, &out.Manifest
		*out = new(BucketObject)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*unstructured.Unstructured, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartTemplate.
func (in *ChartTemplate) DeepCopy() *ChartTemplate {
	if in == nil {
		return nil
	}
	out := new(ChartTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartTemplateOutput) DeepCopyInto(out *ChartTemplateOutput) {
	*out = *in
	out.ChartRef = in.ChartRef
	if in.CRDs != nil {
		in, out := &in.CRDs, &out.CRDs
		*out = make([]BucketFileOutput, len(*in))
		copy(*out, *in)
	}
	if in.Manifest != nil {
		in, out := &in.Manifest, &out.Manifest
		*out = new(BucketObject)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartTemplateOutput.
func (in *ChartTemplateOutput) DeepCopy() *ChartTemplateOutput {
	if in == nil {
		return nil
	}
	out := new(ChartTemplateOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EditResourceOrder) DeepCopyInto(out *EditResourceOrder) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EditResourceOrder.
func (in *EditResourceOrder) DeepCopy() *EditResourceOrder {
	if in == nil {
		return nil
	}
	out := new(EditResourceOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EditorParameters) DeepCopyInto(out *EditorParameters) {
	*out = *in
	if in.ValuesPatch != nil {
		in, out := &in.ValuesPatch, &out.ValuesPatch
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EditorParameters.
func (in *EditorParameters) DeepCopy() *EditorParameters {
	if in == nil {
		return nil
	}
	out := new(EditorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EditorTemplate) DeepCopyInto(out *EditorTemplate) {
	*out = *in
	if in.Manifest != nil {
		in, out := &in.Manifest, &out.Manifest
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = (*in).DeepCopy()
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*unstructured.Unstructured, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EditorTemplate.
func (in *EditorTemplate) DeepCopy() *EditorTemplate {
	if in == nil {
		return nil
	}
	out := new(EditorTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	out.Resource = in.Resource
	out.Release = in.Release
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Model) DeepCopyInto(out *Model) {
	*out = *in
	out.Metadata = in.Metadata
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Model.
func (in *Model) DeepCopy() *Model {
	if in == nil {
		return nil
	}
	out := new(Model)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelMetadata) DeepCopyInto(out *ModelMetadata) {
	*out = *in
	out.Metadata = in.Metadata
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelMetadata.
func (in *ModelMetadata) DeepCopy() *ModelMetadata {
	if in == nil {
		return nil
	}
	out := new(ModelMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectModel) DeepCopyInto(out *ObjectModel) {
	*out = *in
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectModel.
func (in *ObjectModel) DeepCopy() *ObjectModel {
	if in == nil {
		return nil
	}
	out := new(ObjectModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceOutput) DeepCopyInto(out *ResourceOutput) {
	*out = *in
	if in.CRDs != nil {
		in, out := &in.CRDs, &out.CRDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceOutput.
func (in *ResourceOutput) DeepCopy() *ResourceOutput {
	if in == nil {
		return nil
	}
	out := new(ResourceOutput)
	in.DeepCopyInto(out)
	return out
}
