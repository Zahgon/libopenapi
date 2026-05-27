// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/datamodel/low/v3"
)

// TagChanges represents changes made to the Tags object of an OpenAPI document.
type TagChanges struct {
	*PropertyChanges
	ExternalDocs     *ExternalDocChanges `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	ExtensionChanges *ExtensionChanges   `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Tag objects
func (t *TagChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns a count of everything that changed within tags.
func (t *TagChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the number of breaking changes made by Tags
func (t *TagChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareTags will compare a left (original) and a right (new) slice of ValueReference nodes for
// any changes between them. If there are changes, a pointer to TagChanges is returned, if not then
// nil is returned instead.
func CompareTags(l, r []low.ValueReference[*base.Tag]) []*TagChanges {
	_ = "STUB: not implemented"
	return nil
}

// look at the original and then look through the new.

// var changes []*Change

// check for removals, modifications and moves

// if the existing tag exists, let's check it.

// check properties

// compare external docs

// check extensions
