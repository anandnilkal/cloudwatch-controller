package manager

import (
	"context"
	"errors"

	cloudwatchv1alpha1 "github.com/anandnilkal/cloudwatch-controller/api/v1alpha1"
	cloudwatchmanager "github.com/anandnilkal/cloudwatch-controller/pkg/cloudwatch"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type ClientManager struct {
	cloudwatchClient *cloudwatchmanager.CloudwatchClient
}

var clientManager *ClientManager

func Initialization() error {
	var err error
	clientManager.cloudwatchClient, err = cloudwatchmanager.NewCloudwatchClient(context.Background(), "us-west-2")
	return err
}

func CreateCloudwatchAlarm(ctx context.Context, alarm *cloudwatchv1alpha1.Alarms) error {
	logger := log.FromContext(ctx)
	if clientManager.cloudwatchClient == nil {
		return errors.New("CloudWatch Client un-initialized")
		// return fmt.Errorf("CloudWatch manager client not created")
	}
	alarmOut, err := clientManager.cloudwatchClient.CreateCloudwatchAlarm(ctx, alarm)
	if err != nil {
		logger.Error(err, "CloudWatch Alarm Creation: %s", alarm.Spec.Name)
		return err
	}

	logger.V(0).Info("CloudWatch Alarm created: %+v", alarmOut)
	return nil

}
