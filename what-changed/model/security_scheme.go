// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	v2 "github.com/pb33f/libopenapi/datamodel/low/v2"
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
)

// SecuritySchemeChanges represents changes made between Swagger or OpenAPI SecurityScheme Objects.
type SecuritySchemeChanges struct {
	*PropertyChanges
	ExtensionChanges *ExtensionChanges `json:"extensions,omitempty" yaml:"extensions,omitempty"`

	// OpenAPI Version
	OAuthFlowChanges *OAuthFlowsChanges `json:"oAuthFlow,omitempty" yaml:"oAuthFlow,omitempty"`

	// Swagger Version
	ScopesChanges *ScopesChanges `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}

// GetAllChanges returns a slice of all changes made between SecurityRequirement objects
func (ss *SecuritySchemeChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges represents total changes found between two Swagger or OpenAPI SecurityScheme instances.
func (ss *SecuritySchemeChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns total number of breaking changes between two SecurityScheme Objects.
func (ss *SecuritySchemeChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareSecuritySchemesV2 is a Swagger type safe proxy for CompareSecuritySchemes
func CompareSecuritySchemesV2(l, r *v2.SecurityScheme) *SecuritySchemeChanges {
	_ = "STUB: not implemented"
	return nil
}

// CompareSecuritySchemesV3 is an OpenAPI type safe proxt for CompareSecuritySchemes
func CompareSecuritySchemesV3(l, r *v3.SecurityScheme) *SecuritySchemeChanges {
	_ = "STUB: not implemented"
	return nil
}

// CompareSecuritySchemes compares left and right Swagger or OpenAPI Security Scheme objects for changes.
// If anything is found, returns a pointer to *SecuritySchemeChanges or nil if nothing is found.
func CompareSecuritySchemes(l, r any) *SecuritySchemeChanges { _ = "STUB: not implemented"; return nil }

// OpenAPI 3.2+ fields
