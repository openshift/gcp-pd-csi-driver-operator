package deploymentcontroller

import (
	"context"
	"fmt"
	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
	"k8s.io/klog/v2"
	"time"
)

const (
	resyncInterval = 1 * time.Minute
)

type TimeoutController struct {
	operatorClient v1helpers.OperatorClient
	name           string
	timeout        time.Duration
}

func NewTimeoutController(operatorClient v1helpers.OperatorClient, name string, timeout time.Duration, recorder events.Recorder) factory.Controller {
	c := &TimeoutController{
		operatorClient: operatorClient,
		name:           name,
		timeout:        timeout,
	}
	return factory.New().WithInformers(operatorClient.Informer()).WithSync(c.sync).ResyncEvery(resyncInterval).ToController("TimeoutController", recorder)
}

func (c TimeoutController) sync(ctx context.Context, syncCtx factory.SyncContext) error {
	_, currentStatus, _, err := c.operatorClient.GetOperatorState()
	if err != nil {
		klog.Errorf("Error getting operator status: %v", err)
		return err
	}

	operatorCondition := v1helpers.FindOperatorCondition(currentStatus.Conditions, c.name+operatorv1.OperatorStatusTypeProgressing)

	if operatorCondition != nil && operatorCondition.Status == operatorv1.ConditionTrue && operatorCondition.LastTransitionTime.Add(c.timeout).Before(time.Now()) {
		_, _, updateErr := v1helpers.UpdateStatus(ctx, c.operatorClient, v1helpers.UpdateConditionFn(operatorv1.OperatorCondition{
			Type:    c.name + "ProgressingTimeout" + "Degraded",
			Status:  operatorv1.ConditionTrue,
			Reason:  "SyncTimeoutError",
			Message: fmt.Sprintf("Operator condition %s failed to transition to %s in the last %0.f minutes.", operatorCondition.Type, operatorv1.ConditionFalse, c.timeout.Minutes()),
		}))

		if updateErr != nil {
			klog.Warningf("Updating status of %q failed: %v", c.name, updateErr)
			return updateErr
		}
		return nil
	}

	_, _, updateErr := v1helpers.UpdateStatus(ctx, c.operatorClient, v1helpers.UpdateConditionFn(operatorv1.OperatorCondition{
		Type:   c.name + "ProgressingTimeout" + "Degraded",
		Status: operatorv1.ConditionFalse,
	}))
	if updateErr != nil {
		klog.Warningf("Updating status of %q failed: %v", c.name, updateErr)
		return updateErr
	}

	return nil
}
