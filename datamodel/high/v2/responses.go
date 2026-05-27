// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	low "github.com/pb33f/libopenapi/datamodel/low/v2"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Responses is a high-level representation of a Swagger / OpenAPI 2 Responses object, backed by a low level one.
type Responses struct {
	Codes      *orderedmap.Map[string, *Response]
	Default    *Response
	Extensions *orderedmap.Map[string, *yaml.Node]
	low        *low.Responses
}

// NewResponses will create a new high-level instance of Responses from a low-level one.
func NewResponses(responses *low.Responses) *Responses { _ = "STUB: not implemented"; return nil }

// GoLow will return the low-level object used to create the high-level one.
func (r *Responses) GoLow() *low.Responses { _ = "STUB: not implemented"; return nil }
