//go:build !ignore_autogenerated
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

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	json "encoding/json"

	v1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	v1 "k8s.io/api/core/v1"
	prowjobsv1 "k8s.io/test-infra/prow/apis/prowjobs/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Brancher) DeepCopyInto(out *Brancher) {
	*out = *in
	if in.SkipBranches != nil {
		in, out := &in.SkipBranches, &out.SkipBranches
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Branches != nil {
		in, out := &in.Branches, &out.Branches
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.re != nil {
		in, out := &in.re, &out.re
		*out = (*in).DeepCopy()
	}
	if in.reSkip != nil {
		in, out := &in.reSkip, &out.reSkip
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Brancher.
func (in *Brancher) DeepCopy() *Brancher {
	if in == nil {
		return nil
	}
	out := new(Brancher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopyableRegexp.
func (in *CopyableRegexp) DeepCopy() *CopyableRegexp {
	if in == nil {
		return nil
	}
	out := new(CopyableRegexp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobBase) DeepCopyInto(out *JobBase) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(v1.PodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PipelineRunSpec != nil {
		in, out := &in.PipelineRunSpec, &out.PipelineRunSpec
		*out = new(v1alpha1.PipelineRunSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReporterConfig != nil {
		in, out := &in.ReporterConfig, &out.ReporterConfig
		*out = new(prowjobsv1.ReporterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RerunAuthConfig != nil {
		in, out := &in.RerunAuthConfig, &out.RerunAuthConfig
		*out = new(prowjobsv1.RerunAuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ProwJobDefault != nil {
		in, out := &in.ProwJobDefault, &out.ProwJobDefault
		*out = new(prowjobsv1.ProwJobDefault)
		**out = **in
	}
	in.UtilityConfig.DeepCopyInto(&out.UtilityConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobBase.
func (in *JobBase) DeepCopy() *JobBase {
	if in == nil {
		return nil
	}
	out := new(JobBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Postsubmit) DeepCopyInto(out *Postsubmit) {
	*out = *in
	in.JobBase.DeepCopyInto(&out.JobBase)
	if in.AlwaysRun != nil {
		in, out := &in.AlwaysRun, &out.AlwaysRun
		*out = new(bool)
		**out = **in
	}
	in.RegexpChangeMatcher.DeepCopyInto(&out.RegexpChangeMatcher)
	in.Brancher.DeepCopyInto(&out.Brancher)
	out.Reporter = in.Reporter
	if in.JenkinsSpec != nil {
		in, out := &in.JenkinsSpec, &out.JenkinsSpec
		*out = new(JenkinsSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Postsubmit.
func (in *Postsubmit) DeepCopy() *Postsubmit {
	if in == nil {
		return nil
	}
	out := new(Postsubmit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Preset) DeepCopyInto(out *Preset) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Preset.
func (in *Preset) DeepCopy() *Preset {
	if in == nil {
		return nil
	}
	out := new(Preset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Presubmit) DeepCopyInto(out *Presubmit) {
	*out = *in
	in.JobBase.DeepCopyInto(&out.JobBase)
	in.Brancher.DeepCopyInto(&out.Brancher)
	in.RegexpChangeMatcher.DeepCopyInto(&out.RegexpChangeMatcher)
	out.Reporter = in.Reporter
	if in.JenkinsSpec != nil {
		in, out := &in.JenkinsSpec, &out.JenkinsSpec
		*out = new(JenkinsSpec)
		**out = **in
	}
	if in.re != nil {
		in, out := &in.re, &out.re
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Presubmit.
func (in *Presubmit) DeepCopy() *Presubmit {
	if in == nil {
		return nil
	}
	out := new(Presubmit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProwYAML) DeepCopyInto(out *ProwYAML) {
	*out = *in
	if in.Presets != nil {
		in, out := &in.Presets, &out.Presets
		*out = make([]Preset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Presubmits != nil {
		in, out := &in.Presubmits, &out.Presubmits
		*out = make([]Presubmit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Postsubmits != nil {
		in, out := &in.Postsubmits, &out.Postsubmits
		*out = make([]Postsubmit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProwIgnored != nil {
		in, out := &in.ProwIgnored, &out.ProwIgnored
		*out = new(json.RawMessage)
		if **in != nil {
			in, out := *in, *out
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProwYAML.
func (in *ProwYAML) DeepCopy() *ProwYAML {
	if in == nil {
		return nil
	}
	out := new(ProwYAML)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegexpChangeMatcher) DeepCopyInto(out *RegexpChangeMatcher) {
	*out = *in
	if in.reChanges != nil {
		in, out := &in.reChanges, &out.reChanges
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegexpChangeMatcher.
func (in *RegexpChangeMatcher) DeepCopy() *RegexpChangeMatcher {
	if in == nil {
		return nil
	}
	out := new(RegexpChangeMatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UtilityConfig) DeepCopyInto(out *UtilityConfig) {
	*out = *in
	if in.Decorate != nil {
		in, out := &in.Decorate, &out.Decorate
		*out = new(bool)
		**out = **in
	}
	if in.ExtraRefs != nil {
		in, out := &in.ExtraRefs, &out.ExtraRefs
		*out = make([]prowjobsv1.Refs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DecorationConfig != nil {
		in, out := &in.DecorationConfig, &out.DecorationConfig
		*out = new(prowjobsv1.DecorationConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UtilityConfig.
func (in *UtilityConfig) DeepCopy() *UtilityConfig {
	if in == nil {
		return nil
	}
	out := new(UtilityConfig)
	in.DeepCopyInto(out)
	return out
}
