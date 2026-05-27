// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	v2 "github.com/pb33f/libopenapi/datamodel/low/v2"
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
)

// ResponseChanges represents changes found between two Swagger or OpenAPI Response objects.
type ResponseChanges struct {
	*PropertyChanges
	ExtensionChanges *ExtensionChanges         `json:"extensions,omitempty" yaml:"extensions,omitempty"`
	HeadersChanges   map[string]*HeaderChanges `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Swagger Response Properties.
	SchemaChanges   *SchemaChanges   `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	ExamplesChanges *ExamplesChanges `json:"examples,omitempty" yaml:"examples,omitempty"`

	// OpenAPI Response Properties.
	ContentChanges map[string]*MediaTypeChanges `json:"content,omitempty" yaml:"content,omitempty"`
	LinkChanges    map[string]*LinkChanges      `json:"links,omitempty" yaml:"links,omitempty"`
}

// GetAllChanges returns a slice of all changes made between RequestBody objects
func (r *ResponseChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes found between two Swagger or OpenAPI Response Objects
func (r *ResponseChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes found between two swagger or OpenAPI
// Response Objects
func (r *ResponseChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareResponseV2 is a Swagger type safe proxy for CompareResponse
func CompareResponseV2(l, r *v2.Response) *ResponseChanges { _ = "STUB: not implemented"; return nil }

// CompareResponseV3 is an OpenAPI type safe proxy for CompareResponse
func CompareResponseV3(l, r *v3.Response) *ResponseChanges { _ = "STUB: not implemented"; return nil }

// CompareResponse compares a left and right Swagger or OpenAPI Response object. If anything is found
// a pointer to a ResponseChanges is returned, otherwise it returns nil.
func CompareResponse(l, r any) *ResponseChanges { _ = "STUB: not implemented"; return nil }

// perform hash check to avoid further processing

// description

// perform hash check to avoid further processing

// summary (OpenAPI 3.2+)

// description
