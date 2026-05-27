// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	highbase "github.com/pb33f/libopenapi/datamodel/high/base"
	lowmodel "github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	low "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Components represents a high-level OpenAPI 3+ Components Object, that is backed by a low-level one.
//
// Holds a set of reusable objects for different aspects of the OAS. All objects defined within the components object
// will have no effect on the API unless they are explicitly referenced from properties outside the components object.
//   - https://spec.openapis.org/oas/v3.1.0#components-object
type Components struct {
	Schemas         *orderedmap.Map[string, *highbase.SchemaProxy] `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       *orderedmap.Map[string, *Response]             `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      *orderedmap.Map[string, *Parameter]            `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        *orderedmap.Map[string, *highbase.Example]     `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   *orderedmap.Map[string, *RequestBody]          `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         *orderedmap.Map[string, *Header]               `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemes *orderedmap.Map[string, *SecurityScheme]       `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Links           *orderedmap.Map[string, *Link]                 `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       *orderedmap.Map[string, *Callback]             `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	PathItems       *orderedmap.Map[string, *PathItem]             `json:"pathItems,omitempty" yaml:"pathItems,omitempty"`
	MediaTypes      *orderedmap.Map[string, *MediaType]            `json:"mediaTypes,omitempty" yaml:"mediaTypes,omitempty"` // OpenAPI 3.2+ mediaTypes section
	Extensions      *orderedmap.Map[string, *yaml.Node]            `json:"-" yaml:"-"`
	low             *low.Components
}

// NewComponents will create new high-level instance of Components from a low-level one. Components can be considerable
// in scope, with a lot of different properties across different categories. All components are built asynchronously
// in order to keep things fast.
func NewComponents(comp *low.Components) *Components { _ = "STUB: not implemented"; return nil }

// build all components asynchronously.

// contains a component build result.
type componentResult[T any] struct {
	res T
	key string
}

// buildComponent builds component structs from low level structs.
func buildComponent[IN any, OUT any](inMap *orderedmap.Map[lowmodel.KeyReference[string], lowmodel.ValueReference[IN]], outMap *orderedmap.Map[string, OUT], translateItem func(IN) OUT) {
	_ = "STUB: not implemented"
	return
}

// buildSchema builds a schema from low level structs.
func buildSchema(inMap *orderedmap.Map[lowmodel.KeyReference[string], lowmodel.ValueReference[*base.SchemaProxy]], outMap *orderedmap.Map[string, *highbase.SchemaProxy]) {
	_ = "STUB: not implemented"
	return
}

// GoLow returns the low-level Components instance used to create the high-level one.
func (c *Components) GoLow() *low.Components {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level Components instance used to create the high-level one as an interface{}.
	return nil
}

func (c *Components) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Components object as a byte slice.
	return *new(any)
}

func (c *Components) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the Response object.
		nil
}

func (c *Components) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// RenderInline will return a YAML representation of the Components object as a byte slice with references resolved.
func (c *Components) RenderInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAMLInline will create a ready to render YAML representation of the Components object with references resolved.
func (c *Components) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Components) warnPreservedComponentMapRefs() { _ = "STUB: not implemented"; return }

func warnComponentRefEntries[T any](
	logger interface {
		Warn(msg string, args ...any)
	},
	section string,
	m *orderedmap.Map[lowmodel.KeyReference[string], lowmodel.ValueReference[T]],
) {
	_ = "STUB: not implemented"
	return
}

// preserveInvalidComponentMapRefs patches the rendered Components YAML tree so that invalid
// map-level "$ref" entries under component sections survive a render cycle unchanged.
//
// Inputs like:
//
//	components:
//	  parameters:
//	    $ref: "./params.yaml"
//
// are not valid OpenAPI component maps, but they do appear in the wild. The normal high-level
// render path treats "$ref" as a literal component name and can otherwise collapse the scalar
// value into an empty object. For these cases we preserve the original raw YAML nodes and pair
// the behavior with a warning log, rather than silently rewriting the input.
func (c *Components) preserveInvalidComponentMapRefs(rendered *yaml.Node) {
	_ = "STUB: not implemented"
	return
}

// preserveComponentRefEntries re-inserts a scalar "$ref" entry into the rendered YAML for a
// specific component section. Only literal "$ref" keys backed by scalar low-level value nodes
// are preserved; real component entries and malformed non-scalar values are ignored.
func preserveComponentRefEntries[T any](
	rendered *yaml.Node,
	section string,
	m *orderedmap.Map[lowmodel.KeyReference[string], lowmodel.ValueReference[T]],
) {
	_ = "STUB: not implemented"
	return
}

// findMapValueNode returns the mapping value node for key from a YAML mapping node.
func findMapValueNode(m *yaml.Node, key string) *yaml.Node { _ = "STUB: not implemented"; return nil }

// upsertMapNodeEntry replaces or appends a key/value pair in a YAML mapping node.
func upsertMapNodeEntry(m *yaml.Node, keyNode, valueNode *yaml.Node) {
	_ = "STUB: not implemented"
	return
}

// cloneYAMLNode deep-copies a YAML node tree so preserved low-level nodes can be spliced into
// rendered output without mutating the original parsed model.
func cloneYAMLNode(node *yaml.Node) *yaml.Node { _ = "STUB: not implemented"; return nil }
