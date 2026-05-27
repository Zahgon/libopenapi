// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	v3low "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Paths represents a high-level OpenAPI 3+ Paths object, that is backed by a low-level one.
//
// Holds the relative paths to the individual endpoints and their operations. The path is appended to the URL from the
// Server Object in order to construct the full URL. The Paths MAY be empty, due to Access Control List (ACL)
// constraints.
//   - https://spec.openapis.org/oas/v3.1.0#paths-object
type Paths struct {
	PathItems  *orderedmap.Map[string, *PathItem]  `json:"-" yaml:"-"`
	Extensions *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low        *v3low.Paths
}

// NewPaths creates a new high-level instance of Paths from a low-level one.
func NewPaths(paths *v3low.Paths) *Paths { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Paths instance used to create the high-level one.
func (p *Paths) GoLow() *v3low.Paths {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Paths instance that was used to create the high-level one, with no type
	return nil
}

func (p *Paths) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Paths object as a byte slice.
	return *new(any)
}

func (p *Paths) Render() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Paths) RenderInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML will create a ready to render YAML representation of the Paths object.
func (p *Paths) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	// map keys correctly.
	return nil, nil
}

// default to a high value to weight new content to the bottom.

func (p *Paths) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	// map keys correctly.
	return nil, nil
}

// default to a high value to weight new content to the bottom.
