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

package names

import (
	"testing"

	corev1alpha1 "github.com/projectriff/system/pkg/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNames(t *testing.T) {
	tests := []struct {
		name     string
		deployer *corev1alpha1.Deployer
		f        func(*corev1alpha1.Deployer) string
		want     string
	}{{
		name: "Service",
		deployer: &corev1alpha1.Deployer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "default",
			},
		},
		f:    Service,
		want: "foo-deployer",
	}, {
		name: "Deployment",
		deployer: &corev1alpha1.Deployer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "default",
			},
		},
		f:    Deployment,
		want: "foo-deployer",
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.f(test.deployer)
			if got != test.want {
				t.Errorf("%s() = %v, wanted %v", test.name, got, test.want)
			}
		})
	}
}