// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// PayloadReplacement represents a high-level Arazzo Payload Replacement Object.
// https://spec.openapis.org/arazzo/v1.0.1#payload-replacement-object
type PayloadReplacement struct {
	Target     string                              `json:"target,omitempty" yaml:"target,omitempty"`
	Value      *yaml.Node                          `json:"value,omitempty" yaml:"value,omitempty"`
	Extensions *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low        *low.PayloadReplacement
}

// NewPayloadReplacement creates a new high-level PayloadReplacement instance from a low-level one.
func NewPayloadReplacement(pr *low.PayloadReplacement) *PayloadReplacement {
	_ = "STUB: not implemented"
	return nil
}

// GoLow returns the low-level PayloadReplacement instance used to create the high-level one.
func (p *PayloadReplacement) GoLow() *low.PayloadReplacement {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level PayloadReplacement instance with no type.
	return nil
}

func (p *PayloadReplacement) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the PayloadReplacement object as a byte slice.
	return *new(any)
}

func (p *PayloadReplacement) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the PayloadReplacement object.
		nil
}

func (p *PayloadReplacement) MarshalYAML() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
