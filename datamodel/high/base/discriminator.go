// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	lowBase "github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/orderedmap"
)

// Discriminator is only used by OpenAPI 3+ documents, it represents a polymorphic discriminator used for schemas
//
// When request bodies or response payloads may be one of a number of different schemas, a discriminator object can be
// used to aid in serialization, deserialization, and validation. The discriminator is a specific object in a schema
// which is used to inform the consumer of the document of an alternative schema based on the value associated with it.
//
// When using the discriminator, inline schemas will not be considered.
//
//	v3 - https://spec.openapis.org/oas/v3.1.0#discriminator-object
type Discriminator struct {
	PropertyName   string                          `json:"propertyName,omitempty" yaml:"propertyName,omitempty"`
	Mapping        *orderedmap.Map[string, string] `json:"mapping,omitempty" yaml:"mapping,omitempty"`
	DefaultMapping string                          `json:"defaultMapping,omitempty" yaml:"defaultMapping,omitempty"` // OpenAPI 3.2+ defaultMapping for fallback schema
	low            *lowBase.Discriminator
}

// NewDiscriminator will create a new high-level Discriminator from a low-level one.
func NewDiscriminator(disc *lowBase.Discriminator) *Discriminator {
	_ = "STUB: not implemented"
	return nil
}

// GoLow returns the low-level Discriminator used to build the high-level one.
func (d *Discriminator) GoLow() *lowBase.Discriminator {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Discriminator instance that was used to create the high-level one, with no type
	return nil
}

func (d *Discriminator) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Discriminator object as a byte slice.
	return *new(any)
}

func (d *Discriminator) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the Discriminator object.
		nil
}

func (d *Discriminator) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
