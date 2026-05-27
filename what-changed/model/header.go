// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	v2 "github.com/pb33f/libopenapi/datamodel/low/v2"
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
)

// HeaderChanges represents changes made between two Header objects. Supports both Swagger and OpenAPI header
// objects, V2 only property Items is broken out into its own.
type HeaderChanges struct {
	*PropertyChanges
	SchemaChanges    *SchemaChanges               `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	ExamplesChanges  map[string]*ExampleChanges   `json:"examples,omitempty" yaml:"examples,omitempty"`
	ContentChanges   map[string]*MediaTypeChanges `json:"content,omitempty" yaml:"content,omitempty"`
	ExtensionChanges *ExtensionChanges            `json:"extensions,omitempty" yaml:"extensions,omitempty"`

	// Items only supported by Swagger (V2)
	ItemsChanges *ItemsChanges `json:"items,omitempty" yaml:"items,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Header objects
func (h *HeaderChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes made between two Header objects.
func (h *HeaderChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes made between two Header instances.
func (h *HeaderChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// shared header properties
func addOpenAPIHeaderProperties(left, right low.OpenAPIHeader, changes *[]*Change) []*PropertyCheck {
	_ = "STUB: not implemented"
	return nil

	// style
}

// allow reserved

// allow empty value

// explode

// example

// deprecated

// required

// swagger only properties
func addSwaggerHeaderProperties(left, right low.SwaggerHeader, changes *[]*Change) []*PropertyCheck {
	_ = "STUB: not implemented"
	return nil

	// type
}

// format

// collection format

// maximum

// minimum

// exclusive maximum

// exclusive minimum

// max length

// min length

// pattern

// max items

// min items

// unique items

// multiple of

// common header properties
func addCommonHeaderProperties(left, right low.HasDescription, changes *[]*Change) []*PropertyCheck {
	_ = "STUB: not implemented"
	return nil

	// description
}

// CompareHeadersV2 is a Swagger compatible, typed signature used for other generic functions. It simply
// wraps CompareHeaders and provides nothing other that a typed interface.
func CompareHeadersV2(l, r *v2.Header) *HeaderChanges { _ = "STUB: not implemented"; return nil }

// CompareHeadersV3 is an OpenAPI 3+ compatible, typed signature used for other generic functions. It simply
// wraps CompareHeaders and provides nothing other that a typed interface.
func CompareHeadersV3(l, r *v3.Header) *HeaderChanges { _ = "STUB: not implemented"; return nil }

// CompareHeaders will compare left and right Header objects (any version of Swagger or OpenAPI) and return
// a pointer to HeaderChanges with anything that has changed, or nil if nothing changed.
func CompareHeaders(l, r any) *HeaderChanges { _ = "STUB: not implemented"; return nil }

// handle swagger.

// perform hash check to avoid further processing

// enum

// items

// handle OpenAPI

// perform hash check to avoid further processing

// header

// examples

// content
