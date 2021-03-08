// Copyright © 2020 The Knative Authors
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

package command

import (
	"bytes"
	"context"
	camelkapis "github.com/apache/camel-k/pkg/apis/camel/v1alpha1"
	camelkv1alpha1 "github.com/apache/camel-k/pkg/client/camel/clientset/versioned/typed/camel/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/client/pkg/util"
	"knative.dev/kn-plugin-source-kamelet/internal/client"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestListSetup(t *testing.T) {
	p := KameletPluginParams{
		Context: context.TODO(),
	}

	listCmd := NewListCommand(&p)
	assert.Equal(t, listCmd.Use, "list")
	assert.Equal(t, listCmd.Short, "List available Kamelet sources")
	assert.Assert(t, listCmd.RunE != nil)
}

func TestListOutput(t *testing.T) {
	mockClient := client.NewMockKameletClient(t)
	recorder := mockClient.Recorder()

	kamelet1 := createKamelet("k1")
	kamelet2 := createKamelet("k2")
	kamelet3 := createKamelet("k3")
	kameletList := &camelkapis.KameletList{Items: []camelkapis.Kamelet{*kamelet1, *kamelet2, *kamelet3}}
	recorder.List(kameletList, nil)

	output, err := runListCmd(mockClient)
	assert.NilError(t, err)

	outputLines := strings.Split(output, "\n")
	assert.Check(t, util.ContainsAll(outputLines[0], "NAME", "PHASE", "AGE", "CONDITIONS", "READY", "REASON"))
	assert.Check(t, util.ContainsAll(outputLines[1], "k1", "Ready", "True"))
	assert.Check(t, util.ContainsAll(outputLines[2], "k2", "Ready", "True"))
	assert.Check(t, util.ContainsAll(outputLines[3], "k3", "Ready", "True"))

	recorder.Validate()
}

func runListCmd(c *client.MockKameletClient) (string, error) {
	p := KameletPluginParams{
		Context: context.TODO(),
		NewKameletClient: func() (camelkv1alpha1.CamelV1alpha1Interface, error) {
			return c, nil
		},
	}

	listCmd := NewListCommand(&p)

	output := new(bytes.Buffer)
	listCmd.SetOut(output)
	err := listCmd.Execute()
	return output.String(), err
}

func createKamelet(kameletName string) *camelkapis.Kamelet {
	return &camelkapis.Kamelet{
		TypeMeta: v1.TypeMeta{
			Kind: camelkapis.KameletKind,
		},
		ObjectMeta: v1.ObjectMeta{
			Namespace: "default",
			Name:      kameletName,
		},
		Spec: camelkapis.KameletSpec{},
		Status: camelkapis.KameletStatus{
			Phase: camelkapis.KameletPhaseReady,
			Conditions: []camelkapis.KameletCondition{
				camelkapis.KameletCondition{
					Type:   "Ready",
					Status: "True",
				},
			},
		},
	}
}
