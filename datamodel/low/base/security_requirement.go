// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"context"
	"sync"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// SecurityRequirement is a low-level representation of a Swagger / OpenAPI 3 SecurityRequirement object.
//
// SecurityRequirement lists the required security schemes to execute this operation. The object can have multiple
// security schemes declared in it which are all required (that is, there is a logical AND between the schemes).
//
// The name used for each property MUST correspond to a security scheme declared in the Security Definitions
//   - https://swagger.io/specification/v2/#securityDefinitionsObject
//   - https://swagger.io/specification/#security-requirement-object
type SecurityRequirement struct {
	Requirements             low.ValueReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[[]low.ValueReference[string]]]]
	KeyNode                  *yaml.Node
	RootNode                 *yaml.Node
	ContainsEmptyRequirement bool // if a requirement is empty (this means it's optional)
	index                    *index.SpecIndex
	context                  context.Context
	nodeStore                sync.Map
	reference                low.Reference
	*low.Reference
	low.NodeMap
}

// GetContext will return the context.Context instance used when building the SecurityRequirement object
func (s *SecurityRequirement) GetContext() context.Context {
	_ = "STUB: not implemented"

	// GetIndex will return the index.SpecIndex instance attached to the SecurityRequirement object
	return *new(context.Context)
}

func (s *SecurityRequirement) GetIndex() *index.SpecIndex {
	_ = "STUB: not implemented"

	// Build will extract security requirements from the node (the structure is odd, to be honest)
	return nil
}

func (s *SecurityRequirement) Build(ctx context.Context, keyNode, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// reset roles.

// GetRootNode will return the root yaml node of the SecurityRequirement object
func (s *SecurityRequirement) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode will return the key yaml node of the SecurityRequirement object
	return nil
}

func (s *SecurityRequirement) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// FindRequirement will attempt to locate a security requirement string from a supplied name.
	return nil
}

func (s *SecurityRequirement) FindRequirement(name string) []low.ValueReference[string] {
	_ = "STUB: not implemented"
	return nil
}

// GetKeys returns a string slice of all the keys used in the requirement.
func (s *SecurityRequirement) GetKeys() []string { _ = "STUB: not implemented"; return nil }

// Hash will return a consistent hash of the SecurityRequirement object
func (s *SecurityRequirement) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Pre-allocate vals slice
