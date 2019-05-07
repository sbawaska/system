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
	requestv1alpha1 "github.com/projectriff/system/pkg/apis/request/v1alpha1"
	"github.com/projectriff/system/pkg/reconciler/v1alpha1/requestprocessor/resources/names"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MakeApplication creates a Application from an RequestProcessor object.
func MakeApplication(rp *requestv1alpha1.RequestProcessor, i int) (*buildv1alpha1.Application, error) {
	if rp.Spec[i].Build.Application == nil {
		return nil, nil
	}

	Application := &buildv1alpha1.Application{
		ObjectMeta: metav1.ObjectMeta{
			Name:      names.Item(rp, i),
			Namespace: rp.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(rp),
			},
			Labels: makeLabels(rp),
		},
		Spec: *rp.Spec[i].Build.Application,
	}

	return Application, nil
}