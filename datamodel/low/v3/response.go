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

// Response represents a high-level OpenAPI 3+ Response object that is backed by a low-level one.
//
// Describes a single response from an API Operation, including design-time, static links to
// operations based on the response.
//   - https://spec.openapis.org/oas/v3.1.0#response-object
type Response struct {
	Summary     low.NodeReference[string]
	Description low.NodeReference[string]
	Headers     low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Header]]]
	Content     low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*MediaType]]]
	Extensions  *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	Links       low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Link]]]
	KeyNode     *yaml.Node
	RootNode    *yaml.Node
	index       *index.SpecIndex
	context     context.Context
	nodeStore   sync.Map
	reference   low.Reference
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the Response object.
func (r *Response) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Response object.
	return nil
}

func (r *Response) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetRootNode returns the root yaml node of the Response object.
	return *new(context.Context)
}

func (r *Response) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Response object.
	return nil
}

func (r *Response) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// FindExtension will attempt to locate an extension using the supplied key
	return nil
}

func (r *Response) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all OAuthFlow extensions and satisfies the low.HasExtensions interface.
func (r *Response) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// FindContent will attempt to locate a MediaType instance using the supplied key.
}

func (r *Response) FindContent(cType string) *low.ValueReference[*MediaType] {
	_ = "STUB: not implemented"
	return nil
}

// FindHeader will attempt to locate a Header instance using the supplied key.
func (r *Response) FindHeader(hType string) *low.ValueReference[*Header] {
	_ = "STUB: not implemented"
	return nil
}

// FindLink will attempt to locate a Link instance using the supplied key.
func (r *Response) FindLink(hType string) *low.ValueReference[*Link] {
	_ = "STUB: not implemented"
	return nil
}

// Build will extract headers, extensions, content and links from node.
func (r *Response) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// extract headers

// handle links if set

// Hash will return a consistent Hash of the Response object
func (r *Response) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
