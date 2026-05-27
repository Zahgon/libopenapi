// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	lowBase "github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// buildLowExample builds a low-level Example from a resolved YAML node.
func buildLowExample(node *yaml.Node, idx *index.SpecIndex) (*lowBase.Example, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Example represents a high-level Example object as defined by OpenAPI 3+
//
//	v3 - https://spec.openapis.org/oas/v3.1.0#example-object
type Example struct {
	Reference       string                              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary         string                              `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description     string                              `json:"description,omitempty" yaml:"description,omitempty"`
	Value           *yaml.Node                          `json:"value,omitempty" yaml:"value,omitempty"`
	ExternalValue   string                              `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
	DataValue       *yaml.Node                          `json:"dataValue,omitempty" yaml:"dataValue,omitempty"`             // OpenAPI 3.2+ dataValue field
	SerializedValue string                              `json:"serializedValue,omitempty" yaml:"serializedValue,omitempty"` // OpenAPI 3.2+ serializedValue field
	Extensions      *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low             *lowBase.Example
}

// NewExample will create a new instance of an Example, using a low-level Example.
func NewExample(example *lowBase.Example) *Example { _ = "STUB: not implemented"; return nil }

// GoLow will return the low-level Example used to build the high level one.
func (e *Example) GoLow() *lowBase.Example { _ = "STUB: not implemented"; return nil }

// GoLowUntyped will return the low-level Example instance that was used to create the high-level one, with no type
func (e *Example) GoLowUntyped() any { _ = "STUB: not implemented"; return *new(any) }

// IsReference returns true if this Example is a reference to another Example definition.
func (e *Example) IsReference() bool { _ = "STUB: not implemented"; return false }

// GetReference returns the reference string if this is a reference Example.
func (e *Example) GetReference() string {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Example object as a byte slice.
	return ""
}

func (e *Example) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the Example object.
		nil
}

func (e *Example) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	// Handle reference-only example
	return nil, nil
}

// MarshalYAMLInline will create a ready to render YAML representation of the Example object,
// with all references resolved inline.
func (e *Example) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	// reference-only objects render as $ref nodes
	return nil, nil
}

// resolve external reference if present

// buildLowExample never returns an error, so we can ignore it

// MarshalYAMLInlineWithContext will create a ready to render YAML representation of the Example object,
// resolving any references inline where possible. Uses the provided context for cycle detection.
// The ctx parameter should be *InlineRenderContext but is typed as any to satisfy the
// high.RenderableInlineWithContext interface without import cycles.
func (e *Example) MarshalYAMLInlineWithContext(ctx any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolve external reference if present

// buildLowExample never returns an error, so we can ignore it

// CreateExampleRef creates an Example that renders as a $ref to another example definition.
// This is useful when building OpenAPI specs programmatically and you want to reference
// an example defined in components/examples rather than inlining the full definition.
//
// Example:
//
//	ex := base.CreateExampleRef("#/components/examples/UserExample")
//
// Renders as:
//
//	$ref: '#/components/examples/UserExample'
func CreateExampleRef(ref string) *Example { _ = "STUB: not implemented"; return nil }

// MarshalJSON will marshal this into a JSON byte slice
func (e *Example) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractExamples will convert a low-level example map, into a high level one that is simple to navigate.
// no fidelity is lost, everything is still available via GoLow()
func ExtractExamples(elements *orderedmap.Map[low.KeyReference[string], low.ValueReference[*lowBase.Example]]) *orderedmap.Map[string, *Example] {
	_ = "STUB: not implemented"
	return nil
}
