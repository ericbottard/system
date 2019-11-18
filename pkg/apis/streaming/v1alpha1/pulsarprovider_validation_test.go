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
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/google/go-cmp/cmp"

	"github.com/projectriff/system/pkg/validation"
)

func TestValidatePulsarProvider(t *testing.T) {
	for _, c := range []struct {
		name     string
		target   *PulsarProvider
		expected validation.FieldErrors
	}{{
		name:     "empty",
		target:   &PulsarProvider{},
		expected: validation.ErrMissingField("spec"),
	}, {
		name: "valid",
		target: &PulsarProvider{
			Spec: PulsarProviderSpec{
				ServiceURL: "pulsar://localhost:6650",
			},
		},
		expected: validation.FieldErrors{},
	}} {
		t.Run(c.name, func(t *testing.T) {
			actual := c.target.Validate()
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("validatePulsarProvider(%s) (-expected, +actual) = %v", c.name, diff)
			}
		})
	}
}

func TestValidatePulsarProviderSpec(t *testing.T) {
	for _, c := range []struct {
		name     string
		target   *PulsarProviderSpec
		expected validation.FieldErrors
	}{{
		name:     "empty",
		target:   &PulsarProviderSpec{},
		expected: validation.ErrMissingField(validation.CurrentField),
	}, {
		name: "valid",
		target: &PulsarProviderSpec{
			ServiceURL: "pulsar://localhost:6650",
		},
		expected: validation.FieldErrors{},
	}, {
		name: "wrong-scheme",
		target: &PulsarProviderSpec{
			ServiceURL: "localhost:6650",
		},
		expected: validation.FieldErrors{field.Required(field.NewPath("serviceURL"), "serviceURL must use 'pulsar://' or 'pulsar+ssl://' scheme")},
	}} {
		t.Run(c.name, func(t *testing.T) {
			actual := c.target.Validate()
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("validatePulsarProviderSpec(%s) (-expected, +actual) = %v", c.name, diff)
			}
		})
	}
}
