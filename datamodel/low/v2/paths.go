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

// Paths represents a low-level Swagger / OpenAPI Paths object.
type Paths struct {
	PathItems  *orderedmap.Map[low.KeyReference[string], low.ValueReference[*PathItem]]
	Extensions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
}

// GetExtensions returns all Paths extensions and satisfies the low.HasExtensions interface.
func (p *Paths) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// FindPath attempts to locate a PathItem instance, given a path key.
}

func (p *Paths) FindPath(path string) (result *low.ValueReference[*PathItem]) {
	_ = "STUB: not implemented"
	return nil
}

// FindPathAndKey attempts to locate a PathItem instance, given a path key.
func (p *Paths) FindPathAndKey(path string) (key *low.KeyReference[string], value *low.ValueReference[*PathItem]) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindExtension will attempt to locate an extension value given a name.
func (p *Paths) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// Build will extract extensions and paths from node.
func (p *Paths) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Translate YAML nodes to pathsMap using `TranslatePipeline`.

// input and output goroutines.

// TranslatePipeline input.

// TranslatePipeline output.

// Hash will return a consistent Hash of the Paths object
func (p *Paths) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
