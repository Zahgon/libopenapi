// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	lowV2 "github.com/pb33f/libopenapi/datamodel/low/v2"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// PathItem represents a high-level Swagger / OpenAPI 2 PathItem object backed by a low-level one.
//
// Describes the operations available on a single path. A Path Item may be empty, due to ACL constraints.
// The path itself is still exposed to the tooling, but will not know which operations and parameters
// are available.
//   - https://swagger.io/specification/v2/#pathItemObject
type PathItem struct {
	Ref        string
	Get        *Operation
	Put        *Operation
	Post       *Operation
	Delete     *Operation
	Options    *Operation
	Head       *Operation
	Patch      *Operation
	Parameters []*Parameter
	Extensions *orderedmap.Map[string, *yaml.Node]
	low        *lowV2.PathItem
}

// NewPathItem will create a new high-level PathItem from a low-level one. All paths are built out asynchronously.
func NewPathItem(pathItem *lowV2.PathItem) *PathItem { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level PathItem used to create the high-level one.
func (p *PathItem) GoLow() *lowV2.PathItem { _ = "STUB: not implemented"; return nil }

func (p *PathItem) GetOperations() *orderedmap.Map[string, *Operation] {
	_ = "STUB: not implemented"
	return nil
}

// TODO: this is a bit of a hack, but it works for now. We might just want to actually pull the data out of the document as a map and split it into the individual operations
