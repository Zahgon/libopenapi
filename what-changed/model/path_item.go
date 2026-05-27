// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	v2 "github.com/pb33f/libopenapi/datamodel/low/v2"
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
)

// PathItemChanges represents changes found between to Swagger or OpenAPI PathItem object.
type PathItemChanges struct {
	*PropertyChanges
	GetChanges                 *OperationChanges            `json:"get,omitempty" yaml:"get,omitempty"`
	PutChanges                 *OperationChanges            `json:"put,omitempty" yaml:"put,omitempty"`
	PostChanges                *OperationChanges            `json:"post,omitempty" yaml:"post,omitempty"`
	DeleteChanges              *OperationChanges            `json:"delete,omitempty" yaml:"delete,omitempty"`
	OptionsChanges             *OperationChanges            `json:"options,omitempty" yaml:"options,omitempty"`
	HeadChanges                *OperationChanges            `json:"head,omitempty" yaml:"head,omitempty"`
	PatchChanges               *OperationChanges            `json:"patch,omitempty" yaml:"patch,omitempty"`
	TraceChanges               *OperationChanges            `json:"trace,omitempty" yaml:"trace,omitempty"`
	QueryChanges               *OperationChanges            `json:"query,omitempty" yaml:"query,omitempty"`
	AdditionalOperationChanges map[string]*OperationChanges `json:"additionalOperations,omitempty" yaml:"additionalOperations,omitempty"` // OpenAPI 3.2+ additional operations
	ServerChanges              []*ServerChanges             `json:"servers,omitempty" yaml:"servers,omitempty"`
	ParameterChanges           []*ParameterChanges          `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	ExtensionChanges           *ExtensionChanges            `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between PathItem objects
func (p *PathItemChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes found between two Swagger or OpenAPI PathItems
func (p *PathItemChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes found between two Swagger or OpenAPI PathItems
func (p *PathItemChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

type opCheck struct {
	label   string
	changes *OperationChanges
}

// ComparePathItemsV3 is an OpenAPI typesafe proxy method for ComparePathItems
func ComparePathItemsV3(l, r *v3.PathItem) *PathItemChanges { _ = "STUB: not implemented"; return nil }

// ComparePathItems compare a left and right Swagger or OpenAPI PathItem object for changes. If found, returns
// a pointer to PathItemChanges, or returns nil if nothing is found.
func ComparePathItems(l, r any) *PathItemChanges { _ = "STUB: not implemented"; return nil }

// Swagger

// perform hash check to avoid further processing

// OpenAPI

// perform hash check to avoid further processing

// description

// summary

func compareSwaggerPathItem(lPath, rPath *v2.PathItem, changes *[]*Change, pc *PathItemChanges) []*PropertyCheck {
	_ = "STUB: not implemented"
	return nil
}

// get

// put

// post

// delete

// options

// head

// patch

// parameters

// Check configurable breaking rules first

// If config says not breaking, fall back to semantic check (required params are breaking)

// collect up operations changes.

func extractV2ParametersIntoInterface(l, r []low.ValueReference[*v2.Parameter]) ([]low.ValueReference[low.SharedParameters],
	[]low.ValueReference[low.SharedParameters],
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractV3ParametersIntoInterface(l, r []low.ValueReference[*v3.Parameter]) ([]low.ValueReference[low.SharedParameters],
	[]low.ValueReference[low.SharedParameters],
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkParameters(lParams, rParams []low.ValueReference[low.SharedParameters], changes *[]*Change, pc *PathItemChanges) {
	_ = "STUB: not implemented"
	return
}

// Keep the reference wrapper

// Keep the reference wrapper

// Preserve reference information if this parameter is a $ref

func compareOpenAPIPathItem(lPath, rPath *v3.PathItem, changes *[]*Change, pc *PathItemChanges) {
	_ = "STUB: not implemented"
	// var props []*PropertyCheck
	return
}

// get

// put

// post

// delete

// options

// head

// patch

// trace

// query

// additionalOperations (OpenAPI 3.2+)

// check right keys for match

// compare the two operations

// not found, was removed

// check for added operations

// check left keys for match

// not found, was added

// servers

// parameters

// Check configurable breaking rules first

// If config says not breaking, fall back to semantic check (required params are breaking)

// collect up operations changes.

func checkOperation(l, r any, done chan opCheck, method string) { _ = "STUB: not implemented"; return }
