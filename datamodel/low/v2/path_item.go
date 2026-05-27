// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

var buildPathItemOperationModel = low.BuildModel

// PathItem represents a low-level Swagger / OpenAPI 2 PathItem object.
//
// Describes the operations available on a single path. A Path Item may be empty, due to ACL constraints.
// The path itself is still exposed to the tooling, but will not know which operations and parameters
// are available.
//
//   - https://swagger.io/specification/v2/#pathItemObject
type PathItem struct {
	Ref        low.NodeReference[string]
	Get        low.NodeReference[*Operation]
	Put        low.NodeReference[*Operation]
	Post       low.NodeReference[*Operation]
	Delete     low.NodeReference[*Operation]
	Options    low.NodeReference[*Operation]
	Head       low.NodeReference[*Operation]
	Patch      low.NodeReference[*Operation]
	Parameters low.NodeReference[[]low.ValueReference[*Parameter]]
	Extensions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
}

// FindExtension will attempt to locate an extension given a name.
func (p *PathItem) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all PathItem extensions and satisfies the low.HasExtensions interface.
func (p *PathItem) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Build will extract extensions, parameters and operations for all methods. Every method is handled
	// asynchronously, in order to keep things moving quickly for complex operations.
}

func (p *PathItem) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// extract parameters

// because (for some reason) the spec for swagger docs allows for a '$ref' property for path items.
// this is kinda nuts, because '$ref' is a reserved keyword for JSON references, which is ALSO used
// in swagger. Why this choice was made, I do not know.

// the only thing we now care about is handling operations, filter out anything that's not a verb.

// ignore everything else.

// all operations have been superficially built,
// now we need to build out the operation, we will do this asynchronously for speed.

// nothing to do.

// Hash will return a consistent Hash of the PathItem object
func (p *PathItem) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
