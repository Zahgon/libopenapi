// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Step represents a high-level Arazzo Step Object.
// https://spec.openapis.org/arazzo/v1.0.1#step-object
type Step struct {
	StepId          string                              `json:"stepId,omitempty" yaml:"stepId,omitempty"`
	Description     string                              `json:"description,omitempty" yaml:"description,omitempty"`
	OperationId     string                              `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	OperationPath   string                              `json:"operationPath,omitempty" yaml:"operationPath,omitempty"`
	WorkflowId      string                              `json:"workflowId,omitempty" yaml:"workflowId,omitempty"`
	Parameters      []*Parameter                        `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody     *RequestBody                        `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	SuccessCriteria []*Criterion                        `json:"successCriteria,omitempty" yaml:"successCriteria,omitempty"`
	OnSuccess       []*SuccessAction                    `json:"onSuccess,omitempty" yaml:"onSuccess,omitempty"`
	OnFailure       []*FailureAction                    `json:"onFailure,omitempty" yaml:"onFailure,omitempty"`
	Outputs         *orderedmap.Map[string, string]     `json:"outputs,omitempty" yaml:"outputs,omitempty"`
	Extensions      *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low             *low.Step
}

// NewStep creates a new high-level Step instance from a low-level one.
func NewStep(step *low.Step) *Step { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Step instance used to create the high-level one.
func (s *Step) GoLow() *low.Step {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level Step instance with no type.
	return nil
}

func (s *Step) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the Step object as a byte slice.
	return *new(any)
}

func (s *Step) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the Step object.
		nil
}

func (s *Step) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
