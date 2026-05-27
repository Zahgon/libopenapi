// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	low "github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// ExternalDoc represents a high-level External Documentation object as defined by OpenAPI 2 and 3
//
// Allows referencing an external resource for extended documentation.
//
//	v2 - https://swagger.io/specification/v2/#externalDocumentationObject
//	v3 - https://spec.openapis.org/oas/v3.1.0#external-documentation-object
type ExternalDoc struct {
	Description string                              `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string                              `json:"url,omitempty" yaml:"url,omitempty"`
	Extensions  *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low         *low.ExternalDoc
}

// NewExternalDoc will create a new high-level External Documentation object from a low-level one.
func NewExternalDoc(extDoc *low.ExternalDoc) *ExternalDoc { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level ExternalDoc instance used to create the high-level one.
func (e *ExternalDoc) GoLow() *low.ExternalDoc {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level ExternalDoc instance that was used to create the high-level one, with no type
	return nil
}

func (e *ExternalDoc) GoLowUntyped() any { _ = "STUB: not implemented"; return *new(any) }

func (e *ExternalDoc) GetExtensions() *orderedmap.Map[string, *yaml.Node] {
	_ = "STUB: not implemented"
	return nil

	// Render will return a YAML representation of the ExternalDoc object as a byte slice.
}

func (e *ExternalDoc) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the ExternalDoc object.
		nil
}

func (e *ExternalDoc) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
