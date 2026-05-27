// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import v3 "github.com/pb33f/libopenapi/datamodel/low/v3"

// ServerVariableChanges represents changes found between two OpenAPI ServerVariable Objects
type ServerVariableChanges struct {
	*PropertyChanges
}

// GetAllChanges returns a slice of all changes made between SecurityRequirement objects
func (s *ServerVariableChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// CompareServerVariables compares a left and right OpenAPI ServerVariable object for changes.
// If anything is found, returns a pointer to a ServerVariableChanges instance, otherwise returns nil.
func CompareServerVariables(l, r *v3.ServerVariable) *ServerVariableChanges {
	_ = "STUB: not implemented"
	return nil
}
