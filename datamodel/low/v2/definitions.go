// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// ParameterDefinitions is a low-level representation of a Swagger / OpenAPI 2 Parameters Definitions object.
//
// ParameterDefinitions holds parameters to be reused across operations. Parameter definitions can be
// referenced to the ones defined here. It does not define global operation parameters
//   - https://swagger.io/specification/v2/#parametersDefinitionsObject
type ParameterDefinitions struct {
	Definitions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*Parameter]]
}

// ResponsesDefinitions is a low-level representation of a Swagger / OpenAPI 2 Responses Definitions object.
//
// ResponsesDefinitions is an object to hold responses to be reused across operations. Response definitions can be
// referenced to the ones defined here. It does not define global operation responses
//   - https://swagger.io/specification/v2/#responsesDefinitionsObject
type ResponsesDefinitions struct {
	Definitions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*Response]]
}

// SecurityDefinitions is a low-level representation of a Swagger / OpenAPI 2 Security Definitions object.
//
// A declaration of the security schemes available to be used in the specification. This does not enforce the security
// schemes on the operations and only serves to provide the relevant details for each scheme
//   - https://swagger.io/specification/v2/#securityDefinitionsObject
type SecurityDefinitions struct {
	Definitions *orderedmap.Map[low.KeyReference[string], low.ValueReference[*SecurityScheme]]
}

// Definitions is a low-level representation of a Swagger / OpenAPI 2 Definitions object
//
// An object to hold data types that can be consumed and produced by operations. These data types can be primitives,
// arrays or models.
//   - https://swagger.io/specification/v2/#definitionsObject
type Definitions struct {
	Schemas *orderedmap.Map[low.KeyReference[string], low.ValueReference[*base.SchemaProxy]]
}

// FindSchema will attempt to locate a base.SchemaProxy instance using a name.
func (d *Definitions) FindSchema(schema string) *low.ValueReference[*base.SchemaProxy] {
	_ = "STUB: not implemented"
	return nil
}

// FindParameter will attempt to locate a Parameter instance using a name.
func (pd *ParameterDefinitions) FindParameter(parameter string) *low.ValueReference[*Parameter] {
	_ = "STUB: not implemented"
	return nil
}

// FindResponse will attempt to locate a Response instance using a name.
func (r *ResponsesDefinitions) FindResponse(response string) *low.ValueReference[*Response] {
	_ = "STUB: not implemented"
	return nil
}

// FindSecurityDefinition will attempt to locate a SecurityScheme using a name.
func (s *SecurityDefinitions) FindSecurityDefinition(securityDef string) *low.ValueReference[*SecurityScheme] {
	_ = "STUB: not implemented"
	return nil
}

// Build will extract all definitions into SchemaProxy instances.
func (d *Definitions) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// input and output goroutines.

// TranslatePipeline input.

// TranslatePipeline output.

// Hash will return a consistent Hash of the Definitions object
func (d *Definitions) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Build will extract all ParameterDefinitions into Parameter instances.
func (pd *ParameterDefinitions) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// re-usable struct for holding results as k/v pairs.
type definitionResult[T any] struct {
	k *yaml.Node
	v low.ValueReference[T]
}

// Build will extract all ResponsesDefinitions into Response instances.
func (r *ResponsesDefinitions) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Build will extract all SecurityDefinitions into SecurityScheme instances.
func (s *SecurityDefinitions) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}
