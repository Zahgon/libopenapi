// Copyright 2022-2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	lowmodel "github.com/pb33f/libopenapi/datamodel/low"
	lowv3 "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/orderedmap"
)

// Encoding represents an OpenAPI 3+ Encoding object
//   - https://spec.openapis.org/oas/v3.1.0#encoding-object
type Encoding struct {
	ContentType   string                           `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       *orderedmap.Map[string, *Header] `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string                           `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       *bool                            `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                             `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	low           *lowv3.Encoding
}

// NewEncoding creates a new instance of Encoding from a low-level one.
func NewEncoding(encoding *lowv3.Encoding) *Encoding { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Encoding instance used to create the high-level one.
func (e *Encoding) GoLow() *lowv3.Encoding {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Encoding instance that was used to create the high-level one, with no type
	return nil
}

func (e *Encoding) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Encoding object as a byte slice.
	return *new(any)
}

func (e *Encoding) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the Encoding object.
		nil
}

func (e *Encoding) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAMLInline will create a ready to render YAML representation of the Encoding object,
// with all references resolved inline.
func (e *Encoding) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalYAMLInlineWithContext will create a ready to render YAML representation of the Encoding object,
// resolving any references inline where possible. Uses the provided context for cycle detection.
// The ctx parameter should be *base.InlineRenderContext but is typed as any to satisfy the
// high.RenderableInlineWithContext interface without import cycles.
func (e *Encoding) MarshalYAMLInlineWithContext(ctx any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractEncoding converts hard to navigate low-level plumbing Encoding definitions, into a high-level simple map
func ExtractEncoding(elements *orderedmap.Map[lowmodel.KeyReference[string], lowmodel.ValueReference[*lowv3.Encoding]]) *orderedmap.Map[string, *Encoding] {
	_ = "STUB: not implemented"
	return nil
}
