// Copyright 2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"io/fs"
	"log/slog"
	"os"
	"sync"
	"time"

	"context"

	"go.yaml.in/yaml/v4"
)

// CanBeIndexed is an interface that allows a file to be indexed.
type CanBeIndexed interface {
	Index(config *SpecIndexConfig) (*SpecIndex, error)
}

// RolodexFile is an interface that represents a file in the rolodex. It combines multiple `fs` interfaces
// like `fs.FileInfo` and `fs.File` into one interface, so the same struct can be used for everything.
type RolodexFile interface {
	GetContent() string
	GetFileExtension() FileExtension
	GetFullPath() string
	GetErrors() []error
	GetContentAsYAMLNode() (*yaml.Node, error)
	GetIndex() *SpecIndex
	// WaitForIndexing blocks until the file's index is ready.
	// This is used to coordinate between concurrent goroutines when one is loading
	// a file and another needs to use its index.
	WaitForIndexing()
	Name() string
	ModTime() time.Time
	IsDir() bool
	Sys() any
	Size() int64
	Mode() os.FileMode
}

// RolodexFS is an interface that represents a RolodexFS, is the same interface as `fs.FS`, except it
// also exposes a GetFiles() signature, to extract all files in the FS.
type RolodexFS interface {
	Open(name string) (fs.File, error)
	GetFiles() map[string]RolodexFile
}

// Rolodex is a file system abstraction that allows for the indexing of multiple file systems
// and the ability to resolve references across those file systems. It is used to hold references to external
// files, and the indexes they hold. The rolodex is the master lookup for all references.
type Rolodex struct {
	localFS                    map[string]fs.FS
	remoteFS                   map[string]fs.FS
	indexed                    bool
	built                      bool
	manualBuilt                bool
	resolved                   bool
	circChecked                bool
	indexConfig                *SpecIndexConfig
	indexingDuration           time.Duration
	indexes                    []*SpecIndex
	indexMap                   map[string]*SpecIndex
	indexLock                  sync.Mutex
	rootIndex                  *SpecIndex
	rootNode                   *yaml.Node
	caughtErrors               []error
	safeCircularReferences     []*CircularReferenceResult
	infiniteCircularReferences []*CircularReferenceResult
	ignoredCircularReferences  []*CircularReferenceResult
	debouncedSafeCircRefs      []*CircularReferenceResult // cached result from GetSafeCircularReferences
	debouncedIgnoredCircRefs   []*CircularReferenceResult // cached result from GetIgnoredCircularReferences
	circRefCacheLock           sync.Mutex                 // protects debounced cache fields
	logger                     *slog.Logger
	id                         string // unique ID for the rolodex, can be used to identify it in logs or other contexts.
	globalSchemaIdRegistry     map[string]*SchemaIdEntry
	schemaIdRegistryLock       sync.RWMutex
}

// Release nils all fields that can pin YAML node trees, SpecIndex objects, or
// circular reference results in memory. Acquires locks for fields that are
// protected elsewhere. Call this once all consumers of the rolodex are finished.
func (r *Rolodex) Release() { _ = "STUB: not implemented"; return }

// NewRolodex creates a new rolodex with the provided index configuration.
func NewRolodex(indexConfig *SpecIndexConfig) *Rolodex { _ = "STUB: not implemented"; return nil }

// RotateId generates a new unique ID for the rolodex.
func (r *Rolodex) RotateId() string { _ = "STUB: not implemented"; return "" }

// GetId returns the unique ID for the rolodex.
func (r *Rolodex) GetId() string {
	_ = "STUB: not implemented"

	// GetIgnoredCircularReferences returns a list of circular references that were ignored during the indexing process.
	// These can be an array or polymorphic references. Will return an empty slice if no ignored circular references are found.
	return ""
}

func (r *Rolodex) GetIgnoredCircularReferences() []*CircularReferenceResult {
	_ = "STUB: not implemented"
	return nil
}

// GetSafeCircularReferences returns a list of circular references that were found to be safe during the indexing process.
// These can be an array or polymorphic references. Will return an empty slice if no safe circular references are found.
func (r *Rolodex) GetSafeCircularReferences() []*CircularReferenceResult {
	_ = "STUB: not implemented"
	return nil
}

// if this rolodex has not been manually checked for circular references or resolved,
// then we need to perform that check now, looking at all indexes and extracting
// results from the resolvers.

// SetSafeCircularReferences sets the safe circular references for the rolodex.
func (r *Rolodex) SetSafeCircularReferences(refs []*CircularReferenceResult) {
	_ = "STUB: not implemented"
	return
}

// invalidate cache

// GetIndexingDuration returns the duration it took to index the rolodex.
func (r *Rolodex) GetIndexingDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GetRootIndex returns the root index of the rolodex (the entry point, the main document)
func (r *Rolodex) GetRootIndex() *SpecIndex {
	_ = "STUB: not implemented"

	// GetConfig returns the index configuration of the rolodex.
	return nil
}

func (r *Rolodex) GetConfig() *SpecIndexConfig { _ = "STUB: not implemented"; return nil }

// GetRootNode returns the root node of the rolodex (the entry point, the main document)
func (r *Rolodex) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetIndexes returns all the indexes in the rolodex.
	return nil
}

func (r *Rolodex) GetIndexes() []*SpecIndex { _ = "STUB: not implemented"; return nil }

// GetCaughtErrors returns all the errors that were caught during the indexing process.
func (r *Rolodex) GetCaughtErrors() []error { _ = "STUB: not implemented"; return nil }

// AddLocalFS adds a local file system to the rolodex.
func (r *Rolodex) AddLocalFS(baseDir string, fileSystem fs.FS) { _ = "STUB: not implemented"; return }

// SetRootNode sets the root node of the rolodex (the entry point, the main document)
func (r *Rolodex) SetRootNode(node *yaml.Node) {
	_ = "STUB: not implemented"

	// SetRootIndex sets the root index of the rolodex (the entry point, the main document).
	return
}

func (r *Rolodex) SetRootIndex(rootIndex *SpecIndex) { _ = "STUB: not implemented"; return }

func (r *Rolodex) AddExternalIndex(idx *SpecIndex, location string) {
	_ = "STUB: not implemented"
	return
}

// already exists, no need to add again.

// Aggregate $id registrations from this index into the global registry

func (r *Rolodex) AddIndex(idx *SpecIndex) { _ = "STUB: not implemented"; return }

// AddRemoteFS adds a remote file system to the rolodex.
func (r *Rolodex) AddRemoteFS(baseURL string, fileSystem fs.FS) { _ = "STUB: not implemented"; return }

// IndexTheRolodex indexes the rolodex, building out the indexes for each file, and then building the root index.
func (r *Rolodex) IndexTheRolodex(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// copy config and set the

// we will build out everything in two steps.

// Index() does not throw an error anymore.
// for each index, we need a resolver

// check if the config has been set to ignore circular references in arrays and polymorphic schemas

// run through every file system and index every file, fan out as many goroutines as possible.

// now that we have indexed all the files, we can build the index.

// indexed and built every supporting file, we can build the root index (our entry point)

// if there is a base path but no SpecFilePath, then we need to set the root spec config to point to a theoretical root.yaml
// which does not exist, but is used to formulate the absolute path to root references correctly.

// Compute the absolute path to the spec file.
// - If SpecFilePath is already absolute, use it directly.
// - If SpecFilePath is relative, it needs careful handling to avoid path doubling.
//
// The original code used filepath.Base() which incorrectly stripped directory
// segments like /myproject/api-spec/ from nested paths.
//
// Handle cases:
// 1. SpecFilePath = "test_data/nested/doc.yaml", BasePath = "/abs/test_data/nested"
//    -> Should NOT double to /abs/test_data/nested/test_data/nested/doc.yaml
// 2. SpecFilePath = "subdir/doc.yaml", BasePath = "/abs/test_data"
//    -> Should produce /abs/test_data/subdir/doc.yaml

// Check if SpecFilePath starts with the relative basePath or its original value
// This handles cases where SpecFilePath = "test_data/file.yaml" and
// BasePath was originally "test_data" (now absolute)

// Normalize paths to use OS-specific separators for Windows compatibility
// On Windows, paths may use / but os.PathSeparator is \, causing mismatches

// SpecFilePath includes the original basePath, make it absolute directly

// SpecFilePath starts with ".." (parent directory), resolve it from cwd
// Using filepath.Join with basePath would incorrectly double paths
// e.g., basePath="/Users/foo/bar" + "../bar/file.yaml" would give
// "/Users/foo/bar/bar/file.yaml" instead of "/Users/foo/bar/file.yaml"

// SpecFilePath is relative to basePath, join them

// Here we take the root node and also build the index for it.
// This involves extracting references.

// CheckForCircularReferences checks for circular references in the rolodex.
func (r *Rolodex) CheckForCircularReferences() { _ = "STUB: not implemented"; return }

// invalidate debounced caches since underlying slices were mutated

// Resolve resolves references in the rolodex.
func (r *Rolodex) Resolve() { _ = "STUB: not implemented"; return }

// resolve pending nodes

// invalidate debounced caches since underlying slices were mutated

func (r *Rolodex) collectResolvers() []*Resolver { _ = "STUB: not implemented"; return nil }

func (r *Rolodex) mergeResolverResults(res *Resolver) { _ = "STUB: not implemented"; return }

// BuildIndexes builds the indexes in the rolodex, this is generally not required unless manually building a rolodex.
func (r *Rolodex) BuildIndexes() { _ = "STUB: not implemented"; return }

// GetAllReferences  returns all references found in the root and all other indices
func (r *Rolodex) GetAllReferences() map[string]*Reference { _ = "STUB: not implemented"; return nil }

// GetAllMappedReferences returns all mapped references found in the root and all other indices
func (r *Rolodex) GetAllMappedReferences() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil
}

// OpenWithContext opens a file in the rolodex, and returns a RolodexFile - providing a context.
// The method supports both custom file systems (like LocalFS) and standard fs.FS implementations.
// For standard fs.FS implementations, paths are automatically converted to relative paths as required
// by the fs.FS interface specification (which mandates relative, slash-separated paths).
func (r *Rolodex) OpenWithContext(ctx context.Context, location string) (RolodexFile, error) {
	_ = "STUB: not implemented"
	return *new(RolodexFile), nil
}

func (r *Rolodex) openLocalLocation(ctx context.Context, location string) (*LocalFile, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Rolodex) localPathForOpen(baseDir, fileLookup string, fileSystem fs.FS) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *Rolodex) asLocalFile(file fs.File, fileLookup string) (*LocalFile, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wrapExistingRolodexFile(file RolodexFile) (*LocalFile, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Rolodex) openRemoteLocation(ctx context.Context, location string) (*RemoteFile, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Rolodex) asRemoteFile(file fs.File, location string) (*RemoteFile, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func consumeAdaptedFile(file fs.File) ([]byte, fs.FileInfo, []error) {
	_ = "STUB: not implemented"
	return nil, *new(fs.FileInfo), nil
}

func (r *Rolodex) wrapLocalRolodexFile(localFile *LocalFile) (RolodexFile, error) {
	_ = "STUB: not implemented"
	return *new(RolodexFile), nil
}

func (r *Rolodex) wrapRemoteRolodexFile(remoteFile *RemoteFile) (RolodexFile, error) {
	_ = "STUB: not implemented"
	return *new(RolodexFile), nil
}

func openFile(ctx context.Context, location string, v fs.FS) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

// Open opens a file in the rolodex, and returns a RolodexFile.
func (r *Rolodex) Open(location string) (RolodexFile, error) {
	_ = "STUB: not implemented"
	return *new(RolodexFile), nil
}

var suffixes = []string{"B", "KB", "MB", "GB", "TB"}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	_ = "STUB: not implemented"
	return 0
}

func HumanFileSize(size float64) string { _ = "STUB: not implemented"; return "" }

func (r *Rolodex) RolodexFileSizeAsString() string { _ = "STUB: not implemented"; return "" }

func (r *Rolodex) RolodexTotalFiles() int {
	_ = "STUB: not implemented"
	// look through each file system and count the files
	return 0
}

func (r *Rolodex) RolodexFileSize() int64 { _ = "STUB: not implemented"; return 0 }

// GetFullLineCount returns the total number of lines from all files in the Rolodex
func (r *Rolodex) GetFullLineCount() int64 { _ = "STUB: not implemented"; return 0 }

// add in root count

func (r *Rolodex) ClearIndexCaches() { _ = "STUB: not implemented"; return }

// RegisterGlobalSchemaId registers a schema $id in the Rolodex global registry.
// Returns an error if the $id is invalid.
func (r *Rolodex) RegisterGlobalSchemaId(entry *SchemaIdEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// LookupSchemaById looks up a schema by its $id URI across all indexes.
func (r *Rolodex) LookupSchemaById(uri string) *SchemaIdEntry {
	_ = "STUB: not implemented"
	return nil
}

// GetAllGlobalSchemaIds returns a copy of all registered $id entries across all indexes.
func (r *Rolodex) GetAllGlobalSchemaIds() map[string]*SchemaIdEntry {
	_ = "STUB: not implemented"
	return nil
}

// RegisterIdsFromIndex aggregates all $id registrations from an index into the global registry.
// Called after each index is built to populate the Rolodex global registry.
func (r *Rolodex) RegisterIdsFromIndex(idx *SpecIndex) { _ = "STUB: not implemented"; return }
