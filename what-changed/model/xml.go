// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low/base"
)

// XMLChanges represents changes made to the XML object of an OpenAPI document.
type XMLChanges struct {
	*PropertyChanges
	ExtensionChanges *ExtensionChanges `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between XML objects
func (x *XMLChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns a count of everything that was changed within an XML object.
func (x *XMLChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the number of breaking changes made by the XML object.
func (x *XMLChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareXML will compare a left (original) and a right (new) XML instance, and check for
// any changes between them. If changes are found, the function returns a pointer to XMLChanges,
// otherwise, if nothing changed - it will return nil
func CompareXML(l, r *base.XML) *XMLChanges { _ = "STUB: not implemented"; return nil }

// check extensions
