// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

import (
	highbase "github.com/pb33f/libopenapi/datamodel/high/base"
	"go.yaml.in/yaml/v4"
)

func (g *Generator) openapiFromIR(ir *SchemaIR) *highbase.SchemaProxy {
	_ = "STUB: not implemented"
	return nil
}

func applyIRBooleans(schema *highbase.Schema, ir *SchemaIR) { _ = "STUB: not implemented"; return }

func (g *Generator) nullableRefProxy(ir *SchemaIR) *highbase.SchemaProxy {
	_ = "STUB: not implemented"
	return nil
}

func nullableReferenceProxy(target string, dynamic bool, ir *SchemaIR) *highbase.SchemaProxy {
	_ = "STUB: not implemented"
	return nil
}

func referenceProxy(target string, ir *SchemaIR) *highbase.SchemaProxy {
	_ = "STUB: not implemented"
	return nil
}

func refSiblingSchema(ir *SchemaIR) *highbase.Schema { _ = "STUB: not implemented"; return nil }

func (g *Generator) populateOpenAPIObject(schema *highbase.Schema, ir *SchemaIR) {
	_ = "STUB: not implemented"
	return
}

func (g *Generator) populateOpenAPIUnion(schema *highbase.Schema, ir *SchemaIR) {
	_ = "STUB: not implemented"
	return
}

func applySchemaFidelity(schema *highbase.Schema, ir *SchemaIR) { _ = "STUB: not implemented"; return }

func applyNativeNullability(schema *highbase.Schema, ir *SchemaIR) {
	_ = "STUB: not implemented"
	return
}

func schemaNeedsNullAlternative(schema *highbase.Schema) bool {
	_ = "STUB: not implemented"
	return false
}

func stringNode(value string) *yaml.Node { _ = "STUB: not implemented"; return nil }

func nullNode() *yaml.Node { _ = "STUB: not implemented"; return nil }

func schemaTypeContains(values []string, target string) bool {
	_ = "STUB: not implemented"
	return false
}
