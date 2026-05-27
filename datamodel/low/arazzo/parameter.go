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

// Parameter represents a low-level Arazzo Parameter Object.
// A parameter can be a full parameter definition or a Reusable Object with a $components reference.
// https://spec.openapis.org/arazzo/v1.0.1#parameter-object
type Parameter struct {
	Name         low.NodeReference[string]
	In           low.NodeReference[string]
	Value        low.NodeReference[*yaml.Node]
	ComponentRef low.NodeReference[string]
	Extensions   *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode      *yaml.Node
	RootNode     *yaml.Node
	index        *index.SpecIndex
	context      context.Context
	*low.Reference
	low.NodeMap
}

// IsReusable returns true if this parameter is a Reusable Object (has a reference field).
func (p *Parameter) IsReusable() bool { _ = "STUB: not implemented"; return false }

// GetIndex returns the index.SpecIndex instance attached to the Parameter object.
// For Arazzo low models this is typically nil, because Arazzo parsing does not build a SpecIndex.
// The index parameter is still required to satisfy the shared low.Buildable interface and generic extractors.
func (p *Parameter) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Parameter object.
	return nil
}

func (p *Parameter) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (p *Parameter) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the Parameter object.
func (p *Parameter) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Parameter object.
	return nil
}

func (p *Parameter) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract all properties of the Parameter object.
	return nil
}

func (p *Parameter) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all Parameter extensions and satisfies the low.HasExtensions interface.
func (p *Parameter) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent hash of the Parameter object.
}

func (p *Parameter) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
