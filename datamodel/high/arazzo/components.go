// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Components represents a high-level Arazzo Components Object.
// https://spec.openapis.org/arazzo/v1.0.1#components-object
type Components struct {
	Inputs         *orderedmap.Map[string, *yaml.Node]     `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Parameters     *orderedmap.Map[string, *Parameter]     `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	SuccessActions *orderedmap.Map[string, *SuccessAction] `json:"successActions,omitempty" yaml:"successActions,omitempty"`
	FailureActions *orderedmap.Map[string, *FailureAction] `json:"failureActions,omitempty" yaml:"failureActions,omitempty"`
	Extensions     *orderedmap.Map[string, *yaml.Node]     `json:"-" yaml:"-"`
	low            *low.Components
}

// NewComponents creates a new high-level Components instance from a low-level one.
func NewComponents(comp *low.Components) *Components { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Components instance used to create the high-level one.
func (c *Components) GoLow() *low.Components {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level Components instance with no type.
	return nil
}

func (c *Components) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the Components object as a byte slice.
	return *new(any)
}

func (c *Components) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the Components object.
		nil
}

func (c *Components) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
