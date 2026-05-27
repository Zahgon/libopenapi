// Copyright 2022-2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	low "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// buildLowRequestBody builds a low-level RequestBody from a resolved YAML node.
func buildLowRequestBody(node *yaml.Node, idx *index.SpecIndex) (*low.RequestBody, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequestBody represents a high-level OpenAPI 3+ RequestBody object, backed by a low-level one.
//   - https://spec.openapis.org/oas/v3.1.0#request-body-object
type RequestBody struct {
	Reference   string                              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Description string                              `json:"description,omitempty" yaml:"description,omitempty"`
	Content     *orderedmap.Map[string, *MediaType] `json:"content,omitempty" yaml:"content,omitempty"`
	Required    *bool                               `json:"required,omitempty" yaml:"required,renderZero,omitempty"`
	Extensions  *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low         *low.RequestBody
}

// NewRequestBody will create a new high-level RequestBody instance, from a low-level one.
func NewRequestBody(rb *low.RequestBody) *RequestBody { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level RequestBody instance used to create the high-level one.
func (r *RequestBody) GoLow() *low.RequestBody {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level RequestBody instance that was used to create the high-level one, with no type
	return nil
}

func (r *RequestBody) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// IsReference returns true if this RequestBody is a reference to another RequestBody definition.
	return *new(any)
}

func (r *RequestBody) IsReference() bool { _ = "STUB: not implemented"; return false }

// GetReference returns the reference string if this is a reference RequestBody.
func (r *RequestBody) GetReference() string {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the RequestBody object as a byte slice.
	return ""
}

func (r *RequestBody) Render() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *RequestBody) RenderInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML will create a ready to render YAML representation of the RequestBody object.
func (r *RequestBody) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	// Handle reference-only request body
	return nil, nil
}

// MarshalYAMLInline will create a ready to render YAML representation of the RequestBody object,
// resolving any references inline where possible.
func (r *RequestBody) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	// reference-only objects render as $ref nodes
	return nil, nil
}

// resolve external reference if present

// buildLowRequestBody never returns an error, so we can ignore it

// MarshalYAMLInlineWithContext will create a ready to render YAML representation of the RequestBody object,
// resolving any references inline where possible. Uses the provided context for cycle detection.
// The ctx parameter should be *base.InlineRenderContext but is typed as any to satisfy the
// high.RenderableInlineWithContext interface without import cycles.
func (r *RequestBody) MarshalYAMLInlineWithContext(ctx any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolve external reference if present

// buildLowRequestBody never returns an error, so we can ignore it

// CreateRequestBodyRef creates a RequestBody that renders as a $ref to another request body definition.
// This is useful when building OpenAPI specs programmatically and you want to reference
// a request body defined in components/requestBodies rather than inlining the full definition.
//
// Example:
//
//	rb := v3.CreateRequestBodyRef("#/components/requestBodies/UserInput")
//
// Renders as:
//
//	$ref: '#/components/requestBodies/UserInput'
func CreateRequestBodyRef(ref string) *RequestBody { _ = "STUB: not implemented"; return nil }
