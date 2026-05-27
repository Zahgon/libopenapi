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

// FailureAction represents a low-level Arazzo Failure Action Object.
// A failure action can be a full definition or a Reusable Object with a $components reference.
// https://spec.openapis.org/arazzo/v1.0.1#failure-action-object
type FailureAction struct {
	Name         low.NodeReference[string]
	Type         low.NodeReference[string]
	WorkflowId   low.NodeReference[string]
	StepId       low.NodeReference[string]
	RetryAfter   low.NodeReference[float64]
	RetryLimit   low.NodeReference[int64]
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

var extractFailureActionCriteria = extractArray[Criterion]

// IsReusable returns true if this failure action is a Reusable Object (has a reference field).
func (f *FailureAction) IsReusable() bool { _ = "STUB: not implemented"; return false }

// GetIndex returns the index.SpecIndex instance attached to the FailureAction object.
// For Arazzo low models this is typically nil, because Arazzo parsing does not build a SpecIndex.
// The index parameter is still required to satisfy the shared low.Buildable interface and generic extractors.
func (f *FailureAction) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the FailureAction object.
	return nil
}

func (f *FailureAction) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (f *FailureAction) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the FailureAction object.
func (f *FailureAction) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the FailureAction object.
	return nil
}

func (f *FailureAction) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract all properties of the FailureAction object.
	return nil
}

func (f *FailureAction) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Extract numeric fields (retryAfter, retryLimit) which need special parsing

// Extract criteria array

// GetExtensions returns all FailureAction extensions and satisfies the low.HasExtensions interface.
func (f *FailureAction) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent hash of the FailureAction object.
}

func (f *FailureAction) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
