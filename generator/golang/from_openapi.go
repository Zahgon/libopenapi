// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

import (
	highbase "github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/orderedmap"
)

func (g *Generator) irFromOpenAPI(name string, proxy *highbase.SchemaProxy, path string) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) irFromOpenAPIName(name string, nameResolved bool, proxy *highbase.SchemaProxy, path string) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// childIR builds a nested schema. If the nested schema cannot be built it
// records a diagnostic and falls back to an any-typed shape so the surrounding
// field, item, or variant is preserved rather than silently dropped.
func (g *Generator) childIR(name string, proxy *highbase.SchemaProxy, path string) *SchemaIR {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) irFromSchema(name string, nameResolved bool, schema *highbase.Schema, path string) *SchemaIR {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) openapiSchemaTypeName(name string, nameResolved bool, schema *highbase.Schema, path string) string {
	_ = "STUB: not implemented"
	return ""
}

func schemaDeclaresType(schema *highbase.Schema) bool { _ = "STUB: not implemented"; return false }

func (g *Generator) populateSchemaShape(ir *SchemaIR, schema *highbase.Schema, path string) {
	_ = "STUB: not implemented"
	return
}

func (g *Generator) populateMultiTypeUnion(ir *SchemaIR, types []string, path string) {
	_ = "STUB: not implemented"
	return
}

func (g *Generator) populateObject(ir *SchemaIR, schema *highbase.Schema, path string) {
	_ = "STUB: not implemented"
	return
}

func (g *Generator) collectShapeDiagnostics(path string, schema *highbase.Schema) {
	_ = "STUB: not implemented"
	return
}

func (g *Generator) populateUnion(ir *SchemaIR, schema *highbase.Schema, path string) {
	_ = "STUB: not implemented"
	return
}

func (g *Generator) mergeAllOf(ir *SchemaIR) { _ = "STUB: not implemented"; return }

func orderedProperties() *orderedmap.Map[string, *SchemaIR] { _ = "STUB: not implemented"; return nil }

func nonNullTypes(types []string) []string { _ = "STUB: not implemented"; return nil }

func primaryTypeForSchema(schema *highbase.Schema) (string, bool, bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

func kindForJSONType(typ string) Kind { _ = "STUB: not implemented"; return *new(Kind) }

func schemaHasOnlyDynamicRefShape(schema *highbase.Schema) bool {
	_ = "STUB: not implemented"
	return false
}

func hasSchemaMetadata(schema *highbase.Schema) bool { _ = "STUB: not implemented"; return false }

func hasValidationKeyword(schema *highbase.Schema) bool { _ = "STUB: not implemented"; return false }

func hasStringKeyword(schema *highbase.Schema) bool { _ = "STUB: not implemented"; return false }

func hasNumberKeyword(schema *highbase.Schema) bool { _ = "STUB: not implemented"; return false }

func hasArrayKeyword(schema *highbase.Schema) bool { _ = "STUB: not implemented"; return false }

func hasObjectKeyword(schema *highbase.Schema) bool { _ = "STUB: not implemented"; return false }

func nonNullVariants(variants []*SchemaIR) []*SchemaIR { _ = "STUB: not implemented"; return nil }

func isNullOnlyIR(ir *SchemaIR) bool { _ = "STUB: not implemented"; return false }

func schemaOnlyAllowsNull(schema *highbase.Schema) bool { _ = "STUB: not implemented"; return false }

func discriminatorFromSchema(schema *highbase.Schema, variants []*SchemaIR) *Discriminator {
	_ = "STUB: not implemented"
	return nil
}

func inferConstDiscriminator(variants []*SchemaIR) *Discriminator {
	_ = "STUB: not implemented"
	return nil
}
