// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

import (
	"encoding/json"
	"reflect"

	highbase "github.com/pb33f/libopenapi/datamodel/high/base"
)

type SchemaProvider interface {
	OpenAPISchema() *highbase.SchemaProxy
}

type SchemaYAMLProvider interface {
	OpenAPISchemaYAML() string
}

var schemaProviderType = reflect.TypeOf((*SchemaProvider)(nil)).Elem()
var schemaMetadataProviderType = reflect.TypeOf((*SchemaMetadataProvider)(nil)).Elem()
var schemaYAMLProviderType = reflect.TypeOf((*SchemaYAMLProvider)(nil)).Elem()
var rawMessageType = reflect.TypeOf(json.RawMessage{})

func (g *Generator) irFromReflect(t reflect.Type, name, path string) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) irFromReflectName(t reflect.Type, name string, nameResolved bool, path string) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) irFromTypeSchema(t reflect.Type, name, path string, schema *highbase.SchemaProxy, nullable bool) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) irFromSchemaProvider(t reflect.Type, name, path string, nullable bool) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) irFromSchemaMetadataProvider(t reflect.Type, name, path string, nullable bool) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) irFromSchemaYAMLProvider(t reflect.Type, name, path string, nullable bool) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) irFromReflectArray(t reflect.Type, name, path string, nullable bool) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) irFromReflectMap(t reflect.Type, name, path string, nullable bool) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) irFromReflectStruct(t reflect.Type, name, path string, nullable bool) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneIR(ir *SchemaIR) *SchemaIR { _ = "STUB: not implemented"; return nil }

func (g *Generator) irFromReflectField(owner reflect.Type, field reflect.StructField, tag fieldTag, name, path string) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) fieldSchema(owner reflect.Type, field reflect.StructField, jsonName string) *highbase.SchemaProxy {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) irFromFieldSchema(fieldType reflect.Type, name, path string, schema *highbase.SchemaProxy) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) stringEncodedIR(ir *SchemaIR, path string) *SchemaIR {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) irFromReflectInterface(t reflect.Type, name, path string, nullable bool) (*SchemaIR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func implementsOrPointerImplements(t reflect.Type, iface reflect.Type) bool {
	_ = "STUB: not implemented"
	return false
}

func providerValue(t reflect.Type) any { _ = "STUB: not implemented"; return *new(any) }

func schemaProxyFromProviderYAML(name, schemaYAML string) (*highbase.SchemaProxy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func indentSchemaYAML(in, prefix string) string { _ = "STUB: not implemented"; return "" }
