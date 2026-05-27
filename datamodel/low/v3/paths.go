// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"context"
	"sync"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
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
	PathItems  *orderedmap.Map[low.KeyReference[string], low.ValueReference[*PathItem]]
	Extensions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode    *yaml.Node
	RootNode   *yaml.Node
	index      *index.SpecIndex
	context    context.Context
	nodeStore  sync.Map
	reference  low.Reference
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the Paths object.
func (p *Paths) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Paths object.
	return nil
}

func (p *Paths) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetRootNode returns the root yaml node of the Paths object.
	return *new(context.Context)
}

func (p *Paths) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Paths object.
	return nil
}

func (p *Paths) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// FindPath will attempt to locate a PathItem using the provided path string.
	return nil
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

// FindExtension will attempt to locate an extension using the specified string.
func (p *Paths) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all Paths extensions and satisfies the low.HasExtensions interface.
func (p *Paths) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Build will extract extensions and all PathItems. This happens asynchronously for speed.
}

func (p *Paths) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// add path as node to path item, not this path object.

// Hash will return a consistent Hash of the Paths object
func (p *Paths) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

func extractPathItemsMap(ctx context.Context, root *yaml.Node, idx *index.SpecIndex) (*orderedmap.Map[low.KeyReference[string], low.ValueReference[*PathItem]], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
