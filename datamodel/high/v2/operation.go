// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"github.com/pb33f/libopenapi/datamodel/high/base"
	low "github.com/pb33f/libopenapi/datamodel/low/v2"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Operation represents a high-level Swagger / OpenAPI 2 Operation object, backed by a low-level one.
// It describes a single API operation on a path.
//   - https://swagger.io/specification/v2/#operationObject
type Operation struct {
	Tags         []string
	Summary      string
	Description  string
	ExternalDocs *base.ExternalDoc
	OperationId  string
	Consumes     []string
	Produces     []string
	Parameters   []*Parameter
	Responses    *Responses
	Schemes      []string
	Deprecated   bool
	Security     []*base.SecurityRequirement
	Extensions   *orderedmap.Map[string, *yaml.Node]
	low          *low.Operation
}

// NewOperation creates a new high-level Operation instance from a low-level one.
func NewOperation(operation *low.Operation) *Operation { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level operation used to create the high-level one.
func (o *Operation) GoLow() *low.Operation { _ = "STUB: not implemented"; return nil }
