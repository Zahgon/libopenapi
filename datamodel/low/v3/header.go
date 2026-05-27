// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"context"
	"sync"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Header represents a low-level OpenAPI 3+ Header object.
//   - https://spec.openapis.org/oas/v3.1.0#header-object
type Header struct {
	Description     low.NodeReference[string]
	Required        low.NodeReference[bool]
	Deprecated      low.NodeReference[bool]
	AllowEmptyValue low.NodeReference[bool]
	Style           low.NodeReference[string]
	Explode         low.NodeReference[bool]
	AllowReserved   low.NodeReference[bool]
	Schema          low.NodeReference[*base.SchemaProxy]
	Example         low.NodeReference[*yaml.Node]
	Examples        low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*base.Example]]]
	Content         low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*MediaType]]]
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

// GetIndex returns the index.SpecIndex instance attached to the Header object
func (h *Header) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Header object
	return nil
}

func (h *Header) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension will attempt to locate an extension with the supplied name
	return *new(context.Context)
}

func (h *Header) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// FindExample will attempt to locate an Example with a specified name
func (h *Header) FindExample(eType string) *low.ValueReference[*base.Example] {
	_ = "STUB: not implemented"
	return nil
}

// FindContent will attempt to locate a MediaType definition, with a specified name
func (h *Header) FindContent(ext string) *low.ValueReference[*MediaType] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the Header object
func (h *Header) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Header object
	return nil
}

func (h *Header) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetExtensions returns all Header extensions and satisfies the low.HasExtensions interface.
	return nil
}

func (h *Header) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent Hash of the Header object
}

func (h *Header) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Build will extract extensions, examples, schema and content/media types from node.
func (h *Header) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// handle example if set.

// handle examples if set.

// handle schema

// handle content, if set.

// Getter methods to satisfy OpenAPIHeader interface.

func (h *Header) GetDescription() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetRequired() *low.NodeReference[bool] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetDeprecated() *low.NodeReference[bool] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetAllowEmptyValue() *low.NodeReference[bool] {
	_ = "STUB: not implemented"
	return nil
}

func (h *Header) GetSchema() *low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetStyle() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetAllowReserved() *low.NodeReference[bool] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetExplode() *low.NodeReference[bool] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetExample() *low.NodeReference[*yaml.Node] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetExamples() *low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetContent() *low.NodeReference[any] { _ = "STUB: not implemented"; return nil }
