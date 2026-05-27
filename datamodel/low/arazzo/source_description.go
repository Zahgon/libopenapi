// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// SourceDescription represents a low-level Arazzo Source Description Object.
// https://spec.openapis.org/arazzo/v1.0.1#source-description-object
type SourceDescription struct {
	Name       low.NodeReference[string]
	URL        low.NodeReference[string]
	Type       low.NodeReference[string]
	Extensions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode    *yaml.Node
	RootNode   *yaml.Node
	index      *index.SpecIndex
	context    context.Context
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the SourceDescription object.
// For Arazzo low models this is typically nil, because Arazzo parsing does not build a SpecIndex.
// The index parameter is still required to satisfy the shared low.Buildable interface and generic extractors.
func (s *SourceDescription) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the SourceDescription object.
	return nil
}

func (s *SourceDescription) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (s *SourceDescription) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the SourceDescription object.
func (s *SourceDescription) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the SourceDescription object.
	return nil
}

func (s *SourceDescription) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract all properties of the SourceDescription object.
	return nil
}

func (s *SourceDescription) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all SourceDescription extensions and satisfies the low.HasExtensions interface.
func (s *SourceDescription) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent hash of the SourceDescription object.
}

func (s *SourceDescription) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
