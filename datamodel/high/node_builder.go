// Copyright 2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package high

import (
	"github.com/pb33f/libopenapi/datamodel/high/nodes"
	"github.com/pb33f/libopenapi/datamodel/low"
	"go.yaml.in/yaml/v4"
)

// NodeBuilder is a structure used by libopenapi high-level objects, to render themselves back to YAML.
// this allows high-level objects to be 'mutable' because all changes will be rendered out.
type NodeBuilder struct {
	Version       float32
	Nodes         []*nodes.NodeEntry
	High          any
	Low           any
	Resolve       bool // If set to true, all references will be rendered inline
	RenderContext any  // Context for inline rendering cycle detection (*base.InlineRenderContext)
	Errors        []error
}

// RenderableInlineWithContext is an interface that can be implemented by types that support
// context-aware inline rendering for proper cycle detection in concurrent scenarios.
// The context parameter should be *base.InlineRenderContext but is typed as any to avoid import cycles.
type RenderableInlineWithContext interface {
	MarshalYAMLInlineWithContext(ctx any) (interface{}, error)
}

const renderZero = "renderZero"

func originalFloatLexeme(value float64, lowValue any) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// NewNodeBuilder will create a new NodeBuilder instance, this is the only way to create a NodeBuilder.
// The function accepts a high level object and a low level object (need to be siblings/same type).
//
// Using reflection, a map of every field in the high level object is created, ready to be rendered.
func NewNodeBuilder(high any, low any) *NodeBuilder {
	_ = "STUB: not implemented"
	// create a new node builder
	return nil
}

// extract fields from the high level object and add them into our node builder.
// this will allow us to extract the line numbers from the low level object as well.

func (n *NodeBuilder) add(key string, i int) {
	_ = "STUB: not implemented"
	// only operate on exported fields.
	return
}

// if the key is 'Extensions' then we need to extract the keys from the map
// and add them to the node builder.

// If we have low extensions get the original lowest line number so we end up in the same place

// done, extensions are handled separately.

// find the field with the tag supplied.

// extract the value of the field

// create a new node entry

// trim this down

// if there is no low-level object, then we cannot extract line numbers,
// so skip and default to 0, which means a new entry to the spec.
// this will place new content and the top of the rendered object.

// everything else, weight it to the bottom of the rendered object.
// this is things that we have no way of knowing where they should be placed.

func (n *NodeBuilder) renderReference(fg low.IsReferenced) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// Render will render the NodeBuilder back to a YAML node, iterating over every NodeEntry defined
func (n *NodeBuilder) Render() *yaml.Node { _ = "STUB: not implemented"; return nil }

// order nodes by line number, retain original order

// AddYAMLNode will add a new *yaml.Node to the parent node, using the tag, key and value provided.
// If the value is nil, then the node will not be added. This method is recursive, so it will dig down
// into any non-scalar types.
func (n *NodeBuilder) AddYAMLNode(parent *yaml.Node, entry *nodes.NodeEntry) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// check the type

// Always create float node for float64 values, even if they don't contain decimal points
// This handles cases like negative zero (-0.0) which formats as "-0" but should remain float

// Reset skip at the start of each iteration to handle items without low-level models
// (e.g., newly created high-level objects appended to an existing slice)

// check if this is a reference.

// try and render inline, if we can, otherwise treat as normal.
// Prefer a context-aware method when RenderContext is available

// check if this is a pointer or not.

// try an inline render if we can, otherwise there is no option but to default to the
// full render. Prefer a context-aware method when RenderContext is available

// check if the value is a bool, int or float

// Always create float node for float64 values, even if they're whole numbers
// This handles cases like negative zero (-0.0) and ensures type consistency

// check if is a node and it's null

// Renderable is an interface that can be implemented by types that provide a custom MarshalYAML method.
type Renderable interface {
	MarshalYAML() (interface{}, error)
}

// RenderableInline is an interface that can be implemented by types that provide a custom MarshalYAML method.
type RenderableInline interface {
	MarshalYAMLInline() (interface{}, error)
}
