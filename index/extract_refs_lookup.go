// Copyright 2023-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"context"
	"os"
)

// preserveLegacyRefOrder allows opt-out of deterministic ordering if issues arise.
// Set LIBOPENAPI_LEGACY_REF_ORDER=true to use the old non-deterministic ordering.
var preserveLegacyRefOrder = os.Getenv("LIBOPENAPI_LEGACY_REF_ORDER") == "true"

// indexedRef pairs a resolved reference with its original input position for deterministic ordering.
type indexedRef struct {
	ref *Reference
	pos int
}

// ExtractComponentsFromRefs returns located components from references. The returned nodes from here
// can be used for resolving as they contain the actual object properties.
//
// This function uses singleflight to deduplicate concurrent lookups for the same reference,
// channel-based collection to avoid mutex contention during resolution, and sorts results
// by input position for deterministic ordering.

// isExternalReference checks whether a Reference originated from an external $ref.
// ref.Definition may have been transformed (e.g., HTTP URL with fragment becomes "#/fragment"),
// so we also check the original raw ref value.
func isExternalReference(ref *Reference) bool { _ = "STUB: not implemented"; return false }

func (index *SpecIndex) ExtractComponentsFromRefs(ctx context.Context, refs []*Reference) []*Reference {
	_ = "STUB: not implemented"
	return nil
}

func (index *SpecIndex) locateRef(ctx context.Context, ref *Reference) *Reference {
	_ = "STUB: not implemented"
	return nil
}
