// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

import (
	highbase "github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

type SchemaMetadataProvider interface {
	OpenAPISchemaMetadata() any
}

type providerSchemaMetadata struct {
	Ref                   string
	SchemaTypeRef         string
	ExclusiveMaximum      *providerDynamicBoolNumber
	ExclusiveMinimum      *providerDynamicBoolNumber
	Type                  []string
	AllOf                 []*providerSchemaMetadata
	OneOf                 []*providerSchemaMetadata
	AnyOf                 []*providerSchemaMetadata
	Discriminator         *providerDiscriminatorMetadata
	Examples              []*providerYAMLNode
	PrefixItems           []*providerSchemaMetadata
	Contains              *providerSchemaMetadata
	MinContains           *providerInt
	MaxContains           *providerInt
	If                    *providerSchemaMetadata
	Else                  *providerSchemaMetadata
	Then                  *providerSchemaMetadata
	DependentSchemas      []providerNamedSchemaMetadata
	DependentRequired     []providerStringList
	PatternProperties     []providerNamedSchemaMetadata
	PropertyNames         *providerSchemaMetadata
	UnevaluatedItems      *providerSchemaMetadata
	UnevaluatedProperties *providerDynamicSchemaBool
	Items                 *providerDynamicSchemaBool
	ID                    string
	Anchor                string
	DynamicAnchor         string
	DynamicRef            string
	Comment               string
	ContentSchema         *providerSchemaMetadata
	Vocabulary            []providerStringBool
	Not                   *providerSchemaMetadata
	Properties            []providerNamedSchemaMetadata
	Title                 string
	MultipleOf            *providerFloat
	Maximum               *providerFloat
	Minimum               *providerFloat
	MaxLength             *providerInt
	MinLength             *providerInt
	Pattern               string
	Format                string
	MaxItems              *providerInt
	MinItems              *providerInt
	UniqueItems           *providerBool
	MaxProperties         *providerInt
	MinProperties         *providerInt
	Required              []string
	Enum                  []*providerYAMLNode
	AdditionalProperties  *providerDynamicSchemaBool
	Description           string
	ContentEncoding       string
	ContentMediaType      string
	Default               *providerYAMLNode
	Const                 *providerYAMLNode
	Nullable              *providerBool
	ReadOnly              *providerBool
	WriteOnly             *providerBool
	Example               *providerYAMLNode
	Deprecated            *providerBool
	Extensions            []providerNamedYAMLNode
}

type providerDynamicBoolNumber struct {
	Bool   *providerBool
	Number *providerFloat
}

type providerDynamicSchemaBool struct {
	Schema *providerSchemaMetadata
	Bool   *providerBool
}

type providerDiscriminatorMetadata struct {
	PropertyName   string
	Mapping        []providerStringString
	DefaultMapping string
}

type providerNamedSchemaMetadata struct {
	Name   string
	Schema *providerSchemaMetadata
}

type providerNamedYAMLNode struct {
	Name  string
	Value *providerYAMLNode
}

type providerStringBool struct {
	Name  string
	Value bool
}

type providerStringString struct {
	Name  string
	Value string
}

type providerStringList struct {
	Name   string
	Values []string
}

type providerYAMLNode struct {
	Kind    string
	Style   int
	Tag     string
	Value   string
	Anchor  string
	Content []*providerYAMLNode
	Alias   *providerYAMLNode
}

type providerFloat struct {
	Value float64
}

type providerInt struct {
	Value int64
}

type providerBool struct {
	Value bool
}

func schemaProxyFromProviderMetadata(value any) (*highbase.SchemaProxy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func schemaProxyFromMetadata(metadata *providerSchemaMetadata) *highbase.SchemaProxy {
	_ = "STUB: not implemented"
	return nil
}

func schemaFromMetadata(metadata *providerSchemaMetadata) *highbase.Schema {
	_ = "STUB: not implemented"
	return nil
}

func schemaMetadataHasSiblings(metadata *providerSchemaMetadata) bool {
	_ = "STUB: not implemented"
	return false
}

func schemaMetadataEmpty(metadata *providerSchemaMetadata) bool {
	_ = "STUB: not implemented"
	return false
}

func dynamicBoolNumberFromMetadata(metadata *providerDynamicBoolNumber) *highbase.DynamicValue[bool, float64] {
	_ = "STUB: not implemented"
	return nil
}

func dynamicSchemaBoolFromMetadata(metadata *providerDynamicSchemaBool) *highbase.DynamicValue[*highbase.SchemaProxy, bool] {
	_ = "STUB: not implemented"
	return nil
}

func discriminatorFromMetadata(metadata *providerDiscriminatorMetadata) *highbase.Discriminator {
	_ = "STUB: not implemented"
	return nil
}

func schemaSliceFromMetadata(values []*providerSchemaMetadata) []*highbase.SchemaProxy {
	_ = "STUB: not implemented"
	return nil
}

func schemaMapFromMetadata(values []providerNamedSchemaMetadata) *orderedmap.Map[string, *highbase.SchemaProxy] {
	_ = "STUB: not implemented"
	return nil
}

func stringBoolMapFromMetadata(values []providerStringBool) *orderedmap.Map[string, bool] {
	_ = "STUB: not implemented"
	return nil
}

func stringStringMapFromMetadata(values []providerStringString) *orderedmap.Map[string, string] {
	_ = "STUB: not implemented"
	return nil
}

func stringListMapFromMetadata(values []providerStringList) *orderedmap.Map[string, []string] {
	_ = "STUB: not implemented"
	return nil
}

func extensionsFromMetadata(values []providerNamedYAMLNode) *orderedmap.Map[string, *yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

func yamlNodeSliceFromMetadata(values []*providerYAMLNode) []*yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

func yamlNodeFromMetadata(metadata *providerYAMLNode) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

func yamlKindFromMetadata(kind string) yaml.Kind { _ = "STUB: not implemented"; return *new(yaml.Kind) }

func floatFromMetadata(metadata *providerFloat) *float64 { _ = "STUB: not implemented"; return nil }

func intFromMetadata(metadata *providerInt) *int64 { _ = "STUB: not implemented"; return nil }

func boolFromMetadata(metadata *providerBool) *bool { _ = "STUB: not implemented"; return nil }
