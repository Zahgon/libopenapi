// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"context"
	"sync"

	"github.com/pb33f/libopenapi/orderedmap"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"go.yaml.in/yaml/v4"
)

// Callback represents a low-level Callback object for OpenAPI 3+.
//
// A map of possible out-of band callbacks related to the parent operation. Each value in the map is a
// PathItem Object that describes a set of requests that may be initiated by the API provider and the expected
// responses. The key value used to identify the path item object is an expression, evaluated at runtime,
// that identifies a URL to use for the callback operation.
//   - https://spec.openapis.org/oas/v3.1.0#callback-object
type Callback struct {
	Expression *orderedmap.Map[low.KeyReference[string], low.ValueReference[*PathItem]]
	Extensions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode    *yaml.Node
	RootNode   *yaml.Node
	index      *index.SpecIndex
	context    context.Context
	nodeStore  sync.Map
	reference  low.Reference
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the Callback object
func (cb *Callback) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Callback object
	return nil
}

func (cb *Callback) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetExtensions returns all Callback extensions and satisfies the low.HasExtensions interface.
	return *new(context.Context)
}

func (cb *Callback) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// GetRootNode returns the root yaml node of the Callback object
}

func (cb *Callback) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Callback object
	return nil
}

func (cb *Callback) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// FindExpression will locate a string expression and return a ValueReference containing the located PathItem
	return nil
}

func (cb *Callback) FindExpression(exp string) *low.ValueReference[*PathItem] {
	_ = "STUB: not implemented"
	return nil
}

// Build will extract extensions, expressions and PathItem objects for Callback
func (cb *Callback) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Hash will return a consistent Hash of the Callback object
func (cb *Callback) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
