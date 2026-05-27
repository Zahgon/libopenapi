// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// CriterionExpressionType represents a high-level Arazzo Criterion Expression Type Object.
// https://spec.openapis.org/arazzo/v1.0.1#criterion-expression-type-object
type CriterionExpressionType struct {
	Type       string                              `json:"type,omitempty" yaml:"type,omitempty"`
	Version    string                              `json:"version,omitempty" yaml:"version,omitempty"`
	Extensions *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low        *low.CriterionExpressionType
}

// NewCriterionExpressionType creates a new high-level CriterionExpressionType instance from a low-level one.
func NewCriterionExpressionType(cet *low.CriterionExpressionType) *CriterionExpressionType {
	_ = "STUB: not implemented"
	return nil
}

// GoLow returns the low-level CriterionExpressionType instance used to create the high-level one.
func (c *CriterionExpressionType) GoLow() *low.CriterionExpressionType {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level CriterionExpressionType instance with no type.
	return nil
}

func (c *CriterionExpressionType) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the CriterionExpressionType object as a byte slice.
	return *new(any)
}

func (c *CriterionExpressionType) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the CriterionExpressionType object.
		nil
}

func (c *CriterionExpressionType) MarshalYAML() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
