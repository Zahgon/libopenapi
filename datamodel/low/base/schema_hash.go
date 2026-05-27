// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"strings"
	"sync"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Hash will generate a stable hash of the SchemaDynamicValue
func (s *SchemaDynamicValue[A, B]) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// SchemaQuickHashMap is a sync.Map used to store quick hashes of schemas, used by quick hashing to prevent
// over rotation on the same schema. This map is automatically reset each time `CompareDocuments` is called by the
// `what-changed` package and each time a model is built via `BuildV3Model()` etc.
//
// This exists because to ensure deep equality checking when composing schemas using references. However this
// can cause an exhaustive deep hash calculation that chews up compute like crazy, particularly with polymorphic refs.
// The hash map means each schema is hashed once, and then the hash is reused for quick equality checking.
var SchemaQuickHashMap sync.Map

// ClearSchemaQuickHashMap resets the schema quick-hash cache.
// Call this between document lifecycles in long-running processes to bound memory.
func ClearSchemaQuickHashMap() { _ = "STUB: not implemented"; return }

// QuickHash will calculate a hash from the values of the schema, however the hash is not very deep
// and is used for quick equality checking, This method exists because a full hash could end up churning through
// thousands of polymorphic references. With a quick hash, polymorphic properties are not included.
func (s *Schema) QuickHash() uint64 { _ = "STUB: not implemented"; return 0 }

// Hash will calculate a hash from the values of the schema, This allows equality checking against
// Schemas defined inside an OpenAPI document. The only way to know if a schema has changed, is to hash it.
func (s *Schema) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Schema) hash(quick bool) uint64 { _ = "STUB: not implemented"; return 0 }

// Use string builder pool for efficient string concatenation

// calculate a hash from every property in the schema.

func writeSchemaMapHashes[V any](sb *strings.Builder, m *orderedmap.Map[low.KeyReference[string], low.ValueReference[V]]) {
	_ = "STUB: not implemented"
	return
}

func resizeSchemaHashScratch(scratch []string, size int) []string {
	_ = "STUB: not implemented"
	return nil
}

func writeSortedSchemaStrings(sb *strings.Builder, values []string, separate bool) {
	_ = "STUB: not implemented"
	return
}

func writeSchemaBoolMap(sb *strings.Builder, m *orderedmap.Map[low.KeyReference[string], low.ValueReference[bool]]) {
	_ = "STUB: not implemented"
	return
}

func writeSchemaDependentRequired(sb *strings.Builder, m *orderedmap.Map[low.KeyReference[string], low.ValueReference[[]string]]) {
	_ = "STUB: not implemented"
	return
}

func writeSchemaExtensions(sb *strings.Builder, ext *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]) {
	_ = "STUB: not implemented"
	return
}

func (s *Schema) quickHashKey() string { _ = "STUB: not implemented"; return "" }
