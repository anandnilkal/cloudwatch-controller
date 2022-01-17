package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"sigs.k8s.io/controller-runtime/pkg/log"

	cloudwatchv1alpha1 "github.com/anandnilkal/cloudwatch-controller/api/v1alpha1"
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

func (c *CloudwatchClient) CreateCloudwatchAlarm(ctx context.Context, alarm *cloudwatchv1alpha1.Alarms) (*cloudwatch.PutMetricAlarmOutput, error) {
	logger := log.FromContext(ctx)
	logger.V(0).Info("Creating CloudWatch Alarm: %s", alarm.Name)
	return c.Client.PutMetricAlarm(ctx, &cloudwatch.PutMetricAlarmInput{
		AlarmName:                        &alarm.Name,
		ComparisonOperator:               alarm.Spec.ComparisonOperator,
		EvaluationPeriods:                alarm.Spec.EvaluationPeriod,
		ActionsEnabled:                   new(bool),
		AlarmActions:                     []string{alarm.Spec.AlarmActions[0]},
		AlarmDescription:                 alarm.Spec.AlarmDescription,
		DatapointsToAlarm:                alarm.Spec.DatapointsToAlarm,
		Dimensions:                       []types.Dimension{{Name: alarm.Spec.Dimensions[0].Name, Value: alarm.Spec.Dimensions[0].Value}},
		EvaluateLowSampleCountPercentile: new(string),
		ExtendedStatistic:                new(string),
		InsufficientDataActions:          []string{},
		MetricName:                       alarm.Spec.MetricName,
		Metrics:                          []types.MetricDataQuery{},
		Namespace:                        alarm.Spec.Namespace,
		OKActions:                        []string{alarm.Spec.OKActions[0]},
		Period:                           alarm.Spec.Period,
		Statistic:                        alarm.Spec.Statistic,
		Tags:                             []types.Tag{},
		Threshold:                        alarm.Spec.Threshold,
		ThresholdMetricId:                new(string),
		TreatMissingData:                 new(string),
		Unit:                             alarm.Spec.Unit,
	})
}

func (c *CloudwatchClient) DeleteCloudwatchAlarm(ctx context.Context, name string) (*cloudwatch.DeleteAlarmsOutput, error) {
	return c.Client.DeleteAlarms(ctx, &cloudwatch.DeleteAlarmsInput{
		AlarmNames: []string{name},
	})
}

func (c *CloudwatchClient) DescribeCloudwatchAlarm(ctx context.Context, name string) (*cloudwatch.DescribeAlarmsOutput, error) {
	return c.Client.DescribeAlarms(ctx, &cloudwatch.DescribeAlarmsInput{
		AlarmNames: []string{name},
	})
}
