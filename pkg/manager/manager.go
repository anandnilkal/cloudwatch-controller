package manager

import (
	"context"
	"errors"
	"fmt"

	jujuerrors "github.com/juju/errors"

	cloudwatchv1alpha1 "github.com/anandnilkal/cloudwatch-controller/api/v1alpha1"
	cloudwatchmanager "github.com/anandnilkal/cloudwatch-controller/pkg/cloudwatch"
	cloudwatch "github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type ClientManager struct {
	ClientList map[string]*cloudwatchmanager.CloudwatchClient
}

var clientManager ClientManager
var initialized bool

// Regions currently able to create alarms at
var Regions = []string{"us-east-1", "us-west-2"}

func Initialization() error {
	var err error
	clientManager.ClientList = make(map[string]*cloudwatchmanager.CloudwatchClient)
	for _, region := range Regions {
		client, err := cloudwatchmanager.NewCloudwatchClient(context.Background(), region)
		if err != nil {
			return err
		}
		clientManager.ClientList[region] = client
	}
	initialized = true
	return err
}

func Initialized() bool {
	return initialized
}

func CreateCloudwatchAlarm(ctx context.Context, alarm *cloudwatchv1alpha1.Alarms) error {
	logger := log.FromContext(ctx)
	region := "us-west-2"
	if alarm.Spec.Region != nil && len(*alarm.Spec.Region) != 0 {
		region = *alarm.Spec.Region
	}
	if _, ok := clientManager.ClientList[region]; !ok {
		return errors.New("CloudWatch Client un-initialized")
	}
	_, err := clientManager.ClientList[region].CreateCloudwatchAlarm(ctx, alarm)
	if err != nil {
		logger.Error(err, fmt.Sprintf("CloudWatch Alarm Creation failed: %s", alarm.Spec.Name))
		return err
	}

	var alarmDescription *cloudwatch.DescribeAlarmsOutput
	alarmDescription, err = clientManager.ClientList[region].DescribeCloudwatchAlarm(ctx, alarm.Name)
	if err != nil {
		logger.Error(err, fmt.Sprintf("Cloudwatch Alarm describe failed: %s", alarm.Spec.Name))
		return err
	}

	if len(alarmDescription.MetricAlarms) != 0 {
		for _, alarmOut := range alarmDescription.MetricAlarms {
			alarmArn := alarmOut.AlarmArn
			_, err := clientManager.ClientList[region].TagAlarmResource(ctx, *alarmArn, alarm.Spec.Tags)
			if err != nil {
				logger.Error(err, fmt.Sprintf("Cloudwatch alarm resource tagging failed: %s", alarm.Spec.Name))
				return err
			}
		}
	}

	logger.V(0).Info(fmt.Sprintf("CloudWatch Alarm created: %s", alarm.Spec.Name))
	return nil

}

func CheckAndCleanupAlarm(ctx context.Context, name, namespace string) (bool, error) {
	logger := log.FromContext(ctx)
	for _, region := range Regions {
		if _, ok := clientManager.ClientList[region]; !ok {
			return false, errors.New("CloudWatch Client un-initialized")
		}
		var alarmDescription *cloudwatch.DescribeAlarmsOutput
		alarmDescription, err := clientManager.ClientList[region].DescribeCloudwatchAlarm(ctx, name)
		if err != nil {
			logger.Error(err, fmt.Sprintf("Cloudwatch Alarm describe failed: %s", name))
			return true, err
		}
		if len(alarmDescription.MetricAlarms) == 0 {
			continue
		}
		if len(alarmDescription.MetricAlarms) == 1 {
			_, err = clientManager.ClientList[region].DeleteCloudwatchAlarm(ctx, name)
			if err != nil {
				if jujuerrors.IsNotFound(err) {
					return false, nil
				}
				return true, err
			}
			logger.V(0).Info(fmt.Sprintf("Deleted CloudWatch Alarm: %s", name))
			break
		}
	}
	return false, nil
}
