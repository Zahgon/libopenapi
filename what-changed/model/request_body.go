// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import v3 "github.com/pb33f/libopenapi/datamodel/low/v3"

// RequestBodyChanges represents changes made between two OpenAPI RequestBody Objects
type RequestBodyChanges struct {
	*PropertyChanges
	ContentChanges   map[string]*MediaTypeChanges `json:"content,omitempty" yaml:"content,omitempty"`
	ExtensionChanges *ExtensionChanges            `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between RequestBody objects
func (rb *RequestBodyChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes found between two OpenAPI RequestBody objects
func (rb *RequestBodyChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes found between OpenAPI RequestBody objects
func (rb *RequestBodyChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareRequestBodies compares a left and right OpenAPI RequestBody object for changes. If found returns a pointer
// to a RequestBodyChanges instance. Returns nil if nothing was found.
func CompareRequestBodies(l, r *v3.RequestBody) *RequestBodyChanges {
	_ = "STUB: not implemented"
	return nil
}
