// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	v2 "github.com/pb33f/libopenapi/datamodel/low/v2"
)

// ExamplesChanges represents changes made between Swagger Examples objects (Not OpenAPI 3).
type ExamplesChanges struct {
	*PropertyChanges
}

// GetAllChanges returns a slice of all changes made between Examples objects
func (a *ExamplesChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges represents the total number of changes made between Example instances.
func (a *ExamplesChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges will always return 0. Examples cannot break a contract.
func (a *ExamplesChanges) TotalBreakingChanges() int {
	_ = "STUB: not implemented"
	// not supported.
	return 0
}

// CompareExamplesV2 compares two Swagger Examples objects, returning a pointer to
// ExamplesChanges if anything was found.
func CompareExamplesV2(l, r *v2.Examples) *ExamplesChanges { _ = "STUB: not implemented"; return nil }

// check left example hashes

// check right example hashes
