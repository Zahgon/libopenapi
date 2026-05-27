// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"github.com/pb33f/libopenapi/datamodel/high/base"
	lowv2 "github.com/pb33f/libopenapi/datamodel/low/v2"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Response is a representation of a high-level Swagger / OpenAPI 2 Response object, backed by a low-level one.
// Response describes a single response from an API Operation
//   - https://swagger.io/specification/v2/#responseObject
type Response struct {
	Description string
	Schema      *base.SchemaProxy
	Headers     *orderedmap.Map[string, *Header]
	Examples    *Example
	Extensions  *orderedmap.Map[string, *yaml.Node]
	low         *lowv2.Response
}

// NewResponse creates a new high-level instance of Response from a low level one.
func NewResponse(response *lowv2.Response) *Response { _ = "STUB: not implemented"; return nil }

// GoLow will return the low-level Response instance used to create the high level one.
func (r *Response) GoLow() *lowv2.Response { _ = "STUB: not implemented"; return nil }
