// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

import (
	"strings"

	highbase "github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

func (g *Generator) recordSchemaMetadata(typeName string, schema *highbase.Schema) {
	_ = "STUB: not implemented"
	return
}

func (g *Generator) renderSchemaMetadataSidecarDecl() string { _ = "STUB: not implemented"; return "" }

func writeSchemaMetadataTypes(b *strings.Builder) { _ = "STUB: not implemented"; return }

func (g *Generator) schemaMetadataLiteral(schema *highbase.Schema, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) schemaMetadataLiteralWithRef(ref string, schema *highbase.Schema, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) schemaProxyMetadataLiteral(proxy *highbase.SchemaProxy, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func referenceSiblingMetadataSchema(proxy *highbase.SchemaProxy) *highbase.Schema {
	_ = "STUB: not implemented"
	return nil
}

func schemaFromReferenceSiblingNode(refNode *yaml.Node) *highbase.Schema {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) optionalSchemaProxyMetadataLiteral(proxy *highbase.SchemaProxy, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) schemaSliceMetadataLiteral(schemas []*highbase.SchemaProxy, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) schemaMapMetadataLiteral(schemas *orderedmap.Map[string, *highbase.SchemaProxy], depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) metadataDynamicSchemaBoolLiteral(value *highbase.DynamicValue[*highbase.SchemaProxy, bool], depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataDynamicBoolNumberLiteral(value *highbase.DynamicValue[bool, float64]) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataDiscriminatorLiteral(discriminator *highbase.Discriminator, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataStringStringMapLiteral(values *orderedmap.Map[string, string], depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataStringBoolMapLiteral(values *orderedmap.Map[string, bool], depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataStringListMapLiteral(values *orderedmap.Map[string, []string], depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataExtensionsLiteral(values *orderedmap.Map[string, *yaml.Node], depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataYAMLNodeSliceLiteral(nodes []*yaml.Node, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataYAMLNodeLiteral(node *yaml.Node, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func optionalMetadataYAMLNodeLiteral(node *yaml.Node, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataYAMLNodeContentLiteral(nodes []*yaml.Node, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataYAMLKind(kind yaml.Kind) string { _ = "STUB: not implemented"; return "" }

func metadataStringSliceLiteral(values []string, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataStringLiteral(value string) string { _ = "STUB: not implemented"; return "" }

func metadataFloatLiteral(value *float64) string { _ = "STUB: not implemented"; return "" }

func metadataFloatValueLiteral(value float64) string { _ = "STUB: not implemented"; return "" }

func metadataIntLiteral(value *int64) string { _ = "STUB: not implemented"; return "" }

func metadataIntValueLiteral(value int64) string { _ = "STUB: not implemented"; return "" }

func metadataPlainIntLiteral(value int64) string { _ = "STUB: not implemented"; return "" }

func metadataBoolLiteral(value *bool) string { _ = "STUB: not implemented"; return "" }

func metadataBoolValueLiteral(value bool) string { _ = "STUB: not implemented"; return "" }

func writeMetadataField(b *strings.Builder, depth int, name, value string) {
	_ = "STUB: not implemented"
	return
}

func metadataIndent(depth int) string { _ = "STUB: not implemented"; return "" }
