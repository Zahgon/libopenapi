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

// Scopes is a low-level representation of a Swagger / OpenAPI 2 OAuth2 Scopes object.
//
// Scopes lists the available scopes for an OAuth2 security scheme.
//   - https://swagger.io/specification/v2/#scopesObject
type Scopes struct {
	Values     *orderedmap.Map[low.KeyReference[string], low.ValueReference[string]]
	Extensions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
}

// GetExtensions returns all Scopes extensions and satisfies the low.HasExtensions interface.
func (s *Scopes) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// FindScope will attempt to locate a scope string using a key.
}

func (s *Scopes) FindScope(scope string) *low.ValueReference[string] {
	_ = "STUB: not implemented"
	return nil
}

// Build will extract scope values and extensions from node.
func (s *Scopes) Build(_ context.Context, _, root *yaml.Node, _ *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Hash will return a consistent Hash of the Scopes object
func (s *Scopes) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
