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

// License is a low-level representation of a License object as defined by OpenAPI 2 and OpenAPI 3
//
//	v2 - https://swagger.io/specification/v2/#licenseObject
//	v3 - https://spec.openapis.org/oas/v3.1.0#license-object
type License struct {
	Name       low.NodeReference[string]
	URL        low.NodeReference[string]
	Identifier low.NodeReference[string]
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

// Build out a license, complain if both a URL and identifier are present as they are mutually exclusive
func (l *License) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// GetIndex will return the index.SpecIndex instance attached to the License object
func (l *License) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext will return the context.Context instance used when building the License object
	return nil
}

func (l *License) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetRootNode will return the root yaml node of the License object
	return *new(context.Context)
}

func (l *License) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode will return the key yaml node of the License object
	return nil
}

func (l *License) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Hash will return a consistent hash of the License object
	return nil
}

func (l *License) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Note: Extensions are not included in the hash for License

// GetExtensions returns all extensions for License
func (l *License) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil
}
