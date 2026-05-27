// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low/base"
)

// LicenseChanges represent changes to a License object that is a child of Info object. Part of an OpenAPI document
type LicenseChanges struct {
	*PropertyChanges
	ExtensionChanges *ExtensionChanges `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between License objects
func (l *LicenseChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges represents the total number of changes made to a License instance.
func (l *LicenseChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes in License objects.
func (l *LicenseChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareLicense will check a left (original) and right (new) License object for any changes. If there
// were any, a pointer to a LicenseChanges object is returned, otherwise if nothing changed - the function
// returns nil.
func CompareLicense(l, r *base.License) *LicenseChanges { _ = "STUB: not implemented"; return nil }
