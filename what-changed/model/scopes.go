// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	v2 "github.com/pb33f/libopenapi/datamodel/low/v2"
)

// ScopesChanges represents changes between two Swagger Scopes Objects
type ScopesChanges struct {
	*PropertyChanges
	ExtensionChanges *ExtensionChanges `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Scopes objects
func (s *ScopesChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total changes found between two Swagger Scopes objects.
func (s *ScopesChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes between two Swagger Scopes objects.
func (s *ScopesChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareScopes compares a left and right Swagger Scopes objects for changes. If anything is found, returns
// a pointer to ScopesChanges, or returns nil if nothing is found.
func CompareScopes(l, r *v2.Scopes) *ScopesChanges { _ = "STUB: not implemented"; return nil }
