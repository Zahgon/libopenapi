// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Examples represents a low-level Swagger / OpenAPI 2 Example object.
// Allows sharing examples for operation responses
//   - https://swagger.io/specification/v2/#exampleObject
type Examples struct {
	Values *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
}

// FindExample attempts to locate an example value, using a key label.
func (e *Examples) FindExample(name string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// Build will extract all examples and will attempt to unmarshal content into a map or slice based on type.
func (e *Examples) Build(_ context.Context, _, root *yaml.Node, _ *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Hash will return a consistent Hash of the Examples object
func (e *Examples) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
