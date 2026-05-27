// Copyright 2022-2026 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package index

import (
	"context"

	"go.yaml.in/yaml/v4"
)

func visitIndexWithoutDamagingIt(res *Resolver, idx *SpecIndex) { _ = "STUB: not implemented"; return }

type refMap struct {
	ref   *Reference
	nodes []*yaml.Node
}

func visitIndex(res *Resolver, idx *SpecIndex) { _ = "STUB: not implemented"; return }

func (resolver *Resolver) searchReferenceWithContext(sourceRef, searchRef *Reference) (*Reference, *SpecIndex, context.Context) {
	_ = "STUB: not implemented"
	return nil, nil, *new(context.Context)
}

// VisitReference visits a single reference, collecting its relatives (dependencies) and recursively
// visiting them. The seen map prevents infinite loops, journey tracks the path for circular detection,
// and resolve controls whether nodes are actually resolved or just visited for analysis.
func (resolver *Resolver) VisitReference(ref *Reference, seen map[string]bool, journey []*Reference, resolve bool) []*yaml.Node {
	_ = "STUB: not implemented"
	return nil
}
