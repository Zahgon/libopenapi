// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Response is a representation of a high-level Swagger / OpenAPI 2 Response object, backed by a low-level one.
//
// Response describes a single response from an API Operation
//   - https://swagger.io/specification/v2/#responseObject
type Response struct {
	Description low.NodeReference[string]
	Schema      low.NodeReference[*base.SchemaProxy]
	Headers     low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*Header]]]
	Examples    low.NodeReference[*Examples]
	Extensions  *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
}

// FindExtension will attempt to locate an extension value given a key to lookup.
func (r *Response) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all Response extensions and satisfies the low.HasExtensions interface.
func (r *Response) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// FindHeader will attempt to locate a Header value, given a key
}

func (r *Response) FindHeader(hType string) *low.ValueReference[*Header] {
	_ = "STUB: not implemented"
	return nil
}

// Build will extract schema, extensions, examples and headers from node
func (r *Response) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// extract examples

// extract headers

// Hash will return a consistent Hash of the Response object
func (r *Response) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
