// Copyright 2022-2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"github.com/pb33f/libopenapi/datamodel/high/base"
	lowv3 "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Operation is a high-level representation of an OpenAPI 3+ Operation object, backed by a low-level one.
//
// An Operation is perhaps the most important object of the entire specification. Everything of value
// happens here. The entire being for existence of this library and the specification, is this Operation.
//   - https://spec.openapis.org/oas/v3.1.0#operation-object
type Operation struct {
	Tags         []string                            `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                              `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                              `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *base.ExternalDoc                   `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationId  string                              `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   []*Parameter                        `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *RequestBody                        `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    *Responses                          `json:"responses,omitempty" yaml:"responses,omitempty"`
	Callbacks    *orderedmap.Map[string, *Callback]  `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   *bool                               `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     []*base.SecurityRequirement         `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      []*Server                           `json:"servers,omitempty" yaml:"servers,omitempty"`
	Extensions   *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low          *lowv3.Operation
}

// NewOperation will create a new Operation instance from a low-level one.
func NewOperation(operation *lowv3.Operation) *Operation { _ = "STUB: not implemented"; return nil }

// security is defined, but empty.

// GoLow will return the low-level Operation instance that was used to create the high-level one.
func (o *Operation) GoLow() *lowv3.Operation {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Discriminator instance that was used to create the high-level one, with no type
	return nil
}

func (o *Operation) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Operation object as a byte slice.
	return *new(any)
}

func (o *Operation) Render() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (o *Operation) RenderInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML will create a ready to render YAML representation of the Operation object.
func (o *Operation) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (o *Operation) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
