// Copyright 2023-2024 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"go.yaml.in/yaml/v4"
)

// FindNodeOriginWithValue searches all indexes for the origin of a node with a specific value. If the node is found, a NodeOrigin
// is returned, otherwise nil is returned. The key and the value have to be provided. If the refNode and refValue are provided, the
// returned value will be the key origin, not the value origin.
func (r *Rolodex) FindNodeOriginWithValue(key, value, refNode *yaml.Node, refValue string) *NodeOrigin {
	_ = "STUB: not implemented"
	return nil
}

// the value is not in the root index, so we need to search all indexes

// FindNodeOrigin searches all indexes for the origin of a node. If the node is found, a NodeOrigin
// is returned, otherwise nil is returned.
func (r *Rolodex) FindNodeOrigin(node *yaml.Node) *NodeOrigin {
	_ = "STUB: not implemented"
	return nil
}

// FindNodeOrigin searches this index for a matching node. If the node is found, a NodeOrigin
// is returned, otherwise nil is returned.
func (index *SpecIndex) FindNodeOrigin(node *yaml.Node) *NodeOrigin {
	_ = "STUB: not implemented"
	return nil
}

// if the found node is a map. iterate through the content until we locate the node at that position

// hash node and found node

// hash node and found node

// check if the found node is a map and if the first item in the map
// has the same line and column, as well as the same value

type originCheck struct {
	valueOrigin *NodeOrigin
	valueHash   string
	keyOrigin   *NodeOrigin
	rolodex     *Rolodex
	value       *yaml.Node
	ref         string
	refNode     *yaml.Node
}

func checkHash(node, foundNode *yaml.Node) bool { _ = "STUB: not implemented"; return false }

func checkOrigin(check originCheck) (*NodeOrigin, bool) {
	_ = "STUB: not implemented"
	return nil, false

	// hash value and value origin
}

// no hit on the root, but we know the value is in the spec, so we need to search all indexes

// do the hashes match?
