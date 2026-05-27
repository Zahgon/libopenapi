// Copyright 2023-2024 Princess Beef Heavy Industries, LLC / Dave Shanley
// https://pb33f.io
// SPDX-License-Identifier: MIT

package bundler

import (
	"errors"

	"go.yaml.in/yaml/v4"

	"github.com/pb33f/libopenapi/datamodel"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
)

// ErrInvalidModel is returned when the model is not usable.
var ErrInvalidModel = errors.New("invalid model")

func renderBundledModel(model *v3.Document, rootIndex *index.SpecIndex) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateDiscriminatorMappings(rolodex *index.Rolodex) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDiscriminatorMappingsFromIndex(idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDiscriminatorMappingsFromNode(n *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDiscriminatorMappingsFromRootNode(n *yaml.Node, seen map[*yaml.Node]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func isOpenAPIDocumentRoot(n *yaml.Node) bool { _ = "STUB: not implemented"; return false }

func discriminatorValidationNode(n *yaml.Node) *yaml.Node { _ = "STUB: not implemented"; return nil }

func validateDiscriminatorMappingsFromOpenAPIObject(n *yaml.Node, schemaSeen, objectSeen map[*yaml.Node]struct{}, path []string) error {
	_ = "STUB: not implemented"
	return nil
}

func shouldSkipDiscriminatorValidationOpenAPIValue(key string) bool {
	_ = "STUB: not implemented"
	return false
}

func validateDiscriminatorMappingsFromSchemaNode(n *yaml.Node, seen map[*yaml.Node]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDiscriminatorMappingsFromSchemaMap(n *yaml.Node, seen map[*yaml.Node]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDiscriminatorMappingsFromSchemaArray(n *yaml.Node, seen map[*yaml.Node]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func isDiscriminatorValidationSchemaCandidate(n *yaml.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func isDirectSchemaChildKey(key string) bool { _ = "STUB: not implemented"; return false }

func isSchemaMapChildKey(key string) bool { _ = "STUB: not implemented"; return false }

func isSchemaArrayChildKey(key string) bool { _ = "STUB: not implemented"; return false }

type invalidModelBuildError struct {
	cause error
}

func (e *invalidModelBuildError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *invalidModelBuildError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *invalidModelBuildError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// buildV3ModelFromBytes is a helper that parses bytes and builds a v3 model.
// Returns the model and any build errors. The model may be non-nil even when err is non-nil
// (e.g., circular reference warnings), allowing bundling to proceed with warnings.
func buildV3ModelFromBytes(bytes []byte, configuration *datamodel.DocumentConfiguration) (*v3.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return both model and error - caller decides how to handle warnings/errors

// BundleBytes will take a byte slice of an OpenAPI specification and return a bundled version of it.
// This is useful for when you want to take a specification with external references, and you want to bundle it
// into a single document.
//
// This function will 'resolve' all references in the specification and return a single document. The resulting
// document will be a valid OpenAPI specification, containing no references.
//
// Circular references will not be resolved and will be skipped.
func BundleBytes(bytes []byte, configuration *datamodel.DocumentConfiguration) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BundleBytesComposed will take a byte slice of an OpenAPI specification and return a composed bundled version of it.
// this is the same as BundleBytes, but it will compose the bundling instead of inline it.
//
// Composed means that every external file will have references lifted out and added to the `components` section of the document.
// Names will be preserved where possible, conflicts will dealt with by using a delimiter and appending a number.
func BundleBytesComposed(bytes []byte, configuration *datamodel.DocumentConfiguration, compositionConfig *BundleCompositionConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BundleBytesComposedWithOrigins returns a bundled spec with origin tracking for navigation.
// This enables consumers to map bundled components back to their original file locations.
func BundleBytesComposedWithOrigins(bytes []byte, configuration *datamodel.DocumentConfiguration, compositionConfig *BundleCompositionConfig) (*BundleResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BundleDocument will take a v3.Document and return a bundled version of it.
// This is useful for when you want to take a document that has been built
// from a specification with external references, and you want to bundle it
// into a single document.
//
// This function will 'resolve' all references in the specification and return a single document. The resulting
// document will be a valid OpenAPI specification, containing no references.
//
// Circular references will not be resolved and will be skipped.
func BundleDocument(model *v3.Document) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// BundleBytesWithConfig will take a byte slice of an OpenAPI specification and return a bundled version of it,
// with additional configuration options for inline bundling behavior.
//
// Use the BundleInlineConfig to enable features like ResolveDiscriminatorExternalRefs which copies external
// schemas referenced by discriminator mappings to the root document's components section.
func BundleBytesWithConfig(bytes []byte, configuration *datamodel.DocumentConfiguration, bundleConfig *BundleInlineConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BundleDocumentWithConfig will take a v3.Document and return a bundled version of it,
// with additional configuration options for inline bundling behavior.
//
// Use the BundleInlineConfig to enable features like ResolveDiscriminatorExternalRefs which copies external
// schemas referenced by discriminator mappings to the root document's components section.
func BundleDocumentWithConfig(model *v3.Document, bundleConfig *BundleInlineConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BundleCompositionConfig is used to configure the composition of OpenAPI documents when using BundleDocumentComposed.
type BundleCompositionConfig struct {
	Delimiter        string // Delimiter is used to separate clashing names. Defaults to `__`.
	StrictValidation bool   // StrictValidation will cause bundling to fail on invalid OpenAPI specs (e.g. $ref with siblings)
}

// BundleInlineConfig provides configuration options for inline bundling.
//
// Example usage:
//
//	// Inline everything including local refs
//	inlineTrue := true
//	config := &BundleInlineConfig{
//	    InlineLocalRefs: &inlineTrue,
//	}
//	bundled, err := BundleBytesWithConfig(specBytes, docConfig, config)
type BundleInlineConfig struct {
	// ResolveDiscriminatorExternalRefs when true, copies external schemas referenced
	// by discriminator mappings to the root document's components section.
	// This ensures the bundled output is valid and self-contained when discriminators
	// in external files reference other schemas in those external files.
	// Default: false (preserves existing behavior of keeping external refs as-is)
	ResolveDiscriminatorExternalRefs bool

	// InlineLocalRefs controls whether local component references are inlined during bundling.
	// When nil, falls back to DocumentConfiguration.BundleInlineRefs.
	// - false: preserve local refs like #/components/schemas/Pet (discriminator-safe, default behavior)
	// - true: inline all refs including local component refs
	// Default: nil (uses DocumentConfiguration.BundleInlineRefs)
	InlineLocalRefs *bool
}

// BundleDocumentComposed will take a v3.Document and return a composed bundled version of it. Composed means
// that every external file will have references lifted out and added to the `components` section of the document.
// Names will be preserved where possible, conflicts will be appended with a number. If the type of the reference cannot
// be determined, it will be added to the `components` section as a `Schema` type, a warning will be logged.
// The document model will be mutated permanently.
//
// Circular references will not be resolved and will be skipped.
func BundleDocumentComposed(model *v3.Document, compositionConfig *BundleCompositionConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BundleDocumentComposedWithOrigins will take a v3.Document and return a composed bundled version of it
// along with origin tracking information. This allows consumers to map bundled components back to their
// original file locations. The document model will be mutated permanently.
//
// Circular references will not be resolved and will be skipped.
func BundleDocumentComposedWithOrigins(model *v3.Document, compositionConfig *BundleCompositionConfig) (*BundleResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// composeWithOrigins performs composed bundling and returns origin tracking information
func composeWithOrigins(model *v3.Document, compositionConfig *BundleCompositionConfig) (*BundleResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Collect discriminator mappings before ref processing so mapping-only targets can be composed.

// Enqueue mapping targets after cf exists; root-local #/ refs stay in place.

// Refresh indexes in case mapping resolution loaded new ones.

// Remap indexed refs.

// Update discriminator mapping values after component names are final.

// Inline anything that could not be recomposed.

// Rewrite any remaining unindexed refs after mapping resolution loads new indexes.

func compose(model *v3.Document, compositionConfig *BundleCompositionConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Collect discriminator mappings before ref processing so mapping-only targets can be composed.

// Enqueue mapping targets after cf exists; root-local #/ refs stay in place.

// Refresh indexes in case mapping resolution loaded new ones.

// Remap indexed refs.

// Update discriminator mapping values after component names are final.

// Inline anything that could not be recomposed.

// Rewrite any remaining unindexed refs after mapping resolution loads new indexes.

// rewriteInlinedAbsoluteRefs updates absolute $ref values that were resolved by
// the inline fallback after the index's normal rewrite pass has already run.
func rewriteInlinedAbsoluteRefs(rolodex *index.Rolodex, indexes []*index.SpecIndex, inlinedPaths map[string]*yaml.Node) {
	_ = "STUB: not implemented"
	return
}

// inlineRequiredRefs inlines refs that cannot be represented as root components.
func inlineRequiredRefs(required []*processRef, rolodex *index.Rolodex) map[string]*yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// sequencedRefsByFullDefinition buckets refs once for inlineRequiredRefs.
func sequencedRefsByFullDefinition(rolodex *index.Rolodex) map[string][]*index.Reference {
	_ = "STUB: not implemented"
	return nil
}

// inlineProcessRef replaces the source ref node with its resolved target node.
func inlineProcessRef(pr *processRef) *yaml.Node { _ = "STUB: not implemented"; return nil }

// inlineMatchingRefs applies the same inline replacement to repeated matching refs.
func inlineMatchingRefs(pr *processRef, inlinedNode *yaml.Node, refsByDefinition map[string][]*index.Reference) {
	_ = "STUB: not implemented"
	return
}

// resolveBundleInlineConfig resolves the inlineLocalRefs setting from the fallback chain:
// 1. BundleInlineConfig.InlineLocalRefs (explicit per-call)
// 2. DocumentConfiguration.BundleInlineRefs (document-wide default)
// 3. false (system default - preserve local refs)
func resolveBundleInlineConfig(bundleConfig *BundleInlineConfig, docConfig *datamodel.DocumentConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

// system default

func bundleWithConfig(model *v3.Document, config *BundleInlineConfig, docConfig *datamodel.DocumentConfiguration) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// enable bundling mode to preserve local component refs during marshalling
// when inlineLocalRefs is true, skip bundling mode to inline everything

// copy external schemas referenced by discriminator mappings to root components
// ensures bundled output is valid and self-contained

// resolve extension refs before rendering (mutates model's extension nodes in-place)
// extensions are raw yaml nodes that bypass MarshalYAMLInline()

// render inline - discriminator mappings and circular refs are preserved via SchemaProxy.MarshalYAMLInline()

// externalSchemaRef represents an external schema that needs to be copied to the root document's components.
type externalSchemaRef struct {
	idx         *index.SpecIndex // Source index where the schema is defined
	ref         *index.Reference // The reference object
	schemaName  string           // The target name in components
	fullDef     string           // The full definition path (e.g., "/path/to/file.yaml#/components/schemas/Cat")
	originalRef string           // The original reference string (e.g., "#/components/schemas/Cat")
}

// resolveDiscriminatorExternalRefs handles copying external schemas referenced by discriminators
// to the root document's components section and rewrites the references.
func resolveDiscriminatorExternalRefs(model *v3.Document) { _ = "STUB: not implemented"; return }

// Collect all external schemas referenced by discriminators

// Ensure model has Components (buildComponents always succeeds with valid rootIdx,
// and rootIdx must be valid since collectExternalDiscriminatorSchemas would panic otherwise)

// Build existing names map from current components for collision detection

// Copy schemas to components and build ref mapping
// We need to map both local refs (like #/components/schemas/Cat) and
// external refs (like ./external.yaml#/components/schemas/Cat) to the new location

// externalSchemas has unique fullDef values (from map iteration in collectExternalDiscriminatorSchemas)

// Map the local ref format (used in external files)

// Also map external ref formats that might be used in the root document
// e.g., "./vehicles/car.yaml#/components/schemas/Car"
// The external ref format is: relative path from root + JSON pointer

// Calculate relative path from root to external file

// Normalize path separators to forward slashes for cross-platform compatibility
// OpenAPI refs always use forward slashes regardless of OS

// Build external ref format: ./relpath#/components/schemas/Name

// Also try with "./" prefix

// Rewrite discriminator mapping refs and oneOf/anyOf refs

// collectExternalDiscriminatorSchemas identifies external schemas referenced by discriminators
// that need to be copied to the root document's components section.
func collectExternalDiscriminatorSchemas(rolodex *index.Rolodex, rootIdx *index.SpecIndex) []*externalSchemaRef {
	_ = "STUB: not implemented"
	return nil
}

// Use existing infrastructure to collect pinned refs

// Collect from all indexes (root and external)

// Pre-build index lookup map for O(1) lookups instead of O(N) per ref

// Convert pinned refs to externalSchemaRef structs

// Parse the full definition to get the original ref
// Format: "/absolute/path/to/file.yaml#/components/schemas/SchemaName"

// Skip if this is from the root document (not external)

// find the index for this file using pre-built map (O(1) lookup)

// defensive: skip if index not found (shouldn't happen with valid specs)

// find the actual reference - this was already found when pinning

// Extract schema name from the JSON pointer
// e.g., "#/components/schemas/Cat" -> "Cat"

// copySchemaToComponents copies an external schema to the root document's components section.
// Returns the new reference string (e.g., "#/components/schemas/Cat").
// existingNames is updated with the new name to track collisions across multiple calls.
func copySchemaToComponents(model *v3.Document, extSchema *externalSchemaRef, existingNames map[string]bool) string {
	_ = "STUB: not implemented"
	// Build the schema from the YAML node
	// extSchema.ref.Node is always valid (validated when collecting external schemas)
	return ""
}

// Check for naming collisions and get unique name

// Track this name to prevent future collisions

// Add to components

// calculateCollisionNameInline generates a unique name for a schema to avoid collisions.
// It first tries appending the source filename, then falls back to numeric suffixes.
func calculateCollisionNameInline(name, fullDef, delimiter string, existingNames map[string]bool) string {
	_ = "STUB: not implemented"
	// Extract filename from the full definition path
	return ""
}

// Remove extension

// Try filename-based name first

// If filename-based collision exists, try numeric suffixes

// rewriteInlineDiscriminatorRefs updates discriminator mapping refs and oneOf/anyOf refs
// to point to the newly copied component locations.
func rewriteInlineDiscriminatorRefs(rolodex *index.Rolodex, refMapping map[string]string) {
	_ = "STUB: not implemented"
	return
}

// Collect all discriminator mapping nodes

// Update discriminator mapping values

// Also update oneOf/anyOf $ref values in all indexes

// updateOneOfAnyOfRefs recursively walks a YAML node tree to update oneOf/anyOf $ref values.
func updateOneOfAnyOfRefs(n *yaml.Node, refMapping map[string]string) {
	_ = "STUB: not implemented"
	return
}

// First pass: check for discriminator and find oneOf/anyOf

// Update refs in oneOf/anyOf if this schema has a discriminator

// Recursively process all children

// updateUnionRefs updates $ref values in a oneOf or anyOf sequence.
func updateUnionRefs(seq *yaml.Node, refMapping map[string]string) {
	_ = "STUB: not implemented"
	return
}

func collectDiscriminatorMappingValues(idx *index.SpecIndex, n *yaml.Node, pinned map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

func walkDiscriminatorMapping(idx *index.SpecIndex, discriminatorNode *yaml.Node, pinned map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

func walkUnionRefs(idx *index.SpecIndex, seq *yaml.Node, pinned map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

// collectDiscriminatorMappingNodes gathers all discriminator mapping value nodes from the document tree.
func collectDiscriminatorMappingNodes(rolodex *index.Rolodex) []*yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// collectDiscriminatorMappingNodesWithContext gathers all discriminator mapping value nodes
// along with their source index context for proper relative path resolution.
func collectDiscriminatorMappingNodesWithContext(rolodex *index.Rolodex) []*discriminatorMappingWithContext {
	_ = "STUB: not implemented"
	return nil
}

// collectDiscriminatorMappingNodesFromIndexWithContext recursively walks a YAML node tree
// to find discriminator mapping nodes, preserving the source index context.
func collectDiscriminatorMappingNodesFromIndexWithContext(idx *index.SpecIndex, n *yaml.Node, mappings *[]*discriminatorMappingWithContext) {
	_ = "STUB: not implemented"
	return
}

// collectDiscriminatorMappingNodesFromIndex recursively walks a YAML node tree to find discriminator mapping nodes.
func collectDiscriminatorMappingNodesFromIndex(idx *index.SpecIndex, n *yaml.Node, mappingNodes *[]*yaml.Node) {
	_ = "STUB: not implemented"
	return
}

// updateDiscriminatorMappingsComposed updates discriminator mapping references to point to composed component locations.
func updateDiscriminatorMappingsComposed(mappings []*discriminatorMappingWithContext, processedNodes *orderedmap.Map[string, *processRef], rolodex *index.Rolodex) {
	_ = "STUB: not implemented"
	return
}

// Skip external URLs and URNs - they should never be rewritten

// Use the canonicalKey and targetIdx captured before bundling mutates refs.
// Calling SearchIndexForReference again here could return a mutated
// ref.FullDefinition that won't match processedNodes keys.

// If canonicalKey is empty, the mapping wasn't resolved during enqueue.
// Try to resolve it now as a fallback.

// Use the resolved index, not mapping.sourceIdx.

// Gate rewrites on processedNodes presence.
// Only rewrite if the target was actually composed into the bundled output.
// This prevents dangling refs when SearchIndexForReference resolves something
// that never made it into processedNodes (e.g., unprocessed transitive refs).

// Use targetIdx (where the ref actually lives), NOT sourceIdx
