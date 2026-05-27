// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// RequestBody represents a high-level Arazzo Request Body Object.
// https://spec.openapis.org/arazzo/v1.0.1#request-body-object
type RequestBody struct {
	ContentType  string                              `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Payload      *yaml.Node                          `json:"payload,omitempty" yaml:"payload,omitempty"`
	Replacements []*PayloadReplacement               `json:"replacements,omitempty" yaml:"replacements,omitempty"`
	Extensions   *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low          *low.RequestBody
}

// NewRequestBody creates a new high-level RequestBody instance from a low-level one.
func NewRequestBody(rb *low.RequestBody) *RequestBody { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level RequestBody instance used to create the high-level one.
func (r *RequestBody) GoLow() *low.RequestBody {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level RequestBody instance with no type.
	return nil
}

func (r *RequestBody) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the RequestBody object as a byte slice.
	return *new(any)
}

func (r *RequestBody) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the RequestBody object.
		nil
}

func (r *RequestBody) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
