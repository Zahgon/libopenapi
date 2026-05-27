// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	low "github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Tag represents a high-level Tag instance that is backed by a low-level one.
//
// Adds metadata to a single tag that is used by the Operation Object. It is not mandatory to have a Tag Object per
// tag defined in the Operation Object instances.
//   - v2: https://swagger.io/specification/v2/#tagObject
//   - v3: https://swagger.io/specification/#tag-object
//   - v3.2: https://spec.openapis.org/oas/v3.2.0#tag-object
type Tag struct {
	Name         string       `json:"name,omitempty" yaml:"name,omitempty"`
	Summary      string       `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string       `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDoc `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Parent       string       `json:"parent,omitempty" yaml:"parent,omitempty"`
	Kind         string       `json:"kind,omitempty" yaml:"kind,omitempty"`
	Extensions   *orderedmap.Map[string, *yaml.Node]
	low          *low.Tag
}

// NewTag creates a new high-level Tag instance that is backed by a low-level one.
func NewTag(tag *low.Tag) *Tag { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Tag instance used to create the high-level one.
func (t *Tag) GoLow() *low.Tag {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Tag instance that was used to create the high-level one, with no type
	return nil
}

func (t *Tag) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Info object as a byte slice.
	return *new(any)
}

func (t *Tag) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Render will return a YAML representation of the Info object as a byte slice.
		nil
}

func (t *Tag) RenderInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML will create a ready to render YAML representation of the Info object.
func (t *Tag) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tag) MarshalYAMLInline() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
