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

// Link represents a low-level OpenAPI 3+ Link object.
//
// The Link object represents a possible design-time link for a response. The presence of a link does not guarantee the
// caller’s ability to successfully invoke it, rather it provides a known relationship and traversal mechanism between
// responses and other operations.
//
// Unlike dynamic links (i.e. links provided in the response payload), the OAS linking mechanism does not require
// link information in the runtime response.
//
// For computing links, and providing instructions to execute them, a runtime expression is used for accessing values
// in an operation and using them as parameters while invoking the linked operation.
//   - https://spec.openapis.org/oas/v3.1.0#link-object
type Link struct {
	OperationRef low.NodeReference[string]
	OperationId  low.NodeReference[string]
	Parameters   low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[string]]]
	RequestBody  low.NodeReference[string]
	Description  low.NodeReference[string]
	Server       low.NodeReference[*Server]
	Extensions   *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode      *yaml.Node
	RootNode     *yaml.Node
	index        *index.SpecIndex
	context      context.Context
	nodeStore    sync.Map
	reference    low.Reference
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the Link object
func (l *Link) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Link object
	return nil
}

func (l *Link) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetExtensions returns all Link extensions and satisfies the low.HasExtensions interface.
	return *new(context.Context)
}

func (l *Link) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// FindParameter will attempt to locate a parameter string value, using a parameter name input.
}

func (l *Link) FindParameter(pName string) *low.ValueReference[string] {
	_ = "STUB: not implemented"
	return nil
}

// FindExtension will attempt to locate an extension with a specific key
func (l *Link) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the Link object
func (l *Link) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Link object
	return nil
}

func (l *Link) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract extensions and servers from the node.
	return nil
}

func (l *Link) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// extract parameter nodes.

// extract server.

// Hash will return a consistent Hash of the Link object
func (l *Link) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
