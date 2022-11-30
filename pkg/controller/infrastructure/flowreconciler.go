package infrastructure

import (
	"context"
	"fmt"

	azureclient "github.com/gardener/gardener-extension-provider-azure/pkg/azure/client"
	"github.com/gardener/gardener-extension-provider-azure/pkg/controller/infrastructure/infraflow"
	"github.com/gardener/gardener-extension-provider-azure/pkg/internal"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
)

func NewFlowReconciler(ctx context.Context, a *actuator, infra *extensionsv1alpha1.Infrastructure) (*infraflow.FlowReconciler, error) {
	client := a.Client()
	if client == nil {
		return nil, fmt.Errorf("infrastructure actuator has no client set")
	}
	auth, err := internal.GetClientAuthData(ctx, client, infra.Spec.SecretRef, false)
	if err != nil {
		return nil, err
	}
	factory, err := azureclient.NewAzureClientFactoryV2(*auth)
	if err != nil {
		return nil, err
	}
	reconciler := infraflow.NewFlowReconciler(factory)
	return reconciler, nil
}