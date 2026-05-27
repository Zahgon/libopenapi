// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// operationResolver maps source descriptions to attached OpenAPI documents
// and provides operation lookup capabilities. This separates the semantic
// operation lookup concern from the structural validation in the validator.
type operationResolver struct {
	sourceDocs  map[string]*v3high.Document
	sourceOrder []string
	searchDocs  []*v3high.Document
}

// findOperationByID returns true if operationID exists in any attached OpenAPI document.
func (r *operationResolver) findOperationByID(operationID string) bool {
	_ = "STUB: not implemented"
	return false
}

// docForSource returns the OpenAPI document mapped to the given source name, or nil.
func (r *operationResolver) docForSource(sourceName string) *v3high.Document {
	_ = "STUB: not implemented"
	return nil
}

// defaultDoc returns the first available OpenAPI document for fallback lookups.
func (r *operationResolver) defaultDoc() *v3high.Document { _ = "STUB: not implemented"; return nil }
