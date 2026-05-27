// Copyright 2023-2024 Princess Beef Heavy Industries, LLC / Dave Shanley
// https://pb33f.io
// MIT License

package low

import (
	"context"
	"sync"

	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// HasNodes is an interface that defines a method to get a map of nodes
type HasNodes interface {
	GetNodes() map[int][]*yaml.Node
}

// AddNodes is an interface that defined a method to add nodes.
type AddNodes interface {
	AddNode(key int, node *yaml.Node)
}

// NodeMap represents a map of yaml nodes
type NodeMap struct {
	// Nodes is a sync map of nodes for this object, and the key is the line number of the node
	// a line can contain many nodes (in JSON), so the value is a slice of *yaml.Node
	Nodes *sync.Map `yaml:"-" json:"-"`
}

// AddNode will add a node to the NodeMap
func (nm *NodeMap) AddNode(key int, node *yaml.Node) { _ = "STUB: not implemented"; return }

// GetNodes will return the map of nodes
func (nm *NodeMap) GetNodes() map[int][]*yaml.Node { _ = "STUB: not implemented"; return nil }

// return an empty slice if there are no nodes

// ExtractNodes will iterate over a *yaml.Node and extract all nodes with a line number into a map
func (nm *NodeMap) ExtractNodes(node *yaml.Node, recurse bool) { _ = "STUB: not implemented"; return }

// if the node has content, iterate over it and extract every top level line number

// ContainsLine will return true if the NodeMap contains a node with the supplied line number
func (nm *NodeMap) ContainsLine(line int) bool { _ = "STUB: not implemented"; return false }

// ExtractNodes will extract all nodes from a yaml.Node and return them in a map
func ExtractNodes(_ context.Context, root *yaml.Node) *sync.Map {
	_ = "STUB: not implemented"
	return nil
}

// ExtractNodesRecursive will extract all nodes from a yaml.Node and return them in a map, just like ExtractNodes
// however, this version will dive-down the tree and extract all nodes from all child nodes as well until the tree
// is done.
func ExtractNodesRecursive(_ context.Context, root *yaml.Node) *sync.Map {
	_ = "STUB: not implemented"
	return nil
}

// ExtractExtensionNodes will extract all extension nodes from a map of extensions, recursively.
func ExtractExtensionNodes(_ context.Context,
	extensionMap *orderedmap.Map[KeyReference[string],
		ValueReference[*yaml.Node]], nodeMap *sync.Map,
) {
	_ = "STUB: not implemented"
	// range over the extension map and extract all nodes
	return
}
