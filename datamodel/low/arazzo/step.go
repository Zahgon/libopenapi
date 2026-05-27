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

// Step represents a low-level Arazzo Step Object.
// https://spec.openapis.org/arazzo/v1.0.1#step-object
type Step struct {
	StepId          low.NodeReference[string]
	Description     low.NodeReference[string]
	OperationId     low.NodeReference[string]
	OperationPath   low.NodeReference[string]
	WorkflowId      low.NodeReference[string]
	Parameters      low.NodeReference[[]low.ValueReference[*Parameter]]
	RequestBody     low.NodeReference[*RequestBody]
	SuccessCriteria low.NodeReference[[]low.ValueReference[*Criterion]]
	OnSuccess       low.NodeReference[[]low.ValueReference[*SuccessAction]]
	OnFailure       low.NodeReference[[]low.ValueReference[*FailureAction]]
	Outputs         low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[string]]]
	Extensions      *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode         *yaml.Node
	RootNode        *yaml.Node
	index           *index.SpecIndex
	context         context.Context
	*low.Reference
	low.NodeMap
}

var extractStepParameters = extractArray[Parameter]
var extractStepSuccessCriteria = extractArray[Criterion]
var extractStepOnSuccess = extractArray[SuccessAction]

// GetIndex returns the index.SpecIndex instance attached to the Step object.
// For Arazzo low models this is typically nil, because Arazzo parsing does not build a SpecIndex.
// The index parameter is still required to satisfy the shared low.Buildable interface and generic extractors.
func (s *Step) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Step object.
	return nil
}

func (s *Step) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (s *Step) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the Step object.
func (s *Step) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Step object.
	return nil
}

func (s *Step) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract all properties of the Step object.
	return nil
}

func (s *Step) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all Step extensions and satisfies the low.HasExtensions interface.
func (s *Step) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent hash of the Step object.
}

func (s *Step) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
