/*
Copyright 2018 The Knative Authors

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

package resources

import (
	"github.com/knative/pkg/kmeta"
	buildv1alpha1 "github.com/projectriff/system/pkg/apis/build/v1alpha1"
	"github.com/projectriff/system/pkg/reconciler/v1alpha1/functionbuild/resources/names"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MakeBuildCache creates a PersistentVolumeClaim from an FunctionBuild object.
func MakeBuildCache(fb *buildv1alpha1.FunctionBuild) (*corev1.PersistentVolumeClaim, error) {
	if fb.Spec.CacheSize == nil {
		// no cache was requested
		return nil, nil
	}

	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      names.BuildCache(fb),
			Namespace: fb.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(fb),
			},
			Labels: makeLabels(fb),
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: *fb.Spec.CacheSize,
				},
			},
		},
	}

	return pvc, nil
}
