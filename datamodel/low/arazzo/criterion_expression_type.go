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

// CriterionExpressionType represents a low-level Arazzo Criterion Expression Type Object.
// https://spec.openapis.org/arazzo/v1.0.1#criterion-expression-type-object
type CriterionExpressionType struct {
	Type       low.NodeReference[string]
	Version    low.NodeReference[string]
	Extensions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode    *yaml.Node
	RootNode   *yaml.Node
	index      *index.SpecIndex
	context    context.Context
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the CriterionExpressionType object.
// For Arazzo low models this is typically nil, because Arazzo parsing does not build a SpecIndex.
// The index parameter is still required to satisfy the shared low.Buildable interface and generic extractors.
func (c *CriterionExpressionType) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the CriterionExpressionType object.
	return nil
}

func (c *CriterionExpressionType) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (c *CriterionExpressionType) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the CriterionExpressionType object.
func (c *CriterionExpressionType) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the CriterionExpressionType object.
	return nil
}

func (c *CriterionExpressionType) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract all properties of the CriterionExpressionType object.
	return nil
}

func (c *CriterionExpressionType) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all CriterionExpressionType extensions and satisfies the low.HasExtensions interface.
func (c *CriterionExpressionType) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent hash of the CriterionExpressionType object.
}

func (c *CriterionExpressionType) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
