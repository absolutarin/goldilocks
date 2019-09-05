// Copyright 2019 FairwindsOps Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package summary

import (
	"testing"

	"github.com/fairwindsops/goldilocks/pkg/kube"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func TestRun(t *testing.T) {
	// I think that I need to mock out the vpaLabels and excludeContainers
	// that the Run function takes
	fakeVpaLabels := map[string]string{
		"kind": "Deployment",
	}

	// do i need to recreate everything that happens in the function?
	// ie:

	// 	kubeClientVPA := kube.GetVPAInstance()
	// vpaListOptions := metav1.ListOptions{
	// 	LabelSelector: labels.Set(vpaLabels).String(),
	// }

	// this needs to be added to the kube pkg as a fake method
	// this is from test_helpers
	kubeClientVPA := kube.GetMockVPAClient()

	fakeVpaListOptions := metav1.ListOptions{
		LabelSelector: labels.Set(fakeVpaLabels).String(),
	}
	// Not even sure I need these things, but it kept erroring out because
	// it couldn't find corev1 pkg, and since it adds them automatically, this seemed
	// like a thing to do..it makes the test crazy long- not cute
	// I might need them because I think I need Summary and Summary needs Deployment Summary etc...
	type fakeContainerSummary struct {
		LowerBound     corev1.ResourceList `json:"fake"`
		UpperBound     corev1.ResourceList `json:"fake"`
		Target         corev1.ResourceList `json:"fake"`
		UncappedTarget corev1.ResourceList `json:"fake"`
		Limits         corev1.ResourceList `json:"fake"`
		Requests       corev1.ResourceList `json:"fake"`
		ContainerName  string              `json:"fake"`
	}

	type DeploymentSummary struct {
		Containers     []fakeContainerSummary `json:"fake"`
		DeploymentName string                 `json:"fake"`
		Namespace      string                 `json:"fake"`
	}

	// Summary struct is for storing a summary of recommendation data.
	type Summary struct {
		Deployments []deploymentSummary `json:"fake"`
		Namespaces  []string            `json:"fake"`
	}

	// making sure this doesn't error out?
	_, errOk := kubeClientVPA.Client.AutoscalingV1beta2().VerticalPodAutoscalers("").List(fakeVpaListOptions)
	assert.NoError(t, errOk)
	// what even am I testing for? just that the return is in fact
	// type summary?
	// i think this is making an instance of Summary
	var summary Summary

	got, err := Run(fakeVpaLabels, "true")
	assert.NoError(t, err)
	// I don't even think this is really testing anything.
	// BUT I can't figure out how to debug with make test
	assert.EqualValues(t, got, summary)
	// how do i account for errors? how do they even work?
}

// how can you interact with those imported packages(?)
