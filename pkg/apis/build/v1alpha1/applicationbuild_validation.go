/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	"context"

	"github.com/knative/pkg/apis"
	"k8s.io/apimachinery/pkg/api/equality"
)

func (ab *ApplicationBuild) Validate(ctx context.Context) *apis.FieldError {
	errs := &apis.FieldError{}
	errs = errs.Also(validateObjectMetadata(ab.GetObjectMeta()).ViaField("metadata"))
	errs = errs.Also(ab.Spec.Validate(ctx).ViaField("spec"))
	return errs
}

func (abs *ApplicationBuildSpec) Validate(ctx context.Context) *apis.FieldError {
	if equality.Semantic.DeepEqual(abs, &ApplicationBuildSpec{}) {
		return apis.ErrMissingField(apis.CurrentField)
	}

	errs := &apis.FieldError{}

	if abs.Image == "" {
		errs = errs.Also(apis.ErrMissingField("image"))
	}

	errs = errs.Also(abs.Source.Validate(ctx).ViaField("source"))

	return errs
}
