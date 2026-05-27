// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low/base"
)

// ExampleChanges represent changes to an Example object, part of an OpenAPI specification.
type ExampleChanges struct {
	*PropertyChanges
	ExtensionChanges *ExtensionChanges `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Example objects
func (e *ExampleChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes made to Example
func (e *ExampleChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes made to Example
func (e *ExampleChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareExamples returns a pointer to ExampleChanges that contains all changes made between
// left and right Example instances. If l is nil, the example was added. If r is nil, it was removed.
func CompareExamples(l, r *base.Example) *ExampleChanges { _ = "STUB: not implemented"; return nil }

// Example was added - use RootNode for proper line/column location

// Example was removed - use RootNode for proper line/column location

// Value

// if there is no value (value is another map or something else), render the node into yaml and hash it.
// https://github.com/pb33f/libopenapi/issues/61

// if there is no value (value is another map or something else), render the node into yaml and hash it.
// https://github.com/pb33f/libopenapi/issues/61

// ExternalValue

// DataValue (OpenAPI 3.2+)

// SerializedValue (OpenAPI 3.2+)

// check properties

// check extensions
