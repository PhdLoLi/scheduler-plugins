// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"

	config "sigs.k8s.io/scheduler-plugins/pkg/apis/config"
)

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CoschedulingArgs)(nil), (*config.CoschedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(a.(*CoschedulingArgs), b.(*config.CoschedulingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.CoschedulingArgs)(nil), (*CoschedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(a.(*config.CoschedulingArgs), b.(*CoschedulingArgs), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(in *CoschedulingArgs, out *config.CoschedulingArgs, s conversion.Scope) error {
	out.PermitWaitingTimeSeconds = (*int64)(unsafe.Pointer(in.PermitWaitingTimeSeconds))
	out.PodGroupGCIntervalSeconds = (*int64)(unsafe.Pointer(in.PodGroupGCIntervalSeconds))
	out.PodGroupExpirationTimeSeconds = (*int64)(unsafe.Pointer(in.PodGroupExpirationTimeSeconds))
	return nil
}

// Convert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs is an autogenerated conversion function.
func Convert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(in *CoschedulingArgs, out *config.CoschedulingArgs, s conversion.Scope) error {
	return autoConvert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(in, out, s)
}

func autoConvert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(in *config.CoschedulingArgs, out *CoschedulingArgs, s conversion.Scope) error {
	out.PermitWaitingTimeSeconds = (*int64)(unsafe.Pointer(in.PermitWaitingTimeSeconds))
	out.PodGroupGCIntervalSeconds = (*int64)(unsafe.Pointer(in.PodGroupGCIntervalSeconds))
	out.PodGroupExpirationTimeSeconds = (*int64)(unsafe.Pointer(in.PodGroupExpirationTimeSeconds))
	return nil
}

// Convert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs is an autogenerated conversion function.
func Convert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(in *config.CoschedulingArgs, out *CoschedulingArgs, s conversion.Scope) error {
	return autoConvert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(in, out, s)
}
