// Copyright 2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"os"
	"time"

	"go.yaml.in/yaml/v4"
)

type rolodexFile struct {
	location   string
	rolodex    *Rolodex
	index      *SpecIndex
	localFile  *LocalFile
	remoteFile *RemoteFile
}

func (rf *rolodexFile) Name() string { _ = "STUB: not implemented"; return "" }

func (rf *rolodexFile) GetIndex() *SpecIndex { _ = "STUB: not implemented"; return nil }

func (rf *rolodexFile) Index(config *SpecIndexConfig) (*SpecIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// first, we must parse the content of the file

// create a new index for this file and link it to this rolodex.

func (rf *rolodexFile) GetContent() string { _ = "STUB: not implemented"; return "" }

func (rf *rolodexFile) GetContentAsYAMLNode() (*yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rf *rolodexFile) GetFileExtension() FileExtension {
	_ = "STUB: not implemented"
	return *new(FileExtension)
}

func (rf *rolodexFile) GetFullPath() string { _ = "STUB: not implemented"; return "" }

func (rf *rolodexFile) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (rf *rolodexFile) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (rf *rolodexFile) IsDir() bool {
	_ = "STUB: not implemented"
	// always false.
	return false
}

func (rf *rolodexFile) Sys() interface{} {
	_ = "STUB: not implemented"
	// not implemented.
	return nil
}

func (rf *rolodexFile) Mode() os.FileMode { _ = "STUB: not implemented"; return *new(os.FileMode) }

func (rf *rolodexFile) GetErrors() []error { _ = "STUB: not implemented"; return nil }

func (rf *rolodexFile) WaitForIndexing() { _ = "STUB: not implemented"; return }
