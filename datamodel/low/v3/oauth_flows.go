// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"context"
	"sync"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// OAuthFlows represents a low-level OpenAPI 3+ OAuthFlows object.
//   - https://spec.openapis.org/oas/v3.1.0#oauth-flows-object
type OAuthFlows struct {
	Implicit          low.NodeReference[*OAuthFlow]
	Password          low.NodeReference[*OAuthFlow]
	ClientCredentials low.NodeReference[*OAuthFlow]
	AuthorizationCode low.NodeReference[*OAuthFlow]
	Device            low.NodeReference[*OAuthFlow] // OpenAPI 3.2+ device flow
	Extensions        *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode           *yaml.Node
	RootNode          *yaml.Node
	index             *index.SpecIndex
	context           context.Context
	nodeStore         sync.Map
	reference         low.Reference
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the OAuthFlows object.
func (o *OAuthFlows) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the OAuthFlows object.
	return nil
}

func (o *OAuthFlows) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetExtensions returns all OAuthFlows extensions and satisfies the low.HasExtensions interface.
	return *new(context.Context)
}

func (o *OAuthFlows) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// FindExtension will attempt to locate an extension with the supplied name.
}

func (o *OAuthFlows) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the OAuthFlows object.
func (o *OAuthFlows) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the OAuthFlows object.
	return nil
}

func (o *OAuthFlows) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract extensions and all OAuthFlow types from the supplied node.
	return nil
}

func (o *OAuthFlows) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Hash will return a consistent Hash of the OAuthFlows object
func (o *OAuthFlows) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// OAuthFlow represents a low-level OpenAPI 3+ OAuthFlow object.
//   - https://spec.openapis.org/oas/v3.1.0#oauth-flow-object
type OAuthFlow struct {
	AuthorizationUrl low.NodeReference[string]
	TokenUrl         low.NodeReference[string]
	RefreshUrl       low.NodeReference[string]
	Scopes           low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[string]]]
	Extensions       *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	RootNode         *yaml.Node
	index            *index.SpecIndex
	context          context.Context
	nodeStore        sync.Map
	reference        low.Reference
	*low.Reference
	low.NodeMap
}

// GetIndex returns the index.SpecIndex instance attached to the OAuthFlow object.
func (o *OAuthFlow) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// GetContext returns the context.Context instance used when building the OAuthFlow object.
	return nil
}

func (o *OAuthFlow) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetExtensions returns all OAuthFlow extensions and satisfies the low.HasExtensions interface.
	return *new(context.Context)
}

func (o *OAuthFlow) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// FindScope attempts to locate a scope using a specified name.
}

func (o *OAuthFlow) FindScope(scope string) *low.ValueReference[string] {
	_ = "STUB: not implemented"
	return nil
}

// FindExtension attempts to locate an extension with a specified key
func (o *OAuthFlow) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetRootNode returns the root yaml node of the OAuthFlow object.
func (o *OAuthFlow) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// Build will extract extensions from the node.
	return nil
}

func (o *OAuthFlow) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Hash will return a consistent Hash of the OAuthFlow object
func (o *OAuthFlow) Hash() uint64 { _ = "STUB: not implemented"; return 0 }
