/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package v1

import (
	"text/template/parse"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var configconstraintlog = logf.Log.WithName("configconstraint-resource")

func (r *ConfigConstraint) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:path=/validate-apps-kubeblocks-io-v1-configconstraint,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps.kubeblocks.io,resources=configconstraints,verbs=create;update,versions=v1,name=vconfigconstraint.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &ConfigConstraint{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ConfigConstraint) ValidateCreate() (admission.Warnings, error) {
	configconstraintlog.Info("validate create", "name", r.Name)

	if err := r.validateShellTrigger(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *ConfigConstraint) validateShellTrigger() error {
	if !r.Spec.ShellTrigger() {
		return nil
	}
	if r.Spec.BatchReload() {
		return validateBatchParametersTemplate(r.Spec.DynamicReloadAction.ShellTrigger.BatchParametersTemplate)
	}
	return nil
}

func validateBatchParametersTemplate(template string) error {
	if template == "" {
		return field.Invalid(field.NewPath("spec.dynamicReloadAction.shellTrigger.batchParametersTemplate"), nil, "batchParametersTemplate is empty")
	}

	tr := parse.New(template)
	_, err := tr.Parse(template, "", "", make(map[string]*parse.Tree))
	return err
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ConfigConstraint) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	configconstraintlog.Info("validate update", "name", r.Name)

	if err := r.validateShellTrigger(); err != nil {
		return nil, err
	}
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ConfigConstraint) ValidateDelete() (admission.Warnings, error) {
	configconstraintlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
