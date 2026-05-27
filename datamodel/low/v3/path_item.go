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

// PathItem represents a low-level OpenAPI 3+ PathItem object.
//
// Describes the operations available on a single path. A Path Item MAY be empty, due to ACL constraints.
// The path itself is still exposed to the documentation viewer, but they will not know which operations and parameters
// are available.
//   - https://spec.openapis.org/oas/v3.1.0#path-item-object
type PathItem struct {
	Description          low.NodeReference[string]
	Summary              low.NodeReference[string]
	Get                  low.NodeReference[*Operation]
	Put                  low.NodeReference[*Operation]
	Post                 low.NodeReference[*Operation]
	Delete               low.NodeReference[*Operation]
	Options              low.NodeReference[*Operation]
	Head                 low.NodeReference[*Operation]
	Patch                low.NodeReference[*Operation]
	Trace                low.NodeReference[*Operation]
	Query                low.NodeReference[*Operation]
	AdditionalOperations low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.NodeReference[*Operation]]] // OpenAPI 3.2+ additional operations
	Servers              low.NodeReference[[]low.ValueReference[*Server]]
	Parameters           low.NodeReference[[]low.ValueReference[*Parameter]]
	Extensions           *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode              *yaml.Node
	RootNode             *yaml.Node
	index                *index.SpecIndex
	context              context.Context
	nodeStore            sync.Map
	reference            low.Reference
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the PathItem object.
func (p *PathItem) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the PathItem object.
	return nil
}

func (p *PathItem) GetContext() context.Context {
	_ = "STUB: not implemented"

	// Hash will return a consistent Hash of the PathItem object
	return *new(context.Context)
}

func (p *PathItem) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Process AdditionalOperations with pre-allocation and sorting

// Process Parameters with pre-allocation and sorting

// Process Servers with pre-allocation and sorting

// GetRootNode returns the root yaml node of the PathItem object
func (p *PathItem) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the PathItem object
	return nil
}

func (p *PathItem) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// FindExtension attempts to find an extension
	return nil
}

func (p *PathItem) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all PathItem extensions and satisfies the low.HasExtensions interface.
func (p *PathItem) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Build extracts extensions, parameters, servers and each http method defined.
	// everything is extracted asynchronously for speed.
}

func (p *PathItem) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// extract parameters

// https://github.com/pb33f/libopenapi/issues/388
// in the case where a user has an extension with the value 'parameters', make sure we handle
// it correctly, by not skipping.

// this

// check if this is an operation (either standard or additional)

// check if this looks like an HTTP method (and isn't a known non-operation field)

// ignore known non-operation fields

// assume it's an additional operation if it contains a mapping to an operation object

// ignore if not a map

// initialize additionalOps map if this is the first additional operation

// now we need to determine if these are inline additional operations, or just plonked into the root.

// resolve operation reference for each additional operation

// all operations have been superficially built,
// now we need to build out the operation, we will do this asynchronously for speed.

// assign additionalOperations if any were found

// build out each additional operation

// resolveOperationReference handles the resolution of operation references ($ref)
// Returns: foundContext, resolvedPathNode, isRef, refValue, refNode, error
func resolveOperationReference(ctx context.Context, pathNode *yaml.Node, idx *index.SpecIndex) (
	context.Context, *yaml.Node, bool, string, *yaml.Node, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, false, "", nil, nil
}

// According to OpenAPI spec the only valid $ref for paths is
// reference for the whole pathItem. Unfortunately, the internet is full of invalid specs
// even from trusted companies like DigitalOcean where they tend to
// use file $ref for each respective operation:
// /endpoint/call/name:
//   post:
//     $ref: 'file.yaml'
// Check if that is the case and resolve such thing properly too.

// If it's a node from file, tag is empty
