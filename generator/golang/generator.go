// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

import (
	"reflect"

	highbase "github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/orderedmap"
)

// Generator holds immutable configuration for code generation. Each public
// entry point runs against a fresh copy of this configuration (see run), so a
// configured Generator carries no per-invocation state and is safe to reuse for
// many documents and to share across goroutines.
type Generator struct {
	packageName string

	optionalFieldsAsPointers         bool
	omitEmpty                        bool
	nullableAsPointer                bool
	jsonTags                         bool
	yamlTags                         bool
	enumConstants                    bool
	optionalConstDiscriminatorUnions bool
	additionalPropertiesMethods      bool
	generatedComment                 bool
	openapiTags                      bool
	schemaMetadataSidecar            bool
	nestedTypeNameDelimiter          string

	nameResolver          NameResolver
	typeNameResolver      NameResolver
	fieldNameResolver     NameResolver
	enumValueNameResolver NameResolver
	externalRefResolver   ExternalRefResolver
	headerComment         string
	packageComment        string

	formatMappings map[string]formatMapping
	typeSchemas    map[reflect.Type]*highbase.SchemaProxy
	fieldSchemas   map[fieldSchemaKey]*highbase.SchemaProxy
	jsonSchemas    map[fieldSchemaKey]*highbase.SchemaProxy

	diagnostics     []Diagnostic
	imports         map[string]struct{}
	decls           []string
	seenDecls       map[string]struct{}
	metadataSchemas map[string]*highbase.Schema
	metadataOrder   []string

	openapiCache map[*highbase.SchemaProxy]*SchemaIR
	reflectCache map[reflect.Type]*SchemaIR
	reflectStack map[reflect.Type]bool
	typeNames    *nameRegistry

	componentNames     map[string]struct{}
	componentTypeNames map[string]string
	componentKinds     map[string]Kind
	currentComponent   string

	oneOfRegistrations         map[reflect.Type][]reflect.Type
	discriminatorRegistrations map[reflect.Type]discriminatorRegistration
}

// SchemaSet contains OpenAPI schemas generated from one or more Go types.
type SchemaSet struct {
	// Root is the first generated root schema, kept as a convenience for
	// single-root callers.
	Root *highbase.SchemaProxy
	// Roots contains every requested root schema keyed by generated type name.
	Roots *orderedmap.Map[string, *highbase.SchemaProxy]
	// Components contains reusable schemas discovered while walking the root
	// graph.
	Components *orderedmap.Map[string, *highbase.SchemaProxy]
	// Diagnostics reports schema features that required a lossy or notable
	// model-generation decision.
	Diagnostics []Diagnostic
}

const SchemaMetadataFileName = "schema_metadata.go"

// GeneratedFile contains Go source generated from OpenAPI schemas.
type GeneratedFile struct {
	PackageName    string
	Source         []byte
	SchemaMetadata *GeneratedSourceFile
	Types          []*GeneratedType
	Diagnostics    []Diagnostic
}

// GeneratedSourceFile contains a named generated source file.
type GeneratedSourceFile struct {
	Name   string
	Source []byte
}

// GeneratedType describes one top-level generated Go type.
type GeneratedType struct {
	Name string
	Kind Kind
}

// NewGenerator creates a Go model generator.
func NewGenerator(opts ...Option) *Generator { _ = "STUB: not implemented"; return nil }

// run returns a generator carrying fresh per-invocation state. Configuration is
// shared with the receiver and treated as read-only during generation, so a
// configured Generator is safe to reuse across calls and across goroutines.
// renderFile owns the rendering output buffers (imports, decls, metadata), so
// they are reset there rather than duplicated here.
func (g *Generator) run() *Generator { _ = "STUB: not implemented"; return nil }

// RenderSchema renders a single OpenAPI schema as Go source.
func RenderSchema(name string, schema *highbase.SchemaProxy, opts ...Option) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SchemaFromValue generates an OpenAPI schema for the runtime type of value.
func SchemaFromValue(value any, opts ...Option) (*highbase.SchemaProxy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SchemaFromType generates an OpenAPI schema for a Go reflection type.
func SchemaFromType(t reflect.Type, opts ...Option) (*highbase.SchemaProxy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SchemasFromValues generates an OpenAPI component graph for runtime values.
func SchemasFromValues(values ...any) (*SchemaSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SchemasFromValuesWithOptions generates an OpenAPI component graph for runtime
// values using generator options.
func SchemasFromValuesWithOptions(values []any, opts ...Option) (*SchemaSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SchemasFromTypes generates an OpenAPI component graph for Go reflection
// types.
func SchemasFromTypes(types ...reflect.Type) (*SchemaSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SchemasFromTypesWithOptions generates an OpenAPI component graph for Go
// reflection types using generator options.
func SchemasFromTypesWithOptions(types []reflect.Type, opts ...Option) (*SchemaSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RenderSchema renders a single OpenAPI schema as Go source using this
// generator.
func (g *Generator) RenderSchema(name string, schema *highbase.SchemaProxy) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RenderSchemas renders an ordered map of OpenAPI schemas as one Go source
// file.
func (g *Generator) RenderSchemas(schemas *orderedmap.Map[string, *highbase.SchemaProxy]) (*GeneratedFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) resolveComponentTypeNames(schemas *orderedmap.Map[string, *highbase.SchemaProxy]) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) resolveTypeName(original, candidate, path string) string {
	_ = "STUB: not implemented"
	return ""
}

// SchemaFromValue generates an OpenAPI schema for the runtime type of value
// using this generator.
func (g *Generator) SchemaFromValue(value any) (*highbase.SchemaProxy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SchemaFromType generates an OpenAPI schema for a Go reflection type using
// this generator.
func (g *Generator) SchemaFromType(t reflect.Type) (*highbase.SchemaProxy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SchemasFromValues generates an OpenAPI component graph for runtime values
// using this generator.
func (g *Generator) SchemasFromValues(values ...any) (*SchemaSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SchemasFromTypes generates an OpenAPI component graph for Go reflection types
// using this generator.
func (g *Generator) SchemasFromTypes(types ...reflect.Type) (*SchemaSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) rootProxy(ir *SchemaIR) *highbase.SchemaProxy {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) addDiagnostic(code, path, message string) { _ = "STUB: not implemented"; return }

func (g *Generator) addImport(path string) { _ = "STUB: not implemented"; return }

func isComponentKind(kind Kind) bool { _ = "STUB: not implemented"; return false }

func sortIRsByName(irs []*SchemaIR) { _ = "STUB: not implemented"; return }
