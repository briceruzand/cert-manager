//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The cert-manager Authors.

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

package v1alpha1

import (
	unsafe "unsafe"

	controller "github.com/cert-manager/cert-manager/internal/apis/config/controller"
	sharedv1alpha1 "github.com/cert-manager/cert-manager/internal/apis/config/shared/v1alpha1"
	v1alpha1 "github.com/cert-manager/cert-manager/pkg/apis/config/controller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ACMEDNS01Config)(nil), (*controller.ACMEDNS01Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ACMEDNS01Config_To_controller_ACMEDNS01Config(a.(*v1alpha1.ACMEDNS01Config), b.(*controller.ACMEDNS01Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*controller.ACMEDNS01Config)(nil), (*v1alpha1.ACMEDNS01Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_controller_ACMEDNS01Config_To_v1alpha1_ACMEDNS01Config(a.(*controller.ACMEDNS01Config), b.(*v1alpha1.ACMEDNS01Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ACMEHTTP01Config)(nil), (*controller.ACMEHTTP01Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ACMEHTTP01Config_To_controller_ACMEHTTP01Config(a.(*v1alpha1.ACMEHTTP01Config), b.(*controller.ACMEHTTP01Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*controller.ACMEHTTP01Config)(nil), (*v1alpha1.ACMEHTTP01Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_controller_ACMEHTTP01Config_To_v1alpha1_ACMEHTTP01Config(a.(*controller.ACMEHTTP01Config), b.(*v1alpha1.ACMEHTTP01Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ControllerConfiguration)(nil), (*controller.ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerConfiguration_To_controller_ControllerConfiguration(a.(*v1alpha1.ControllerConfiguration), b.(*controller.ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*controller.ControllerConfiguration)(nil), (*v1alpha1.ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_controller_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(a.(*controller.ControllerConfiguration), b.(*v1alpha1.ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.IngressShimConfig)(nil), (*controller.IngressShimConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IngressShimConfig_To_controller_IngressShimConfig(a.(*v1alpha1.IngressShimConfig), b.(*controller.IngressShimConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*controller.IngressShimConfig)(nil), (*v1alpha1.IngressShimConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_controller_IngressShimConfig_To_v1alpha1_IngressShimConfig(a.(*controller.IngressShimConfig), b.(*v1alpha1.IngressShimConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.LeaderElectionConfig)(nil), (*controller.LeaderElectionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_LeaderElectionConfig_To_controller_LeaderElectionConfig(a.(*v1alpha1.LeaderElectionConfig), b.(*controller.LeaderElectionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*controller.LeaderElectionConfig)(nil), (*v1alpha1.LeaderElectionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_controller_LeaderElectionConfig_To_v1alpha1_LeaderElectionConfig(a.(*controller.LeaderElectionConfig), b.(*v1alpha1.LeaderElectionConfig), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ACMEDNS01Config_To_controller_ACMEDNS01Config(in *v1alpha1.ACMEDNS01Config, out *controller.ACMEDNS01Config, s conversion.Scope) error {
	out.RecursiveNameservers = *(*[]string)(unsafe.Pointer(&in.RecursiveNameservers))
	if err := v1.Convert_Pointer_bool_To_bool(&in.RecursiveNameserversOnly, &out.RecursiveNameserversOnly, s); err != nil {
		return err
	}
	if err := sharedv1alpha1.Convert_Pointer_v1alpha1_Duration_To_time_Duration(&in.CheckRetryPeriod, &out.CheckRetryPeriod, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ACMEDNS01Config_To_controller_ACMEDNS01Config is an autogenerated conversion function.
func Convert_v1alpha1_ACMEDNS01Config_To_controller_ACMEDNS01Config(in *v1alpha1.ACMEDNS01Config, out *controller.ACMEDNS01Config, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMEDNS01Config_To_controller_ACMEDNS01Config(in, out, s)
}

func autoConvert_controller_ACMEDNS01Config_To_v1alpha1_ACMEDNS01Config(in *controller.ACMEDNS01Config, out *v1alpha1.ACMEDNS01Config, s conversion.Scope) error {
	out.RecursiveNameservers = *(*[]string)(unsafe.Pointer(&in.RecursiveNameservers))
	if err := v1.Convert_bool_To_Pointer_bool(&in.RecursiveNameserversOnly, &out.RecursiveNameserversOnly, s); err != nil {
		return err
	}
	if err := sharedv1alpha1.Convert_time_Duration_To_Pointer_v1alpha1_Duration(&in.CheckRetryPeriod, &out.CheckRetryPeriod, s); err != nil {
		return err
	}
	return nil
}

// Convert_controller_ACMEDNS01Config_To_v1alpha1_ACMEDNS01Config is an autogenerated conversion function.
func Convert_controller_ACMEDNS01Config_To_v1alpha1_ACMEDNS01Config(in *controller.ACMEDNS01Config, out *v1alpha1.ACMEDNS01Config, s conversion.Scope) error {
	return autoConvert_controller_ACMEDNS01Config_To_v1alpha1_ACMEDNS01Config(in, out, s)
}

func autoConvert_v1alpha1_ACMEHTTP01Config_To_controller_ACMEHTTP01Config(in *v1alpha1.ACMEHTTP01Config, out *controller.ACMEHTTP01Config, s conversion.Scope) error {
	out.SolverImage = in.SolverImage
	out.SolverResourceRequestCPU = in.SolverResourceRequestCPU
	out.SolverResourceRequestMemory = in.SolverResourceRequestMemory
	out.SolverResourceLimitsCPU = in.SolverResourceLimitsCPU
	out.SolverResourceLimitsMemory = in.SolverResourceLimitsMemory
	if err := v1.Convert_Pointer_bool_To_bool(&in.SolverRunAsNonRoot, &out.SolverRunAsNonRoot, s); err != nil {
		return err
	}
	out.SolverNameservers = *(*[]string)(unsafe.Pointer(&in.SolverNameservers))
	return nil
}

// Convert_v1alpha1_ACMEHTTP01Config_To_controller_ACMEHTTP01Config is an autogenerated conversion function.
func Convert_v1alpha1_ACMEHTTP01Config_To_controller_ACMEHTTP01Config(in *v1alpha1.ACMEHTTP01Config, out *controller.ACMEHTTP01Config, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMEHTTP01Config_To_controller_ACMEHTTP01Config(in, out, s)
}

func autoConvert_controller_ACMEHTTP01Config_To_v1alpha1_ACMEHTTP01Config(in *controller.ACMEHTTP01Config, out *v1alpha1.ACMEHTTP01Config, s conversion.Scope) error {
	out.SolverImage = in.SolverImage
	out.SolverResourceRequestCPU = in.SolverResourceRequestCPU
	out.SolverResourceRequestMemory = in.SolverResourceRequestMemory
	out.SolverResourceLimitsCPU = in.SolverResourceLimitsCPU
	out.SolverResourceLimitsMemory = in.SolverResourceLimitsMemory
	if err := v1.Convert_bool_To_Pointer_bool(&in.SolverRunAsNonRoot, &out.SolverRunAsNonRoot, s); err != nil {
		return err
	}
	out.SolverNameservers = *(*[]string)(unsafe.Pointer(&in.SolverNameservers))
	return nil
}

// Convert_controller_ACMEHTTP01Config_To_v1alpha1_ACMEHTTP01Config is an autogenerated conversion function.
func Convert_controller_ACMEHTTP01Config_To_v1alpha1_ACMEHTTP01Config(in *controller.ACMEHTTP01Config, out *v1alpha1.ACMEHTTP01Config, s conversion.Scope) error {
	return autoConvert_controller_ACMEHTTP01Config_To_v1alpha1_ACMEHTTP01Config(in, out, s)
}

func autoConvert_v1alpha1_ControllerConfiguration_To_controller_ControllerConfiguration(in *v1alpha1.ControllerConfiguration, out *controller.ControllerConfiguration, s conversion.Scope) error {
	out.KubeConfig = in.KubeConfig
	out.APIServerHost = in.APIServerHost
	if err := sharedv1alpha1.Convert_Pointer_float32_To_float32(&in.KubernetesAPIQPS, &out.KubernetesAPIQPS, s); err != nil {
		return err
	}
	if err := sharedv1alpha1.Convert_Pointer_int32_To_int(&in.KubernetesAPIBurst, &out.KubernetesAPIBurst, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	out.ClusterResourceNamespace = in.ClusterResourceNamespace
	if err := Convert_v1alpha1_LeaderElectionConfig_To_controller_LeaderElectionConfig(&in.LeaderElectionConfig, &out.LeaderElectionConfig, s); err != nil {
		return err
	}
	out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
	if err := v1.Convert_Pointer_bool_To_bool(&in.IssuerAmbientCredentials, &out.IssuerAmbientCredentials, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.ClusterIssuerAmbientCredentials, &out.ClusterIssuerAmbientCredentials, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableCertificateOwnerRef, &out.EnableCertificateOwnerRef, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableGatewayAPI, &out.EnableGatewayAPI, s); err != nil {
		return err
	}
	out.CopiedAnnotationPrefixes = *(*[]string)(unsafe.Pointer(&in.CopiedAnnotationPrefixes))
	if err := sharedv1alpha1.Convert_Pointer_int32_To_int(&in.NumberOfConcurrentWorkers, &out.NumberOfConcurrentWorkers, s); err != nil {
		return err
	}
	if err := sharedv1alpha1.Convert_Pointer_int32_To_int(&in.MaxConcurrentChallenges, &out.MaxConcurrentChallenges, s); err != nil {
		return err
	}
	out.MetricsListenAddress = in.MetricsListenAddress
	if err := sharedv1alpha1.Convert_v1alpha1_TLSConfig_To_shared_TLSConfig(&in.MetricsTLSConfig, &out.MetricsTLSConfig, s); err != nil {
		return err
	}
	out.HealthzListenAddress = in.HealthzListenAddress
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnablePprof, &out.EnablePprof, s); err != nil {
		return err
	}
	out.PprofAddress = in.PprofAddress
	out.Logging = in.Logging
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	if err := Convert_v1alpha1_IngressShimConfig_To_controller_IngressShimConfig(&in.IngressShimConfig, &out.IngressShimConfig, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ACMEHTTP01Config_To_controller_ACMEHTTP01Config(&in.ACMEHTTP01Config, &out.ACMEHTTP01Config, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ACMEDNS01Config_To_controller_ACMEDNS01Config(&in.ACMEDNS01Config, &out.ACMEDNS01Config, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ControllerConfiguration_To_controller_ControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ControllerConfiguration_To_controller_ControllerConfiguration(in *v1alpha1.ControllerConfiguration, out *controller.ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerConfiguration_To_controller_ControllerConfiguration(in, out, s)
}

func autoConvert_controller_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *controller.ControllerConfiguration, out *v1alpha1.ControllerConfiguration, s conversion.Scope) error {
	out.APIServerHost = in.APIServerHost
	out.KubeConfig = in.KubeConfig
	if err := sharedv1alpha1.Convert_float32_To_Pointer_float32(&in.KubernetesAPIQPS, &out.KubernetesAPIQPS, s); err != nil {
		return err
	}
	if err := sharedv1alpha1.Convert_int_To_Pointer_int32(&in.KubernetesAPIBurst, &out.KubernetesAPIBurst, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	out.ClusterResourceNamespace = in.ClusterResourceNamespace
	if err := Convert_controller_LeaderElectionConfig_To_v1alpha1_LeaderElectionConfig(&in.LeaderElectionConfig, &out.LeaderElectionConfig, s); err != nil {
		return err
	}
	out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
	if err := v1.Convert_bool_To_Pointer_bool(&in.IssuerAmbientCredentials, &out.IssuerAmbientCredentials, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.ClusterIssuerAmbientCredentials, &out.ClusterIssuerAmbientCredentials, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableCertificateOwnerRef, &out.EnableCertificateOwnerRef, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableGatewayAPI, &out.EnableGatewayAPI, s); err != nil {
		return err
	}
	out.CopiedAnnotationPrefixes = *(*[]string)(unsafe.Pointer(&in.CopiedAnnotationPrefixes))
	if err := sharedv1alpha1.Convert_int_To_Pointer_int32(&in.NumberOfConcurrentWorkers, &out.NumberOfConcurrentWorkers, s); err != nil {
		return err
	}
	if err := sharedv1alpha1.Convert_int_To_Pointer_int32(&in.MaxConcurrentChallenges, &out.MaxConcurrentChallenges, s); err != nil {
		return err
	}
	out.MetricsListenAddress = in.MetricsListenAddress
	if err := sharedv1alpha1.Convert_shared_TLSConfig_To_v1alpha1_TLSConfig(&in.MetricsTLSConfig, &out.MetricsTLSConfig, s); err != nil {
		return err
	}
	out.HealthzListenAddress = in.HealthzListenAddress
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnablePprof, &out.EnablePprof, s); err != nil {
		return err
	}
	out.PprofAddress = in.PprofAddress
	out.Logging = in.Logging
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	if err := Convert_controller_IngressShimConfig_To_v1alpha1_IngressShimConfig(&in.IngressShimConfig, &out.IngressShimConfig, s); err != nil {
		return err
	}
	if err := Convert_controller_ACMEHTTP01Config_To_v1alpha1_ACMEHTTP01Config(&in.ACMEHTTP01Config, &out.ACMEHTTP01Config, s); err != nil {
		return err
	}
	if err := Convert_controller_ACMEDNS01Config_To_v1alpha1_ACMEDNS01Config(&in.ACMEDNS01Config, &out.ACMEDNS01Config, s); err != nil {
		return err
	}
	return nil
}

// Convert_controller_ControllerConfiguration_To_v1alpha1_ControllerConfiguration is an autogenerated conversion function.
func Convert_controller_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *controller.ControllerConfiguration, out *v1alpha1.ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_controller_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_IngressShimConfig_To_controller_IngressShimConfig(in *v1alpha1.IngressShimConfig, out *controller.IngressShimConfig, s conversion.Scope) error {
	out.DefaultIssuerName = in.DefaultIssuerName
	out.DefaultIssuerKind = in.DefaultIssuerKind
	out.DefaultIssuerGroup = in.DefaultIssuerGroup
	out.DefaultAutoCertificateAnnotations = *(*[]string)(unsafe.Pointer(&in.DefaultAutoCertificateAnnotations))
	out.ExtraCertificateAnnotations = *(*[]string)(unsafe.Pointer(&in.ExtraCertificateAnnotations))
	return nil
}

// Convert_v1alpha1_IngressShimConfig_To_controller_IngressShimConfig is an autogenerated conversion function.
func Convert_v1alpha1_IngressShimConfig_To_controller_IngressShimConfig(in *v1alpha1.IngressShimConfig, out *controller.IngressShimConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_IngressShimConfig_To_controller_IngressShimConfig(in, out, s)
}

func autoConvert_controller_IngressShimConfig_To_v1alpha1_IngressShimConfig(in *controller.IngressShimConfig, out *v1alpha1.IngressShimConfig, s conversion.Scope) error {
	out.DefaultIssuerName = in.DefaultIssuerName
	out.DefaultIssuerKind = in.DefaultIssuerKind
	out.DefaultIssuerGroup = in.DefaultIssuerGroup
	out.DefaultAutoCertificateAnnotations = *(*[]string)(unsafe.Pointer(&in.DefaultAutoCertificateAnnotations))
	out.ExtraCertificateAnnotations = *(*[]string)(unsafe.Pointer(&in.ExtraCertificateAnnotations))
	return nil
}

// Convert_controller_IngressShimConfig_To_v1alpha1_IngressShimConfig is an autogenerated conversion function.
func Convert_controller_IngressShimConfig_To_v1alpha1_IngressShimConfig(in *controller.IngressShimConfig, out *v1alpha1.IngressShimConfig, s conversion.Scope) error {
	return autoConvert_controller_IngressShimConfig_To_v1alpha1_IngressShimConfig(in, out, s)
}

func autoConvert_v1alpha1_LeaderElectionConfig_To_controller_LeaderElectionConfig(in *v1alpha1.LeaderElectionConfig, out *controller.LeaderElectionConfig, s conversion.Scope) error {
	if err := sharedv1alpha1.Convert_v1alpha1_LeaderElectionConfig_To_shared_LeaderElectionConfig(&in.LeaderElectionConfig, &out.LeaderElectionConfig, s); err != nil {
		return err
	}
	if err := sharedv1alpha1.Convert_Pointer_v1alpha1_Duration_To_time_Duration(&in.HealthzTimeout, &out.HealthzTimeout, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_LeaderElectionConfig_To_controller_LeaderElectionConfig is an autogenerated conversion function.
func Convert_v1alpha1_LeaderElectionConfig_To_controller_LeaderElectionConfig(in *v1alpha1.LeaderElectionConfig, out *controller.LeaderElectionConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_LeaderElectionConfig_To_controller_LeaderElectionConfig(in, out, s)
}

func autoConvert_controller_LeaderElectionConfig_To_v1alpha1_LeaderElectionConfig(in *controller.LeaderElectionConfig, out *v1alpha1.LeaderElectionConfig, s conversion.Scope) error {
	if err := sharedv1alpha1.Convert_shared_LeaderElectionConfig_To_v1alpha1_LeaderElectionConfig(&in.LeaderElectionConfig, &out.LeaderElectionConfig, s); err != nil {
		return err
	}
	if err := sharedv1alpha1.Convert_time_Duration_To_Pointer_v1alpha1_Duration(&in.HealthzTimeout, &out.HealthzTimeout, s); err != nil {
		return err
	}
	return nil
}

// Convert_controller_LeaderElectionConfig_To_v1alpha1_LeaderElectionConfig is an autogenerated conversion function.
func Convert_controller_LeaderElectionConfig_To_v1alpha1_LeaderElectionConfig(in *controller.LeaderElectionConfig, out *v1alpha1.LeaderElectionConfig, s conversion.Scope) error {
	return autoConvert_controller_LeaderElectionConfig_To_v1alpha1_LeaderElectionConfig(in, out, s)
}
