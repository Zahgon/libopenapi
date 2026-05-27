// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

// ResponsesChanges represents changes made between two Swagger or OpenAPI Responses objects.
type ResponsesChanges struct {
	*PropertyChanges
	ResponseChanges  map[string]*ResponseChanges `json:"response,omitempty" yaml:"response,omitempty"`
	DefaultChanges   *ResponseChanges            `json:"default,omitempty" yaml:"default,omitempty"`
	ExtensionChanges *ExtensionChanges           `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Responses objects
func (r *ResponsesChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes found between two Swagger or OpenAPI Responses objects
func (r *ResponsesChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of changes found between two Swagger or OpenAPI
// Responses Objects
func (r *ResponsesChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareResponses compares a left and right Swagger or OpenAPI Responses object for any changes. If found
// returns a pointer to ResponsesChanges, or returns nil.
func CompareResponses(l, r any) *ResponsesChanges { _ = "STUB: not implemented"; return nil }

// swagger

// perform hash check to avoid further processing

// openapi

// perform hash check to avoid further processing
