// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// SourceDescription represents a high-level Arazzo Source Description Object.
// https://spec.openapis.org/arazzo/v1.0.1#source-description-object
type SourceDescription struct {
	Name       string                              `json:"name,omitempty" yaml:"name,omitempty"`
	URL        string                              `json:"url,omitempty" yaml:"url,omitempty"`
	Type       string                              `json:"type,omitempty" yaml:"type,omitempty"`
	Extensions *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low        *low.SourceDescription
}

// NewSourceDescription creates a new high-level SourceDescription instance from a low-level one.
func NewSourceDescription(sd *low.SourceDescription) *SourceDescription {
	_ = "STUB: not implemented"
	return nil
}

// GoLow returns the low-level SourceDescription instance used to create the high-level one.
func (s *SourceDescription) GoLow() *low.SourceDescription {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level SourceDescription instance with no type.
	return nil
}

func (s *SourceDescription) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the SourceDescription object as a byte slice.
	return *new(any)
}

func (s *SourceDescription) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the SourceDescription object.
		nil
}

func (s *SourceDescription) MarshalYAML() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
