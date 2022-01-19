package cloudwatch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
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
