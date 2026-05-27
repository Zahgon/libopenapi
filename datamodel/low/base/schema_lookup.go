// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// GetIndex will return the index.SpecIndex instance attached to the Schema object
func (s *Schema) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext will return the context.Context instance used when building the Schema object
	return nil
}

func (s *Schema) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindProperty will return a ValueReference pointer containing a SchemaProxy pointer
	// from a property key name. if found
	return *new(context.Context)
}

func (s *Schema) FindProperty(name string) *low.ValueReference[*SchemaProxy] {
	_ = "STUB: not implemented"
	return nil
}

// FindDependentSchema will return a ValueReference pointer containing a SchemaProxy pointer
// from a dependent schema key name. if found (3.1+ only)
func (s *Schema) FindDependentSchema(name string) *low.ValueReference[*SchemaProxy] {
	_ = "STUB: not implemented"
	return nil
}

// FindPatternProperty will return a ValueReference pointer containing a SchemaProxy pointer
// from a pattern property key name. if found (3.1+ only)
func (s *Schema) FindPatternProperty(name string) *low.ValueReference[*SchemaProxy] {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all extensions for Schema
func (s *Schema) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// GetRootNode will return the root yaml node of the Schema object
}

func (s *Schema) GetRootNode() *yaml.Node { _ = "STUB: not implemented"; return nil }
