// Copyright 2022-2025 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package overlay

import (
	low "github.com/pb33f/libopenapi/datamodel/low/overlay"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Action represents a high-level Overlay Action Object.
// https://spec.openapis.org/overlay/v1.1.0#action-object
type Action struct {
	Target      string                              `json:"target,omitempty" yaml:"target,omitempty"`
	Description string                              `json:"description,omitempty" yaml:"description,omitempty"`
	Update      *yaml.Node                          `json:"update,omitempty" yaml:"update,omitempty"`
	Remove      bool                                `json:"remove,omitempty" yaml:"remove,omitempty"`
	Copy        string                              `json:"copy,omitempty" yaml:"copy,omitempty"`
	Extensions  *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low         *low.Action
}

// NewAction creates a new high-level Action instance from a low-level one.
func NewAction(action *low.Action) *Action { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Action instance used to create the high-level one.
func (a *Action) GoLow() *low.Action {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level Action instance with no type.
	return nil
}

func (a *Action) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the Action object as a byte slice.
	return *new(any)
}

func (a *Action) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the Action object.
		nil
}

func (a *Action) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
