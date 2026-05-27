// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	low "github.com/pb33f/libopenapi/datamodel/low/v2"
	"go.yaml.in/yaml/v4"
)

// Items is a high-level representation of a Swagger / OpenAPI 2 Items object, backed by a low level one.
// Items is a limited subset of JSON-Schema's items object. It is used by parameter definitions that are not
// located in "body"
//   - https://swagger.io/specification/v2/#itemsObject
type Items struct {
	Type             string
	Format           string
	CollectionFormat string
	Items            *Items
	Default          *yaml.Node
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
	Enum             []*yaml.Node
	MultipleOf       int
	low              *low.Items
}

// NewItems creates a new high-level Items instance from a low-level one.
func NewItems(items *low.Items) *Items { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Items object that was used to create the high-level one.
func (i *Items) GoLow() *low.Items { _ = "STUB: not implemented"; return nil }
