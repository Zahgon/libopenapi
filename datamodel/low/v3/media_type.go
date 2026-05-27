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

// MediaType represents a low-level OpenAPI MediaType object.
//
// Each Media Type Object provides schema and examples for the media type identified by its key.
//   - https://spec.openapis.org/oas/v3.1.0#media-type-object
type MediaType struct {
	Schema       low.NodeReference[*base.SchemaProxy]
	ItemSchema   low.NodeReference[*base.SchemaProxy]
	Example      low.NodeReference[*yaml.Node]
	Examples     low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*base.Example]]]
	Encoding     low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Encoding]]]
	ItemEncoding low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Encoding]]]
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

// GetIndex returns the index.SpecIndex instance attached to the MediaType object.
func (mt *MediaType) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the MediaType object.
	return nil
}

func (mt *MediaType) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetExtensions returns all MediaType extensions and satisfies the low.HasExtensions interface.
	return *new(context.Context)
}

func (mt *MediaType) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// FindExtension will attempt to locate an extension with the supplied name.
}

func (mt *MediaType) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// FindPropertyEncoding will attempt to locate an Encoding value with a specific name.
func (mt *MediaType) FindPropertyEncoding(eType string) *low.ValueReference[*Encoding] {
	_ = "STUB: not implemented"
	return nil
}

// FindExample will attempt to locate an Example with a specific name.
func (mt *MediaType) FindExample(eType string) *low.ValueReference[*base.Example] {
	_ = "STUB: not implemented"
	return nil
}

// GetAllExamples will extract all examples from the MediaType instance.
func (mt *MediaType) GetAllExamples() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*base.Example]] {
	_ = "STUB: not implemented"
	return nil

	// GetRootNode returns the root yaml node of the MediaType object.
}

func (mt *MediaType) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the MediaType object.
	return nil
}

func (mt *MediaType) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract examples, extensions, schema and encoding from node.
	return nil
}

func (mt *MediaType) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// handle example if set.

// handle schema

// handle examples if set.

// handle encoding

// handle itemSchema

// handle itemEncoding

// Hash will return a consistent Hash of the MediaType object
func (mt *MediaType) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
