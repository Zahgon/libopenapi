// Copyright 2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

// RefType identifies where a reference points to: within the same file (Local),
// to another file on disk (File), or to a remote URL (HTTP).
const (
	Local RefType = iota
	File
	HTTP
)

// RefType is an enum identifying the location type of a reference.
type RefType int

// ExtractedRef represents a parsed reference with its resolved location and type.
type ExtractedRef struct {
	Location string
	Type     RefType
}

// GetFile returns the file path of the reference.
func (r *ExtractedRef) GetFile() string { _ = "STUB: not implemented"; return "" }

// GetReference returns the reference path of the reference.
func (r *ExtractedRef) GetReference() string { _ = "STUB: not implemented"; return "" }

// ExtractFileType returns the file extension of the reference.
func ExtractFileType(ref string) FileExtension {
	_ = "STUB: not implemented"
	return *new(FileExtension)
}
