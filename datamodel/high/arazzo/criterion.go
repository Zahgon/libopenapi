// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Criterion represents a high-level Arazzo Criterion Object.
// https://spec.openapis.org/arazzo/v1.0.1#criterion-object
type Criterion struct {
	Context        string                              `json:"context,omitempty" yaml:"context,omitempty"`
	Condition      string                              `json:"condition,omitempty" yaml:"condition,omitempty"`
	Type           string                              `json:"-" yaml:"-"`
	ExpressionType *CriterionExpressionType            `json:"-" yaml:"-"`
	Extensions     *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low            *low.Criterion
}

// GetEffectiveType returns the effective criterion type. Returns "simple" when Type is empty,
// the string value when set as a scalar, or ExpressionType.Type when the type field is an object.
func (c *Criterion) GetEffectiveType() string { _ = "STUB: not implemented"; return "" }

// NewCriterion creates a new high-level Criterion instance from a low-level one.
func NewCriterion(criterion *low.Criterion) *Criterion { _ = "STUB: not implemented"; return nil }

// Type is a union: scalar string or CriterionExpressionType mapping

// GoLow returns the low-level Criterion instance used to create the high-level one.
func (c *Criterion) GoLow() *low.Criterion {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level Criterion instance with no type.
	return nil
}

func (c *Criterion) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render returns a YAML representation of the Criterion object as a byte slice.
	return *new(any)
}

func (c *Criterion) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the Criterion object.
		nil
}

func (c *Criterion) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
