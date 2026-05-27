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

// Action represents a low-level Overlay Action Object.
// https://spec.openapis.org/overlay/v1.1.0#action-object
type Action struct {
	Target      low.NodeReference[string]
	Description low.NodeReference[string]
	Update      low.NodeReference[*yaml.Node]
	Remove      low.NodeReference[bool]
	Copy        low.NodeReference[string]
	Extensions  *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode     *yaml.Node
	RootNode    *yaml.Node
	index       *index.SpecIndex
	context     context.Context
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the Action object
func (a *Action) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Action object
	return nil
}

func (a *Action) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (a *Action) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the Action object
func (a *Action) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Action object
	return nil
}

func (a *Action) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract extensions for the Action object.
	return nil
}

func (a *Action) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Extract the update node directly if present

// GetExtensions returns all Action extensions and satisfies the low.HasExtensions interface.
func (a *Action) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent Hash of the Action object
}

func (a *Action) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
