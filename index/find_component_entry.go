// Copyright 2023-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"context"

	"go.yaml.in/yaml/v4"
)

// FindComponent locates a component in the index by reference.
//
// It resolves local references directly from the current document first, then recurses through
// rolodex-backed file and remote references as needed. It returns nil when the target cannot be found.
func (index *SpecIndex) FindComponent(ctx context.Context, componentId string) *Reference {
	_ = "STUB: not implemented"
	return nil
}

// FindComponent locates a component within a specific root YAML node.
//
// The lookup prefers direct fragment navigation and direct component maps first, and falls back to
// JSONPath traversal for legacy or non-direct component identifiers.
func FindComponent(_ context.Context, root *yaml.Node, componentID, absoluteFilePath string, index *SpecIndex) *Reference {
	_ = "STUB: not implemented"
	return nil
}

// FindComponentInRoot locates a component reference in the current root document only.
//
// It normalizes file-prefixed local references back to root-document fragments before delegating
// to FindComponent.
func (index *SpecIndex) FindComponentInRoot(ctx context.Context, componentID string) *Reference {
	_ = "STUB: not implemented"
	return nil
}
