// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import v3 "github.com/pb33f/libopenapi/datamodel/low/v3"

// MediaTypeChanges represent changes made between two OpenAPI MediaType instances.
type MediaTypeChanges struct {
	*PropertyChanges
	SchemaChanges       *SchemaChanges              `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	ItemSchemaChanges   *SchemaChanges              `json:"itemSchemas,omitempty" yaml:"itemSchemas,omitempty"`
	ExtensionChanges    *ExtensionChanges           `json:"extensions,omitempty" yaml:"extensions,omitempty"`
	ExampleChanges      map[string]*ExampleChanges  `json:"examples,omitempty" yaml:"examples,omitempty"`
	EncodingChanges     map[string]*EncodingChanges `json:"encoding,omitempty" yaml:"encoding,omitempty"`
	ItemEncodingChanges map[string]*EncodingChanges `json:"itemEncoding,omitempty" yaml:"itemEncoding,omitempty"`
}

// GetAllChanges returns a slice of all changes made between MediaType objects
func (m *MediaTypeChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes between two MediaType instances.
func (m *MediaTypeChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes made between two MediaType instances.
func (m *MediaTypeChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareMediaTypes compares a left and a right MediaType object for any changes. If found, a pointer to a
// MediaTypeChanges instance is returned; otherwise nothing is returned.
func CompareMediaTypes(l, r *v3.MediaType) *MediaTypeChanges { _ = "STUB: not implemented"; return nil }

// Example

// schema

// examples - use nil-aware version so added/removed examples appear in the map for tree rendering

// encoding

// itemSchema

// itemEncoding
