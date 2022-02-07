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

package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"sigs.k8s.io/controller-runtime/pkg/predicate"

	cloudwatchv1alpha1 "github.com/anandnilkal/cloudwatch-controller/api/v1alpha1"
	"github.com/anandnilkal/cloudwatch-controller/pkg/manager"
)

// AlarmsReconciler reconciles a Alarms object
type AlarmsReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cloudwatch.anandnilkal.io,resources=alarms,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cloudwatch.anandnilkal.io,resources=alarms/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cloudwatch.anandnilkal.io,resources=alarms/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Alarms object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *AlarmsReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here

	if !manager.Initialized() {
		return ctrl.Result{Requeue: true}, nil
	}

	var alarm cloudwatchv1alpha1.Alarms
	err := r.Get(ctx, req.NamespacedName, &alarm)
	if err != nil {
		if !errors.IsNotFound(err) {
			logger.Error(err, fmt.Sprintf("error (%s)", err.Error()))
			return ctrl.Result{}, nil
		}
		requeue, _ := manager.CheckAndCleanupAlarm(ctx, req.NamespacedName.Name, req.NamespacedName.Namespace, *alarm.Spec.Region)
		if requeue {
			return ctrl.Result{Requeue: requeue}, nil
		}
		return ctrl.Result{}, nil
	}

	// logger.V(0).Info(fmt.Sprintf("alarm data: %+v", alarm))

	if alarm.Status.Configured {
		return ctrl.Result{}, nil
	}
	err = manager.CreateCloudwatchAlarm(ctx, &alarm)
	if err != nil {
		logger.Error(err, fmt.Sprintf("failed: %s", err))
		errmessage := err.Error()
		alarm.Status.Configured = false
		alarm.Status.Error = &errmessage
		alarm.Status.ErroMessage = &errmessage
		err = r.Status().Update(ctx, &alarm)
		if err != nil {
			logger.Error(err, fmt.Sprintf("Status update failed: %s", err))
		}
		return ctrl.Result{Requeue: true}, err
	}

	alarm.Status.Configured = true
	err = r.Status().Update(ctx, &alarm)
	if err != nil {
		logger.Error(err, fmt.Sprintf("Status update failed: %s", err))
		return ctrl.Result{Requeue: true}, nil
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AlarmsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	logger := log.FromContext(context.Background())
	err := manager.Initialization()
	if err != nil {
		logger.Error(err, "SetupWithManager failed: %s", err)
		return err
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&cloudwatchv1alpha1.Alarms{}).WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
