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

// RequestBody represents a low-level Arazzo Request Body Object.
// https://spec.openapis.org/arazzo/v1.0.1#request-body-object
type RequestBody struct {
	ContentType  low.NodeReference[string]
	Payload      low.NodeReference[*yaml.Node]
	Replacements low.NodeReference[[]low.ValueReference[*PayloadReplacement]]
	Extensions   *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode      *yaml.Node
	RootNode     *yaml.Node
	index        *index.SpecIndex
	context      context.Context
	*low.Reference
	low.NodeMap
}

var extractRequestBodyReplacements = extractArray[PayloadReplacement]

// GetIndex returns the index.SpecIndex instance attached to the RequestBody object.
// For Arazzo low models this is typically nil, because Arazzo parsing does not build a SpecIndex.
// The index parameter is still required to satisfy the shared low.Buildable interface and generic extractors.
func (r *RequestBody) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the RequestBody object.
	return nil
}

func (r *RequestBody) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (r *RequestBody) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the RequestBody object.
func (r *RequestBody) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the RequestBody object.
	return nil
}

func (r *RequestBody) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract all properties of the RequestBody object.
	return nil
}

func (r *RequestBody) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all RequestBody extensions and satisfies the low.HasExtensions interface.
func (r *RequestBody) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent hash of the RequestBody object.
}

func (r *RequestBody) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
