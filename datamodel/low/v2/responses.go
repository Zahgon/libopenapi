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

// Responses is a low-level representation of a Swagger / OpenAPI 2 Responses object.
type Responses struct {
	Codes      *orderedmap.Map[low.KeyReference[string], low.ValueReference[*Response]]
	Default    low.NodeReference[*Response]
	Extensions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
}

// GetExtensions returns all Responses extensions and satisfies the low.HasExtensions interface.
func (r *Responses) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Build will extract default value and extensions from node.
}

func (r *Responses) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
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

// FindResponseByCode will attempt to locate a Response instance using an HTTP response code string.
func (r *Responses) FindResponseByCode(code string) *low.ValueReference[*Response] {
	_ = "STUB: not implemented"
	return nil
}

// Hash will return a consistent Hash of the Responses object
func (r *Responses) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
