// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"context"
	"sync"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Example represents a low-level Example object as defined by OpenAPI 3+
//
//	v3 - https://spec.openapis.org/oas/v3.1.0#example-object
type Example struct {
	Summary         low.NodeReference[string]
	Description     low.NodeReference[string]
	Value           low.NodeReference[*yaml.Node]
	ExternalValue   low.NodeReference[string]
	DataValue       low.NodeReference[*yaml.Node] // OpenAPI 3.2+ dataValue field (mutually exclusive with value/externalValue)
	SerializedValue low.NodeReference[string]     // OpenAPI 3.2+ serializedValue field (mutually exclusive with value/externalValue)
	Extensions      *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode         *yaml.Node
	RootNode        *yaml.Node
	index           *index.SpecIndex
	context         context.Context
	nodeStore       sync.Map
	reference       low.Reference
	*low.Reference
	low.NodeMap
}

// FindExtension returns a ValueReference containing the extension value, if found.
func (ex *Example) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode will return the root yaml node of the Example object
func (ex *Example) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode will return the key yaml node of the Example object
	return nil
}

func (ex *Example) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Hash will return a consistent hash of the Example object
	return nil
}

func (ex *Example) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Build extracts extensions and example value
func (ex *Example) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// OpenAPI 3.2+ dataValue field

// OpenAPI 3.2+ serializedValue field

// GetExtensions will return Example extensions to satisfy the HasExtensions interface.
func (ex *Example) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// GetIndex will return the index.SpecIndex instance attached to the Example object
}

func (ex *Example) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext will return the context.Context instance used when building the Example object
	return nil
}

func (ex *Example) GetContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
