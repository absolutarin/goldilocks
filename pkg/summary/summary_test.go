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
	"github.com/fairwindsops/goldilocks/pkg/utils"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func TestRun(t *testing.T) {
	// this always has to be present when the run func is called now
	kubeClientVPA := kube.GetMockVPAClient()

	fakeVpaListOptions := metav1.ListOptions{
		LabelSelector: labels.Set(utils.VpaLabels).String(),
	}

	_, errOk := kubeClientVPA.Client.AutoscalingV1beta2().VerticalPodAutoscalers("").List(fakeVpaListOptions)
	assert.NoError(t, errOk)

	var summary Summary

	got, err := Run(kubeClientVPA, utils.VpaLabels, "true")
	assert.NoError(t, err)

	assert.EqualValues(t, got, summary)
}

// called on line 41 cmd/summary.go
// called on line 165 in dashboard.go
