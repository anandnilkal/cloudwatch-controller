//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
func (in *Alarms) DeepCopyInto(out *Alarms) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alarms.
func (in *Alarms) DeepCopy() *Alarms {
	if in == nil {
		return nil
	}
	out := new(Alarms)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alarms) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmsList) DeepCopyInto(out *AlarmsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alarms, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmsList.
func (in *AlarmsList) DeepCopy() *AlarmsList {
	if in == nil {
		return nil
	}
	out := new(AlarmsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlarmsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmsSpec) DeepCopyInto(out *AlarmsSpec) {
	*out = *in
	if in.EvaluationPeriod != nil {
		in, out := &in.EvaluationPeriod, &out.EvaluationPeriod
		*out = new(int32)
		**out = **in
	}
	if in.AlarmActions != nil {
		in, out := &in.AlarmActions, &out.AlarmActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AlarmDescription != nil {
		in, out := &in.AlarmDescription, &out.AlarmDescription
		*out = new(string)
		**out = **in
	}
	if in.DatapointsToAlarm != nil {
		in, out := &in.DatapointsToAlarm, &out.DatapointsToAlarm
		*out = new(int32)
		**out = **in
	}
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]Dimension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EvaluateLowSampleCountPercentile != nil {
		in, out := &in.EvaluateLowSampleCountPercentile, &out.EvaluateLowSampleCountPercentile
		*out = new(string)
		**out = **in
	}
	if in.ExtendedStatistic != nil {
		in, out := &in.ExtendedStatistic, &out.ExtendedStatistic
		*out = new(string)
		**out = **in
	}
	if in.InsufficientDataActions != nil {
		in, out := &in.InsufficientDataActions, &out.InsufficientDataActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]MetricDataQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.OKActions != nil {
		in, out := &in.OKActions, &out.OKActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(int32)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdMetricId != nil {
		in, out := &in.ThresholdMetricId, &out.ThresholdMetricId
		*out = new(string)
		**out = **in
	}
	if in.TreatMissingData != nil {
		in, out := &in.TreatMissingData, &out.TreatMissingData
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmsSpec.
func (in *AlarmsSpec) DeepCopy() *AlarmsSpec {
	if in == nil {
		return nil
	}
	out := new(AlarmsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmsStatus) DeepCopyInto(out *AlarmsStatus) {
	*out = *in
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(string)
		**out = **in
	}
	if in.ErroMessage != nil {
		in, out := &in.ErroMessage, &out.ErroMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmsStatus.
func (in *AlarmsStatus) DeepCopy() *AlarmsStatus {
	if in == nil {
		return nil
	}
	out := new(AlarmsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dimension) DeepCopyInto(out *Dimension) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dimension.
func (in *Dimension) DeepCopy() *Dimension {
	if in == nil {
		return nil
	}
	out := new(Dimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metric) DeepCopyInto(out *Metric) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]Dimension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metric.
func (in *Metric) DeepCopy() *Metric {
	if in == nil {
		return nil
	}
	out := new(Metric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricDataQuery) DeepCopyInto(out *MetricDataQuery) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.AccountId != nil {
		in, out := &in.AccountId, &out.AccountId
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.MetricStat != nil {
		in, out := &in.MetricStat, &out.MetricStat
		*out = new(MetricStat)
		(*in).DeepCopyInto(*out)
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(int32)
		**out = **in
	}
	if in.ReturnData != nil {
		in, out := &in.ReturnData, &out.ReturnData
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricDataQuery.
func (in *MetricDataQuery) DeepCopy() *MetricDataQuery {
	if in == nil {
		return nil
	}
	out := new(MetricDataQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricStat) DeepCopyInto(out *MetricStat) {
	*out = *in
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(Metric)
		(*in).DeepCopyInto(*out)
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(int32)
		**out = **in
	}
	if in.Stat != nil {
		in, out := &in.Stat, &out.Stat
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricStat.
func (in *MetricStat) DeepCopy() *MetricStat {
	if in == nil {
		return nil
	}
	out := new(MetricStat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}