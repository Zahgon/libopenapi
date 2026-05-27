// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"github.com/pb33f/libopenapi/datamodel/high/base"
	low "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// buildLowParameter builds a low-level Parameter from a resolved YAML node.
func buildLowParameter(node *yaml.Node, idx *index.SpecIndex) (*low.Parameter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parameter represents a high-level OpenAPI 3+ Parameter object, that is backed by a low-level one.
//
// A unique parameter is defined by a combination of a name and location.
//   - https://spec.openapis.org/oas/v3.1.0#parameter-object
type Parameter struct {
	Reference       string                                 `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Name            string                                 `json:"name,omitempty" yaml:"name,omitempty"`
	In              string                                 `json:"in,omitempty" yaml:"in,omitempty"`
	Description     string                                 `json:"description,omitempty" yaml:"description,omitempty"`
	Required        *bool                                  `json:"required,renderZero,omitempty" yaml:"required,renderZero,omitempty"`
	Deprecated      bool                                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool                                   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	Style           string                                 `json:"style,omitempty" yaml:"style,omitempty"`
	Explode         *bool                                  `json:"explode,renderZero,omitempty" yaml:"explode,renderZero,omitempty"`
	AllowReserved   bool                                   `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema          *base.SchemaProxy                      `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example         *yaml.Node                             `json:"example,omitempty" yaml:"example,omitempty"`
	Examples        *orderedmap.Map[string, *base.Example] `json:"examples,omitempty" yaml:"examples,omitempty"`
	Content         *orderedmap.Map[string, *MediaType]    `json:"content,omitempty" yaml:"content,omitempty"`
	Extensions      *orderedmap.Map[string, *yaml.Node]    `json:"-" yaml:"-"`
	low             *low.Parameter
}

// NewParameter will create a new high-level instance of a Parameter, using a low-level one.
func NewParameter(param *low.Parameter) *Parameter { _ = "STUB: not implemented"; return nil }

// CreateParameterRef creates a Parameter that renders as a $ref to another parameter definition.
// This is useful when building OpenAPI specs programmatically and you want to reference
// a parameter defined in components/parameters rather than inlining the full definition.
//
// Example:
//
//	param := v3.CreateParameterRef("#/components/parameters/limitParam")
//
// Renders as:
//
//	$ref: '#/components/parameters/limitParam'
func CreateParameterRef(ref string) *Parameter { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Parameter used to create the high-level one.
func (p *Parameter) GoLow() *low.Parameter {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Discriminator instance that was used to create the high-level one, with no type
	return nil
}

func (p *Parameter) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// IsReference returns true if this Parameter is a reference to another Parameter definition.
	return *new(any)
}

func (p *Parameter) IsReference() bool { _ = "STUB: not implemented"; return false }

// GetReference returns the reference string if this is a reference Parameter.
func (p *Parameter) GetReference() string {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Encoding object as a byte slice.
	return ""
}

func (p *Parameter) Render() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Parameter) RenderInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML will create a ready to render YAML representation of the Parameter object.
func (p *Parameter) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	// Handle reference-only parameter
	return nil, nil
}

// MarshalYAMLInline will create a ready to render YAML representation of the Parameter object,
// resolving any references inline where possible.
func (p *Parameter) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	// reference-only objects render as $ref nodes
	return nil, nil
}

// resolve external reference if present

// MarshalYAMLInlineWithContext will create a ready to render YAML representation of the Parameter object,
// resolving any references inline where possible. Uses the provided context for cycle detection.
// The ctx parameter should be *base.InlineRenderContext but is typed as any to satisfy the
// high.RenderableInlineWithContext interface without import cycles.
func (p *Parameter) MarshalYAMLInlineWithContext(ctx any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolve external reference if present

// IsExploded will return true if the parameter is exploded, false otherwise.
func (p *Parameter) IsExploded() bool { _ = "STUB: not implemented"; return false }

// IsDefaultFormEncoding will return true if the parameter has no exploded value, or has exploded set to true, and no style
// or a style set to form. This combination is the default encoding/serialization style for parameters for OpenAPI 3+
func (p *Parameter) IsDefaultFormEncoding() bool { _ = "STUB: not implemented"; return false }

// IsDefaultHeaderEncoding will return true if the parameter has no exploded value, or has exploded set to false, and no style
// or a style set to simple. This combination is the default encoding/serialization style for header parameters for OpenAPI 3+
func (p *Parameter) IsDefaultHeaderEncoding() bool { _ = "STUB: not implemented"; return false }

// IsDefaultPathEncoding will return true if the parameter has no exploded value, or has exploded set to false, and no style
// or a style set to simple. This combination is the default encoding/serialization style for path parameters for OpenAPI 3+
func (p *Parameter) IsDefaultPathEncoding() bool { _ = "STUB: not implemented"; return false }

// header default encoding and path default encoding are the same
