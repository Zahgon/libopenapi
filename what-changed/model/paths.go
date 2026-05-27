// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

// PathsChanges represents changes found between two Swagger or OpenAPI Paths Objects.
type PathsChanges struct {
	*PropertyChanges
	PathItemsChanges map[string]*PathItemChanges `json:"pathItems,omitempty" yaml:"pathItems,omitempty"`
	ExtensionChanges *ExtensionChanges           `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Paths objects
func (p *PathsChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes between two Swagger or OpenAPI Paths Objects
func (p *PathsChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns tht total number of changes found between two Swagger or OpenAPI Path Objects
func (p *PathsChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// ComparePaths compares a left and right Swagger or OpenAPI Paths Object for changes. If found, returns a pointer
// to a PathsChanges instance. Returns nil if nothing is found.
func ComparePaths(l, r any) *PathsChanges { _ = "STUB: not implemented"; return nil }

// Swagger

// perform hash check to avoid further processing

// run every comparison in a thread.

// wait for the things to be done.

// OpenAPI

// perform hash check to avoid further processing

// run every comparison in a thread.

// wait for the things to be done.
