// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"context"
	"sync"

	"github.com/pb33f/libopenapi/orderedmap"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"go.yaml.in/yaml/v4"
)

// Info represents a low-level Info object as defined by both OpenAPI 2 and OpenAPI 3.
//
// The object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented
// in editing or documentation generation tools for convenience.
//
//	v2 - https://swagger.io/specification/v2/#infoObject
//	v3 - https://spec.openapis.org/oas/v3.1.0#info-object
type Info struct {
	Title          low.NodeReference[string]
	Summary        low.NodeReference[string]
	Description    low.NodeReference[string]
	TermsOfService low.NodeReference[string]
	Contact        low.NodeReference[*Contact]
	License        low.NodeReference[*License]
	Version        low.NodeReference[string]
	Extensions     *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode        *yaml.Node
	RootNode       *yaml.Node
	index          *index.SpecIndex
	context        context.Context
	nodeStore      sync.Map
	reference      low.Reference
	*low.Reference
	low.NodeMap
}

// FindExtension attempts to locate an extension with the supplied key
func (i *Info) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode will return the root yaml node of the Info object
func (i *Info) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode will return the key yaml node of the Info object
	return nil
}

func (i *Info) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetExtensions returns all extensions for Info
	return nil
}

func (i *Info) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Build will extract out the Contact and Info objects from the supplied root node.
}

func (i *Info) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// extract contact

// extract license

// GetIndex will return the index.SpecIndex instance attached to the Info object
func (i *Info) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext will return the context.Context instance used when building the Info object
	return nil
}

func (i *Info) GetContext() context.Context {
	_ = "STUB: not implemented"

	// Hash will return a consistent hash of the Info object
	return *new(context.Context)
}

func (i *Info) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
