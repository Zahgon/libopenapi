// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Server represents a low-level OpenAPI 3+ Server object.
//   - https://spec.openapis.org/oas/v3.1.0#server-object
type Server struct {
	Name        low.NodeReference[string] // OpenAPI 3.2+ name field for documentation
	URL         low.NodeReference[string]
	Description low.NodeReference[string]
	Variables   low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*ServerVariable]]]
	Extensions  *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode     *yaml.Node
	RootNode    *yaml.Node
	index       *index.SpecIndex
	context     context.Context
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the Server object.
func (s *Server) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the Server object.
	return nil
}

func (s *Server) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetRootNode returns the root yaml node of the Server object.
	return *new(context.Context)
}

func (s *Server) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetExtensions returns all Paths extensions and satisfies the low.HasExtensions interface.
	return nil
}

func (s *Server) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// FindVariable attempts to locate a ServerVariable instance using the supplied key.
}

func (s *Server) FindVariable(serverVar string) *low.ValueReference[*ServerVariable] {
	_ = "STUB: not implemented"
	return nil
}

// Build will extract server variables from the supplied node.
func (s *Server) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Hash will return a consistent Hash of the Server object
func (s *Server) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
