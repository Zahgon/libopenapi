// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// SecurityRequirementChanges represents changes found between two SecurityRequirement Objects.
type SecurityRequirementChanges struct {
	*PropertyChanges
}

// GetAllChanges returns a slice of all changes made between SecurityRequirement objects
func (s *SecurityRequirementChanges) GetAllChanges() []*Change {
	_ = "STUB: not implemented"
	return nil
}

// TotalChanges returns the total number of changes between two SecurityRequirement Objects.
func (s *SecurityRequirementChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes between two SecurityRequirement Objects.
func (s *SecurityRequirementChanges) TotalBreakingChanges() int {
	_ = "STUB: not implemented"
	return 0
}

// CompareSecurityRequirement compares left and right SecurityRequirement objects for changes. If anything
// is found, then a pointer to SecurityRequirementChanges is returned, otherwise nil.
func CompareSecurityRequirement(l, r *base.SecurityRequirement) *SecurityRequirementChanges {
	_ = "STUB: not implemented"
	return nil
}

func removedSecurityRequirement(vn *yaml.Node, schemeName, scopeName string, changes *[]*Change) {
	_ = "STUB: not implemented"
	return
}

// entire scheme was removed, use scheme name as value

// Don't use the node for entire scheme removal, as it may be an empty array []

// scope was removed

func addedSecurityRequirement(vn *yaml.Node, schemeName, scopeName string, changes *[]*Change) {
	_ = "STUB: not implemented"
	return
}

// entire scheme was added, use scheme name as value

// Don't use the node for entire scheme addition, as it may be an empty array []

// scope was added

// tricky to do this correctly, this is my solution.
func checkSecurityRequirement(lSec, rSec *orderedmap.Map[low.KeyReference[string], low.ValueReference[[]low.ValueReference[string]]],
	changes *[]*Change,
) {
	_ = "STUB: not implemented"
	return
}

// check if actual values match up

// Skip empty scope values (from malformed YAML)

// Trim to actual size

// Skip empty scope values (from malformed YAML)

// Trim to actual size
