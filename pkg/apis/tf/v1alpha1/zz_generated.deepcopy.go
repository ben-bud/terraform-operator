//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright isaaguilar.

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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/rbac/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSCredentials) DeepCopyInto(out *AWSCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSCredentials.
func (in *AWSCredentials) DeepCopy() *AWSCredentials {
	if in == nil {
		return nil
	}
	out := new(AWSCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapSelector) DeepCopyInto(out *ConfigMapSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapSelector.
func (in *ConfigMapSelector) DeepCopy() *ConfigMapSelector {
	if in == nil {
		return nil
	}
	out := new(ConfigMapSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credentials) DeepCopyInto(out *Credentials) {
	*out = *in
	out.SecretNameRef = in.SecretNameRef
	out.AWSCredentials = in.AWSCredentials
	if in.ServiceAccountAnnotations != nil {
		in, out := &in.ServiceAccountAnnotations, &out.ServiceAccountAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credentials.
func (in *Credentials) DeepCopy() *Credentials {
	if in == nil {
		return nil
	}
	out := new(Credentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportRepo) DeepCopyInto(out *ExportRepo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportRepo.
func (in *ExportRepo) DeepCopy() *ExportRepo {
	if in == nil {
		return nil
	}
	out := new(ExportRepo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHTTPS) DeepCopyInto(out *GitHTTPS) {
	*out = *in
	if in.TokenSecretRef != nil {
		in, out := &in.TokenSecretRef, &out.TokenSecretRef
		*out = new(TokenSecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHTTPS.
func (in *GitHTTPS) DeepCopy() *GitHTTPS {
	if in == nil {
		return nil
	}
	out := new(GitHTTPS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSCM) DeepCopyInto(out *GitSCM) {
	*out = *in
	if in.SSH != nil {
		in, out := &in.SSH, &out.SSH
		*out = new(GitSSH)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTPS != nil {
		in, out := &in.HTTPS, &out.HTTPS
		*out = new(GitHTTPS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSCM.
func (in *GitSCM) DeepCopy() *GitSCM {
	if in == nil {
		return nil
	}
	out := new(GitSCM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSSH) DeepCopyInto(out *GitSSH) {
	*out = *in
	if in.SSHKeySecretRef != nil {
		in, out := &in.SSHKeySecretRef, &out.SSHKeySecretRef
		*out = new(SSHKeySecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSSH.
func (in *GitSSH) DeepCopy() *GitSSH {
	if in == nil {
		return nil
	}
	out := new(GitSSH)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Inline) DeepCopyInto(out *Inline) {
	*out = *in
	if in.ConfigMapFiles != nil {
		in, out := &in.ConfigMapFiles, &out.ConfigMapFiles
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Inline.
func (in *Inline) DeepCopy() *Inline {
	if in == nil {
		return nil
	}
	out := new(Inline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyOpts) DeepCopyInto(out *ProxyOpts) {
	*out = *in
	out.SSHKeySecretRef = in.SSHKeySecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyOpts.
func (in *ProxyOpts) DeepCopy() *ProxyOpts {
	if in == nil {
		return nil
	}
	out := new(ProxyOpts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReconcileTerraformDeployment) DeepCopyInto(out *ReconcileTerraformDeployment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReconcileTerraformDeployment.
func (in *ReconcileTerraformDeployment) DeepCopy() *ReconcileTerraformDeployment {
	if in == nil {
		return nil
	}
	out := new(ReconcileTerraformDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceDownload) DeepCopyInto(out *ResourceDownload) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceDownload.
func (in *ResourceDownload) DeepCopy() *ResourceDownload {
	if in == nil {
		return nil
	}
	out := new(ResourceDownload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SCMAuthMethod) DeepCopyInto(out *SCMAuthMethod) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitSCM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCMAuthMethod.
func (in *SCMAuthMethod) DeepCopy() *SCMAuthMethod {
	if in == nil {
		return nil
	}
	out := new(SCMAuthMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKeySecretRef) DeepCopyInto(out *SSHKeySecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKeySecretRef.
func (in *SSHKeySecretRef) DeepCopy() *SSHKeySecretRef {
	if in == nil {
		return nil
	}
	out := new(SSHKeySecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretNameRef) DeepCopyInto(out *SecretNameRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretNameRef.
func (in *SecretNameRef) DeepCopy() *SecretNameRef {
	if in == nil {
		return nil
	}
	out := new(SecretNameRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stage) DeepCopyInto(out *Stage) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.StopTime.DeepCopyInto(&out.StopTime)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Terraform) DeepCopyInto(out *Terraform) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Terraform.
func (in *Terraform) DeepCopy() *Terraform {
	if in == nil {
		return nil
	}
	out := new(Terraform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Terraform) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformList) DeepCopyInto(out *TerraformList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Terraform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformList.
func (in *TerraformList) DeepCopy() *TerraformList {
	if in == nil {
		return nil
	}
	out := new(TerraformList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerraformList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformSpec) DeepCopyInto(out *TerraformSpec) {
	*out = *in
	if in.PersistentVolumeSize != nil {
		in, out := &in.PersistentVolumeSize, &out.PersistentVolumeSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.RunnerRules != nil {
		in, out := &in.RunnerRules, &out.RunnerRules
		*out = make([]v1.PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RunnerAnnotations != nil {
		in, out := &in.RunnerAnnotations, &out.RunnerAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RunnerLabels != nil {
		in, out := &in.RunnerLabels, &out.RunnerLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TerraformRunnerExecutionScriptConfigMap != nil {
		in, out := &in.TerraformRunnerExecutionScriptConfigMap, &out.TerraformRunnerExecutionScriptConfigMap
		*out = new(corev1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ScriptRunnerExecutionScriptConfigMap != nil {
		in, out := &in.ScriptRunnerExecutionScriptConfigMap, &out.ScriptRunnerExecutionScriptConfigMap
		*out = new(corev1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SetupRunnerExecutionScriptConfigMap != nil {
		in, out := &in.SetupRunnerExecutionScriptConfigMap, &out.SetupRunnerExecutionScriptConfigMap
		*out = new(corev1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.TerraformModuleConfigMap != nil {
		in, out := &in.TerraformModuleConfigMap, &out.TerraformModuleConfigMap
		*out = new(ConfigMapSelector)
		**out = **in
	}
	if in.OutputsToInclude != nil {
		in, out := &in.OutputsToInclude, &out.OutputsToInclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OutputsToOmit != nil {
		in, out := &in.OutputsToOmit, &out.OutputsToOmit
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceDownloads != nil {
		in, out := &in.ResourceDownloads, &out.ResourceDownloads
		*out = make([]*ResourceDownload, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ResourceDownload)
				**out = **in
			}
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]Credentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Reconcile != nil {
		in, out := &in.Reconcile, &out.Reconcile
		*out = new(ReconcileTerraformDeployment)
		**out = **in
	}
	if in.ExportRepo != nil {
		in, out := &in.ExportRepo, &out.ExportRepo
		*out = new(ExportRepo)
		**out = **in
	}
	if in.SSHTunnel != nil {
		in, out := &in.SSHTunnel, &out.SSHTunnel
		*out = new(ProxyOpts)
		**out = **in
	}
	if in.SCMAuthMethods != nil {
		in, out := &in.SCMAuthMethods, &out.SCMAuthMethods
		*out = make([]SCMAuthMethod, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformSpec.
func (in *TerraformSpec) DeepCopy() *TerraformSpec {
	if in == nil {
		return nil
	}
	out := new(TerraformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformStatus) DeepCopyInto(out *TerraformStatus) {
	*out = *in
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]Stage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformStatus.
func (in *TerraformStatus) DeepCopy() *TerraformStatus {
	if in == nil {
		return nil
	}
	out := new(TerraformStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenSecretRef) DeepCopyInto(out *TokenSecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenSecretRef.
func (in *TokenSecretRef) DeepCopy() *TokenSecretRef {
	if in == nil {
		return nil
	}
	out := new(TokenSecretRef)
	in.DeepCopyInto(out)
	return out
}
