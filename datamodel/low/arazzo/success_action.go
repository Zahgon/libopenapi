// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// SuccessAction represents a low-level Arazzo Success Action Object.
// A success action can be a full definition or a Reusable Object with a $components reference.
// https://spec.openapis.org/arazzo/v1.0.1#success-action-object
type SuccessAction struct {
	Name         low.NodeReference[string]
	Type         low.NodeReference[string]
	WorkflowId   low.NodeReference[string]
	StepId       low.NodeReference[string]
	Criteria     low.NodeReference[[]low.ValueReference[*Criterion]]
	ComponentRef low.NodeReference[string]
	Extensions   *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode      *yaml.Node
	RootNode     *yaml.Node
	index        *index.SpecIndex
	context      context.Context
	*low.Reference
	low.NodeMap
}

var extractSuccessActionCriteria = extractArray[Criterion]

// IsReusable returns true if this success action is a Reusable Object (has a reference field).
func (s *SuccessAction) IsReusable() bool { _ = "STUB: not implemented"; return false }

// GetIndex returns the index.SpecIndex instance attached to the SuccessAction object.
// For Arazzo low models this is typically nil, because Arazzo parsing does not build a SpecIndex.
// The index parameter is still required to satisfy the shared low.Buildable interface and generic extractors.
func (s *SuccessAction) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the SuccessAction object.
	return nil
}

func (s *SuccessAction) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (s *SuccessAction) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the SuccessAction object.
func (s *SuccessAction) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the SuccessAction object.
	return nil
}

func (s *SuccessAction) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract all properties of the SuccessAction object.
	return nil
}

func (s *SuccessAction) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Extract criteria array

// GetExtensions returns all SuccessAction extensions and satisfies the low.HasExtensions interface.
func (s *SuccessAction) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent hash of the SuccessAction object.
}

func (s *SuccessAction) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
