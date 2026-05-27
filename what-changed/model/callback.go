// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
)

// CallbackChanges represents all changes made between two Callback OpenAPI objects.
type CallbackChanges struct {
	*PropertyChanges
	ExpressionChanges map[string]*PathItemChanges `json:"expressions,omitempty" yaml:"expressions,omitempty"`
	ExtensionChanges  *ExtensionChanges           `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// TotalChanges returns a total count of all changes made between Callback objects
func (c *CallbackChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// GetAllChanges returns a slice of all changes made between Callback objects
func (c *CallbackChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalBreakingChanges returns a total count of all changes made between Callback objects
func (c *CallbackChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareCallback will compare two Callback objects and return a pointer to CallbackChanges with all the things
// that have changed between them. Handles nil inputs for added/removed callback scenarios.
func CompareCallback(l, r *v3.Callback) *CallbackChanges { _ = "STUB: not implemented"; return nil }

// whole callback added - use operation.callbacks breaking rules

// whole callback removed - use operation.callbacks breaking rules

// Both exist - compare them

// check left path item hashes

// run comparison.

// check right path item hashes
