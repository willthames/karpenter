//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
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

package v1beta1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Deprovisioning) DeepCopyInto(out *Deprovisioning) {
	*out = *in
	out.ConsolidationTTL = in.ConsolidationTTL
	out.ExpirationTTL = in.ExpirationTTL
	if in.EmptinessTTL != nil {
		in, out := &in.EmptinessTTL, &out.EmptinessTTL
		*out = new(metav1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Deprovisioning.
func (in *Deprovisioning) DeepCopy() *Deprovisioning {
	if in == nil {
		return nil
	}
	out := new(Deprovisioning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfiguration) DeepCopyInto(out *KubeletConfiguration) {
	*out = *in
	if in.ClusterDNS != nil {
		in, out := &in.ClusterDNS, &out.ClusterDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ContainerRuntime != nil {
		in, out := &in.ContainerRuntime, &out.ContainerRuntime
		*out = new(string)
		**out = **in
	}
	if in.MaxPods != nil {
		in, out := &in.MaxPods, &out.MaxPods
		*out = new(int32)
		**out = **in
	}
	if in.PodsPerCore != nil {
		in, out := &in.PodsPerCore, &out.PodsPerCore
		*out = new(int32)
		**out = **in
	}
	if in.SystemReserved != nil {
		in, out := &in.SystemReserved, &out.SystemReserved
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.KubeReserved != nil {
		in, out := &in.KubeReserved, &out.KubeReserved
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.EvictionHard != nil {
		in, out := &in.EvictionHard, &out.EvictionHard
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionSoft != nil {
		in, out := &in.EvictionSoft, &out.EvictionSoft
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionSoftGracePeriod != nil {
		in, out := &in.EvictionSoftGracePeriod, &out.EvictionSoftGracePeriod
		*out = make(map[string]metav1.Duration, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionMaxPodGracePeriod != nil {
		in, out := &in.EvictionMaxPodGracePeriod, &out.EvictionMaxPodGracePeriod
		*out = new(int32)
		**out = **in
	}
	if in.ImageGCHighThresholdPercent != nil {
		in, out := &in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent
		*out = new(int32)
		**out = **in
	}
	if in.ImageGCLowThresholdPercent != nil {
		in, out := &in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent
		*out = new(int32)
		**out = **in
	}
	if in.CPUCFSQuota != nil {
		in, out := &in.CPUCFSQuota, &out.CPUCFSQuota
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfiguration.
func (in *KubeletConfiguration) DeepCopy() *KubeletConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeletConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Limits) DeepCopyInto(out *Limits) {
	{
		in := &in
		*out = make(Limits, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Limits.
func (in Limits) DeepCopy() Limits {
	if in == nil {
		return nil
	}
	out := new(Limits)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeClaim) DeepCopyInto(out *NodeClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeClaim.
func (in *NodeClaim) DeepCopy() *NodeClaim {
	if in == nil {
		return nil
	}
	out := new(NodeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeClaimList) DeepCopyInto(out *NodeClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeClaimList.
func (in *NodeClaimList) DeepCopy() *NodeClaimList {
	if in == nil {
		return nil
	}
	out := new(NodeClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeClaimSpec) DeepCopyInto(out *NodeClaimSpec) {
	*out = *in
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartupTaints != nil {
		in, out := &in.StartupTaints, &out.StartupTaints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Requirements != nil {
		in, out := &in.Requirements, &out.Requirements
		*out = make([]v1.NodeSelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.KubeletConfiguration != nil {
		in, out := &in.KubeletConfiguration, &out.KubeletConfiguration
		*out = new(KubeletConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeClass != nil {
		in, out := &in.NodeClass, &out.NodeClass
		*out = new(NodeClassRef)
		**out = **in
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeClaimSpec.
func (in *NodeClaimSpec) DeepCopy() *NodeClaimSpec {
	if in == nil {
		return nil
	}
	out := new(NodeClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeClaimStatus) DeepCopyInto(out *NodeClaimStatus) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Allocatable != nil {
		in, out := &in.Allocatable, &out.Allocatable
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apis.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeClaimStatus.
func (in *NodeClaimStatus) DeepCopy() *NodeClaimStatus {
	if in == nil {
		return nil
	}
	out := new(NodeClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeClaimTemplate) DeepCopyInto(out *NodeClaimTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeClaimTemplate.
func (in *NodeClaimTemplate) DeepCopy() *NodeClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(NodeClaimTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeClassRef) DeepCopyInto(out *NodeClassRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeClassRef.
func (in *NodeClassRef) DeepCopy() *NodeClassRef {
	if in == nil {
		return nil
	}
	out := new(NodeClassRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePool) DeepCopyInto(out *NodePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePool.
func (in *NodePool) DeepCopy() *NodePool {
	if in == nil {
		return nil
	}
	out := new(NodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolList) DeepCopyInto(out *NodePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolList.
func (in *NodePoolList) DeepCopy() *NodePoolList {
	if in == nil {
		return nil
	}
	out := new(NodePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpec) DeepCopyInto(out *NodePoolSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Deprovisioning.DeepCopyInto(&out.Deprovisioning)
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(Limits, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpec.
func (in *NodePoolSpec) DeepCopy() *NodePoolSpec {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolStatus) DeepCopyInto(out *NodePoolStatus) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolStatus.
func (in *NodePoolStatus) DeepCopy() *NodePoolStatus {
	if in == nil {
		return nil
	}
	out := new(NodePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirements.
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirements)
	in.DeepCopyInto(out)
	return out
}
