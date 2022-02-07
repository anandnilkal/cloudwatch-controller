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
	return err
}

func CreateCloudwatchAlarm(ctx context.Context, alarm *cloudwatchv1alpha1.Alarms) error {
	logger := log.FromContext(ctx)
	if _, ok := clientManager.ClientList[*alarm.Spec.Region]; !ok {
		return errors.New("CloudWatch Client un-initialized")
	}
	_, err := clientManager.ClientList[*alarm.Spec.Region].CreateCloudwatchAlarm(ctx, alarm)
	if err != nil {
		logger.Error(err, fmt.Sprintf("CloudWatch Alarm Creation failed: %s", alarm.Spec.Name))
		return err
	}

	var alarmDescription *cloudwatch.DescribeAlarmsOutput
	alarmDescription, err = clientManager.ClientList[*alarm.Spec.Region].DescribeCloudwatchAlarm(ctx, alarm.Name)
	if err != nil {
		logger.Error(err, fmt.Sprintf("Cloudwatch Alarm describe failed: %s", alarm.Spec.Name))
		return err
	}

	if len(alarmDescription.MetricAlarms) != 0 {
		for _, alarmOut := range alarmDescription.MetricAlarms {
			alarmArn := alarmOut.AlarmArn
			_, err := clientManager.ClientList[*alarm.Spec.Region].TagAlarmResource(ctx, *alarmArn, alarm.Spec.Tags)
			if err != nil {
				logger.Error(err, fmt.Sprintf("Cloudwatch alarm resource tagging failed: %s", alarm.Spec.Name))
				return err
			}
		}
	}

	logger.V(0).Info(fmt.Sprintf("CloudWatch Alarm created: %s", alarm.Spec.Name))
	return nil

}

func CheckAndCleanupAlarm(ctx context.Context, name, namespace, region string) (bool, error) {
	logger := log.FromContext(ctx)
	if _, ok := clientManager.ClientList[region]; !ok {
		return false, errors.New("CloudWatch Client un-initialized")
	}
	_, err := clientManager.ClientList[region].DeleteCloudwatchAlarm(ctx, name)
	if err != nil {
		if jujuerrors.IsNotFound(err) {
			return false, nil
		}
		return true, err
	}
	logger.V(0).Info(fmt.Sprintf("Deleted CloudWatch Alarm: %s", name))
	return false, err
}
