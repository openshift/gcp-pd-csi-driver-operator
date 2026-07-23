package operator

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"

	"github.com/openshift/gcp-pd-csi-driver-operator/assets"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resource/resourceapply"
)

// sccPrerequisiteAssets lists ServiceAccounts and SCC RBAC that must exist
// before CSIControllerSet.Run starts the Deployment/DaemonSet controllers.
// Without this, the Deployment controller can create pods before the
// StaticResourcesController applies the ClusterRole+Binding, causing SCC
// admission to deny pod creation (hostnetwork-v2 / privileged).
var sccPrerequisiteAssets = []string{
	"controller_sa.yaml",
	"node_sa.yaml",
	"rbac/hostnetwork_role.yaml",
	"rbac/controller_hostnetwork_binding.yaml",
	"rbac/privileged_role.yaml",
	"rbac/node_privileged_binding.yaml",
}

func ensureSCCPrerequisites(ctx context.Context, kubeClient kubernetes.Interface, recorder events.Recorder) error {
	klog.Info("Applying SCC prerequisite resources before starting CSI controllers")
	results := resourceapply.ApplyDirectly(
		ctx,
		resourceapply.NewKubeClientHolder(kubeClient),
		recorder,
		resourceapply.NewResourceCache(),
		assets.ReadFile,
		sccPrerequisiteAssets...,
	)
	var errs []error
	for _, result := range results {
		if result.Error != nil {
			errs = append(errs, fmt.Errorf("%s: %w", result.File, result.Error))
			continue
		}
		klog.V(2).Infof("Applied SCC prerequisite %s (changed=%v)", result.File, result.Changed)
	}
	if len(errs) > 0 {
		return fmt.Errorf("failed to apply SCC prerequisites: %v", errs)
	}
	return nil
}
