// Copyright 2022-2025 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package overlay

import (
	low "github.com/pb33f/libopenapi/datamodel/low/overlay"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Overlay represents a high-level OpenAPI Overlay document.
// https://spec.openapis.org/overlay/v1.0.0
type Overlay struct {
	Overlay    string                              `json:"overlay,omitempty" yaml:"overlay,omitempty"`
	Info       *Info                               `json:"info,omitempty" yaml:"info,omitempty"`
	Extends    string                              `json:"extends,omitempty" yaml:"extends,omitempty"`
	Actions    []*Action                           `json:"actions,omitempty" yaml:"actions,omitempty"`
	Extensions *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low        *low.Overlay
}

// NewOverlay creates a new high-level Overlay instance from a low-level one.
func NewOverlay(overlay *low.Overlay) *Overlay { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Overlay instance used to create the high-level one.
func (o *Overlay) GoLow() *low.Overlay {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level Overlay instance with no type.
	return nil
}

func (o *Overlay) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the Overlay object as a byte slice.
	return *new(any)
}

func (o *Overlay) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the Overlay object.
		nil
}

func (o *Overlay) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
