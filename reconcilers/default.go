package reconcilers

import (
	apibase "github.com/byuoitav/av-api/base"
)

/*
DefaultReconciler is the default reconciler. It does nothing.
*/
type DefaultReconciler struct {
}

/*
Fulfils the requirement for the interface
*/
func (r *DefaultReconciler) Reconcile(actions []apibase.ActionStructure) ([]apibase.ActionStructure, error) {
	return actions, nil
}
