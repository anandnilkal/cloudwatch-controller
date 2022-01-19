package v1alpha1

import (
	cloudwatchtypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

type Tag struct {

	// Key string that you can use to assign a value.
	Key *string `json:"key"`

	// Value for the specified tag key.
	Value *string `json:"value"`
}

type Dimension struct {
	// Name name of the dimension
	Name *string `json:"name"`

	// Value of the dimension
	Value *string `json:"value"`
}

type Metric struct {

	// The dimensions for the metric.
	Dimensions []Dimension `json:"dimensions"`

	// The name of the metric. This is a required field.
	MetricName *string `json:"metricName"`

	// The namespace of the metric.
	Namespace *string `json:"namespace"`
}

type MetricStat struct {

	// The metric to return, including the metric name, namespace, and dimensions.
	Metric *Metric `json:"metric"`

	// The granularity, in seconds, of the returned data points.
	Period *int32 `json:"period"`

	// The statistic to return. It can include any CloudWatch statistic or extended statistic.
	Stat *string `json:"stat"`

	// When you are using a Put operation, this defines what unit you want to use when storing the metric.
	Unit cloudwatchtypes.StandardUnit `json:"unit,omitempty"`
}

type MetricDataQuery struct {

	// Id short name
	Id *string `json:"id"`

	// AccountId ID of the account where the metrics are located
	AccountId *string `json:"accountId,omitempty"`

	// Expression math expression to be performed on the returned data, if this object is performing a math expression.
	Expression *string `json:"expression,omitempty"`

	// Label human-readable label for this metric or expression.
	Label *string `json:"label,omitempty"`

	// MetricStat The metric to be returned, along with statistics, period, and units.
	MetricStat *MetricStat `json:"metricStat,omitempty"`

	// Period The granularity, in seconds, of the returned data points.
	Period *int32 `json:"period,omitempty"`

	// ReturnData When used in GetMetricData, this option indicates whether to return the
	// timestamps and raw data values of this metric.
	ReturnData *bool `json:"returnData"`
}
