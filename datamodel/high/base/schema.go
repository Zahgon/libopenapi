// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Schema represents a JSON Schema that support Swagger, OpenAPI 3 and OpenAPI 3.1
//
// Until 3.1 OpenAPI had a strange relationship with JSON Schema. It's been a super-set/sub-set
// mix, which has been confusing. So, instead of building a bunch of different models, we have compressed
// all variations into a single model that makes it easy to support multiple spec types.
//
//   - v2 schema: https://swagger.io/specification/v2/#schemaObject
//   - v3 schema: https://swagger.io/specification/#schema-object
//   - v3.1 schema: https://spec.openapis.org/oas/v3.1.0#schema-object
type Schema struct {
	// 3.1 only, used to define a dialect for this schema, label is '$schema'.
	SchemaTypeRef string `json:"$schema,omitempty" yaml:"$schema,omitempty"`

	// In versions 2 and 3.0, this ExclusiveMaximum can only be a boolean.
	// In version 3.1, ExclusiveMaximum is a number.
	ExclusiveMaximum *DynamicValue[bool, float64] `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`

	// In versions 2 and 3.0, this ExclusiveMinimum can only be a boolean.
	// In version 3.1, ExclusiveMinimum is a number.
	ExclusiveMinimum *DynamicValue[bool, float64] `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`

	// In versions 2 and 3.0, this Type is a single value, so array will only ever have one value
	// in version 3.1, Type can be multiple values
	Type []string `json:"type,omitempty" yaml:"type,omitempty"`

	// Schemas are resolved on demand using a SchemaProxy
	AllOf []*SchemaProxy `json:"allOf,omitempty" yaml:"allOf,omitempty"`

	// Polymorphic Schemas are only available in version 3+
	OneOf         []*SchemaProxy `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf         []*SchemaProxy `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	Discriminator *Discriminator `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`

	// in 3.1 examples can be an array (which is recommended)
	Examples []*yaml.Node `json:"examples,omitempty" yaml:"examples,omitempty"`

	// in 3.1 prefixItems provides tuple validation support.
	PrefixItems []*SchemaProxy `json:"prefixItems,omitempty" yaml:"prefixItems,omitempty"`

	// 3.1 Specific properties
	Contains          *SchemaProxy                          `json:"contains,omitempty" yaml:"contains,omitempty"`
	MinContains       *int64                                `json:"minContains,renderZero,omitempty" yaml:"minContains,renderZero,omitempty"`
	MaxContains       *int64                                `json:"maxContains,renderZero,omitempty" yaml:"maxContains,renderZero,omitempty"`
	If                *SchemaProxy                          `json:"if,omitempty" yaml:"if,omitempty"`
	Else              *SchemaProxy                          `json:"else,omitempty" yaml:"else,omitempty"`
	Then              *SchemaProxy                          `json:"then,omitempty" yaml:"then,omitempty"`
	DependentSchemas  *orderedmap.Map[string, *SchemaProxy] `json:"dependentSchemas,omitempty" yaml:"dependentSchemas,omitempty"`
	DependentRequired *orderedmap.Map[string, []string]     `json:"dependentRequired,omitempty" yaml:"dependentRequired,omitempty"`
	PatternProperties *orderedmap.Map[string, *SchemaProxy] `json:"patternProperties,omitempty" yaml:"patternProperties,omitempty"`
	PropertyNames     *SchemaProxy                          `json:"propertyNames,omitempty" yaml:"propertyNames,omitempty"`
	UnevaluatedItems  *SchemaProxy                          `json:"unevaluatedItems,omitempty" yaml:"unevaluatedItems,omitempty"`

	// in 3.1 UnevaluatedProperties can be a Schema or a boolean
	// https://github.com/pb33f/libopenapi/issues/118
	UnevaluatedProperties *DynamicValue[*SchemaProxy, bool] `json:"unevaluatedProperties,omitempty" yaml:"unevaluatedProperties,omitempty"`

	// in 3.1 Items can be a Schema or a boolean
	Items *DynamicValue[*SchemaProxy, bool] `json:"items,omitempty" yaml:"items,omitempty"`

	// 3.1+ only, JSON Schema 2020-12 $id - declares this schema as a schema resource with a URI identifier
	Id string `json:"$id,omitempty" yaml:"$id,omitempty"`

	// 3.1 only, part of the JSON Schema spec provides a way to identify a sub-schema
	Anchor string `json:"$anchor,omitempty" yaml:"$anchor,omitempty"`

	// 3.1+ only, JSON Schema 2020-12 dynamic anchor for recursive schema resolution
	DynamicAnchor string `json:"$dynamicAnchor,omitempty" yaml:"$dynamicAnchor,omitempty"`

	// 3.1+ only, JSON Schema 2020-12 dynamic reference for recursive schema resolution
	DynamicRef string `json:"$dynamicRef,omitempty" yaml:"$dynamicRef,omitempty"`

	// 3.1+ only, JSON Schema 2020-12 $comment - explanatory notes without affecting validation
	Comment string `json:"$comment,omitempty" yaml:"$comment,omitempty"`

	// 3.1+ only, JSON Schema 2020-12 contentSchema - describes structure of decoded content
	ContentSchema *SchemaProxy `json:"contentSchema,omitempty" yaml:"contentSchema,omitempty"`

	// 3.1+ only, JSON Schema 2020-12 $vocabulary - defines available vocabularies in meta-schemas
	Vocabulary *orderedmap.Map[string, bool] `json:"$vocabulary,omitempty" yaml:"$vocabulary,omitempty"`

	// Compatible with all versions
	Not                  *SchemaProxy                          `json:"not,omitempty" yaml:"not,omitempty"`
	Properties           *orderedmap.Map[string, *SchemaProxy] `json:"properties,omitempty" yaml:"properties,omitempty"`
	Title                string                                `json:"title,omitempty" yaml:"title,omitempty"`
	MultipleOf           *float64                              `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Maximum              *float64                              `json:"maximum,renderZero,omitempty" yaml:"maximum,renderZero,omitempty"`
	Minimum              *float64                              `json:"minimum,renderZero,omitempty," yaml:"minimum,renderZero,omitempty"`
	MaxLength            *int64                                `json:"maxLength,renderZero,omitempty" yaml:"maxLength,renderZero,omitempty"`
	MinLength            *int64                                `json:"minLength,renderZero,omitempty" yaml:"minLength,renderZero,omitempty"`
	Pattern              string                                `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	Format               string                                `json:"format,omitempty" yaml:"format,omitempty"`
	MaxItems             *int64                                `json:"maxItems,renderZero,omitempty" yaml:"maxItems,renderZero,omitempty"`
	MinItems             *int64                                `json:"minItems,renderZero,omitempty" yaml:"minItems,renderZero,omitempty"`
	UniqueItems          *bool                                 `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	MaxProperties        *int64                                `json:"maxProperties,renderZero,omitempty" yaml:"maxProperties,renderZero,omitempty"`
	MinProperties        *int64                                `json:"minProperties,renderZero,omitempty" yaml:"minProperties,renderZero,omitempty"`
	Required             []string                              `json:"required,omitempty" yaml:"required,omitempty"`
	Enum                 []*yaml.Node                          `json:"enum,omitempty" yaml:"enum,omitempty"`
	AdditionalProperties *DynamicValue[*SchemaProxy, bool]     `json:"additionalProperties,renderZero,omitempty" yaml:"additionalProperties,renderZero,omitempty"`
	Description          string                                `json:"description,omitempty" yaml:"description,omitempty"`
	ContentEncoding      string                                `json:"contentEncoding,omitempty" yaml:"contentEncoding,omitempty"`
	ContentMediaType     string                                `json:"contentMediaType,omitempty" yaml:"contentMediaType,omitempty"`
	Default              *yaml.Node                            `json:"default,omitempty" yaml:"default,renderZero,omitempty"`
	Const                *yaml.Node                            `json:"const,omitempty" yaml:"const,renderZero,omitempty"`
	Nullable             *bool                                 `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	ReadOnly             *bool                                 `json:"readOnly,renderZero,omitempty" yaml:"readOnly,renderZero,omitempty"`   // https://github.com/pb33f/libopenapi/issues/30
	WriteOnly            *bool                                 `json:"writeOnly,renderZero,omitempty" yaml:"writeOnly,renderZero,omitempty"` // https://github.com/pb33f/libopenapi/issues/30
	XML                  *XML                                  `json:"xml,omitempty" yaml:"xml,omitempty"`
	ExternalDocs         *ExternalDoc                          `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Example              *yaml.Node                            `json:"example,omitempty" yaml:"example,omitempty"`
	Deprecated           *bool                                 `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Extensions           *orderedmap.Map[string, *yaml.Node]   `json:"-" yaml:"-"`
	low                  *base.Schema

	// Parent Proxy refers back to the low level SchemaProxy that is proxying this schema.
	ParentProxy *SchemaProxy `json:"-" yaml:"-"`
}

// NewSchema will create a new high-level schema from a low-level one.
func NewSchema(schema *base.Schema) *Schema { _ = "STUB: not implemented"; return nil }

// if we're dealing with a 3.0 spec using a bool

// if we're dealing with a 3.1 spec using an int

// if we're dealing with a 3.0 spec using a bool

// if we're dealing with a 3.1 spec, using an int

// 3.0 spec is a single value

// 3.1 spec may have multiple values

// async work.
// any polymorphic properties need to be handled in their own threads
// any properties each need to be processed in their own thread.
// we go as fast as we can.

// for every item, build schema async

// schema async

// props async

// Handle DependentRequired

// GoLow will return the low-level instance of Schema that was used to create the high level one.
func (s *Schema) GoLow() *base.Schema {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Schema instance that was used to create the high-level one, with no type
	return nil
}

func (s *Schema) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Schema object as a byte slice.
	return *new(any)
}

func (s *Schema) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// RenderInlineWithContext will return a YAML representation of the Schema object as a byte slice
		// using the provided InlineRenderContext for cycle detection.
		// Use this when multiple goroutines may render the same schemas concurrently.
		// The ctx parameter should be *InlineRenderContext but is typed as any to avoid import cycles.
		nil
}

func (s *Schema) RenderInlineWithContext(ctx any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RenderInline will return a YAML representation of the Schema object as a byte slice.
// All the $ref values will be inlined, as in resolved in place.
// This method creates a fresh InlineRenderContext internally.
//
// Make sure you don't have any circular references!
func (s *Schema) RenderInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML will create a ready to render YAML representation of the Schema object.
func (s *Schema) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// determine index version

// MarshalJSON will create a ready to render JSON representation of the Schema object.
func (s *Schema) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// determine index version

// render node

// marshal into struct

// return JSON bytes

// MarshalYAMLInlineWithContext will render out the Schema pointer as YAML using the provided
// InlineRenderContext for cycle detection. All refs will be inlined fully.
// Use this when multiple goroutines may render the same schemas concurrently.
// The ctx parameter should be *InlineRenderContext but is typed as any to satisfy the
// high.RenderableInlineWithContext interface without import cycles.
func (s *Schema) MarshalYAMLInlineWithContext(ctx any) (interface{}, error) {
	_ = "STUB: not implemented"
	// ensure we have a valid render context; create default bundle mode context if nil.
	// this ensures backward compatibility where nil context = bundle mode behavior.
	return nil, nil
}

// determine if we should preserve discriminator refs based on rendering mode.
// in validation mode, we need to fully inline all refs for the JSON schema compiler.
// in bundle mode (default), we preserve discriminator refs for mapping compatibility.

// mark oneOf/anyOf refs as preserved in the context (not on the SchemaProxy).
// this avoids mutating shared state and prevents race conditions.

// determine index version

// MarshalYAMLInline will render out the Schema pointer as YAML, and all refs will be inlined fully.
// This method creates a fresh InlineRenderContext internally.
func (s *Schema) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalJSONInline will render out the Schema pointer as JSON, and all refs will be inlined fully
func (s *Schema) MarshalJSONInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// determine index version

// render node

// marshal into struct

// return JSON bytes
