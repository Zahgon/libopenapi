// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// FailureAction represents a high-level Arazzo Failure Action Object.
// A failure action can be a full definition or a Reusable Object with a $components reference.
// https://spec.openapis.org/arazzo/v1.0.1#failure-action-object
type FailureAction struct {
	Name       string                              `json:"name,omitempty" yaml:"name,omitempty"`
	Type       string                              `json:"type,omitempty" yaml:"type,omitempty"`
	WorkflowId string                              `json:"workflowId,omitempty" yaml:"workflowId,omitempty"`
	StepId     string                              `json:"stepId,omitempty" yaml:"stepId,omitempty"`
	RetryAfter *float64                            `json:"retryAfter,omitempty" yaml:"retryAfter,omitempty"`
	RetryLimit *int64                              `json:"retryLimit,omitempty" yaml:"retryLimit,omitempty"`
	Criteria   []*Criterion                        `json:"criteria,omitempty" yaml:"criteria,omitempty"`
	Reference  string                              `json:"reference,omitempty" yaml:"reference,omitempty"`
	Extensions *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low        *low.FailureAction
}

// IsReusable returns true if this failure action is a Reusable Object (has a reference field).
func (f *FailureAction) IsReusable() bool { _ = "STUB: not implemented"; return false }

// NewFailureAction creates a new high-level FailureAction instance from a low-level one.
func NewFailureAction(fa *low.FailureAction) *FailureAction { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level FailureAction instance used to create the high-level one.
func (f *FailureAction) GoLow() *low.FailureAction {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level FailureAction instance with no type.
	return nil
}

func (f *FailureAction) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the FailureAction object as a byte slice.
	return *new(any)
}

func (f *FailureAction) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the FailureAction object.
		nil
}

func (f *FailureAction) MarshalYAML() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
