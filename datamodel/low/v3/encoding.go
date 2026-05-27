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

// Encoding represents a low-level OpenAPI 3+ Encoding object
//   - https://spec.openapis.org/oas/v3.1.0#encoding-object
type Encoding struct {
	ContentType   low.NodeReference[string]
	Headers       low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Header]]]
	Style         low.NodeReference[string]
	Explode       low.NodeReference[bool]
	AllowReserved low.NodeReference[bool]
	KeyNode       *yaml.Node
	RootNode      *yaml.Node
	index         *index.SpecIndex
	context       context.Context
	nodeStore     sync.Map
	reference     low.Reference
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the Encoding object
func (en *Encoding) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Encoding object
	return nil
}

func (en *Encoding) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindHeader attempts to locate a Header with the supplied name
	return *new(context.Context)
}

func (en *Encoding) FindHeader(hType string) *low.ValueReference[*Header] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the Encoding object
func (en *Encoding) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Encoding object
	return nil
}

func (en *Encoding) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Hash will return a consistent Hash of the Encoding object
	return nil
}

func (en *Encoding) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Build will extract all Header objects from supplied node.
func (en *Encoding) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}
