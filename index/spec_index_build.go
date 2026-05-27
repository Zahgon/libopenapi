// Copyright 2022-2026 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package index

import (
	"context"
	"log/slog"

	"go.yaml.in/yaml/v4"
)

const (
	// theoreticalRoot is the name of the theoretical spec file used when a root spec file does not exist
	theoreticalRoot = "root.yaml"
)

// NewSpecIndexWithConfigAndContext creates a new SpecIndex from the given root YAML node and configuration.
// The context is passed through to reference extraction for schema ID scope tracking.
func NewSpecIndexWithConfigAndContext(ctx context.Context, rootNode *yaml.Node, config *SpecIndexConfig) *SpecIndex {
	_ = "STUB: not implemented"
	return nil
}

// NewSpecIndexWithConfig creates a new SpecIndex from the given root YAML node and configuration,
// using a background context.
func NewSpecIndexWithConfig(rootNode *yaml.Node, config *SpecIndexConfig) *SpecIndex {
	_ = "STUB: not implemented"
	return nil
}

// NewSpecIndex creates a new SpecIndex with default configuration from the given root YAML node.
func NewSpecIndex(rootNode *yaml.Node) *SpecIndex { _ = "STUB: not implemented"; return nil }

func createNewIndex(ctx context.Context, rootNode *yaml.Node, index *SpecIndex, avoidBuildOut bool) *SpecIndex {
	_ = "STUB: not implemented"
	return nil
}

// BuildIndex runs all count and extraction functions concurrently to populate the index.
// This is called automatically during construction unless AvoidBuildIndex is set in the config.
func (index *SpecIndex) BuildIndex() { _ = "STUB: not implemented"; return }

// GetLogger returns the structured logger used by this index.
func (index *SpecIndex) GetLogger() *slog.Logger { _ = "STUB: not implemented"; return nil }

// GetRootNode returns the root YAML node of the specification document.
func (index *SpecIndex) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// SetRootNode sets the root YAML node for this index.
	return nil
}

func (index *SpecIndex) SetRootNode(node *yaml.Node) {
	_ = "STUB: not implemented"

	// GetRolodex returns the Rolodex file system abstraction associated with this index.
	return
}

func (index *SpecIndex) GetRolodex() *Rolodex { _ = "STUB: not implemented"; return nil }

// SetRolodex sets the Rolodex file system abstraction for this index.
func (index *SpecIndex) SetRolodex(rolodex *Rolodex) { _ = "STUB: not implemented"; return }

// GetSpecFileName returns the base filename of the specification (e.g. "openapi.yaml").
// Falls back to "root.yaml" if no file path is configured.
func (index *SpecIndex) GetSpecFileName() string { _ = "STUB: not implemented"; return "" }

// GetGlobalTagsNode returns the raw YAML node for the top-level "tags" array.
func (index *SpecIndex) GetGlobalTagsNode() *yaml.Node { _ = "STUB: not implemented"; return nil }
