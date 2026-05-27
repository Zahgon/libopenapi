// Copyright 2022-2026 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package index

import "go.yaml.in/yaml/v4"

func (resolver *Resolver) visitReferenceShortCircuit(ref *Reference, resolve bool) ([]*yaml.Node, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (resolver *Resolver) collectReferenceRelatives(
	ref *Reference,
	seen map[string]bool,
	journey []*Reference,
	resolve bool,
) []*Reference {
	_ = "STUB: not implemented"
	return nil
}

func (resolver *Resolver) visitReferenceRelatives(
	ref *Reference,
	relatives []*Reference,
	seen map[string]bool,
	journey []*Reference,
	resolve bool,
) {
	_ = "STUB: not implemented"
	return
}

func (resolver *Resolver) resolveRelativeReference(
	ref, relative *Reference,
	seen map[string]bool,
	journey []*Reference,
	resolve bool,
) {
	_ = "STUB: not implemented"
	return
}
