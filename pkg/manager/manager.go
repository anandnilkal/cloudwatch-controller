package manager

import (
	"context"
	"errors"
	"fmt"

	jujuerrors "github.com/juju/errors"

	cloudwatchv1alpha1 "github.com/anandnilkal/cloudwatch-controller/api/v1alpha1"
	cloudwatchmanager "github.com/anandnilkal/cloudwatch-controller/pkg/cloudwatch"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type ClientManager struct {
	cloudwatchClient *cloudwatchmanager.CloudwatchClient
}

var clientManager ClientManager

func Initialization() error {
	var err error
	clientManager.cloudwatchClient, err = cloudwatchmanager.NewCloudwatchClient(context.Background(), "us-west-2")
	return err
}

func CreateCloudwatchAlarm(ctx context.Context, alarm *cloudwatchv1alpha1.Alarms) error {
	logger := log.FromContext(ctx)
	if clientManager.cloudwatchClient == nil {
		return errors.New("CloudWatch Client un-initialized")
	}
	_, err := clientManager.cloudwatchClient.CreateCloudwatchAlarm(ctx, alarm)
	if err != nil {
		logger.Error(err, fmt.Sprintf("CloudWatch Alarm Creation failed: %s", alarm.Spec.Name))
		return err
	}

	logger.V(0).Info(fmt.Sprintf("CloudWatch Alarm created: %s", alarm.Spec.Name))
	return nil

}

func CheckAndCleanupAlarm(ctx context.Context, name string, namespace string) (bool, error) {
	logger := log.FromContext(ctx)
	if clientManager.cloudwatchClient == nil {
		return false, errors.New("CloudWatch Client un-initialized")
	}
	_, err := clientManager.cloudwatchClient.DeleteCloudwatchAlarm(ctx, name)
	if err != nil {
		if jujuerrors.IsNotFound(err) {
			return false, nil
		}
		return true, err
	}
	logger.V(0).Info(fmt.Sprintf("Deleted CloudWatch Alarm: %s", name))
	return false, err
}
