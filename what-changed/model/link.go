// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
)

// LinkChanges represent changes made between two OpenAPI Link Objects.
type LinkChanges struct {
	*PropertyChanges
	ExtensionChanges *ExtensionChanges `json:"extensions,omitempty" yaml:"extensions,omitempty"`
	ServerChanges    *ServerChanges    `json:"server,omitempty" yaml:"server,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Link objects
func (l *LinkChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total changes made between OpenAPI Link objects
func (l *LinkChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the number of breaking changes made between two OpenAPI Link Objects
func (l *LinkChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareLinks checks a left and right OpenAPI Link for any changes. If they are found, returns a pointer to
// LinkChanges, and returns nil if nothing is found.
func CompareLinks(l, r *v3.Link) *LinkChanges { _ = "STUB: not implemented"; return nil }

// server

// parameters
