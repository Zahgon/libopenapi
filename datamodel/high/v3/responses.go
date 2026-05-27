// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	low "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Responses represents a high-level OpenAPI 3+ Responses object that is backed by a low-level one.
//
// It's a container for the expected responses of an operation. The container maps a HTTP response code to the
// expected response.
//
// The specification is not necessarily expected to cover all possible HTTP response codes because they may not be
// known in advance. However, documentation is expected to cover a successful operation response and any known errors.
//
// The default MAY be used as a default response object for all HTTP codes that are not covered individually by
// the Responses Object.
//
// The Responses Object MUST contain at least one response code, and if only one response code is provided it SHOULD
// be the response for a successful operation call.
//   - https://spec.openapis.org/oas/v3.1.0#responses-object
type Responses struct {
	Codes      *orderedmap.Map[string, *Response]  `json:"-" yaml:"-"`
	Default    *Response                           `json:"default,omitempty" yaml:"default,omitempty"`
	Extensions *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low        *low.Responses
}

// NewResponses will create a new high-level Responses instance from a low-level one. It operates asynchronously
// internally, as each response may be considerable in complexity.
func NewResponses(responses *low.Responses) *Responses { _ = "STUB: not implemented"; return nil }

// FindResponseByCode is a shortcut for looking up code by an integer vs. a string
func (r *Responses) FindResponseByCode(code int) *Response { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Response object used to create the high-level one.
func (r *Responses) GoLow() *low.Responses {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Responses instance that was used to create the high-level one, with no type
	return nil
}

func (r *Responses) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Responses object as a byte slice.
	return *new(any)
}

func (r *Responses) Render() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Responses) RenderInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML will create a ready to render YAML representation of the Responses object.
func (r *Responses) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	// map keys correctly.
	return nil, nil
}

// default to a high value to weight new content to the bottom.

// extract extensions

func (r *Responses) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	// map keys correctly.
	return nil, nil
}

// default to a high value to weight new content to the bottom.

// extract extensions
