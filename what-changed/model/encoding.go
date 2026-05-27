// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
)

// EncodingChanges represent all the changes made to an Encoding object
type EncodingChanges struct {
	*PropertyChanges
	HeaderChanges map[string]*HeaderChanges `json:"headers,omitempty" yaml:"headers,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Encoding objects
func (e *EncodingChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes made between two Encoding objects
func (e *EncodingChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the number of changes made between two Encoding objects that were breaking.
func (e *EncodingChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareEncoding returns a pointer to *EncodingChanges that contain all changes made between a left and right
// set of Encoding objects.
func CompareEncoding(l, r *v3.Encoding) *EncodingChanges { _ = "STUB: not implemented"; return nil }

// check everything.

// headers
