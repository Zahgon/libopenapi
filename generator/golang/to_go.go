// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

import (
	"go/format"
	"strings"
)

var formatSource = format.Source

func (g *Generator) renderFile(irs []*SchemaIR) (*GeneratedFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) renderSchemaMetadataSource() (*GeneratedSourceFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) writeImports(b *strings.Builder) { _ = "STUB: not implemented"; return }

func (g *Generator) renderDecl(ir *SchemaIR) { _ = "STUB: not implemented"; return }

func (g *Generator) rememberDecl(name string) bool { _ = "STUB: not implemented"; return false }

func (g *Generator) renderObjectDecl(ir *SchemaIR) { _ = "STUB: not implemented"; return }

func (g *Generator) renderChildren(ir *SchemaIR) { _ = "STUB: not implemented"; return }

func (g *Generator) renderNested(ir *SchemaIR) { _ = "STUB: not implemented"; return }

func (g *Generator) renderAliasDecl(ir *SchemaIR) { _ = "STUB: not implemented"; return }

func shouldRenderObjectAlias(ir *SchemaIR) bool { _ = "STUB: not implemented"; return false }

func (g *Generator) renderEnumDecl(ir *SchemaIR) { _ = "STUB: not implemented"; return }

func writeAdditionalPropertiesMethods(b *strings.Builder, ir *SchemaIR, fieldName, valueType string) {
	_ = "STUB: not implemented"
	return
}

func (g *Generator) goType(ir *SchemaIR, required bool, field bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) formatType(format, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

func shouldPointer(typ string, ir *SchemaIR, required, optionalPointers, nullablePointer bool) bool {
	_ = "STUB: not implemented"
	return false
}

func writeComment(b *strings.Builder, name, text string) { _ = "STUB: not implemented"; return }

func writeIRComments(b *strings.Builder, ir *SchemaIR) { _ = "STUB: not implemented"; return }

func writeFieldComments(b *strings.Builder, fieldName string, ir *SchemaIR) {
	_ = "STUB: not implemented"
	return
}

func writeLineCommentBlock(b *strings.Builder, text string) { _ = "STUB: not implemented"; return }

func writeLineComment(b *strings.Builder, line string) { _ = "STUB: not implemented"; return }
