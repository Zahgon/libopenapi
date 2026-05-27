// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package low

import (
	"sync"

	"go.yaml.in/yaml/v4"
)

// MergeRecursiveNodesIfLineAbsent walks a node tree and adds each discovered node to dst
// unless that line already exists in the destination map.
func MergeRecursiveNodesIfLineAbsent(dst *sync.Map, node *yaml.Node) {
	_ = "STUB: not implemented"
	return
}

// AppendRecursiveNodes walks a node tree and appends each discovered node to dst.
func AppendRecursiveNodes(dst AddNodes, node *yaml.Node) { _ = "STUB: not implemented"; return }

func walkRecursiveNodes(node *yaml.Node, visit func(*yaml.Node)) { _ = "STUB: not implemented"; return }
