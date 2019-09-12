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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func TestRun(t *testing.T) {
	fakeVpaLabels := map[string]string{
		"kind": "Deployment",
	}
	// this always has to be present when the run func is called now
	kubeClientVPA := kube.GetMockVPAClient()

	fakeVpaListOptions := metav1.ListOptions{
		LabelSelector: labels.Set(fakeVpaLabels).String(),
	}
	// do I actually need this?
	_, errOk := kubeClientVPA.Client.AutoscalingV1beta2().VerticalPodAutoscalers("").List(fakeVpaListOptions)
	assert.NoError(t, errOk)

	var summary Summary

	got, err := Run(kubeClientVPA, fakeVpaLabels, "true")
	assert.NoError(t, err)

	assert.EqualValues(t, got, summary)
}
