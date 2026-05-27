// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low/base"
)

// ExternalDocChanges represents changes made to any ExternalDoc object from an OpenAPI document.
type ExternalDocChanges struct {
	*PropertyChanges
	ExtensionChanges *ExtensionChanges `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Example objects
func (e *ExternalDocChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns a count of everything that changed
func (e *ExternalDocChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes in ExternalDoc objects.
func (e *ExternalDocChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareExternalDocs will compare a left (original) and a right (new) slice of ValueReference
// nodes for any changes between them. If there are changes, then a pointer to ExternalDocChanges
// is returned, otherwise if nothing changed - then nil is returned.
func CompareExternalDocs(l, r *base.ExternalDoc) *ExternalDocChanges {
	_ = "STUB: not implemented"
	return nil
}

// check extensions
