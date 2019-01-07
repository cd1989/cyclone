// +build !ignore_autogenerated

/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Argument) DeepCopyInto(out *Argument) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Argument.
func (in *Argument) DeepCopy() *Argument {
	if in == nil {
		return nil
	}
	out := new(Argument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgumentValue) DeepCopyInto(out *ArgumentValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgumentValue.
func (in *ArgumentValue) DeepCopy() *ArgumentValue {
	if in == nil {
		return nil
	}
	out := new(ArgumentValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactItem) DeepCopyInto(out *ArtifactItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactItem.
func (in *ArtifactItem) DeepCopy() *ArtifactItem {
	if in == nil {
		return nil
	}
	out := new(ArtifactItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRDSpec) DeepCopyInto(out *CRDSpec) {
	*out = *in
	out.StatusRef = in.StatusRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRDSpec.
func (in *CRDSpec) DeepCopy() *CRDSpec {
	if in == nil {
		return nil
	}
	out := new(CRDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRDWorkload) DeepCopyInto(out *CRDWorkload) {
	*out = *in
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRDWorkload.
func (in *CRDWorkload) DeepCopy() *CRDWorkload {
	if in == nil {
		return nil
	}
	out := new(CRDWorkload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Inputs) DeepCopyInto(out *Inputs) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceItem, len(*in))
		copy(*out, *in)
	}
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make([]ArgumentValue, len(*in))
		copy(*out, *in)
	}
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = make([]ArtifactItem, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Inputs.
func (in *Inputs) DeepCopy() *Inputs {
	if in == nil {
		return nil
	}
	out := new(Inputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyValue) DeepCopyInto(out *KeyValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyValue.
func (in *KeyValue) DeepCopy() *KeyValue {
	if in == nil {
		return nil
	}
	out := new(KeyValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Outputs) DeepCopyInto(out *Outputs) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceItem, len(*in))
		copy(*out, *in)
	}
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = make([]ArtifactItem, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Outputs.
func (in *Outputs) DeepCopy() *Outputs {
	if in == nil {
		return nil
	}
	out := new(Outputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterConfig) DeepCopyInto(out *ParameterConfig) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ParameterItem, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterConfig.
func (in *ParameterConfig) DeepCopy() *ParameterConfig {
	if in == nil {
		return nil
	}
	out := new(ParameterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterItem) DeepCopyInto(out *ParameterItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterItem.
func (in *ParameterItem) DeepCopy() *ParameterItem {
	if in == nil {
		return nil
	}
	out := new(ParameterItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistent) DeepCopyInto(out *Persistent) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistent.
func (in *Persistent) DeepCopy() *Persistent {
	if in == nil {
		return nil
	}
	out := new(Persistent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodInfo) DeepCopyInto(out *PodInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodInfo.
func (in *PodInfo) DeepCopy() *PodInfo {
	if in == nil {
		return nil
	}
	out := new(PodInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodWorkload) DeepCopyInto(out *PodWorkload) {
	*out = *in
	in.Inputs.DeepCopyInto(&out.Inputs)
	in.Outputs.DeepCopyInto(&out.Outputs)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodWorkload.
func (in *PodWorkload) DeepCopy() *PodWorkload {
	if in == nil {
		return nil
	}
	out := new(PodWorkload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Resource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceItem) DeepCopyInto(out *ResourceItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceItem.
func (in *ResourceItem) DeepCopy() *ResourceItem {
	if in == nil {
		return nil
	}
	out := new(ResourceItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceList) DeepCopyInto(out *ResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Resource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceList.
func (in *ResourceList) DeepCopy() *ResourceList {
	if in == nil {
		return nil
	}
	out := new(ResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
	if in.Persistent != nil {
		in, out := &in.Persistent, &out.Persistent
		*out = new(Persistent)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ParameterItem, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stage) DeepCopyInto(out *Stage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stage.
func (in *Stage) DeepCopy() *Stage {
	if in == nil {
		return nil
	}
	out := new(Stage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageItem) DeepCopyInto(out *StageItem) {
	*out = *in
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = make([]ArtifactItem, len(*in))
		copy(*out, *in)
	}
	if in.Depends != nil {
		in, out := &in.Depends, &out.Depends
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageItem.
func (in *StageItem) DeepCopy() *StageItem {
	if in == nil {
		return nil
	}
	out := new(StageItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageList) DeepCopyInto(out *StageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageList.
func (in *StageList) DeepCopy() *StageList {
	if in == nil {
		return nil
	}
	out := new(StageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageSpec) DeepCopyInto(out *StageSpec) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(TemplateRef)
		(*in).DeepCopyInto(*out)
	}
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(PodWorkload)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageSpec.
func (in *StageSpec) DeepCopy() *StageSpec {
	if in == nil {
		return nil
	}
	out := new(StageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageStatus) DeepCopyInto(out *StageStatus) {
	*out = *in
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(PodInfo)
		**out = **in
	}
	in.Status.DeepCopyInto(&out.Status)
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]KeyValue, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageStatus.
func (in *StageStatus) DeepCopy() *StageStatus {
	if in == nil {
		return nil
	}
	out := new(StageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageTemplate) DeepCopyInto(out *StageTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageTemplate.
func (in *StageTemplate) DeepCopy() *StageTemplate {
	if in == nil {
		return nil
	}
	out := new(StageTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StageTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageTemplateList) DeepCopyInto(out *StageTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StageTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageTemplateList.
func (in *StageTemplateList) DeepCopy() *StageTemplateList {
	if in == nil {
		return nil
	}
	out := new(StageTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StageTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageTemplateSpec) DeepCopyInto(out *StageTemplateSpec) {
	*out = *in
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make([]Argument, len(*in))
		copy(*out, *in)
	}
	in.Pod.DeepCopyInto(&out.Pod)
	out.CRD = in.CRD
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageTemplateSpec.
func (in *StageTemplateSpec) DeepCopy() *StageTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(StageTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusRef) DeepCopyInto(out *StatusRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusRef.
func (in *StatusRef) DeepCopy() *StatusRef {
	if in == nil {
		return nil
	}
	out := new(StatusRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateRef) DeepCopyInto(out *TemplateRef) {
	*out = *in
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make([]ArgumentValue, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateRef.
func (in *TemplateRef) DeepCopy() *TemplateRef {
	if in == nil {
		return nil
	}
	out := new(TemplateRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workflow) DeepCopyInto(out *Workflow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workflow.
func (in *Workflow) DeepCopy() *Workflow {
	if in == nil {
		return nil
	}
	out := new(Workflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workflow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowList) DeepCopyInto(out *WorkflowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workflow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowList.
func (in *WorkflowList) DeepCopy() *WorkflowList {
	if in == nil {
		return nil
	}
	out := new(WorkflowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowParam) DeepCopyInto(out *WorkflowParam) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowParam.
func (in *WorkflowParam) DeepCopy() *WorkflowParam {
	if in == nil {
		return nil
	}
	out := new(WorkflowParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowParam) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowParamList) DeepCopyInto(out *WorkflowParamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkflowParam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowParamList.
func (in *WorkflowParamList) DeepCopy() *WorkflowParamList {
	if in == nil {
		return nil
	}
	out := new(WorkflowParamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowParamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowParamSpec) DeepCopyInto(out *WorkflowParamSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowParamSpec.
func (in *WorkflowParamSpec) DeepCopy() *WorkflowParamSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowParamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowRun) DeepCopyInto(out *WorkflowRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowRun.
func (in *WorkflowRun) DeepCopy() *WorkflowRun {
	if in == nil {
		return nil
	}
	out := new(WorkflowRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowRunList) DeepCopyInto(out *WorkflowRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkflowRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowRunList.
func (in *WorkflowRunList) DeepCopy() *WorkflowRunList {
	if in == nil {
		return nil
	}
	out := new(WorkflowRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowRunSpec) DeepCopyInto(out *WorkflowRunSpec) {
	*out = *in
	if in.WorkflowRef != nil {
		in, out := &in.WorkflowRef, &out.WorkflowRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.StartStages != nil {
		in, out := &in.StartStages, &out.StartStages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EndStages != nil {
		in, out := &in.EndStages, &out.EndStages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ParameterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]ParameterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowRunSpec.
func (in *WorkflowRunSpec) DeepCopy() *WorkflowRunSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowRunStatus) DeepCopyInto(out *WorkflowRunStatus) {
	*out = *in
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make(map[string]*StageStatus, len(*in))
		for key, val := range *in {
			var outVal *StageStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(StageStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	in.Overall.DeepCopyInto(&out.Overall)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowRunStatus.
func (in *WorkflowRunStatus) DeepCopy() *WorkflowRunStatus {
	if in == nil {
		return nil
	}
	out := new(WorkflowRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowSpec) DeepCopyInto(out *WorkflowSpec) {
	*out = *in
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]StageItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowSpec.
func (in *WorkflowSpec) DeepCopy() *WorkflowSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTrigger) DeepCopyInto(out *WorkflowTrigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowTrigger.
func (in *WorkflowTrigger) DeepCopy() *WorkflowTrigger {
	if in == nil {
		return nil
	}
	out := new(WorkflowTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowTrigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTriggerList) DeepCopyInto(out *WorkflowTriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkflowTrigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowTriggerList.
func (in *WorkflowTriggerList) DeepCopy() *WorkflowTriggerList {
	if in == nil {
		return nil
	}
	out := new(WorkflowTriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowTriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTriggerSpec) DeepCopyInto(out *WorkflowTriggerSpec) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ParameterItem, len(*in))
		copy(*out, *in)
	}
	in.WorkflowRunSpec.DeepCopyInto(&out.WorkflowRunSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowTriggerSpec.
func (in *WorkflowTriggerSpec) DeepCopy() *WorkflowTriggerSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowTriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTriggerStatus) DeepCopyInto(out *WorkflowTriggerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowTriggerStatus.
func (in *WorkflowTriggerStatus) DeepCopy() *WorkflowTriggerStatus {
	if in == nil {
		return nil
	}
	out := new(WorkflowTriggerStatus)
	in.DeepCopyInto(out)
	return out
}
