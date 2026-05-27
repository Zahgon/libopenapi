// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	low "github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// License is a high-level representation of a License object as defined by OpenAPI 2 and OpenAPI 3
//
//	v2 - https://swagger.io/specification/v2/#licenseObject
//	v3 - https://spec.openapis.org/oas/v3.1.0#license-object
type License struct {
	Name       string                              `json:"name,omitempty" yaml:"name,omitempty"`
	URL        string                              `json:"url,omitempty" yaml:"url,omitempty"`
	Identifier string                              `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	Extensions *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low        *low.License
}

// NewLicense will create a new high-level License instance from a low-level one.
func NewLicense(license *low.License) *License { _ = "STUB: not implemented"; return nil }

// GoLow will return the low-level License used to create the high-level one.
func (l *License) GoLow() *low.License {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level License instance that was used to create the high-level one, with no type
	return nil
}

func (l *License) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the License object as a byte slice.
	return *new(any)
}

func (l *License) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the License object.
		nil
}

func (l *License) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
