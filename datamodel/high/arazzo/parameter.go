// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Parameter represents a high-level Arazzo Parameter Object.
// A parameter can be a full parameter definition or a Reusable Object with a $components reference.
// https://spec.openapis.org/arazzo/v1.0.1#parameter-object
type Parameter struct {
	Name       string                              `json:"name,omitempty" yaml:"name,omitempty"`
	In         string                              `json:"in,omitempty" yaml:"in,omitempty"`
	Value      *yaml.Node                          `json:"value,omitempty" yaml:"value,omitempty"`
	Reference  string                              `json:"reference,omitempty" yaml:"reference,omitempty"`
	Extensions *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low        *low.Parameter
}

// IsReusable returns true if this parameter is a Reusable Object (has a reference field).
func (p *Parameter) IsReusable() bool { _ = "STUB: not implemented"; return false }

// NewParameter creates a new high-level Parameter instance from a low-level one.
func NewParameter(param *low.Parameter) *Parameter { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Parameter instance used to create the high-level one.
func (p *Parameter) GoLow() *low.Parameter {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level Parameter instance with no type.
	return nil
}

func (p *Parameter) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the Parameter object as a byte slice.
	return *new(any)
}

func (p *Parameter) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the Parameter object.
		nil
}

func (p *Parameter) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
