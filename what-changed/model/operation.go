// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
)

// OperationChanges represent changes made between two Swagger or OpenAPI Operation objects.
type OperationChanges struct {
	*PropertyChanges
	ExternalDocChanges         *ExternalDocChanges           `json:"externalDoc,omitempty" yaml:"externalDoc,omitempty"`
	ParameterChanges           []*ParameterChanges           `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	ResponsesChanges           *ResponsesChanges             `json:"responses,omitempty" yaml:"responses,omitempty"`
	SecurityRequirementChanges []*SecurityRequirementChanges `json:"securityRequirements,omitempty" yaml:"securityRequirements,omitempty"`

	// OpenAPI 3+ only changes
	RequestBodyChanges *RequestBodyChanges         `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	ServerChanges      []*ServerChanges            `json:"servers,omitempty" yaml:"servers,omitempty"`
	ExtensionChanges   *ExtensionChanges           `json:"extensions,omitempty" yaml:"extensions,omitempty"`
	CallbackChanges    map[string]*CallbackChanges `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Operation objects
func (o *OperationChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes made between two Swagger or OpenAPI Operation objects.
func (o *OperationChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes made between two Swagger
// or OpenAPI Operation objects.
func (o *OperationChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// check for properties shared between operations objects.
func addSharedOperationProperties(left, right low.SharedOperations, changes *[]*Change) []*PropertyCheck {
	_ = "STUB: not implemented"
	return nil

	// tags
}

// summary

// description

// deprecated

// operation id

// check shared objects
func compareSharedOperationObjects(l, r low.SharedOperations, changes *[]*Change, opChanges *OperationChanges) {
	_ = "STUB: not implemented"
	// external docs
	return
}

// responses

// CompareOperations compares a left and right Swagger or OpenAPI Operation object. If changes are found, returns
// a pointer to an OperationChanges instance, or nil if nothing is found.
func CompareOperations(l, r any) *OperationChanges { _ = "STUB: not implemented"; return nil }

// Swagger

// perform hash check to avoid further processing

// parameters

// Keep the reference wrapper

// Keep the reference wrapper

// Preserve reference information if this parameter is a $ref

// security

// produces

// consumes

// schemes

// OpenAPI

// perform hash check to avoid further processing

// parameters

// Keep the reference wrapper

// Keep the reference wrapper

// Preserve reference information if this parameter is a $ref

// Check configurable breaking rules first

// If config doesn't say breaking, fall back to semantic check (required parameter)

// Check configurable breaking rules first

// If config doesn't say breaking, fall back to semantic check (required parameter)

// security

// request body

// callbacks - use CheckMapForChangesWithNilSupport to properly populate CallbackChanges
// for added/removed callbacks, enabling proper tree hierarchy rendering

// servers

// check servers property
// component and property are used for breaking rules lookup (e.g., CompOperation/PropServers or CompServers/"")
func checkServers(lServers, rServers low.NodeReference[[]low.ValueReference[*v3.Server]], component, property string) []*ServerChanges {
	_ = "STUB: not implemented"
	return nil
}

// check security property.
func checkSecurity(lSecurity, rSecurity low.NodeReference[[]low.ValueReference[*base.SecurityRequirement]],
	changes *[]*Change, oc any,
) {
	_ = "STUB: not implemented"
	return
}

// Determine breaking rules based on type (zero allocations using type switch)

// Whole security requirement was removed - create SecurityRequirementChanges
// so it appears under "Security Requirements" section

// Whole security requirement was added - create SecurityRequirementChanges
// so it appears under "Security Requirements" section

// Assign to correct type using type switch (zero allocations)
