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

// Components represents a low-level Arazzo Components Object.
// https://spec.openapis.org/arazzo/v1.0.1#components-object
type Components struct {
	Inputs         low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]]
	Parameters     low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Parameter]]]
	SuccessActions low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*SuccessAction]]]
	FailureActions low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*FailureAction]]]
	Extensions     *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode        *yaml.Node
	RootNode       *yaml.Node
	index          *index.SpecIndex
	context        context.Context
	*low.Reference
	low.NodeMap
}

var extractComponentsParametersMap = extractObjectMap[Parameter]
var extractComponentsSuccessActionsMap = extractObjectMap[SuccessAction]

// GetIndex returns the index.SpecIndex instance attached to the Components object.
// For Arazzo low models this is typically nil, because Arazzo parsing does not build a SpecIndex.
// The index parameter is still required to satisfy the shared low.Buildable interface and generic extractors.
func (c *Components) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Components object.
	return nil
}

func (c *Components) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (c *Components) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the Components object.
func (c *Components) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Components object.
	return nil
}

func (c *Components) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract all properties of the Components object.
	return nil
}

func (c *Components) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Extract inputs as raw node map (JSON Schema objects keyed by name)

// Extract parameters map

// Extract successActions map

// Extract failureActions map

// GetExtensions returns all Components extensions and satisfies the low.HasExtensions interface.
func (c *Components) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent hash of the Components object.
}

func (c *Components) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
