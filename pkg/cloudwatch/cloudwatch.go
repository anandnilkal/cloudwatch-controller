package cloudwatch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"sigs.k8s.io/controller-runtime/pkg/log"

	jujuerrors "github.com/juju/errors"

	cloudwatchv1alpha1 "github.com/anandnilkal/cloudwatch-controller/api/v1alpha1"
	cloudwatchtypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

type CloudwatchClient struct {
	Client *cloudwatch.Client
	Region string
}

func NewCloudwatchClient(ctx context.Context, region string) (*CloudwatchClient, error) {
	logger := log.FromContext(ctx)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		logger.Error(err, "failed to load SDK configuration, %v", err)
		return nil, err
	}
	return &CloudwatchClient{
		Client: cloudwatch.NewFromConfig(cfg),
		Region: region,
	}, nil
}

// default returns GreaterThanThreshold
func getOperator(operator string) cloudwatchtypes.ComparisonOperator {
	switch operator {
	case "GreaterThanOrEqualToThreshold":
		return cloudwatchtypes.ComparisonOperatorGreaterThanOrEqualToThreshold
	case "GreaterThanThreshold":
		return cloudwatchtypes.ComparisonOperatorGreaterThanThreshold
	case "LessThanThreshold":
		return cloudwatchtypes.ComparisonOperatorLessThanThreshold
	case "LessThanOrEqualToThreshold":
		return cloudwatchtypes.ComparisonOperatorLessThanOrEqualToThreshold
	case "LessThanLowerOrGreaterThanUpperThreshold":
		return cloudwatchtypes.ComparisonOperatorLessThanLowerOrGreaterThanUpperThreshold
	case "LessThanLowerThreshold":
		return cloudwatchtypes.ComparisonOperatorLessThanLowerThreshold
	case "GreaterThanUpperThreshold":
		return cloudwatchtypes.ComparisonOperatorGreaterThanUpperThreshold
	}
	return cloudwatchtypes.ComparisonOperatorGreaterThanThreshold
}

func getStatistic(stat string) cloudwatchtypes.Statistic {
	switch stat {
	case "SampleCount":
		return cloudwatchtypes.StatisticSampleCount
	case "Average":
		return cloudwatchtypes.StatisticAverage
	case "Sum":
		return cloudwatchtypes.StatisticSum
	case "Minimum":
		return cloudwatchtypes.StatisticMinimum
	case "Maximum":
		return cloudwatchtypes.StatisticMaximum
	}
	return cloudwatchtypes.StatisticAverage
}

func (c *CloudwatchClient) validateAlarmInput(alarm *cloudwatchv1alpha1.Alarms) error {
	if string(alarm.Spec.Statistic) != "" && alarm.Spec.ExtendedStatistic != nil {
		return jujuerrors.NotValidf("Statistics and ExtendedStatistic both provided")
	}
	if len(alarm.Spec.Metrics) != 0 && alarm.Spec.MetricName != nil {
		return jujuerrors.NotValidf("MetricName and Metrics both provided")
	}
	return nil
}

func (c *CloudwatchClient) populateAlarmInput(alarm *cloudwatchv1alpha1.Alarms) (*cloudwatch.PutMetricAlarmInput, error) {
	err := c.validateAlarmInput(alarm)
	if err != nil {
		return nil, err
	}
	putMetricAlarmInput := &cloudwatch.PutMetricAlarmInput{}
	if alarm.Name != "" {
		putMetricAlarmInput.AlarmName = &alarm.Name
	}
	if string(alarm.Spec.ComparisonOperator) != "" {
		putMetricAlarmInput.ComparisonOperator = getOperator(string(alarm.Spec.ComparisonOperator))
	}
	if alarm.Spec.EvaluationPeriod != nil {
		putMetricAlarmInput.EvaluationPeriods = alarm.Spec.EvaluationPeriod
	}
	if len(alarm.Spec.AlarmActions) != 0 {
		putMetricAlarmInput.AlarmActions = append(putMetricAlarmInput.AlarmActions, alarm.Spec.AlarmActions...)
	}
	if alarm.Spec.AlarmDescription != nil {
		putMetricAlarmInput.AlarmDescription = alarm.Spec.AlarmDescription
	}
	if alarm.Spec.DatapointsToAlarm != nil {
		putMetricAlarmInput.DatapointsToAlarm = alarm.Spec.DatapointsToAlarm
	}
	if len(alarm.Spec.Dimensions) != 0 {
		for i := 0; i < len(alarm.Spec.Dimensions); i++ {
			putMetricAlarmInput.Dimensions = append(putMetricAlarmInput.Dimensions, cloudwatchtypes.Dimension{
				Name:  alarm.Spec.Dimensions[i].Name,
				Value: alarm.Spec.Dimensions[i].Value,
			})
		}
	}
	if alarm.Spec.EvaluateLowSampleCountPercentile != nil {
		putMetricAlarmInput.EvaluateLowSampleCountPercentile = alarm.Spec.EvaluateLowSampleCountPercentile
	}
	if alarm.Spec.ExtendedStatistic != nil {
		putMetricAlarmInput.ExtendedStatistic = alarm.Spec.ExtendedStatistic
	}
	if len(alarm.Spec.InsufficientDataActions) != 0 {
		putMetricAlarmInput.InsufficientDataActions = append(putMetricAlarmInput.InsufficientDataActions, alarm.Spec.InsufficientDataActions...)
	}
	if alarm.Spec.MetricName != nil {
		putMetricAlarmInput.MetricName = alarm.Spec.MetricName
	}
	if len(alarm.Spec.Metrics) != 0 {
		for i, _ := range alarm.Spec.Metrics {
			var dimensions []cloudwatchtypes.Dimension
			for j, _ := range alarm.Spec.Metrics[i].MetricStat.Metric.Dimensions {
				dimensions = append(dimensions, cloudwatchtypes.Dimension{
					Name:  alarm.Spec.Metrics[i].MetricStat.Metric.Dimensions[j].Name,
					Value: alarm.Spec.Metrics[i].MetricStat.Metric.Dimensions[j].Value,
				})
			}
			putMetricAlarmInput.Metrics = append(putMetricAlarmInput.Metrics, cloudwatchtypes.MetricDataQuery{
				Id:         alarm.Spec.Metrics[i].Id,
				AccountId:  alarm.Spec.Metrics[i].AccountId,
				Expression: alarm.Spec.Metrics[i].Expression,
				Label:      alarm.Spec.Metrics[i].Label,
				MetricStat: &cloudwatchtypes.MetricStat{
					Metric: &cloudwatchtypes.Metric{
						Dimensions: dimensions,
					},
					Period: alarm.Spec.Metrics[i].MetricStat.Period,
					Stat:   alarm.Spec.Metrics[i].MetricStat.Stat,
					Unit:   alarm.Spec.Metrics[i].MetricStat.Unit,
				},
				Period:     alarm.Spec.Metrics[i].Period,
				ReturnData: alarm.Spec.Metrics[i].ReturnData,
			})
		}
	}
	if alarm.Spec.Namespace != nil {
		putMetricAlarmInput.Namespace = alarm.Spec.Namespace
	}
	if len(alarm.Spec.OKActions) != 0 {
		putMetricAlarmInput.OKActions = append(putMetricAlarmInput.OKActions, alarm.Spec.OKActions...)
	}
	if alarm.Spec.Period != nil {
		putMetricAlarmInput.Period = alarm.Spec.Period
	}
	if string(alarm.Spec.Statistic) != "" {
		putMetricAlarmInput.Statistic = getStatistic(string(alarm.Spec.Statistic))
	}
	if len(alarm.Spec.Tags) != 0 {
		for i, _ := range alarm.Spec.Tags {
			putMetricAlarmInput.Tags = append(putMetricAlarmInput.Tags, cloudwatchtypes.Tag{
				Key:   alarm.Spec.Tags[i].Key,
				Value: alarm.Spec.Tags[i].Value,
			})
		}
	}
	if alarm.Spec.Threshold != nil {
		putMetricAlarmInput.Threshold = alarm.Spec.Threshold
	}
	if alarm.Spec.ThresholdMetricId != nil {
		putMetricAlarmInput.ThresholdMetricId = alarm.Spec.ThresholdMetricId
	}
	if alarm.Spec.TreatMissingData != nil {
		putMetricAlarmInput.TreatMissingData = alarm.Spec.TreatMissingData
	}
	if string(alarm.Spec.Unit) != "" {
		putMetricAlarmInput.Unit = alarm.Spec.Unit
	}

	return putMetricAlarmInput, nil
}

func (c *CloudwatchClient) CreateCloudwatchAlarm(ctx context.Context, alarm *cloudwatchv1alpha1.Alarms) (*cloudwatch.PutMetricAlarmOutput, error) {
	logger := log.FromContext(ctx)
	logger.V(0).Info(fmt.Sprintf("Creating CloudWatch Alarm: %s", alarm.Name))
	putMetricAlarmInput, err := c.populateAlarmInput(alarm)
	if err != nil {
		return nil, err
	}

	return c.Client.PutMetricAlarm(ctx, putMetricAlarmInput)
}

func (c *CloudwatchClient) DeleteCloudwatchAlarm(ctx context.Context, name string) (*cloudwatch.DeleteAlarmsOutput, error) {
	logger := log.FromContext(ctx)
	logger.V(0).Info(fmt.Sprintf("Deleting CloudWatch Alarm: %s", name))
	return c.Client.DeleteAlarms(ctx, &cloudwatch.DeleteAlarmsInput{
		AlarmNames: []string{name},
	})
}

func (c *CloudwatchClient) DescribeCloudwatchAlarm(ctx context.Context, name string) (*cloudwatch.DescribeAlarmsOutput, error) {
	return c.Client.DescribeAlarms(ctx, &cloudwatch.DescribeAlarmsInput{
		AlarmNames: []string{name},
	})
}
