// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import v3 "github.com/pb33f/libopenapi/datamodel/low/v3"

// ServerChanges represents changes found between two OpenAPI Server Objects
type ServerChanges struct {
	*PropertyChanges
	Server                *v3.Server
	ServerVariableChanges map[string]*ServerVariableChanges `json:"serverVariables,omitempty" yaml:"serverVariables,omitempty"`
	ExtensionChanges      *ExtensionChanges                 `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between SecurityRequirement objects
func (s *ServerChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns total changes found between two OpenAPI Server Objects
func (s *ServerChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes found between two OpenAPI Server objects.
func (s *ServerChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareServers compares two OpenAPI Server objects for any changes. If anything is found, returns a pointer
// to a ServerChanges instance, or returns nil if nothing is found.
func CompareServers(l, r *v3.Server) *ServerChanges { _ = "STUB: not implemented"; return nil }
