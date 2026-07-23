package operator

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/utils/clock"

	"github.com/openshift/library-go/pkg/operator/events"
)

func TestEnsureSCCPrerequisites(t *testing.T) {
	ctx := context.Background()
	client := fake.NewSimpleClientset()
	recorder := events.NewInMemoryRecorder("test", clock.RealClock{})

	if err := ensureSCCPrerequisites(ctx, client, recorder); err != nil {
		t.Fatalf("ensureSCCPrerequisites failed: %v", err)
	}

	checks := []struct {
		name  string
		check func() error
	}{
		{
			name: "controller service account",
			check: func() error {
				_, err := client.CoreV1().ServiceAccounts(defaultNamespace).Get(ctx, "gcp-pd-csi-driver-controller-sa", metav1.GetOptions{})
				return err
			},
		},
		{
			name: "node service account",
			check: func() error {
				_, err := client.CoreV1().ServiceAccounts(defaultNamespace).Get(ctx, "gcp-pd-csi-driver-node-sa", metav1.GetOptions{})
				return err
			},
		},
		{
			name: "hostnetwork role",
			check: func() error {
				_, err := client.RbacV1().ClusterRoles().Get(ctx, "gcp-pd-hostnetwork-role", metav1.GetOptions{})
				return err
			},
		},
		{
			name: "controller hostnetwork binding",
			check: func() error {
				_, err := client.RbacV1().ClusterRoleBindings().Get(ctx, "gcp-pd-controller-hostnetwork-binding", metav1.GetOptions{})
				return err
			},
		},
		{
			name: "privileged role",
			check: func() error {
				_, err := client.RbacV1().ClusterRoles().Get(ctx, "gcp-pd-privileged-role", metav1.GetOptions{})
				return err
			},
		},
		{
			name: "node privileged binding",
			check: func() error {
				_, err := client.RbacV1().ClusterRoleBindings().Get(ctx, "gcp-pd-node-privileged-binding", metav1.GetOptions{})
				return err
			},
		},
	}
	for _, c := range checks {
		if err := c.check(); err != nil {
			t.Errorf("%s missing after ensureSCCPrerequisites: %v", c.name, err)
		}
	}

	// Idempotent: second call must also succeed.
	if err := ensureSCCPrerequisites(ctx, client, recorder); err != nil {
		t.Fatalf("ensureSCCPrerequisites second call failed: %v", err)
	}
}
