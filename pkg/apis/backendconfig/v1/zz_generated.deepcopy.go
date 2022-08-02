//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Kubernetes Authors.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfig) DeepCopyInto(out *BackendConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfig.
func (in *BackendConfig) DeepCopy() *BackendConfig {
	if in == nil {
		return nil
	}
	out := new(BackendConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfigList) DeepCopyInto(out *BackendConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackendConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfigList.
func (in *BackendConfigList) DeepCopy() *BackendConfigList {
	if in == nil {
		return nil
	}
	out := new(BackendConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfigSpec) DeepCopyInto(out *BackendConfigSpec) {
	*out = *in
	if in.Iap != nil {
		in, out := &in.Iap, &out.Iap
		*out = new(IAPConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Cdn != nil {
		in, out := &in.Cdn, &out.Cdn
		*out = new(CDNConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityPolicy != nil {
		in, out := &in.SecurityPolicy, &out.SecurityPolicy
		*out = new(SecurityPolicyConfig)
		**out = **in
	}
	if in.TimeoutSec != nil {
		in, out := &in.TimeoutSec, &out.TimeoutSec
		*out = new(int64)
		**out = **in
	}
	if in.ConnectionDraining != nil {
		in, out := &in.ConnectionDraining, &out.ConnectionDraining
		*out = new(ConnectionDrainingConfig)
		**out = **in
	}
	if in.SessionAffinity != nil {
		in, out := &in.SessionAffinity, &out.SessionAffinity
		*out = new(SessionAffinityConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomRequestHeaders != nil {
		in, out := &in.CustomRequestHeaders, &out.CustomRequestHeaders
		*out = new(CustomRequestHeadersConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheckConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(LogConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfigSpec.
func (in *BackendConfigSpec) DeepCopy() *BackendConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BackendConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfigStatus) DeepCopyInto(out *BackendConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfigStatus.
func (in *BackendConfigStatus) DeepCopy() *BackendConfigStatus {
	if in == nil {
		return nil
	}
	out := new(BackendConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BypassCacheOnRequestHeader) DeepCopyInto(out *BypassCacheOnRequestHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BypassCacheOnRequestHeader.
func (in *BypassCacheOnRequestHeader) DeepCopy() *BypassCacheOnRequestHeader {
	if in == nil {
		return nil
	}
	out := new(BypassCacheOnRequestHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDNConfig) DeepCopyInto(out *CDNConfig) {
	*out = *in
	if in.BypassCacheOnRequestHeaders != nil {
		in, out := &in.BypassCacheOnRequestHeaders, &out.BypassCacheOnRequestHeaders
		*out = make([]*BypassCacheOnRequestHeader, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BypassCacheOnRequestHeader)
				**out = **in
			}
		}
	}
	if in.CachePolicy != nil {
		in, out := &in.CachePolicy, &out.CachePolicy
		*out = new(CacheKeyPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.CacheMode != nil {
		in, out := &in.CacheMode, &out.CacheMode
		*out = new(string)
		**out = **in
	}
	if in.ClientTtl != nil {
		in, out := &in.ClientTtl, &out.ClientTtl
		*out = new(int64)
		**out = **in
	}
	if in.DefaultTtl != nil {
		in, out := &in.DefaultTtl, &out.DefaultTtl
		*out = new(int64)
		**out = **in
	}
	if in.MaxTtl != nil {
		in, out := &in.MaxTtl, &out.MaxTtl
		*out = new(int64)
		**out = **in
	}
	if in.NegativeCaching != nil {
		in, out := &in.NegativeCaching, &out.NegativeCaching
		*out = new(bool)
		**out = **in
	}
	if in.NegativeCachingPolicy != nil {
		in, out := &in.NegativeCachingPolicy, &out.NegativeCachingPolicy
		*out = make([]*NegativeCachingPolicy, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NegativeCachingPolicy)
				**out = **in
			}
		}
	}
	if in.RequestCoalescing != nil {
		in, out := &in.RequestCoalescing, &out.RequestCoalescing
		*out = new(bool)
		**out = **in
	}
	if in.ServeWhileStale != nil {
		in, out := &in.ServeWhileStale, &out.ServeWhileStale
		*out = new(int64)
		**out = **in
	}
	if in.SignedUrlCacheMaxAgeSec != nil {
		in, out := &in.SignedUrlCacheMaxAgeSec, &out.SignedUrlCacheMaxAgeSec
		*out = new(int64)
		**out = **in
	}
	if in.SignedUrlKeys != nil {
		in, out := &in.SignedUrlKeys, &out.SignedUrlKeys
		*out = make([]*SignedUrlKey, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SignedUrlKey)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDNConfig.
func (in *CDNConfig) DeepCopy() *CDNConfig {
	if in == nil {
		return nil
	}
	out := new(CDNConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheKeyPolicy) DeepCopyInto(out *CacheKeyPolicy) {
	*out = *in
	if in.QueryStringBlacklist != nil {
		in, out := &in.QueryStringBlacklist, &out.QueryStringBlacklist
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.QueryStringWhitelist != nil {
		in, out := &in.QueryStringWhitelist, &out.QueryStringWhitelist
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheKeyPolicy.
func (in *CacheKeyPolicy) DeepCopy() *CacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := new(CacheKeyPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionDrainingConfig) DeepCopyInto(out *ConnectionDrainingConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionDrainingConfig.
func (in *ConnectionDrainingConfig) DeepCopy() *ConnectionDrainingConfig {
	if in == nil {
		return nil
	}
	out := new(ConnectionDrainingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRequestHeadersConfig) DeepCopyInto(out *CustomRequestHeadersConfig) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRequestHeadersConfig.
func (in *CustomRequestHeadersConfig) DeepCopy() *CustomRequestHeadersConfig {
	if in == nil {
		return nil
	}
	out := new(CustomRequestHeadersConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckConfig) DeepCopyInto(out *HealthCheckConfig) {
	*out = *in
	if in.CheckIntervalSec != nil {
		in, out := &in.CheckIntervalSec, &out.CheckIntervalSec
		*out = new(int64)
		**out = **in
	}
	if in.TimeoutSec != nil {
		in, out := &in.TimeoutSec, &out.TimeoutSec
		*out = new(int64)
		**out = **in
	}
	if in.HealthyThreshold != nil {
		in, out := &in.HealthyThreshold, &out.HealthyThreshold
		*out = new(int64)
		**out = **in
	}
	if in.UnhealthyThreshold != nil {
		in, out := &in.UnhealthyThreshold, &out.UnhealthyThreshold
		*out = new(int64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.RequestPath != nil {
		in, out := &in.RequestPath, &out.RequestPath
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckConfig.
func (in *HealthCheckConfig) DeepCopy() *HealthCheckConfig {
	if in == nil {
		return nil
	}
	out := new(HealthCheckConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAPConfig) DeepCopyInto(out *IAPConfig) {
	*out = *in
	if in.OAuthClientCredentials != nil {
		in, out := &in.OAuthClientCredentials, &out.OAuthClientCredentials
		*out = new(OAuthClientCredentials)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAPConfig.
func (in *IAPConfig) DeepCopy() *IAPConfig {
	if in == nil {
		return nil
	}
	out := new(IAPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogConfig) DeepCopyInto(out *LogConfig) {
	*out = *in
	if in.SampleRate != nil {
		in, out := &in.SampleRate, &out.SampleRate
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogConfig.
func (in *LogConfig) DeepCopy() *LogConfig {
	if in == nil {
		return nil
	}
	out := new(LogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NegativeCachingPolicy) DeepCopyInto(out *NegativeCachingPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NegativeCachingPolicy.
func (in *NegativeCachingPolicy) DeepCopy() *NegativeCachingPolicy {
	if in == nil {
		return nil
	}
	out := new(NegativeCachingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuthClientCredentials) DeepCopyInto(out *OAuthClientCredentials) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuthClientCredentials.
func (in *OAuthClientCredentials) DeepCopy() *OAuthClientCredentials {
	if in == nil {
		return nil
	}
	out := new(OAuthClientCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicyConfig) DeepCopyInto(out *SecurityPolicyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicyConfig.
func (in *SecurityPolicyConfig) DeepCopy() *SecurityPolicyConfig {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionAffinityConfig) DeepCopyInto(out *SessionAffinityConfig) {
	*out = *in
	if in.AffinityCookieTtlSec != nil {
		in, out := &in.AffinityCookieTtlSec, &out.AffinityCookieTtlSec
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionAffinityConfig.
func (in *SessionAffinityConfig) DeepCopy() *SessionAffinityConfig {
	if in == nil {
		return nil
	}
	out := new(SessionAffinityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignedUrlKey) DeepCopyInto(out *SignedUrlKey) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignedUrlKey.
func (in *SignedUrlKey) DeepCopy() *SignedUrlKey {
	if in == nil {
		return nil
	}
	out := new(SignedUrlKey)
	in.DeepCopyInto(out)
	return out
}
