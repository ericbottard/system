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
	"strings"

	"k8s.io/apimachinery/pkg/api/equality"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"github.com/projectriff/system/pkg/validation"
)

// +kubebuilder:webhook:path=/validate-streaming-projectriff-io-v1alpha1-pulsarprovider,mutating=false,failurePolicy=fail,groups=streaming.projectriff.io,resources=pulsarproviders,verbs=create;update,versions=v1alpha1,name=pulsarproviders.streaming.projectriff.io

var (
	_ webhook.Validator         = &PulsarProvider{}
	_ validation.FieldValidator = &PulsarProvider{}
)

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *PulsarProvider) ValidateCreate() error {
	return r.Validate().ToAggregate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *PulsarProvider) ValidateUpdate(old runtime.Object) error {
	// TODO check for immutable fields
	return r.Validate().ToAggregate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *PulsarProvider) ValidateDelete() error {
	return nil
}

func (r *PulsarProvider) Validate() validation.FieldErrors {
	errs := validation.FieldErrors{}

	errs = errs.Also(r.Spec.Validate().ViaField("spec"))

	return errs
}

func (s *PulsarProviderSpec) Validate() validation.FieldErrors {
	if equality.Semantic.DeepEqual(s, &PulsarProviderSpec{}) {
		return validation.ErrMissingField(validation.CurrentField)
	}

	errs := validation.FieldErrors{}

	if s.ServiceURL == "" {
		errs = errs.Also(validation.ErrMissingField("serviceURL"))
	} else if !(strings.HasPrefix(s.ServiceURL, "pulsar://") || strings.HasPrefix(s.ServiceURL, "pulsar+ssl://")) {
		errs = errs.Also(validation.FieldErrors{
			field.Invalid(field.NewPath("serviceURL"), "must use 'pulsar://' or 'pulsar+ssl://' scheme"),
		})
	}

	return errs
}
