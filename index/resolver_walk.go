// Copyright 2022-2026 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package index

import "go.yaml.in/yaml/v4"

type relativeWalkState struct {
	foundRelatives map[string]bool
	journey        []*Reference
	resolve        bool
	depth          int
	schemaIDBase   string
}

func newRelativeWalkState(
	foundRelatives map[string]bool,
	journey []*Reference,
	resolve bool,
	depth int,
	schemaIDBase string,
) relativeWalkState {
	_ = "STUB: not implemented"
	return *new(relativeWalkState)
}

func (state relativeWalkState) withNodeBase(resolver *Resolver, node *yaml.Node) relativeWalkState {
	_ = "STUB: not implemented"
	return *new(relativeWalkState)
}

func (state relativeWalkState) descend() relativeWalkState {
	_ = "STUB: not implemented"
	return *new(relativeWalkState)
}
