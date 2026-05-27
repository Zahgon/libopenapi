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

// Workflow represents a low-level Arazzo Workflow Object.
// https://spec.openapis.org/arazzo/v1.0.1#workflow-object
type Workflow struct {
	WorkflowId     low.NodeReference[string]
	Summary        low.NodeReference[string]
	Description    low.NodeReference[string]
	Inputs         low.NodeReference[*yaml.Node]
	DependsOn      low.NodeReference[[]low.ValueReference[string]]
	Steps          low.NodeReference[[]low.ValueReference[*Step]]
	SuccessActions low.NodeReference[[]low.ValueReference[*SuccessAction]]
	FailureActions low.NodeReference[[]low.ValueReference[*FailureAction]]
	Outputs        low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[string]]]
	Parameters     low.NodeReference[[]low.ValueReference[*Parameter]]
	Extensions     *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode        *yaml.Node
	RootNode       *yaml.Node
	index          *index.SpecIndex
	context        context.Context
	*low.Reference
	low.NodeMap
}

var extractWorkflowSuccessActions = extractArray[SuccessAction]
var extractWorkflowParameters = extractArray[Parameter]

// GetIndex returns the index.SpecIndex instance attached to the Workflow object.
// For Arazzo low models this is typically nil, because Arazzo parsing does not build a SpecIndex.
// The index parameter is still required to satisfy the shared low.Buildable interface and generic extractors.
func (w *Workflow) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Workflow object.
	return nil
}

func (w *Workflow) GetContext() context.Context {
	_ = "STUB: not implemented"

	// FindExtension returns a ValueReference containing the extension value, if found.
	return *new(context.Context)
}

func (w *Workflow) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the Workflow object.
func (w *Workflow) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the Workflow object.
	return nil
}

func (w *Workflow) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract all properties of the Workflow object.
	return nil
}

func (w *Workflow) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// raw node: JSON Schema

// GetExtensions returns all Workflow extensions and satisfies the low.HasExtensions interface.
func (w *Workflow) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent hash of the Workflow object.
}

func (w *Workflow) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
