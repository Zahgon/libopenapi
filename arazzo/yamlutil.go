// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	"go.yaml.in/yaml/v4"
)

func toYAMLNode(value any) (*yaml.Node, error) { _ = "STUB: not implemented"; return nil, nil }

// directYAMLNode converts a Go value to a *yaml.Node for expression evaluation.
// Map key ordering is not deterministic since the output is used for JSONPath
// and expression evaluation, not for rendering.
func directYAMLNode(value any) (*yaml.Node, error) { _ = "STUB: not implemented"; return nil, nil }

// sprintMapKey converts a map key to a string, fast-pathing the common string case.
func sprintMapKey(k any) string { _ = "STUB: not implemented"; return "" }
