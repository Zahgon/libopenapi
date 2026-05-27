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

// XML represents a low-level representation of an XML object defined by all versions of OpenAPI.
//
// A metadata object that allows for more fine-tuned XML model definitions.
//
// When using arrays, XML element names are not inferred (for singular/plural forms) and the name property SHOULD be
// used to add that information. See examples for expected behavior.
//
//	v2 - https://swagger.io/specification/v2/#xmlObject
//	v3 - https://swagger.io/specification/#xml-object
type XML struct {
	Name       low.NodeReference[string]
	Namespace  low.NodeReference[string]
	Prefix     low.NodeReference[string]
	Attribute  low.NodeReference[bool]
	NodeType   low.NodeReference[string] // OpenAPI 3.2+ nodeType field (replaces deprecated attribute field)
	Wrapped    low.NodeReference[bool]
	Extensions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	RootNode   *yaml.Node
	index      *index.SpecIndex
	context    context.Context
	nodeStore  sync.Map
	reference  low.Reference
	*low.Reference
	low.NodeMap
}

// Build will extract extensions from the XML instance.
func (x *XML) Build(root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all Tag extensions and satisfies the low.HasExtensions interface.
func (x *XML) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// GetRootNode returns the root yaml node of the Tag object
}

func (x *XML) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetIndex returns the index of the XML object
	return nil
}

func (x *XML) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// Hash generates a hash of the XML object using properties
	return nil
}

func (x *XML) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
