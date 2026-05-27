// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package renderer

import (
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"go.yaml.in/yaml/v4"
)

const mockSchemaCacheRoleDefault = "schema"

type mockSchemaKey struct {
	node   *yaml.Node
	ref    string
	schema *base.Schema
}

type mockSchemaCacheKey struct {
	schema mockSchemaKey
	role   string
}

type mockRenderContext struct {
	renderer       *SchemaRenderer
	options        MockGenerationOptions
	active         map[mockSchemaKey]int
	completed      map[mockSchemaCacheKey]any
	enforceBudgets bool
	nodes          int
	props          int
	refs           int
	bytes          int
	err            error
}

func newMockRenderContext(renderer *SchemaRenderer) *mockRenderContext {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *mockRenderContext) diveIntoSchema(schema *base.Schema, key string, structure map[string]any, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (ctx *mockRenderContext) renderString(schema *base.Schema, key string, structure map[string]any) bool {
	_ = "STUB: not implemented"
	return false
}

func (ctx *mockRenderContext) renderNumber(schema *base.Schema, key string, structure map[string]any) bool {
	_ = "STUB: not implemented"
	return false
}

func (ctx *mockRenderContext) renderObject(schema *base.Schema, key string, structure map[string]any, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (ctx *mockRenderContext) renderArray(schema *base.Schema, key string, structure map[string]any, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (ctx *mockRenderContext) isObjectSchema(schema *base.Schema) bool {
	_ = "STUB: not implemented"
	return false
}

func (ctx *mockRenderContext) enterSchema(schema *base.Schema, key string, structure map[string]any) (mockSchemaKey, mockSchemaCacheKey, bool, bool, bool) {
	_ = "STUB: not implemented"
	return *new(mockSchemaKey), *new(mockSchemaCacheKey), false, false, false
}

func (ctx *mockRenderContext) leaveSchema(schemaKey mockSchemaKey, cacheKey mockSchemaCacheKey, hasSchemaKey bool, value any) {
	_ = "STUB: not implemented"
	return
}

func mockCacheRole(key string) string { _ = "STUB: not implemented"; return "" }

func (ctx *mockRenderContext) schemaKey(schema *base.Schema) (mockSchemaKey, bool) {
	_ = "STUB: not implemented"
	return *new(mockSchemaKey), false
}

func (ctx *mockRenderContext) noteNode() bool { _ = "STUB: not implemented"; return false }

func (ctx *mockRenderContext) noteProperty(name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (ctx *mockRenderContext) noteRefExpansion() bool { _ = "STUB: not implemented"; return false }

func (ctx *mockRenderContext) noteValue(value any) bool { _ = "STUB: not implemented"; return false }

func (ctx *mockRenderContext) checkBudget(name string, limit int, actual int) bool {
	_ = "STUB: not implemented"
	return false
}

func copyMockValue(value any) any { _ = "STUB: not implemented"; return *new(any) }

func estimatedMockValueBytes(value any) int { _ = "STUB: not implemented"; return 0 }
