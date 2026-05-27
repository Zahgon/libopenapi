// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	lowv3 "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// buildLowResponse builds a low-level Response from a resolved YAML node.
func buildLowResponse(node *yaml.Node, idx *index.SpecIndex) (*lowv3.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Response represents a high-level OpenAPI 3+ Response object that is backed by a low-level one.
//
// Describes a single response from an API Operation, including design-time, static links to
// operations based on the response.
//   - https://spec.openapis.org/oas/v3.1.0#response-object
type Response struct {
	Reference   string                              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string                              `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string                              `json:"description" yaml:"description"`
	Headers     *orderedmap.Map[string, *Header]    `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     *orderedmap.Map[string, *MediaType] `json:"content,omitempty" yaml:"content,omitempty"`
	Links       *orderedmap.Map[string, *Link]      `json:"links,omitempty" yaml:"links,omitempty"`
	Extensions  *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low         *lowv3.Response
}

// NewResponse creates a new high-level Response object that is backed by a low-level one.
func NewResponse(response *lowv3.Response) *Response { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Response object that was used to create the high-level one.
func (r *Response) GoLow() *lowv3.Response {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Response instance that was used to create the high-level one, with no type
	return nil
}

func (r *Response) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// IsReference returns true if this Response is a reference to another Response definition.
	return *new(any)
}

func (r *Response) IsReference() bool { _ = "STUB: not implemented"; return false }

// GetReference returns the reference string if this is a reference Response.
func (r *Response) GetReference() string {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Response object as a byte slice.
	return ""
}

func (r *Response) Render() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Response) RenderInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML will create a ready to render YAML representation of the Response object.
func (r *Response) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	// Handle reference-only response
	return nil, nil
}

// MarshalYAMLInline will create a ready to render YAML representation of the Response object,
// resolving any references inline where possible.
func (r *Response) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	// reference-only objects render as $ref nodes
	return nil, nil
}

// resolve external reference if present

// MarshalYAMLInlineWithContext will create a ready to render YAML representation of the Response object,
// resolving any references inline where possible. Uses the provided context for cycle detection.
// The ctx parameter should be *base.InlineRenderContext but is typed as any to satisfy the
// high.RenderableInlineWithContext interface without import cycles.
func (r *Response) MarshalYAMLInlineWithContext(ctx any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolve external reference if present

// CreateResponseRef creates a Response that renders as a $ref to another response definition.
// This is useful when building OpenAPI specs programmatically and you want to reference
// a response defined in components/responses rather than inlining the full definition.
//
// Example:
//
//	resp := v3.CreateResponseRef("#/components/responses/NotFound")
//
// Renders as:
//
//	$ref: '#/components/responses/NotFound'
func CreateResponseRef(ref string) *Response { _ = "STUB: not implemented"; return nil }
