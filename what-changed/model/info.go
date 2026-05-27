// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low/base"
)

// InfoChanges represents the number of changes to an Info object. Part of an OpenAPI document
type InfoChanges struct {
	*PropertyChanges
	ContactChanges   *ContactChanges   `json:"contact,omitempty" yaml:"contact,omitempty"`
	LicenseChanges   *LicenseChanges   `json:"license,omitempty" yaml:"license,omitempty"`
	ExtensionChanges *ExtensionChanges `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Info objects
func (i *InfoChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges represents the total number of changes made to an Info object.
func (i *InfoChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes in Info objects.
func (i *InfoChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareInfo will compare a left (original) and a right (new) Info object. Any changes
// will be returned in a pointer to InfoChanges, otherwise if nothing is found, then nil is
// returned instead.
func CompareInfo(l, r *base.Info) *InfoChanges { _ = "STUB: not implemented"; return nil }

// check properties

// compare contact.

// compare license.

// check extensions.
