// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	lowv2 "github.com/pb33f/libopenapi/datamodel/low/v2"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Example represents a high-level Swagger / OpenAPI 2 Example object, backed by a low level one.
// Allows sharing examples for operation responses
//   - https://swagger.io/specification/v2/#exampleObject
type Example struct {
	Values *orderedmap.Map[string, *yaml.Node]
	low    *lowv2.Examples
}

// NewExample creates a new high-level Example instance from a low-level one.
func NewExample(examples *lowv2.Examples) *Example { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Example used to create the high-level one.
func (e *Example) GoLow() *lowv2.Examples { _ = "STUB: not implemented"; return nil }
