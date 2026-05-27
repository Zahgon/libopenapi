// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"context"
	"hash/maphash"
	"sync"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Components represents a low-level OpenAPI 3+ Components Object, that is backed by a low-level one.
//
// Holds a set of reusable objects for different aspects of the OAS. All objects defined within the components object
// will have no effect on the API unless they are explicitly referenced from properties outside the components object.
//   - https://spec.openapis.org/oas/v3.1.0#components-object
type Components struct {
	Schemas         low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*base.SchemaProxy]]]
	Responses       low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Response]]]
	Parameters      low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Parameter]]]
	Examples        low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*base.Example]]]
	RequestBodies   low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*RequestBody]]]
	Headers         low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Header]]]
	SecuritySchemes low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*SecurityScheme]]]
	Links           low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Link]]]
	Callbacks       low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Callback]]]
	PathItems       low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*PathItem]]]
	MediaTypes      low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*MediaType]]] // OpenAPI 3.2+ mediaTypes section
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

type componentBuildResult[T any] struct {
	key   low.KeyReference[string]
	value low.ValueReference[T]
}

type componentInput struct {
	node         *yaml.Node
	currentLabel *yaml.Node
}

// GetIndex returns the index.SpecIndex instance attached to the Components object
func (co *Components) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Components object
	return nil
}

func (co *Components) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetExtensions returns all Components extensions and satisfies the low.HasExtensions interface.
	return *new(context.Context)
}

func (co *Components) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// GetRootNode returns the root yaml node of the Components object
}

func (co *Components) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Components object
	return nil
}

func (co *Components) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Hash will return a consistent Hash of the Components object
	return nil
}

func (co *Components) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

func generateHashForObjectMap[T any](collection *orderedmap.Map[low.KeyReference[string], low.ValueReference[T]], h *maphash.Hash) {
	_ = "STUB: not implemented"
	return
}

// FindExtension attempts to locate an extension with the supplied key
func (co *Components) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// FindSchema attempts to locate a SchemaProxy from 'schemas' with a specific name
func (co *Components) FindSchema(schema string) *low.ValueReference[*base.SchemaProxy] {
	_ = "STUB: not implemented"
	return nil
}

// FindResponse attempts to locate a Response from 'responses' with a specific name
func (co *Components) FindResponse(response string) *low.ValueReference[*Response] {
	_ = "STUB: not implemented"
	return nil
}

// FindParameter attempts to locate a Parameter from 'parameters' with a specific name
func (co *Components) FindParameter(response string) *low.ValueReference[*Parameter] {
	_ = "STUB: not implemented"
	return nil
}

// FindSecurityScheme attempts to locate a SecurityScheme from 'securitySchemes' with a specific name
func (co *Components) FindSecurityScheme(sScheme string) *low.ValueReference[*SecurityScheme] {
	_ = "STUB: not implemented"
	return nil
}

// FindExample attempts tp
func (co *Components) FindExample(example string) *low.ValueReference[*base.Example] {
	_ = "STUB: not implemented"
	return nil
}

func (co *Components) FindRequestBody(requestBody string) *low.ValueReference[*RequestBody] {
	_ = "STUB: not implemented"
	return nil
}

func (co *Components) FindHeader(header string) *low.ValueReference[*Header] {
	_ = "STUB: not implemented"
	return nil
}

func (co *Components) FindLink(link string) *low.ValueReference[*Link] {
	_ = "STUB: not implemented"
	return nil
}

func (co *Components) FindPathItem(path string) *low.ValueReference[*PathItem] {
	_ = "STUB: not implemented"
	return nil
}

func (co *Components) FindCallback(callback string) *low.ValueReference[*Callback] {
	_ = "STUB: not implemented"
	return nil
}

// FindMediaType attempts to locate a MediaType from 'mediaTypes' with a specific name
func (co *Components) FindMediaType(mediaType string) *low.ValueReference[*MediaType] {
	_ = "STUB: not implemented"
	return nil
}

// Build converts root YAML node containing components to low level model.
// Process each component in parallel.
func (co *Components) Build(ctx context.Context, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// extractComponentValues converts all the YAML nodes of a component type to
// low level model.
// Process each node in parallel.
func extractComponentValues[T low.Buildable[N], N any](ctx context.Context, label string, root *yaml.Node, idx *index.SpecIndex, co *Components) (low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[T]]], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// build.

// for SchemaProxy, use the transformed node from sp.vn instead of original node

// Check if the type implements low.HasKeyNode

// use transformed node if available
