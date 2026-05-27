// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Workflow represents a high-level Arazzo Workflow Object.
// https://spec.openapis.org/arazzo/v1.0.1#workflow-object
type Workflow struct {
	WorkflowId     string                              `json:"workflowId,omitempty" yaml:"workflowId,omitempty"`
	Summary        string                              `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description    string                              `json:"description,omitempty" yaml:"description,omitempty"`
	Inputs         *yaml.Node                          `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	DependsOn      []string                            `json:"dependsOn,omitempty" yaml:"dependsOn,omitempty"`
	Steps          []*Step                             `json:"steps,omitempty" yaml:"steps,omitempty"`
	SuccessActions []*SuccessAction                    `json:"successActions,omitempty" yaml:"successActions,omitempty"`
	FailureActions []*FailureAction                    `json:"failureActions,omitempty" yaml:"failureActions,omitempty"`
	Outputs        *orderedmap.Map[string, string]     `json:"outputs,omitempty" yaml:"outputs,omitempty"`
	Parameters     []*Parameter                        `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Extensions     *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low            *low.Workflow
}

// NewWorkflow creates a new high-level Workflow instance from a low-level one.
func NewWorkflow(wf *low.Workflow) *Workflow { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Workflow instance used to create the high-level one.
func (w *Workflow) GoLow() *low.Workflow {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level Workflow instance with no type.
	return nil
}

func (w *Workflow) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the Workflow object as a byte slice.
	return *new(any)
}

func (w *Workflow) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the Workflow object.
		nil
}

func (w *Workflow) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
