// Copyright 2025 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"github.com/pb33f/libopenapi/datamodel"
	"go.yaml.in/yaml/v4"
)

// PropertyMerger handles merging of local properties with referenced schema properties
type PropertyMerger struct {
	strategy datamodel.PropertyMergeStrategy
}

// NewPropertyMerger creates a new property merger with the specified strategy
func NewPropertyMerger(strategy datamodel.PropertyMergeStrategy) *PropertyMerger {
	_ = "STUB: not implemented"
	return nil
}

// MergeProperties merges local properties with referenced schema properties based on strategy
// localNode contains properties that should be preserved (e.g., examples, descriptions)
// referencedNode contains the resolved reference content
func (pm *PropertyMerger) MergeProperties(localNode, referencedNode *yaml.Node) (*yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extract properties from both nodes

// create merged node starting with referenced content

// apply merge strategy for each local property

// property exists in both - apply strategy

// keep referenced value (already in merged)

// property only exists locally - always preserve

// rebuild the merged node content

// extractProperties extracts key-value pairs from a yaml mapping node
func (pm *PropertyMerger) extractProperties(node *yaml.Node) map[string]*yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// rebuildNodeFromProperties reconstructs a yaml mapping node from property map
func (pm *PropertyMerger) rebuildNodeFromProperties(baseNode *yaml.Node, props map[string]*yaml.Node) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// rebuild content from properties

// copyNode creates a deep copy of a yaml node
func (pm *PropertyMerger) copyNode(node *yaml.Node) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// ShouldMergeProperties determines if property merging should be applied based on configuration
func (pm *PropertyMerger) ShouldMergeProperties(localNode, referencedNode *yaml.Node, config *datamodel.DocumentConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

// only merge if both nodes have properties to merge
