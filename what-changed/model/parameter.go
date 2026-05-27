// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
	"go.yaml.in/yaml/v4"
)

// ParameterChanges represents changes found between Swagger or OpenAPI Parameter objects.
type ParameterChanges struct {
	*PropertyChanges
	Name             string            `json:"name,omitempty" yaml:"name,omitempty"`
	SchemaChanges    *SchemaChanges    `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	ExtensionChanges *ExtensionChanges `json:"extensions,omitempty" yaml:"extensions,omitempty"`

	// Swagger supports Items.
	ItemsChanges *ItemsChanges `json:"items,omitempty" yaml:"items,omitempty"`

	// OpenAPI supports examples and content types.
	ExamplesChanges map[string]*ExampleChanges   `json:"examples,omitempty" yaml:"examples,omitempty"`
	ContentChanges  map[string]*MediaTypeChanges `json:"content,omitempty" yaml:"content,omitempty"`
}

// GetAllChanges returns a slice of all changes made between Parameter objects
func (p *ParameterChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns a count of everything that changed
func (p *ParameterChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges always returns 0 for ExternalDoc objects, they are non-binding.
func (p *ParameterChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

func addPropertyCheck(props *[]*PropertyCheck,
	lvn, rvn *yaml.Node, lv, rv any, changes *[]*Change, label string, breaking bool,
	component, property string,
) {
	_ = "STUB: not implemented"
	return
}

func addOpenAPIParameterProperties(left, right low.OpenAPIParameter, changes *[]*Change) []*PropertyCheck {
	_ = "STUB: not implemented"
	return nil

	// style
}

// allow reserved

// explode

// deprecated

// example

func addSwaggerParameterProperties(left, right low.SwaggerParameter, changes *[]*Change) []*PropertyCheck {
	_ = "STUB: not implemented"
	return nil

	// type
}

// format

// collection format

// maximum

// minimum

// exclusive maximum

// exclusive minimum

// max length

// min length

// pattern

// max items

// min items

// unique items

// default

// multiple of

func addCommonParameterProperties(left, right low.SharedParameters, changes *[]*Change) []*PropertyCheck {
	_ = "STUB: not implemented"
	return nil
}

// in

// description

// required

// allow empty value

// CompareParametersV3 is an OpenAPI type safe proxy for CompareParameters
func CompareParametersV3(l, r *v3.Parameter) *ParameterChanges {
	_ = "STUB: not implemented"
	return nil
}

// CompareParameters compares a left and right Swagger or OpenAPI Parameter object for any changes. If found returns
// a pointer to ParameterChanges. If nothing is found, returns nil.
func CompareParameters(l, r any) *ParameterChanges { _ = "STUB: not implemented"; return nil }

// perform hash check to avoid further processing

// extract schema

// items

// enum

// OpenAPI

// perform hash check to avoid further processing

// example

// examples

// content

func checkParameterExample(expLeft, expRight low.NodeReference[*yaml.Node], changes []*Change) {
	_ = "STUB: not implemented"
	return
}
