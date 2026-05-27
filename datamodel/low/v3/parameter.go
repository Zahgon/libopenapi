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

// Parameter represents a high-level OpenAPI 3+ Parameter object, that is backed by a low-level one.
//
// A unique parameter is defined by a combination of a name and location.
//   - https://spec.openapis.org/oas/v3.1.0#parameter-object
type Parameter struct {
	KeyNode         *yaml.Node
	RootNode        *yaml.Node
	Name            low.NodeReference[string]
	In              low.NodeReference[string]
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
	index           *index.SpecIndex
	context         context.Context
	nodeStore       sync.Map
	reference       low.Reference
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the Parameter object.
func (p *Parameter) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Parameter
	return nil
}

func (p *Parameter) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetRootNode returns the root yaml node of the Parameter object.
	return *new(context.Context)
}

func (p *Parameter) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Parameter object.
	return nil
}

func (p *Parameter) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// FindContent will attempt to locate a MediaType instance using the specified name.
	return nil
}

func (p *Parameter) FindContent(cType string) *low.ValueReference[*MediaType] {
	_ = "STUB: not implemented"
	return nil
}

// FindExample will attempt to locate a base.Example instance using the specified name.
func (p *Parameter) FindExample(eType string) *low.ValueReference[*base.Example] {
	_ = "STUB: not implemented"
	return nil
}

// FindExtension attempts to locate an extension using the specified name.
func (p *Parameter) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all extensions for Parameter.
func (p *Parameter) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Build will extract examples, extensions and content/media types.
}

func (p *Parameter) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// handle example if set.

// handle schema

// handle examples if set.

// Only consider examples if they are defined in the root node.

// handle content, if set.

// Hash will return a consistent Hash of the Parameter object
func (p *Parameter) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// IsParameter compliance methods.

func (p *Parameter) GetName() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (p *Parameter) GetIn() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (p *Parameter) GetDescription() *low.NodeReference[string] {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parameter) GetRequired() *low.NodeReference[bool] { _ = "STUB: not implemented"; return nil }

func (p *Parameter) GetDeprecated() *low.NodeReference[bool] { _ = "STUB: not implemented"; return nil }

func (p *Parameter) GetAllowEmptyValue() *low.NodeReference[bool] {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parameter) GetSchema() *low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (p *Parameter) GetStyle() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (p *Parameter) GetAllowReserved() *low.NodeReference[bool] {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parameter) GetExplode() *low.NodeReference[bool] { _ = "STUB: not implemented"; return nil }

func (p *Parameter) GetExample() *low.NodeReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parameter) GetExamples() *low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (p *Parameter) GetContent() *low.NodeReference[any] { _ = "STUB: not implemented"; return nil }
