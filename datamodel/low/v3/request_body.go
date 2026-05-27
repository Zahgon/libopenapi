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

// RequestBody represents a low-level OpenAPI 3+ RequestBody object.
//   - https://spec.openapis.org/oas/v3.1.0#request-body-object
type RequestBody struct {
	Description low.NodeReference[string]
	Content     low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*MediaType]]]
	Required    low.NodeReference[bool]
	Extensions  *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode     *yaml.Node
	RootNode    *yaml.Node
	index       *index.SpecIndex
	context     context.Context
	nodeStore   sync.Map
	reference   low.Reference
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the RequestBody object.
func (rb *RequestBody) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the RequestBody object.
	return nil
}

func (rb *RequestBody) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetRootNode returns the root yaml node of the RequestBody object.
	return *new(context.Context)
}

func (rb *RequestBody) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the RequestBody object.
	return nil
}

func (rb *RequestBody) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// FindExtension attempts to locate an extension using the provided name.
	return nil
}

func (rb *RequestBody) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all RequestBody extensions and satisfies the low.HasExtensions interface.
func (rb *RequestBody) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// FindContent attempts to find content/MediaType defined using a specified name.
}

func (rb *RequestBody) FindContent(cType string) *low.ValueReference[*MediaType] {
	_ = "STUB: not implemented"
	return nil
}

// Build will extract extensions and MediaType objects from the node.
func (rb *RequestBody) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// handle content, if set.

// Hash will return a consistent Hash of the RequestBody object
func (rb *RequestBody) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
