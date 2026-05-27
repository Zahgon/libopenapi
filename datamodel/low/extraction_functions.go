// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package low

import (
	"context"
	"errors"
	"hash/maphash"
	"strings"
	"sync"

	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// stringBuilderPool is a sync.Pool that reuses strings.Builder instances to reduce memory allocations
// when generating hashes across the codebase.
var stringBuilderPool = sync.Pool{
	New: func() interface{} {
		return new(strings.Builder)
	},
}

// hashCache is a global cache for computed hash values to avoid redundant calculations.
// Uses sync.Map for thread-safe concurrent access.
var hashCache sync.Map

// ErrExternalRefSkipped is returned by LocateRefNodeWithContext when
// SkipExternalRefResolution is enabled and the reference is external.
var ErrExternalRefSkipped = errors.New("external reference resolution skipped")

// ClearHashCache clears the global hash cache. This should be called before
// starting a new document comparison to ensure clean state.
func ClearHashCache() { _ = "STUB: not implemented"; return }

// GetStringBuilder retrieves a strings.Builder from the pool, resets it, and returns it.
// The caller must call PutStringBuilder when done to return it to the pool.
func GetStringBuilder() *strings.Builder { _ = "STUB: not implemented"; return nil }

// PutStringBuilder returns a strings.Builder to the pool for reuse.
func PutStringBuilder(sb *strings.Builder) { _ = "STUB: not implemented"; return }

// FindItemInOrderedMap accepts a string key and a collection of KeyReference[string] and ValueReference[T].
// Every KeyReference will have its value checked against the string key and if there is a match, it will be
// returned.
func FindItemInOrderedMap[T any](item string, collection *orderedmap.Map[KeyReference[string], ValueReference[T]]) *ValueReference[T] {
	_ = "STUB: not implemented"
	return nil
}

// FindItemInOrderedMapWithKey is the same as FindItemInOrderedMap, except this code returns the key as well as the value.
func FindItemInOrderedMapWithKey[T any](item string, collection *orderedmap.Map[KeyReference[string], ValueReference[T]]) (*KeyReference[string], *ValueReference[T]) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HashExtensions will generate a hash from the low representation of extensions.
func HashExtensions(ext *orderedmap.Map[KeyReference[string], ValueReference[*yaml.Node]]) []string {
	_ = "STUB: not implemented"
	return nil
}

// Collect key-value entries and sort by key, avoiding a full map copy via SortAlpha.

// indexCollectionCache caches the result of generateIndexCollection per SpecIndex.
var indexCollectionCache sync.Map

// helper function to generate a list of all the things an index should be searched for.
// Cached per SpecIndex instance to avoid repeated slice+closure allocations.
func generateIndexCollection(idx *index.SpecIndex) []func() map[string]*index.Reference {
	_ = "STUB: not implemented"
	return nil
}

func LocateRefNodeWithContext(ctx context.Context, root *yaml.Node, idx *index.SpecIndex) (*yaml.Node, *index.SpecIndex, error, context.Context) {
	_ = "STUB: not implemented"
	return nil, nil, nil, *new(context.Context)
}

// run through everything and return as soon as we find a match.
// this operates as fast as possible as ever

// if this is a ref node, we need to keep diving
// until we hit something that isn't a ref.

// if this node is circular, stop drop and roll.

// Obtain the absolute filepath/URL of the spec in which we are trying to
// resolve the reference value [rv] from. It's either available from the
// index or passed down through context.

// explodedRefValue contains both the path to the file containing the
// reference value at index 0 and the path within that file to a specific
// sub-schema, should it exist, at index 1.

// The ref points to a component within either this file or another file.

// The ref is not an absolute URL.

// The ref is not an absolute local file path.

// The schema containing the ref is itself a remote file.

// p is the directory the referenced file is expected to be in.

// We are using the path of the resolved URL from the rolodex to
// obtain the "folder" or base of the file URL.

// We are resolving the relative URL against the absolute URL of
// the spec containing the reference.

// Turn the reference value [rv] into the absolute filepath/URL we
// resolved.

// The schema containing the ref is a local file or doesn't have an
// absolute URL.

// We have _some_ path for the schema containing the reference.

// Reference is made within the schema file, so we are using the
// same absolute local filepath.

// break off any fragments from the spec path

// Create a clean (absolute?) path to the file containing the
// referenced value.

// We don't have a path for the schema we are trying to resolve
// relative references from. This likely happens when the schema
// is the root schema, i.e., the file given to libopenapi as an entry.
//

// check for a config BaseURL and use that if it exists.

// check for a config baseURL and use that if it exists.

// let's try something else to find our references.

// cant be found? last resort is to try a path lookup

func applyResolvedSchemaIdScope(ctx context.Context, ref *index.Reference, idx *index.SpecIndex) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// LocateRefNode will perform a complete lookup for a $ref node. This function searches the entire index for
// the reference being supplied. If there is a match found, the reference *yaml.Node is returned.
func LocateRefNode(root *yaml.Node, idx *index.SpecIndex) (*yaml.Node, *index.SpecIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ExtractObjectRaw will extract a typed Buildable[N] object from a root yaml.Node. The 'raw' aspect is
// that there is no NodeReference wrapper around the result returned, just the raw object.
func ExtractObjectRaw[T Buildable[N], N any](ctx context.Context, key, root *yaml.Node, idx *index.SpecIndex) (T, error, bool, string) {
	_ = "STUB: not implemented"
	return *new(T), nil, false, ""
}

// if this is a reference, keep track of the reference in the value

// do we want to throw an error as well if circular error reporting is on?

// ExtractObject will extract a typed Buildable[N] object from a root yaml.Node. The result is wrapped in a
// NodeReference[T] that contains the key node found and value node found when looking up the reference.
func ExtractObject[T Buildable[N], N any](ctx context.Context, label string, root *yaml.Node, idx *index.SpecIndex) (NodeReference[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if this is a reference, keep track of the reference in the value

// do we want to throw an error as well if circular error reporting is on?

func extractArrayValueReferences[T Buildable[N], N any](
	ctx context.Context,
	label string,
	labelNode, valueNode *yaml.Node,
	idx *index.SpecIndex,
	isRef bool,
) ([]ValueReference[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if this was pulled from a ref, but it's not a sequence, check the label and see if anything comes out,
// and then check that is a sequence, if not, fail it.

func SetReference(obj any, ref string, refNode *yaml.Node) { _ = "STUB: not implemented"; return }

// Ensure the embedded *Reference is initialized before calling SetReference.
// Buildable types embed *Reference (a pointer) which is nil after new(T).
// Calling SetReference on a nil *Reference would panic.

// initEmbeddedReference uses reflection to find and initialize a nil *Reference
// field embedded in obj. This is needed when objects are created via new(T) without
// calling Build(), which normally initializes the embedded *Reference.
func initEmbeddedReference(obj any) { _ = "STUB: not implemented"; return }

// ExtractArray will extract a slice of []ValueReference[T] from a root yaml.Node that is defined as a sequence.
// Used when the value being extracted is an array.
func ExtractArray[T Buildable[N], N any](ctx context.Context, label string, root *yaml.Node, idx *index.SpecIndex) ([]ValueReference[T],
	*yaml.Node, *yaml.Node, error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// ExtractArrayNoLookup builds an array of low-level values from an already-located YAML sequence node.
func ExtractArrayNoLookup[T Buildable[N], N any](
	ctx context.Context,
	labelNode, valueNode *yaml.Node,
	idx *index.SpecIndex,
) ([]ValueReference[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractMapNoLookupExtensions will extract a map of KeyReference and ValueReference from a root yaml.Node. The 'NoLookup' part
// refers to the fact that there is no key supplied as part of the extraction, there  is no lookup performed and the
// root yaml.Node pointer is used directly. Pass a true bit to includeExtensions to include extension keys in the map.
//
// This is useful when the node to be extracted, is already known and does not require a search.
func ExtractMapNoLookupExtensions[PT Buildable[N], N any](
	ctx context.Context,
	root *yaml.Node,
	idx *index.SpecIndex,
	includeExtensions bool,
) (*orderedmap.Map[KeyReference[string], ValueReference[PT]], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if value is a reference, we have to look it up in the index!

// ExtractMapNoLookup will extract a map of KeyReference and ValueReference from a root yaml.Node. The 'NoLookup' part
// refers to the fact that there is no key supplied as part of the extraction, there  is no lookup performed and the
// root yaml.Node pointer is used directly.
//
// This is useful when the node to be extracted, is already known and does not require a search.
func ExtractMapNoLookup[PT Buildable[N], N any](
	ctx context.Context,
	root *yaml.Node,
	idx *index.SpecIndex,
) (*orderedmap.Map[KeyReference[string], ValueReference[PT]], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type mappingResult[T any] struct {
	k KeyReference[string]
	v ValueReference[T]
	e error
}

type buildInput struct {
	label *yaml.Node
	value *yaml.Node
}

func findExtractLabelNode(label string, root *yaml.Node) (keyNode *yaml.Node, labelNode *yaml.Node, valueNode *yaml.Node) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func collectMapBuildInputs(valueNode *yaml.Node, extensions bool) []buildInput {
	_ = "STUB: not implemented"
	return nil
}

// ExtractMapExtensions will extract a map of KeyReference and ValueReference from a root yaml.Node. The 'label' is
// used to locate the node to be extracted from the root node supplied. Supply a bit to decide if extensions should
// be included or not. required in some use cases.
//
// The second return value is the yaml.Node found for the 'label' and the third return value is the yaml.Node
// found for the value extracted from the label node.
func ExtractMapExtensions[PT Buildable[N], N any](
	ctx context.Context,
	label string,
	root *yaml.Node,
	idx *index.SpecIndex,
	extensions bool,
) (*orderedmap.Map[KeyReference[string], ValueReference[PT]], *yaml.Node, *yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// locate reference in index.

// check our valueNode isn't a reference still.

// ExtractMap will extract a map of KeyReference and ValueReference from a root yaml.Node. The 'label' is
// used to locate the node to be extracted from the root node supplied.
//
// The second return value is the yaml.Node found for the 'label' and the third return value is the yaml.Node
// found for the value extracted from the label node.
func ExtractMap[PT Buildable[N], N any](
	ctx context.Context,
	label string,
	root *yaml.Node,
	idx *index.SpecIndex,
) (*orderedmap.Map[KeyReference[string], ValueReference[PT]], *yaml.Node, *yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// ExtractExtensions will extract any 'x-' prefixed key nodes from a root node into a map. Requirements have been pre-cast:
//
// Maps
//
//	*orderedmap.Map[string, *yaml.Node] for maps
//
// Slices
//
//	[]interface{}
//
// int, float, bool, string
//
//	int64, float64, bool, string
func ExtractExtensions(root *yaml.Node) *orderedmap.Map[KeyReference[string], ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil
}

// AreEqual returns true if two Hashable objects are equal or not.
func AreEqual(l, r Hashable) bool { _ = "STUB: not implemented"; return false }

// GenerateHashString will generate a SHA256 hash of any object passed in. If the object is Hashable
// then the underlying Hash() method will be called. Optimized to avoid excessive allocations and
// uses caching to eliminate redundant calculations.
func GenerateHashString(v any) string { _ = "STUB: not implemented"; return "" }

// Try cache first using the pointer as key for non-primitives
// However, skip caching for types with mutable hash state like SchemaProxy

// Check if this is a type that has mutable hash state or complex comparison logic

// Format uint64 hash as hex string

// Fast path for common YAML node types to avoid marshaling

// Primitive types
// if we get here, we're a primitive, check if we're a pointer and de-point

// Convert to string efficiently using strconv instead of fmt.Sprintf

// Store in cache if we have a valid pointer and caching is enabled for this type

// hashYamlNodeFast provides fast hashing for YAML nodes without ANY marshaling
func hashYamlNodeFast(n *yaml.Node) string { _ = "STUB: not implemented"; return "" }

// Try cache first for complex nodes
// Use pointer directly as key - *yaml.Node pointers are stable and comparable

// Cache complex nodes

// hashNodeTree walks the YAML tree and hashes it without marshaling
func hashNodeTree(h *maphash.Hash, n *yaml.Node, visited map[*yaml.Node]bool) {
	_ = "STUB: not implemented"
	return
}

func hashNodeTreeWithNumericNormalization(h *maphash.Hash, n *yaml.Node, visited map[*yaml.Node]bool, normalizeNumericScalars bool) {
	_ = "STUB: not implemented"
	return
}

// Prevent circular reference infinite loops

// Hash node metadata. Numeric scalars are normalized so semantically equivalent
// values like `1e-08` and `1e-8` compare equal.

// CRITICAL: Snapshot Content to prevent TOCTOU races
// This captures the slice header (pointer, len, cap) atomically.
// Even if another goroutine reassigns n.Content later, our local
// 'content' variable still refers to the original backing array.

// Hash based on node type

// Already hashed value above

// Guard against empty mapping nodes

// For maps, we need consistent ordering
// Collect key-value pairs and sort by key hash

// Sort for consistent hashing

// Hash in sorted order

func comparableScalarTagAndValue(n *yaml.Node) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func scalarTagAndValueForHash(n *yaml.Node, normalizeNumericScalars bool) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// CompareYAMLNodes compares two YAML nodes for equality without marshaling to YAML.
// This reuses the hashNodeTree logic to generate consistent hashes for comparison,
// avoiding the expensive yaml.Marshal operations that cause massive allocations.
func CompareYAMLNodes(left, right *yaml.Node) bool { _ = "STUB: not implemented"; return false }

// YAMLNodeToBytes converts a YAML node to bytes in a more efficient way than yaml.Marshal
// This function should be used when you actually need the marshaled bytes (like for JSON conversion)
// rather than just comparing nodes (use CompareYAMLNodes for that)
func YAMLNodeToBytes(n *yaml.Node) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// For now, we still use yaml.Marshal for cases that actually need the bytes
// This can be optimized further in the future with a custom serializer

// HashYAMLNodeSlice creates a hash for a slice of YAML nodes efficiently
// This replaces the pattern of yaml.Marshal + sha256 that's used in example comparisons
func HashYAMLNodeSlice(nodes []*yaml.Node) string { _ = "STUB: not implemented"; return "" }

// AppendMapHashes will append all the hashes of a map to a slice of strings.
// Optimized to avoid creating sorted copies on every call.
func AppendMapHashes[v any](a []string, m *orderedmap.Map[KeyReference[string], ValueReference[v]]) []string {
	_ = "STUB: not implemented"
	return nil

	// Pre-allocate slice for better performance when we know the size
}

// Collect entries and sort them by key for consistent hashing
// This is more efficient than orderedmap.SortAlpha() which creates a full copy

// Sort entries by key for consistent hash ordering
// Use a simple insertion sort for small maps, quicksort for larger ones

// Insertion sort for small maps

// Use Go's built-in sort for larger maps

// For small maps, avoid string builder overhead and use direct string concatenation

// Use string builder for larger maps with pre-allocated capacity

// Pre-size for this specific entry to avoid growth
// key + hash + separator

func ValueToString(v any) string { _ = "STUB: not implemented"; return "" }

// For simple scalar nodes, return the value directly

// For complex nodes, still need to marshal for string representation

// LocateRefEnd will perform a complete lookup for a $ref node. This function searches the entire index for
// the reference being supplied. If there is a match found, the reference *yaml.Node is returned.
// the function operates recursively and will keep iterating through references until it finds a non-reference
// node.
func LocateRefEnd(ctx context.Context, root *yaml.Node, idx *index.SpecIndex, depth int) (*yaml.Node, *index.SpecIndex, error, context.Context) {
	_ = "STUB: not implemented"
	return nil, nil, nil, *new(context.Context)
}

// FromReferenceMap will convert a *orderedmap.Map[KeyReference[K], ValueReference[V]] to a *orderedmap.Map[K, V]
//
//go:noinline
func FromReferenceMap[K comparable, V any](refMap *orderedmap.Map[KeyReference[K], ValueReference[V]]) *orderedmap.Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// FromReferenceMapWithFunc will convert a *orderedmap.Map[KeyReference[K], ValueReference[V]] to a *orderedmap.Map[K, VOut] using a transform function
//
//go:noinline
func FromReferenceMapWithFunc[K comparable, V any, VOut any](refMap *orderedmap.Map[KeyReference[K], ValueReference[V]], transform func(v V) VOut) *orderedmap.Map[K, VOut] {
	_ = "STUB: not implemented"
	return nil
}
