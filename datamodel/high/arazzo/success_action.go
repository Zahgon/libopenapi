// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// SuccessAction represents a high-level Arazzo Success Action Object.
// A success action can be a full definition or a Reusable Object with a $components reference.
// https://spec.openapis.org/arazzo/v1.0.1#success-action-object
type SuccessAction struct {
	Name       string                              `json:"name,omitempty" yaml:"name,omitempty"`
	Type       string                              `json:"type,omitempty" yaml:"type,omitempty"`
	WorkflowId string                              `json:"workflowId,omitempty" yaml:"workflowId,omitempty"`
	StepId     string                              `json:"stepId,omitempty" yaml:"stepId,omitempty"`
	Criteria   []*Criterion                        `json:"criteria,omitempty" yaml:"criteria,omitempty"`
	Reference  string                              `json:"reference,omitempty" yaml:"reference,omitempty"`
	Extensions *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low        *low.SuccessAction
}

// IsReusable returns true if this success action is a Reusable Object (has a reference field).
func (s *SuccessAction) IsReusable() bool { _ = "STUB: not implemented"; return false }

// NewSuccessAction creates a new high-level SuccessAction instance from a low-level one.
func NewSuccessAction(sa *low.SuccessAction) *SuccessAction { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level SuccessAction instance used to create the high-level one.
func (s *SuccessAction) GoLow() *low.SuccessAction {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level SuccessAction instance with no type.
	return nil
}

func (s *SuccessAction) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the SuccessAction object as a byte slice.
	return *new(any)
}

func (s *SuccessAction) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the SuccessAction object.
		nil
}

func (s *SuccessAction) MarshalYAML() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
