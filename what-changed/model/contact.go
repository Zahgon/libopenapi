// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low/base"
)

// ContactChanges Represent changes to a Contact object that is a child of Info, part of an OpenAPI document.
type ContactChanges struct {
	*PropertyChanges
}

// GetAllChanges returns a slice of all changes made between Callback objects
func (c *ContactChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges represents the total number of changes that have occurred to a Contact object
func (c *ContactChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes in Contact objects.
func (c *ContactChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareContact will check a left (original) and right (new) Contact object for any changes. If there
// were any, a pointer to a ContactChanges object is returned, otherwise if nothing changed - the function
// returns nil.
func CompareContact(l, r *base.Contact) *ContactChanges { _ = "STUB: not implemented"; return nil }
