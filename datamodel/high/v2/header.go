// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	low "github.com/pb33f/libopenapi/datamodel/low/v2"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Header Represents a high-level Swagger / OpenAPI 2 Header object, backed by a low-level one.
// A Header is essentially identical to a Parameter, except it does not contain 'name' or 'in' properties.
//   - https://swagger.io/specification/v2/#headerObject
type Header struct {
	Type             string
	Format           string
	Description      string
	Items            *Items
	CollectionFormat string
	Default          any
	Maximum          int
	ExclusiveMaximum bool
	Minimum          int
	ExclusiveMinimum bool
	MaxLength        int
	MinLength        int
	Pattern          string
	MaxItems         int
	MinItems         int
	UniqueItems      bool
	Enum             []any
	MultipleOf       int
	Extensions       *orderedmap.Map[string, *yaml.Node]
	low              *low.Header
}

// NewHeader will create a new high-level Swagger / OpenAPI 2 Header instance, from a low-level one.
func NewHeader(header *low.Header) *Header { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level header used to create the high-level one.
func (h *Header) GoLow() *low.Header { _ = "STUB: not implemented"; return nil }
