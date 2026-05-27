// Copyright 2023-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package utils

import "go.yaml.in/yaml/v4"

type simpleJSONPathStepKind uint8

const (
	simpleJSONPathProperty simpleJSONPathStepKind = iota
	simpleJSONPathIndex
)

type simpleJSONPathStep struct {
	kind     simpleJSONPathStepKind
	property string
	index    int
}

func findNodesWithoutDeserializingFastPath(node *yaml.Node, jsonPath string) ([]*yaml.Node, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func parseSimpleJSONPath(path string) ([]simpleJSONPathStep, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func navigateJSONPathProperty(node *yaml.Node, property string) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

func navigateJSONPathIndex(node *yaml.Node, index int) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}
