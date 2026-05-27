// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"context"
	"sync"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Contact represents a low-level representation of the Contact definitions found at
//
//	v2 - https://swagger.io/specification/v2/#contactObject
//	v3 - https://spec.openapis.org/oas/v3.1.0#contact-object
type Contact struct {
	Name       low.NodeReference[string]
	URL        low.NodeReference[string]
	Email      low.NodeReference[string]
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

func (c *Contact) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// GetIndex will return the index.SpecIndex instance attached to the Contact object
func (c *Contact) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext will return the context.Context instance used when building the Contact object
	return nil
}

func (c *Contact) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetRootNode will return the root yaml node of the Contact object
	return *new(context.Context)
}

func (c *Contact) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode will return the key yaml node of the Contact object
	return nil
}

func (c *Contact) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Hash will return a consistent hash of the Contact object
	return nil
}

func (c *Contact) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Note: Extensions are not included in the hash for Contact

// GetExtensions returns all extensions for Contact
func (c *Contact) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil
}
