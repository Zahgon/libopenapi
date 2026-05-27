// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// ExtensionChanges represents any changes to custom extensions defined for an OpenAPI object.
type ExtensionChanges struct {
	*PropertyChanges
}

// GetAllChanges returns a slice of all changes made between Extension objects
func (e *ExtensionChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of object extensions that were made.
func (e *ExtensionChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes in Extension objects.
func (e *ExtensionChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareExtensions will compare a left and right map of Tag/ValueReference models for any changes to
// anything. This function does not try and cast the value of an extension to perform checks, it
// will perform a basic value check.
//
// A current limitation relates to extensions being objects and a property of the object changes,
// there is currently no support for knowing anything changed - so it is ignored.
func CompareExtensions(l, r *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]) *ExtensionChanges {
	_ = "STUB: not implemented"
	// look at the original and then look through the new.
	return nil
}

// check properties with encoding for extensions

// CheckExtensions is a helper method to un-pack a left and right model that contains extensions. Once unpacked
// the extensions are compared and returns a pointer to ExtensionChanges. If nothing changed, nil is returned.
func CheckExtensions[T low.HasExtensions[T]](l, r T) *ExtensionChanges {
	_ = "STUB: not implemented"
	return nil
}
