// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"io/fs"
)

// resolveRelativeFilePath resolves a relative file reference against a base directory.
// It prefers paths that actually exist in the configured local file systems, falling
// back to defRoot if no match is found.
func (index *SpecIndex) resolveRelativeFilePath(defRoot, ref string) string {
	_ = "STUB: not implemented"
	return ""
}

// Prefer the path relative to the current file if it exists.

// Prefer the configured BasePath if present and it yields an existing file.

// Otherwise, try each registered local filesystem base directory.

// ResolveRelativeFilePath is a public wrapper for resolving local file references.
func (index *SpecIndex) ResolveRelativeFilePath(defRoot, ref string) string {
	_ = "STUB: not implemented"
	return ""
}

func pathExistsInFS(baseDir string, fsys fs.FS, absPath string) bool {
	_ = "STUB: not implemented"
	return false
}
