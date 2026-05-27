// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

const conflictNameDelimiter = "__"

type nameRegistry struct {
	used map[string]string
}

func newNameRegistry() *nameRegistry { _ = "STUB: not implemented"; return nil }

func (r *nameRegistry) resolve(original, candidate string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
