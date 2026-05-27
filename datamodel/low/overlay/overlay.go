// Copyright 2022-2025 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package overlay

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Overlay represents a low-level OpenAPI Overlay document.
// https://spec.openapis.org/overlay/v1.0.0
type Overlay struct {
	Overlay    low.NodeReference[string]
	Info       low.NodeReference[*Info]
	Extends    low.NodeReference[string]
	Actions    low.NodeReference[[]low.ValueReference[*Action]]
	Extensions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode    *yaml.Node
	RootNode   *yaml.Node
	index      *index.SpecIndex
	context    context.Context
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the Overlay object
func (o *Overlay) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Overlay object
	return nil
}

func (o *Overlay) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (o *Overlay) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the Overlay object
func (o *Overlay) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Overlay object
	return nil
}

func (o *Overlay) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract all properties of the Overlay document.
	return nil
}

func (o *Overlay) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Extract info object

// Extract actions array

func (o *Overlay) extractActions(ctx context.Context, root *yaml.Node, idx *index.SpecIndex) low.NodeReference[[]low.ValueReference[*Action]] {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all Overlay extensions and satisfies the low.HasExtensions interface.
func (o *Overlay) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent Hash of the Overlay object
}

func (o *Overlay) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
