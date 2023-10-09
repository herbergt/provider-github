/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	branch "github.com/herbergt/provider-github/internal/controller/branch/branch"
	providerconfig "github.com/herbergt/provider-github/internal/controller/providerconfig"
	repository "github.com/herbergt/provider-github/internal/controller/repository/repository"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		branch.Setup,
		providerconfig.Setup,
		repository.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
