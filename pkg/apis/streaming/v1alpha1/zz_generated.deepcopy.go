// +build !ignore_autogenerated

/*
Copyright 2019 the original author or authors.

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
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingReference) DeepCopyInto(out *BindingReference) {
	*out = *in
	out.MetadataRef = in.MetadataRef
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingReference.
func (in *BindingReference) DeepCopy() *BindingReference {
	if in == nil {
		return nil
	}
	out := new(BindingReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Build) DeepCopyInto(out *Build) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Build.
func (in *Build) DeepCopy() *Build {
	if in == nil {
		return nil
	}
	out := new(Build)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryProvider) DeepCopyInto(out *InMemoryProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryProvider.
func (in *InMemoryProvider) DeepCopy() *InMemoryProvider {
	if in == nil {
		return nil
	}
	out := new(InMemoryProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InMemoryProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryProviderList) DeepCopyInto(out *InMemoryProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InMemoryProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryProviderList.
func (in *InMemoryProviderList) DeepCopy() *InMemoryProviderList {
	if in == nil {
		return nil
	}
	out := new(InMemoryProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InMemoryProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryProviderSpec) DeepCopyInto(out *InMemoryProviderSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryProviderSpec.
func (in *InMemoryProviderSpec) DeepCopy() *InMemoryProviderSpec {
	if in == nil {
		return nil
	}
	out := new(InMemoryProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryProviderStatus) DeepCopyInto(out *InMemoryProviderStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.GatewayDeploymentRef != nil {
		in, out := &in.GatewayDeploymentRef, &out.GatewayDeploymentRef
		*out = (*in).DeepCopy()
	}
	if in.GatewayServiceRef != nil {
		in, out := &in.GatewayServiceRef, &out.GatewayServiceRef
		*out = (*in).DeepCopy()
	}
	if in.ProvisionerDeploymentRef != nil {
		in, out := &in.ProvisionerDeploymentRef, &out.ProvisionerDeploymentRef
		*out = (*in).DeepCopy()
	}
	if in.ProvisionerServiceRef != nil {
		in, out := &in.ProvisionerServiceRef, &out.ProvisionerServiceRef
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryProviderStatus.
func (in *InMemoryProviderStatus) DeepCopy() *InMemoryProviderStatus {
	if in == nil {
		return nil
	}
	out := new(InMemoryProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputStreamBinding) DeepCopyInto(out *InputStreamBinding) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputStreamBinding.
func (in *InputStreamBinding) DeepCopy() *InputStreamBinding {
	if in == nil {
		return nil
	}
	out := new(InputStreamBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaProvider) DeepCopyInto(out *KafkaProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaProvider.
func (in *KafkaProvider) DeepCopy() *KafkaProvider {
	if in == nil {
		return nil
	}
	out := new(KafkaProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaProviderList) DeepCopyInto(out *KafkaProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaProviderList.
func (in *KafkaProviderList) DeepCopy() *KafkaProviderList {
	if in == nil {
		return nil
	}
	out := new(KafkaProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaProviderSpec) DeepCopyInto(out *KafkaProviderSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaProviderSpec.
func (in *KafkaProviderSpec) DeepCopy() *KafkaProviderSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaProviderStatus) DeepCopyInto(out *KafkaProviderStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.GatewayDeploymentRef != nil {
		in, out := &in.GatewayDeploymentRef, &out.GatewayDeploymentRef
		*out = (*in).DeepCopy()
	}
	if in.GatewayServiceRef != nil {
		in, out := &in.GatewayServiceRef, &out.GatewayServiceRef
		*out = (*in).DeepCopy()
	}
	if in.ProvisionerDeploymentRef != nil {
		in, out := &in.ProvisionerDeploymentRef, &out.ProvisionerDeploymentRef
		*out = (*in).DeepCopy()
	}
	if in.ProvisionerServiceRef != nil {
		in, out := &in.ProvisionerServiceRef, &out.ProvisionerServiceRef
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaProviderStatus.
func (in *KafkaProviderStatus) DeepCopy() *KafkaProviderStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputStreamBinding) DeepCopyInto(out *OutputStreamBinding) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputStreamBinding.
func (in *OutputStreamBinding) DeepCopy() *OutputStreamBinding {
	if in == nil {
		return nil
	}
	out := new(OutputStreamBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Processor) DeepCopyInto(out *Processor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Processor.
func (in *Processor) DeepCopy() *Processor {
	if in == nil {
		return nil
	}
	out := new(Processor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Processor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessorList) DeepCopyInto(out *ProcessorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Processor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessorList.
func (in *ProcessorList) DeepCopy() *ProcessorList {
	if in == nil {
		return nil
	}
	out := new(ProcessorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProcessorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessorSpec) DeepCopyInto(out *ProcessorSpec) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(Build)
		**out = **in
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]InputStreamBinding, len(*in))
		copy(*out, *in)
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]OutputStreamBinding, len(*in))
		copy(*out, *in)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(v1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessorSpec.
func (in *ProcessorSpec) DeepCopy() *ProcessorSpec {
	if in == nil {
		return nil
	}
	out := new(ProcessorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessorStatus) DeepCopyInto(out *ProcessorStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.DeprecatedInputAddresses != nil {
		in, out := &in.DeprecatedInputAddresses, &out.DeprecatedInputAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeprecatedOutputAddresses != nil {
		in, out := &in.DeprecatedOutputAddresses, &out.DeprecatedOutputAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeprecatedOutputContentTypes != nil {
		in, out := &in.DeprecatedOutputContentTypes, &out.DeprecatedOutputContentTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeploymentRef != nil {
		in, out := &in.DeploymentRef, &out.DeploymentRef
		*out = (*in).DeepCopy()
	}
	if in.ScaledObjectRef != nil {
		in, out := &in.ScaledObjectRef, &out.ScaledObjectRef
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessorStatus.
func (in *ProcessorStatus) DeepCopy() *ProcessorStatus {
	if in == nil {
		return nil
	}
	out := new(ProcessorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarProvider) DeepCopyInto(out *PulsarProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarProvider.
func (in *PulsarProvider) DeepCopy() *PulsarProvider {
	if in == nil {
		return nil
	}
	out := new(PulsarProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarProviderList) DeepCopyInto(out *PulsarProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PulsarProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarProviderList.
func (in *PulsarProviderList) DeepCopy() *PulsarProviderList {
	if in == nil {
		return nil
	}
	out := new(PulsarProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarProviderSpec) DeepCopyInto(out *PulsarProviderSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarProviderSpec.
func (in *PulsarProviderSpec) DeepCopy() *PulsarProviderSpec {
	if in == nil {
		return nil
	}
	out := new(PulsarProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarProviderStatus) DeepCopyInto(out *PulsarProviderStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.GatewayDeploymentRef != nil {
		in, out := &in.GatewayDeploymentRef, &out.GatewayDeploymentRef
		*out = (*in).DeepCopy()
	}
	if in.GatewayServiceRef != nil {
		in, out := &in.GatewayServiceRef, &out.GatewayServiceRef
		*out = (*in).DeepCopy()
	}
	if in.ProvisionerDeploymentRef != nil {
		in, out := &in.ProvisionerDeploymentRef, &out.ProvisionerDeploymentRef
		*out = (*in).DeepCopy()
	}
	if in.ProvisionerServiceRef != nil {
		in, out := &in.ProvisionerServiceRef, &out.ProvisionerServiceRef
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarProviderStatus.
func (in *PulsarProviderStatus) DeepCopy() *PulsarProviderStatus {
	if in == nil {
		return nil
	}
	out := new(PulsarProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stream) DeepCopyInto(out *Stream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stream.
func (in *Stream) DeepCopy() *Stream {
	if in == nil {
		return nil
	}
	out := new(Stream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamAddress) DeepCopyInto(out *StreamAddress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamAddress.
func (in *StreamAddress) DeepCopy() *StreamAddress {
	if in == nil {
		return nil
	}
	out := new(StreamAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamList) DeepCopyInto(out *StreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamList.
func (in *StreamList) DeepCopy() *StreamList {
	if in == nil {
		return nil
	}
	out := new(StreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamSpec) DeepCopyInto(out *StreamSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamSpec.
func (in *StreamSpec) DeepCopy() *StreamSpec {
	if in == nil {
		return nil
	}
	out := new(StreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamStatus) DeepCopyInto(out *StreamStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	out.Address = in.Address
	out.Binding = in.Binding
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamStatus.
func (in *StreamStatus) DeepCopy() *StreamStatus {
	if in == nil {
		return nil
	}
	out := new(StreamStatus)
	in.DeepCopyInto(out)
	return out
}
