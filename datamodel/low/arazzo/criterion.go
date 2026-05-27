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

// Criterion represents a low-level Arazzo Criterion Object.
// https://spec.openapis.org/arazzo/v1.0.1#criterion-object
type Criterion struct {
	Context    low.NodeReference[string]
	Condition  low.NodeReference[string]
	Type       low.NodeReference[*yaml.Node]
	Extensions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode    *yaml.Node
	RootNode   *yaml.Node
	index      *index.SpecIndex
	context    context.Context
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the Criterion object.
// For Arazzo low models this is typically nil, because Arazzo parsing does not build a SpecIndex.
// The index parameter is still required to satisfy the shared low.Buildable interface and generic extractors.
func (c *Criterion) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Criterion object.
	return nil
}

func (c *Criterion) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (c *Criterion) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the Criterion object.
func (c *Criterion) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Criterion object.
	return nil
}

func (c *Criterion) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract all properties of the Criterion object.
	// The Type field is a union: it can be a scalar string ("simple", "regex") or a mapping node
	// (CriterionExpressionType). We store it as a raw *yaml.Node for the high-level to interpret.
	return nil
}

func (c *Criterion) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Extract type as raw node since it's a union type

// GetExtensions returns all Criterion extensions and satisfies the low.HasExtensions interface.
func (c *Criterion) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent hash of the Criterion object.
}

func (c *Criterion) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
