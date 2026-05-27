// Copyright 2022-2026 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package index

import (
	"go.yaml.in/yaml/v4"
)

func (resolver *Resolver) extractPolymorphicRelatives(
	ref *Reference,
	node, keywordNode *yaml.Node,
	state relativeWalkState,
	index int,
) []*Reference {
	_ = "STUB: not implemented"
	return nil
}

func (resolver *Resolver) visitPolymorphicReference(
	ref *Reference,
	polymorphicType string,
	parentNode *yaml.Node,
	lookup string,
	state relativeWalkState,
	loopIndex int,
) {
	_ = "STUB: not implemented"
	return
}
