// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	v2low "github.com/pb33f/libopenapi/datamodel/low/v2"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Paths represents a high-level Swagger / OpenAPI Paths object, backed by a low-level one.
type Paths struct {
	PathItems  *orderedmap.Map[string, *PathItem]
	Extensions *orderedmap.Map[string, *yaml.Node]
	low        *v2low.Paths
}

// NewPaths creates a new high-level instance of Paths from a low-level one.
func NewPaths(paths *v2low.Paths) *Paths { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Paths instance that backs the high level one.
func (p *Paths) GoLow() *v2low.Paths { _ = "STUB: not implemented"; return nil }
