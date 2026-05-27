// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	low "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// ServerVariable represents a high-level OpenAPI 3+ ServerVariable object, that is backed by a low-level one.
//
// ServerVariable is an object representing a Server Variable for server URL template substitution.
// - https://spec.openapis.org/oas/v3.1.0#server-variable-object
type ServerVariable struct {
	Enum        []string                            `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string                              `json:"default,omitempty" yaml:"default,omitempty"`
	Description string                              `json:"description,omitempty" yaml:"description,omitempty"`
	Extensions  *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low         *low.ServerVariable
}

// NewServerVariable will return a new high-level instance of a ServerVariable from a low-level one.
func NewServerVariable(variable *low.ServerVariable) *ServerVariable {
	_ = "STUB: not implemented"
	return nil
}

// GoLow returns the low-level ServerVariable used to create the high\-level one.
func (s *ServerVariable) GoLow() *low.ServerVariable {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level ServerVariable instance that was used to create the high-level one, with no type
	return nil
}

func (s *ServerVariable) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the ServerVariable object as a byte slice.
	return *new(any)
}

func (s *ServerVariable) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the ServerVariable object.
		nil
}

func (s *ServerVariable) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
