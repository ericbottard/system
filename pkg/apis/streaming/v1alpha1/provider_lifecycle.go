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
	"github.com/projectriff/system/pkg/apis"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	ProviderConditionReady                                         = apis.ConditionReady
	ProviderConditionLiiklusDeploymentReady     apis.ConditionType = "LiiklusDeploymentReady"
	ProviderConditionLiiklusServiceReady        apis.ConditionType = "LiiklusServiceReady"
	ProviderConditionProvisionerDeploymentReady apis.ConditionType = "ProvisionerDeploymentReady"
	ProviderConditionProvisionerServiceReady    apis.ConditionType = "ProvisionerServiceReady"
)

var providerCondSet = apis.NewLivingConditionSet(
	ProviderConditionLiiklusDeploymentReady,
	ProviderConditionLiiklusServiceReady,
	ProviderConditionProvisionerDeploymentReady,
	ProviderConditionProvisionerServiceReady,
)

func (ds *ProviderStatus) GetObservedGeneration() int64 {
	return ds.ObservedGeneration
}

func (ds *ProviderStatus) IsReady() bool {
	return providerCondSet.Manage(ds).IsHappy()
}

func (*ProviderStatus) GetReadyConditionType() apis.ConditionType {
	return ProviderConditionReady
}

func (ds *ProviderStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return providerCondSet.Manage(ds).GetCondition(t)
}

func (ds *ProviderStatus) InitializeConditions() {
	providerCondSet.Manage(ds).InitializeConditions()
}

func (ds *ProviderStatus) MarkLiiklusDeploymentNotOwned() {
	providerCondSet.Manage(ds).MarkFalse(ProviderConditionLiiklusDeploymentReady, "NotOwned",
		"There is an existing Deployment %q that we do not own.", ds.LiiklusDeploymentName)
}

func (ds *ProviderStatus) PropagateLiiklusDeploymentStatus(cds *appsv1.DeploymentStatus) {
	var available, progressing *appsv1.DeploymentCondition
	for i := range cds.Conditions {
		switch cds.Conditions[i].Type {
		case appsv1.DeploymentAvailable:
			available = &cds.Conditions[i]
		case appsv1.DeploymentProgressing:
			progressing = &cds.Conditions[i]
		}
	}
	if available == nil || progressing == nil {
		return
	}
	if progressing.Status == corev1.ConditionTrue && available.Status == corev1.ConditionFalse {
		// DeploymentAvailable is False while progressing, avoid reporting ProviderConditionReady as False
		providerCondSet.Manage(ds).MarkUnknown(ProviderConditionLiiklusDeploymentReady, progressing.Reason, progressing.Message)
		return
	}
	switch {
	case available.Status == corev1.ConditionUnknown:
		providerCondSet.Manage(ds).MarkUnknown(ProviderConditionLiiklusDeploymentReady, available.Reason, available.Message)
	case available.Status == corev1.ConditionTrue:
		providerCondSet.Manage(ds).MarkTrue(ProviderConditionLiiklusDeploymentReady)
	case available.Status == corev1.ConditionFalse:
		providerCondSet.Manage(ds).MarkFalse(ProviderConditionLiiklusDeploymentReady, available.Reason, available.Message)
	}
}

func (ds *ProviderStatus) MarkLiiklusServiceNotOwned() {
	providerCondSet.Manage(ds).MarkFalse(ProviderConditionLiiklusServiceReady, "NotOwned",
		"There is an existing Service %q that we do not own.", ds.LiiklusServiceName)
}

func (ds *ProviderStatus) PropagateLiiklusServiceStatus(ss *corev1.ServiceStatus) {
	// services don't have meaningful status
	providerCondSet.Manage(ds).MarkTrue(ProviderConditionLiiklusServiceReady)
}

func (ds *ProviderStatus) MarkProvisionerDeploymentNotOwned() {
	providerCondSet.Manage(ds).MarkFalse(ProviderConditionProvisionerDeploymentReady, "NotOwned",
		"There is an existing Deployment %q that we do not own.", ds.ProvisionerDeploymentName)
}

func (ds *ProviderStatus) PropagateProvisionerDeploymentStatus(cds *appsv1.DeploymentStatus) {
	var available, progressing *appsv1.DeploymentCondition
	for i := range cds.Conditions {
		switch cds.Conditions[i].Type {
		case appsv1.DeploymentAvailable:
			available = &cds.Conditions[i]
		case appsv1.DeploymentProgressing:
			progressing = &cds.Conditions[i]
		}
	}
	if available == nil || progressing == nil {
		return
	}
	if progressing.Status == corev1.ConditionTrue && available.Status == corev1.ConditionFalse {
		// DeploymentAvailable is False while progressing, avoid reporting ProviderConditionReady as False
		providerCondSet.Manage(ds).MarkUnknown(ProviderConditionProvisionerDeploymentReady, progressing.Reason, progressing.Message)
		return
	}
	switch {
	case available.Status == corev1.ConditionUnknown:
		providerCondSet.Manage(ds).MarkUnknown(ProviderConditionProvisionerDeploymentReady, available.Reason, available.Message)
	case available.Status == corev1.ConditionTrue:
		providerCondSet.Manage(ds).MarkTrue(ProviderConditionProvisionerDeploymentReady)
	case available.Status == corev1.ConditionFalse:
		providerCondSet.Manage(ds).MarkFalse(ProviderConditionProvisionerDeploymentReady, available.Reason, available.Message)
	}
}

func (ds *ProviderStatus) MarkProvisionerServiceNotOwned() {
	providerCondSet.Manage(ds).MarkFalse(ProviderConditionProvisionerServiceReady, "NotOwned",
		"There is an existing Service %q that we do not own.", ds.ProvisionerServiceName)
}

func (ds *ProviderStatus) PropagateProvisionerServiceStatus(ss *corev1.ServiceStatus) {
	// services don't have meaningful status
	providerCondSet.Manage(ds).MarkTrue(ProviderConditionProvisionerServiceReady)
}
