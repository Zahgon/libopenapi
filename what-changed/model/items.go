// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	v2 "github.com/pb33f/libopenapi/datamodel/low/v2"
)

// ItemsChanges represent changes found between a left (original) and right (modified) object. Items is only
// used by Swagger documents.
type ItemsChanges struct {
	*PropertyChanges
	ItemsChanges *ItemsChanges `json:"items,omitempty" yaml:"items,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Items objects
func (i *ItemsChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes found between two Items objects
// This is a recursive function because Items can contain Items. Be careful!
func (i *ItemsChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes found between two Swagger Items objects
// This is a recursive method, Items are recursive, be careful!
func (i *ItemsChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareItems compares two sets of Swagger Item objects. If there are any changes found then a pointer to
// ItemsChanges will be returned, otherwise nil is returned.
//
// It is worth nothing that Items can contain Items. This means recursion is possible and has the potential for
// runaway code if not using the resolver's circular reference checking.
func CompareItems(l, r *v2.Items) *ItemsChanges { _ = "STUB: not implemented"; return nil }

// header is identical to items, except for a description.

// inline, check hashes, if they don't match, compare.

// compare.

// added items

// removed items
