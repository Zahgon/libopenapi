// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Info represents a high-level Arazzo Info Object.
// https://spec.openapis.org/arazzo/v1.0.1#info-object
type Info struct {
	Title       string                              `json:"title,omitempty" yaml:"title,omitempty"`
	Summary     string                              `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string                              `json:"description,omitempty" yaml:"description,omitempty"`
	Version     string                              `json:"version,omitempty" yaml:"version,omitempty"`
	Extensions  *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low         *low.Info
}

// NewInfo creates a new high-level Info instance from a low-level one.
func NewInfo(info *low.Info) *Info { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Info instance used to create the high-level one.
func (i *Info) GoLow() *low.Info {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level Info instance with no type.
	return nil
}

func (i *Info) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the Info object as a byte slice.
	return *new(any)
}

func (i *Info) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the Info object.
		nil
}

func (i *Info) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
