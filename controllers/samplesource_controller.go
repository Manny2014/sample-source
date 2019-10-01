package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sourcesv1alpha1 "github.com/manny2014/sample-source/api/v1alpha1"
)

// SampleSourceReconciler reconciles a SampleSource object
type SampleSourceReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=sources.manny2014,resources=samplesources,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sources.manny2014,resources=samplesources/status,verbs=get;update;patch

func (r *SampleSourceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("samplesource", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *SampleSourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sourcesv1alpha1.SampleSource{}).
		Complete(r)
}
