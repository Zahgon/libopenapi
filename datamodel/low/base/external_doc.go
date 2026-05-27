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

// ExternalDoc represents a low-level External Documentation object as defined by OpenAPI 2 and 3
//
// Allows referencing an external resource for extended documentation.
//
//	v2 - https://swagger.io/specification/v2/#externalDocumentationObject
//	v3 - https://spec.openapis.org/oas/v3.1.0#external-documentation-object
type ExternalDoc struct {
	Description low.NodeReference[string]
	URL         low.NodeReference[string]
	Extensions  *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode     *yaml.Node
	RootNode    *yaml.Node
	index       *index.SpecIndex
	context     context.Context
	nodeStore   sync.Map
	reference   low.Reference
	*low.Reference
	low.NodeMap
}

// FindExtension returns a ValueReference containing the extension value, if found.
func (ex *ExternalDoc) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode will return the root yaml node of the ExternalDoc object
func (ex *ExternalDoc) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode will return the key yaml node of the ExternalDoc object
	return nil
}

func (ex *ExternalDoc) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract extensions from the ExternalDoc instance.
	return nil
}

func (ex *ExternalDoc) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all ExternalDoc extensions and satisfies the low.HasExtensions interface.
func (ex *ExternalDoc) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil
}

func (ex *ExternalDoc) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// GetIndex returns the index.SpecIndex instance attached to the ExternalDoc object
func (ex *ExternalDoc) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the ExternalDoc object
	return nil
}

func (ex *ExternalDoc) GetContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
