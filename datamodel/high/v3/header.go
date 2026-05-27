// Copyright 2022-2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	highbase "github.com/pb33f/libopenapi/datamodel/high/base"
	lowmodel "github.com/pb33f/libopenapi/datamodel/low"
	lowv3 "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// buildLowHeader builds a low-level Header from a resolved YAML node.
func buildLowHeader(node *yaml.Node, idx *index.SpecIndex) (*lowv3.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Header represents a high-level OpenAPI 3+ Header object backed by a low-level one.
//   - https://spec.openapis.org/oas/v3.1.0#header-object
type Header struct {
	Reference       string                                     `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Description     string                                     `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool                                       `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      bool                                       `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool                                       `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	Style           string                                     `json:"style,omitempty" yaml:"style,omitempty"`
	Explode         bool                                       `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved   bool                                       `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema          *highbase.SchemaProxy                      `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example         *yaml.Node                                 `json:"example,omitempty" yaml:"example,omitempty"`
	Examples        *orderedmap.Map[string, *highbase.Example] `json:"examples,omitempty" yaml:"examples,omitempty"`
	Content         *orderedmap.Map[string, *MediaType]        `json:"content,omitempty" yaml:"content,omitempty"`
	Extensions      *orderedmap.Map[string, *yaml.Node]        `json:"-" yaml:"-"`
	low             *lowv3.Header
}

// NewHeader creates a new high-level Header instance from a low-level one.
func NewHeader(header *lowv3.Header) *Header { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Header instance used to create the high-level one.
func (h *Header) GoLow() *lowv3.Header {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Header instance that was used to create the high-level one, with no type
	return nil
}

func (h *Header) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// IsReference returns true if this Header is a reference to another Header definition.
	return *new(any)
}

func (h *Header) IsReference() bool { _ = "STUB: not implemented"; return false }

// GetReference returns the reference string if this is a reference Header.
func (h *Header) GetReference() string {
	_ = "STUB: not implemented"

	// ExtractHeaders will extract a hard to navigate low-level Header map, into simple high-level one.
	return ""
}

func ExtractHeaders(elements *orderedmap.Map[lowmodel.KeyReference[string], lowmodel.ValueReference[*lowv3.Header]]) *orderedmap.Map[string, *Header] {
	_ = "STUB: not implemented"
	return nil
}

// Render will return a YAML representation of the Header object as a byte slice.
func (h *Header) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// RenderInline will return a YAML representation of the Header object as a byte slice with references resolved.
		nil
}

func (h *Header) RenderInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML will create a ready to render YAML representation of the Header object.
func (h *Header) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	// Handle reference-only header
	return nil, nil
}

// MarshalYAMLInline will create a ready to render YAML representation of the Header object with references resolved.
func (h *Header) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	// reference-only objects render as $ref nodes
	return nil, nil
}

// resolve external reference if present

// MarshalYAMLInlineWithContext will create a ready to render YAML representation of the Header object,
// resolving any references inline where possible. Uses the provided context for cycle detection.
// The ctx parameter should be *base.InlineRenderContext but is typed as any to satisfy the
// high.RenderableInlineWithContext interface without import cycles.
func (h *Header) MarshalYAMLInlineWithContext(ctx any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolve external reference if present

// CreateHeaderRef creates a Header that renders as a $ref to another header definition.
// This is useful when building OpenAPI specs programmatically, and you want to reference
// a header defined in components/headers rather than inlining the full definition.
//
// Example:
//
//	header := v3.CreateHeaderRef("#/components/headers/X-Rate-Limit")
//
// Renders as:
//
//	$ref: '#/components/headers/X-Rate-Limit'
func CreateHeaderRef(ref string) *Header { _ = "STUB: not implemented"; return nil }
