// Copyright 2022-2026 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package index

import (
	"go.yaml.in/yaml/v4"
)

// SetCircularReferences sets the circular reference results for this index.
func (index *SpecIndex) SetCircularReferences(refs []*CircularReferenceResult) {
	_ = "STUB: not implemented"
	return
}

// GetCircularReferences returns all circular references found during resolution.
func (index *SpecIndex) GetCircularReferences() []*CircularReferenceResult {
	_ = "STUB: not implemented"
	return nil
}

// GetTagCircularReferences returns circular references found in tag parent hierarchies.
func (index *SpecIndex) GetTagCircularReferences() []*CircularReferenceResult {
	_ = "STUB: not implemented"
	return nil
}

// SetIgnoredPolymorphicCircularReferences sets circular references that were ignored because they
// involve polymorphic keywords (allOf, oneOf, anyOf).
func (index *SpecIndex) SetIgnoredPolymorphicCircularReferences(refs []*CircularReferenceResult) {
	_ = "STUB: not implemented"
	return
}

// SetIgnoredArrayCircularReferences sets circular references that were ignored because they
// involve array items.
func (index *SpecIndex) SetIgnoredArrayCircularReferences(refs []*CircularReferenceResult) {
	_ = "STUB: not implemented"
	return
}

// GetIgnoredPolymorphicCircularReferences returns circular references that were ignored because
// they involve polymorphic keywords.
func (index *SpecIndex) GetIgnoredPolymorphicCircularReferences() []*CircularReferenceResult {
	_ = "STUB: not implemented"
	return nil
}

// GetIgnoredArrayCircularReferences returns circular references that were ignored because they
// involve array items.
func (index *SpecIndex) GetIgnoredArrayCircularReferences() []*CircularReferenceResult {
	_ = "STUB: not implemented"
	return nil
}

// GetPathsNode returns the raw YAML node for the top-level "paths" object.
func (index *SpecIndex) GetPathsNode() *yaml.Node { _ = "STUB: not implemented"; return nil }

// GetDiscoveredReferences returns all deduplicated references found during extraction,
// keyed by their full definition path.
func (index *SpecIndex) GetDiscoveredReferences() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil

	// GetPolyReferences returns all polymorphic references (allOf, oneOf, anyOf) keyed by definition.
}

func (index *SpecIndex) GetPolyReferences() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetPolyAllOfReferences returns all references found under allOf keywords.
func (index *SpecIndex) GetPolyAllOfReferences() []*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetPolyAnyOfReferences returns all references found under anyOf keywords.
func (index *SpecIndex) GetPolyAnyOfReferences() []*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetPolyOneOfReferences returns all references found under oneOf keywords.
func (index *SpecIndex) GetPolyOneOfReferences() []*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetAllCombinedReferences returns a merged map of all standard and polymorphic references.
func (index *SpecIndex) GetAllCombinedReferences() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetRefsByLine returns a map of reference definition to the set of line numbers where it appears.
func (index *SpecIndex) GetRefsByLine() map[string]map[int]bool {
	_ = "STUB: not implemented"
	return nil

	// GetLinesWithReferences returns a set of line numbers that contain at least one reference.
}

func (index *SpecIndex) GetLinesWithReferences() map[int]bool {
	_ = "STUB: not implemented"
	return nil

	// GetMappedReferences returns all resolved component references keyed by definition path.
}

func (index *SpecIndex) GetMappedReferences() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil

	// SetMappedReferences replaces the mapped references for this index.
}

func (index *SpecIndex) SetMappedReferences(mappedRefs map[string]*Reference) {
	_ = "STUB: not implemented"
	return
}

// GetRawReferencesSequenced returns all raw references in the order they were scanned.
func (index *SpecIndex) GetRawReferencesSequenced() []*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensionRefsSequenced returns only references that appear under x-* extension paths,
// in scan order.
func (index *SpecIndex) GetExtensionRefsSequenced() []*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetMappedReferencesSequenced returns all resolved component references in deterministic order.
func (index *SpecIndex) GetMappedReferencesSequenced() []*ReferenceMapped {
	_ = "STUB: not implemented"
	return nil
}

// GetOperationParameterReferences returns parameters keyed by path, then HTTP method, then parameter name.
func (index *SpecIndex) GetOperationParameterReferences() map[string]map[string]map[string][]*Reference {
	_ = "STUB: not implemented"
	return nil

	// GetAllSchemas returns all schemas (component, inline, and reference) sorted by line number.
}

func (index *SpecIndex) GetAllSchemas() []*Reference { _ = "STUB: not implemented"; return nil }

// GetAllInlineSchemaObjects returns all inline schema definitions that are objects.
func (index *SpecIndex) GetAllInlineSchemaObjects() []*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetAllInlineSchemas returns all inline schema definitions found during extraction.
func (index *SpecIndex) GetAllInlineSchemas() []*Reference { _ = "STUB: not implemented"; return nil }

// GetAllReferenceSchemas returns all schema definitions that are $ref references.
func (index *SpecIndex) GetAllReferenceSchemas() []*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetAllComponentSchemas returns all component schema definitions, converting from the
// internal sync.Map on first access and caching the result.
func (index *SpecIndex) GetAllComponentSchemas() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetAllSecuritySchemes returns all security scheme definitions from the components section.
func (index *SpecIndex) GetAllSecuritySchemes() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetAllHeaders returns all header definitions from the components section.
func (index *SpecIndex) GetAllHeaders() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil

	// GetAllExternalDocuments returns all external document references found in the specification.
}

func (index *SpecIndex) GetAllExternalDocuments() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetAllExamples returns all example definitions from the components section.
func (index *SpecIndex) GetAllExamples() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil

	// GetAllDescriptions returns all description nodes found during indexing.
}

func (index *SpecIndex) GetAllDescriptions() []*DescriptionReference {
	_ = "STUB: not implemented"
	return nil
}

// GetAllEnums returns all enum definitions found during indexing.
func (index *SpecIndex) GetAllEnums() []*EnumReference { _ = "STUB: not implemented"; return nil }

// GetAllObjectsWithProperties returns all objects that have a "properties" keyword.
func (index *SpecIndex) GetAllObjectsWithProperties() []*ObjectReference {
	_ = "STUB: not implemented"
	return nil
}

// GetAllSummaries returns all summary nodes found during indexing.
func (index *SpecIndex) GetAllSummaries() []*DescriptionReference {
	_ = "STUB: not implemented"
	return nil

	// GetAllRequestBodies returns all request body definitions from the components section.
}

func (index *SpecIndex) GetAllRequestBodies() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetAllLinks returns all link definitions from the components section.
func (index *SpecIndex) GetAllLinks() map[string]*Reference { _ = "STUB: not implemented"; return nil }

// GetAllParameters returns all parameter definitions from the components section.
func (index *SpecIndex) GetAllParameters() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil

	// GetAllResponses returns all response definitions from the components section.
}

func (index *SpecIndex) GetAllResponses() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil

	// GetAllCallbacks returns all callback definitions from the components section.
}

func (index *SpecIndex) GetAllCallbacks() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil

	// GetAllComponentPathItems returns all path item definitions from the components section.
}

func (index *SpecIndex) GetAllComponentPathItems() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetInlineOperationDuplicateParameters returns parameters with duplicate names found inline in operations.
func (index *SpecIndex) GetInlineOperationDuplicateParameters() map[string][]*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetReferencesWithSiblings returns references that have sibling properties alongside the $ref keyword.
func (index *SpecIndex) GetReferencesWithSiblings() map[string]Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetAllReferences returns all deduplicated references found during extraction.
func (index *SpecIndex) GetAllReferences() map[string]*Reference {
	_ = "STUB: not implemented"
	return nil

	// GetAllSequencedReferences returns all raw references in scan order.
}

func (index *SpecIndex) GetAllSequencedReferences() []*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetSchemasNode returns the raw YAML node for the components/schemas (or definitions) section.
func (index *SpecIndex) GetSchemasNode() *yaml.Node { _ = "STUB: not implemented"; return nil }

// GetParametersNode returns the raw YAML node for the components/parameters section.
func (index *SpecIndex) GetParametersNode() *yaml.Node { _ = "STUB: not implemented"; return nil }

// GetReferenceIndexErrors returns any errors that occurred during reference extraction.
func (index *SpecIndex) GetReferenceIndexErrors() []error { _ = "STUB: not implemented"; return nil }

// GetOperationParametersIndexErrors returns any errors found when scanning operation parameters.
func (index *SpecIndex) GetOperationParametersIndexErrors() []error {
	_ = "STUB: not implemented"
	return nil
}

// GetAllPaths returns all path items keyed by path, then HTTP method.
func (index *SpecIndex) GetAllPaths() map[string]map[string]*Reference {
	_ = "STUB: not implemented"
	return nil

	// GetOperationTags returns tags keyed by path, then HTTP method.
}

func (index *SpecIndex) GetOperationTags() map[string]map[string][]*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetAllParametersFromOperations returns all parameters keyed by path, HTTP method, then parameter name.
func (index *SpecIndex) GetAllParametersFromOperations() map[string]map[string]map[string][]*Reference {
	_ = "STUB: not implemented"
	return nil

	// GetRootSecurityReferences returns references from the top-level security requirement array.
}

func (index *SpecIndex) GetRootSecurityReferences() []*Reference {
	_ = "STUB: not implemented"
	return nil

	// GetSecurityRequirementReferences returns security requirements keyed by security scheme name.
}

func (index *SpecIndex) GetSecurityRequirementReferences() map[string]map[string][]*Reference {
	_ = "STUB: not implemented"
	return nil
}

// GetRootSecurityNode returns the raw YAML node for the top-level "security" array.
func (index *SpecIndex) GetRootSecurityNode() *yaml.Node { _ = "STUB: not implemented"; return nil }

// GetRootServersNode returns the raw YAML node for the top-level "servers" array.
func (index *SpecIndex) GetRootServersNode() *yaml.Node { _ = "STUB: not implemented"; return nil }

// GetAllRootServers returns all server references from the top-level "servers" array.
func (index *SpecIndex) GetAllRootServers() []*Reference { _ = "STUB: not implemented"; return nil }

// GetAllOperationsServers returns server references keyed by path, then HTTP method.
func (index *SpecIndex) GetAllOperationsServers() map[string]map[string][]*Reference {
	_ = "STUB: not implemented"
	return nil

	// SetAllowCircularReferenceResolving sets whether circular references should be resolved
	// instead of returning an error.
}

func (index *SpecIndex) SetAllowCircularReferenceResolving(allow bool) {
	_ = "STUB: not implemented"
	return
}

// AllowCircularReferenceResolving returns whether circular reference resolving is enabled.
func (index *SpecIndex) AllowCircularReferenceResolving() bool {
	_ = "STUB: not implemented"
	return false
}

func (index *SpecIndex) checkPolymorphicNode(name string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// RegisterSchemaId registers a JSON Schema $id entry in this index's local registry.
func (index *SpecIndex) RegisterSchemaId(entry *SchemaIdEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSchemaById looks up a schema by its resolved $id URI in this index's local registry.
func (index *SpecIndex) GetSchemaById(uri string) *SchemaIdEntry {
	_ = "STUB: not implemented"
	return nil
}

// GetAllSchemaIds returns a copy of all $id entries registered in this index.
func (index *SpecIndex) GetAllSchemaIds() map[string]*SchemaIdEntry {
	_ = "STUB: not implemented"
	return nil
}
