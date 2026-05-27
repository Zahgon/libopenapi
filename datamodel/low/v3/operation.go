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

// Operation is a low-level representation of an OpenAPI 3+ Operation object.
//
// An Operation is perhaps the most important object of the entire specification. Everything of value
// happens here. The entire being for existence of this library and the specification, is this Operation.
//   - https://spec.openapis.org/oas/v3.1.0#operation-object
type Operation struct {
	Tags         low.NodeReference[[]low.ValueReference[string]]
	Summary      low.NodeReference[string]
	Description  low.NodeReference[string]
	ExternalDocs low.NodeReference[*base.ExternalDoc]
	OperationId  low.NodeReference[string]
	Parameters   low.NodeReference[[]low.ValueReference[*Parameter]]
	RequestBody  low.NodeReference[*RequestBody]
	Responses    low.NodeReference[*Responses]
	Callbacks    low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Callback]]]
	Deprecated   low.NodeReference[bool]
	Security     low.NodeReference[[]low.ValueReference[*base.SecurityRequirement]]
	Servers      low.NodeReference[[]low.ValueReference[*Server]]
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

// GetIndex returns the index.SpecIndex instance attached to the Operation object.
func (o *Operation) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Operation object.
	return nil
}

func (o *Operation) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindCallback will attempt to locate a Callback instance by the supplied name.
	return *new(context.Context)
}

func (o *Operation) FindCallback(callback string) *low.ValueReference[*Callback] {
	_ = "STUB: not implemented"
	return nil
}

// FindSecurityRequirement will attempt to locate a security requirement string from a supplied name.
func (o *Operation) FindSecurityRequirement(name string) []low.ValueReference[string] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the Operation object
func (o *Operation) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Operation object
	return nil
}

func (o *Operation) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract external docs, parameters, request body, responses, callbacks, security and servers.
	return nil
}

func (o *Operation) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// extract externalDocs

// extract parameters

// extract request body

// extract tags, but only extract nodes, the model has already been built

// extract responses

// extract callbacks

// extract security

// if security is defined and requirements are provided.

// if security is set, but no requirements are defined.
// https://github.com/pb33f/libopenapi/issues/111

// extract servers

// Hash will return a consistent Hash of the Operation object
func (o *Operation) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Pre-allocate keys for sorting

// Tags array - pre-allocate and sort

// Servers array - pre-allocate and sort

// Parameters array - pre-allocate and sort

// Callbacks

// Extensions

// methods to satisfy swagger operations interface

func (o *Operation) GetTags() low.NodeReference[[]low.ValueReference[string]] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operation) GetSummary() low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetDescription() low.NodeReference[string] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operation) GetExternalDocs() low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetOperationId() low.NodeReference[string] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operation) GetDeprecated() low.NodeReference[bool] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operation) GetResponses() low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetRequestBody() low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetParameters() low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetSecurity() low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetServers() low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetCallbacks() low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Callback]]] {
	_ = "STUB: not implemented"
	return nil
}
