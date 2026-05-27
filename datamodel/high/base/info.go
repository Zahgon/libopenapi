// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	low "github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Info represents a high-level Info object as defined by both OpenAPI 2 and OpenAPI 3.
//
// The object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented
// in editing or documentation generation tools for convenience.
//
//	v2 - https://swagger.io/specification/v2/#infoObject
//	v3 - https://spec.openapis.org/oas/v3.1.0#info-object
type Info struct {
	Summary        string                              `json:"summary,omitempty" yaml:"summary,omitempty"`
	Title          string                              `json:"title,omitempty" yaml:"title,omitempty"`
	Description    string                              `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string                              `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact                            `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License                            `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string                              `json:"version,omitempty" yaml:"version,omitempty"`
	Extensions     *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low            *low.Info
}

// NewInfo will create a new high-level Info instance from a low-level one.
func NewInfo(info *low.Info) *Info { _ = "STUB: not implemented"; return nil }

// GoLow will return the low-level Info instance that was used to create the high-level one.
func (i *Info) GoLow() *low.Info {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Info instance that was used to create the high-level one, with no type
	return nil
}

func (i *Info) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Info object as a byte slice.
	return *new(any)
}

func (i *Info) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the Info object.
		nil
}

func (i *Info) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
