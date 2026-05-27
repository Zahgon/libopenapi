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

// Responses represents a low-level OpenAPI 3+ Responses object.
//
// It's a container for the expected responses of an operation. The container maps an HTTP response code to the
// expected response.
//
// The specification is not necessarily expected to cover all possible HTTP response codes because they may not be
// known in advance. However, documentation is expected to cover a successful operation response and any known errors.
//
// The default MAY be used as a default response object for all HTTP codes that are not covered individually by
// the Responses Object.
//
// The Responses Object MUST contain at least one response code, and if only one response code is provided it SHOULD
// be the response for a successful operation call.
//   - https://spec.openapis.org/oas/v3.1.0#responses-object
//
// This structure is identical to the v2 version, however they use different response types, hence
// the duplication. Perhaps in the future we could use generics here, but for now to keep things
// simple, they are broken out into individual versions.
type Responses struct {
	Codes      *orderedmap.Map[low.KeyReference[string], low.ValueReference[*Response]]
	Default    low.NodeReference[*Response]
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

// GetIndex returns the index.SpecIndex instance attached to the Responses object.
func (r *Responses) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Responses object.
	return nil
}

func (r *Responses) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetRootNode returns the root yaml node of the Responses object.
	return *new(context.Context)
}

func (r *Responses) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Responses object.
	return nil
}

func (r *Responses) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetExtensions returns all Responses extensions and satisfies the low.HasExtensions interface.
	return nil
}

func (r *Responses) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Build will extract default response and all Response objects for each code
}

func (r *Responses) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// default is bundled into codes, pull it out

// remove default from codes

func (r *Responses) getDefault() *low.NodeReference[*Response] {
	_ = "STUB: not implemented"
	return nil
}

// used to remove default from codes extracted by Build()
func (r *Responses) deleteCode(code string) { _ = "STUB: not implemented"; return }

// should never be nil, but, you never know... science and all that!

// FindResponseByCode will attempt to locate a Response using an HTTP response code.
func (r *Responses) FindResponseByCode(code string) *low.ValueReference[*Response] {
	_ = "STUB: not implemented"
	return nil
}

// Hash will return a consistent Hash of the Responses object
func (r *Responses) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
