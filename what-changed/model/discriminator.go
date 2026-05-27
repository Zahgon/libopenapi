// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low/base"
)

// DiscriminatorChanges represents changes made to a Discriminator OpenAPI object
type DiscriminatorChanges struct {
	*PropertyChanges
	MappingChanges []*Change `json:"mappings,omitempty" yaml:"mappings,omitempty"`
}

// TotalChanges returns a count of everything changed within the Discriminator object
func (d *DiscriminatorChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// GetAllChanges returns a slice of all changes made between Callback objects
func (c *DiscriminatorChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalBreakingChanges returns the number of breaking changes made by the Discriminator
func (d *DiscriminatorChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareDiscriminator will check a left (original) and right (new) Discriminator object for changes
// and will return a pointer to DiscriminatorChanges
func CompareDiscriminator(l, r *base.Discriminator) *DiscriminatorChanges {
	_ = "STUB: not implemented"
	return nil
}

// flatten maps

// check for removals, modifications and moves

// if the existing tag exists, let's check it.
