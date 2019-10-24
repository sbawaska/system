/*
Copyright 2019 the original author or authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *Function) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-build-projectriff-io-v1alpha1-function,mutating=true,failurePolicy=fail,groups=build.projectriff.io,resources=functions,verbs=create;update,versions=v1alpha1,name=functions.build.projectriff.io

var _ webhook.Defaulter = &Function{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Function) Default() {
	// TODO implement
}

// +kubebuilder:webhook:path=/validate-build-projectriff-io-v1alpha1-function,mutating=false,failurePolicy=fail,groups=build.projectriff.io,resources=functions,verbs=create;update,versions=v1alpha1,name=functions.build.projectriff.io

var _ webhook.Validator = &Function{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Function) ValidateCreate() error {
	// TODO implement
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Function) ValidateUpdate(old runtime.Object) error {
	// TODO implement
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Function) ValidateDelete() error {
	// TODO implement
	return nil
}
