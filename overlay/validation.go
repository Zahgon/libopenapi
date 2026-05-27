// Copyright 2022-2025 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package overlay

import (
	highoverlay "github.com/pb33f/libopenapi/datamodel/high/overlay"
	"go.yaml.in/yaml/v4"
)

// validateOverlay checks that the overlay has all required fields.
func validateOverlay(overlay *highoverlay.Overlay) error { _ = "STUB: not implemented"; return nil }

// validateTarget checks that a target node is a valid target (object or array).
// Per the Overlay Spec, primitive/null targets are invalid.
func validateTarget(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }
