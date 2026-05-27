// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/orderedmap"
)

// SecurityRequirement is a high-level representation of a Swagger / OpenAPI 3 SecurityRequirement object.
//
// SecurityRequirement lists the required security schemes to execute this operation. The object can have multiple
// security schemes declared in it which are all required (that is, there is a logical AND between the schemes).
//
// The name used for each property MUST correspond to a security scheme declared in the Security Definitions
//   - https://swagger.io/specification/v2/#securityDefinitionsObject
type SecurityRequirement struct {
	Requirements             *orderedmap.Map[string, []string] `json:"-" yaml:"-"`
	ContainsEmptyRequirement bool                              // if a requirement is empty (this means it's optional)
	low                      *base.SecurityRequirement
}

// NewSecurityRequirement creates a new high-level SecurityRequirement from a low-level one.
func NewSecurityRequirement(req *base.SecurityRequirement) *SecurityRequirement {
	_ = "STUB: not implemented"
	return nil
}

// to keep things fast, avoiding copying anything - makes it a little hard to read.

// GoLow returns the low-level SecurityRequirement used to create the high-level one.
func (s *SecurityRequirement) GoLow() *base.SecurityRequirement {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Discriminator instance that was used to create the high-level one, with no type
	return nil
}

func (s *SecurityRequirement) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the SecurityRequirement object as a byte slice.
	return *new(any)
}

func (s *SecurityRequirement) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the SecurityRequirement object.
		nil
}

func (s *SecurityRequirement) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for each key, extract all the values and order them.
