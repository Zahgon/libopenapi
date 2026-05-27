// Copyright 2023-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"context"

	"go.yaml.in/yaml/v4"
)

type extractRefsState struct {
	ctx           context.Context
	scope         *SchemaIdScope
	parentBaseURI string
	seenPath      []string
	lastAppended  bool
	level         int
	poly          bool
	polyName      string
	prev          string
}

func (index *SpecIndex) initializeExtractRefsState(
	ctx context.Context,
	node *yaml.Node,
	seenPath []string,
	level int,
	poly bool,
	pName string,
) extractRefsState {
	_ = "STUB: not implemented"
	return *new(extractRefsState)
}

func (index *SpecIndex) walkExtractRefs(node, parent *yaml.Node, state *extractRefsState) []*Reference {
	_ = "STUB: not implemented"
	return nil
}

// In YAML mapping nodes, Content alternates key-value: even indices (0, 2, 4...)
// are keys, odd indices (1, 3, 5...) are values.

func (index *SpecIndex) walkChildExtractRefs(node, parent *yaml.Node, state *extractRefsState) []*Reference {
	_ = "STUB: not implemented"
	return nil
}

func isDirectOpenAPIExampleRefNode(node *yaml.Node) bool { _ = "STUB: not implemented"; return false }

func (index *SpecIndex) handleExtractRefsKey(
	node, parent *yaml.Node,
	state *extractRefsState,
	keyIndex int,
	found *[]*Reference,
) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldSkipMapSchemaCollection(seenPath []string) bool { _ = "STUB: not implemented"; return false }

func (index *SpecIndex) unwindExtractRefsPath(node *yaml.Node, state *extractRefsState, currentIndex int) {
	_ = "STUB: not implemented"
	return
}
