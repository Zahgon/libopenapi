// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

import (
	"reflect"

	highbase "github.com/pb33f/libopenapi/datamodel/high/base"
)

// Option configures a Generator.
type Option func(*Generator)

// NameResolver maps OpenAPI names to Go identifiers. Returning an empty string
// falls back to the generator's default naming.
type NameResolver func(string) string

// ExternalRefResolver maps an external OpenAPI $ref to a Go type name.
// Returning an empty string falls back to deriving the type name from the
// reference tail.
type ExternalRefResolver func(ref string) string

// Diagnostic describes a notable generator decision.
type Diagnostic struct {
	Code    string
	Path    string
	Message string
}

const (
	DiagnosticComponentNameCollision     = "componentNameCollision"
	DiagnosticChildSchema                = "childSchema"
	DiagnosticAdditionalPropertiesFalse  = "additionalPropertiesFalse"
	DiagnosticArrayContains              = "arrayContains"
	DiagnosticBooleanItems               = "booleanItems"
	DiagnosticConstKeyword               = "constKeyword"
	DiagnosticContentSchema              = "contentSchema"
	DiagnosticDependentRequired          = "dependentRequired"
	DiagnosticDependentSchemas           = "dependentSchemas"
	DiagnosticDynamicReference           = "dynamicReference"
	DiagnosticExternalReference          = "externalReference"
	DiagnosticFieldNameCollision         = "fieldNameCollision"
	DiagnosticConditionalSchema          = "conditionalSchema"
	DiagnosticImplicitType               = "implicitType"
	DiagnosticMixedEnum                  = "mixedEnum"
	DiagnosticMultiTypeSchema            = "multiTypeSchema"
	DiagnosticNotSchema                  = "notSchema"
	DiagnosticNullEnum                   = "nullEnum"
	DiagnosticOptionalConstDiscriminator = "optionalConstDiscriminator"
	DiagnosticPatternProperties          = "patternProperties"
	DiagnosticPrefixItems                = "prefixItems"
	DiagnosticPropertyNames              = "propertyNames"
	DiagnosticSchemaMetadata             = "schemaMetadata"
	DiagnosticStringEncoded              = "stringEncoded"
	DiagnosticTypeNameCollision          = "typeNameCollision"
	DiagnosticUnevaluatedItems           = "unevaluatedItems"
	DiagnosticRootNameCollision          = "rootNameCollision"
	DiagnosticUnevaluatedProperties      = "unevaluatedProperties"
	DiagnosticValidationKeyword          = "validationKeyword"
)

type formatMapping struct {
	goType     string
	importPath string
}

type discriminatorRegistration struct {
	property string
	mapping  map[string]string
}

type fieldSchemaKey struct {
	owner reflect.Type
	name  string
}

// WithPackageName sets the generated Go package name.
func WithPackageName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOptionalFieldsAsPointers controls whether optional scalar fields render
// as pointers.
func WithOptionalFieldsAsPointers(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithOmitEmpty controls omitempty on optional generated tags.
func WithOmitEmpty(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNullableAsPointer controls whether nullable scalar fields render as
// pointers.
func WithNullableAsPointer(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGenerateJSONTags controls generated json tags.
func WithGenerateJSONTags(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGenerateYAMLTags controls generated yaml tags.
func WithGenerateYAMLTags(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnumConstants controls whether enum values generate Go constants.
func WithEnumConstants(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHeaderComment writes a file header comment before the package clause.
func WithHeaderComment(text string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPackageComment writes a package doc comment before the package clause.
func WithPackageComment(text string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGeneratedComment writes a standard generated-code comment.
func WithGeneratedComment(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOpenAPITags controls whether generated struct fields include compact
// openapi tags for metadata that cannot be recovered from Go reflection alone.
func WithOpenAPITags(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSchemaMetadataSidecar controls whether generated named types include a
// typed OpenAPISchemaMetadata sidecar. Enabling the sidecar preserves original
// OpenAPI schema fidelity for Go reflection round trips. Disabling it keeps the
// generated model code leaner, but OpenAPI -> Go -> OpenAPI reconstruction is
// intentionally lossy and falls back to Go type shape plus tags.
func WithSchemaMetadataSidecar(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithFormatMapping maps an OpenAPI string format to a Go type and optional
// import path.
func WithFormatMapping(format, goType, importPath string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithNameResolver sets a broad fallback resolver for generated Go names.
func WithNameResolver(resolver NameResolver) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTypeNameResolver sets a resolver for generated Go type names.
func WithTypeNameResolver(resolver NameResolver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFieldNameResolver sets a resolver for generated Go struct field names.
func WithFieldNameResolver(resolver NameResolver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEnumValueNameResolver sets a resolver for generated enum constant suffixes.
func WithEnumValueNameResolver(resolver NameResolver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithOptionalConstDiscriminatorUnions allows optional shared const
// discriminator properties to produce typed oneOf unions.
func WithOptionalConstDiscriminatorUnions(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAdditionalPropertiesMethods controls whether schema-valued
// additionalProperties generates JSON marshal/unmarshal methods that round-trip
// unknown fields through the AdditionalProperties map.
func WithAdditionalPropertiesMethods(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithNestedTypeNameDelimiter sets the separator inserted between generated
// parent and child type names for inline schemas. The default is "_"; passing
// an empty delimiter restores compact names like ParentChild.
func WithNestedTypeNameDelimiter(delimiter string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithExternalRefTypeResolver sets a resolver for external OpenAPI $ref values
// when rendering Go type names. The resolver is not used for local component
// references.
func WithExternalRefTypeResolver(resolver ExternalRefResolver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTypeSchema overrides reflected schema generation for a specific Go type.
// This is useful for project scalar aliases that need a custom OpenAPI format,
// enum, or extension without implementing SchemaProvider on the type.
func WithTypeSchema(t reflect.Type, schema *highbase.SchemaProxy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFieldSchema overrides reflected schema generation for a specific Go
// struct field name while keeping the surrounding model reflected normally.
func WithFieldSchema(t reflect.Type, fieldName string, schema *highbase.SchemaProxy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFieldSchemaByJSONName overrides reflected schema generation for a
// specific JSON field name while keeping the surrounding model reflected
// normally.
func WithFieldSchemaByJSONName(t reflect.Type, jsonName string, schema *highbase.SchemaProxy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithOneOfTypes registers concrete variants for a Go interface when producing
// OpenAPI oneOf schemas from reflection.
func WithOneOfTypes(target any, variants ...any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDiscriminatorMapping registers discriminator metadata for a reflected
// interface union.
func WithDiscriminatorMapping(target any, property string, mapping map[string]string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
