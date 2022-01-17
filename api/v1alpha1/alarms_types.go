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

package v1alpha1

import (
	cloudwatchtypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AlarmsSpec defines the desired state of Alarms
type AlarmsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name name of the Alarm to be created
	Name string `json:"name"`

	// ComparisonOperator comparison operator used for deciding alarm status
	ComparisonOperator cloudwatchtypes.ComparisonOperator `json:"operator"`

	// EvaluationPeriod number of period over which alarm data is compared to threshold value
	EvaluationPeriod *int32 `json:"evaluationPeriod"`

	// AlarmActions action to take when alarm is generated
	AlarmActions []string `json:"alarmActions"`

	// AlarmDescription description for the alarm.
	AlarmDescription *string `json:"description,omitempty"`

	// DatapointsToAlarm number of data points that must be breaching to trigger the alarm. This is
	DatapointsToAlarm *int32 `json:"numDataPointsToAlarm"`

	// Dimensions dimensions for the metric specified in MetricName.
	Dimensions []Dimension `json:"dimensions"`

	// EvaluateLowSampleCountPercentile Used only for alarms based on percentiles. Valid Values: evaluate | ignore
	EvaluateLowSampleCountPercentile *string `json:"evaluateLowSampleCountPercentile,omitempty"`

	// ExtendedStatistic The percentile statistic for the metric specified in MetricName.
	ExtendedStatistic *string `json:"extendedStatistic,omitempty"`

	// InsufficentDataActions The actions to execute when this alarm transitions to the INSUFFICIENT_DATA
	InsufficientDataActions []string `json:"insufficientDataActions,omitempty"`

	// MetricName name for the metric associated with the alarm.
	MetricName *string `json:"metricName"`

	// Metrics An array of MetricDataQuery structures that enable you to create an alarm based
	// on the result of a metric math expression.
	Metrics []MetricDataQuery `json:"metrics"`

	// The namespace for the metric associated specified in MetricName.
	Namespace *string `json:"namespace"`

	// OKActions The actions to execute when this alarm transitions to an OK state from any other
	// state.
	OKActions []string `json:"okActions,omitempty"`

	// Period The length, in seconds, used each time the metric specified in MetricName is
	// evaluated. Valid values are 10, 30, and any multiple of 60.
	Period *int32 `json:"period"`

	// Statistic The statistic for the metric specified in MetricName.
	Statistic cloudwatchtypes.Statistic `json:"statistics,omitempty"`

	// A list of key-value pairs to associate with the alarm.
	Tags []Tag `json:"tags,omitempty"`

	// The value against which the specified statistic is compared.
	Threshold *float64 `json:"threshold"`

	// ThreshouldMetricId If this is an alarm based on an anomaly detection model.
	ThresholdMetricId *string `json:"thresholdMetricId,omitempty"`

	// TreatMissingData Sets how this alarm is to handle missing data points.
	// Valid Values: breaching | notBreaching | ignore | missing
	TreatMissingData *string `json:"treatMissingData,omitempty"`

	// Unit The unit of measure for the statistic.
	Unit cloudwatchtypes.StandardUnit `json:"unit,omitempty"`
}

// AlarmsStatus defines the observed state of Alarms
type AlarmsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Configured  bool    `json:"configured"`
	Error       *string `json:"error"`
	ErroMessage *string `json:"errorMessage"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Alarms is the Schema for the alarms API
type Alarms struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlarmsSpec   `json:"spec,omitempty"`
	Status AlarmsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AlarmsList contains a list of Alarms
type AlarmsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alarms `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Alarms{}, &AlarmsList{})
}
