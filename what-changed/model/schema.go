// Copyright 2022-2025 Princess Beef Heavy Industries, LLC / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"sync"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// SchemaChanges represent all changes to a base.Schema OpenAPI object. These changes are represented
// by all versions of OpenAPI.
//
// Any additions or removals to slice based results will be recorded in the PropertyChanges of the parent
// changes, and not the child for example, adding a new schema to `anyOf` will create a new change result in
// PropertyChanges.Changes, and not in the AnyOfChanges property.
type SchemaChanges struct {
	*PropertyChanges
	DiscriminatorChanges        *DiscriminatorChanges     `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	AllOfChanges                []*SchemaChanges          `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	AnyOfChanges                []*SchemaChanges          `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	OneOfChanges                []*SchemaChanges          `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	PrefixItemsChanges          []*SchemaChanges          `json:"prefixItems,omitempty" yaml:"prefixItems,omitempty"`
	NotChanges                  *SchemaChanges            `json:"not,omitempty" yaml:"not,omitempty"`
	ItemsChanges                *SchemaChanges            `json:"items,omitempty" yaml:"items,omitempty"`
	SchemaPropertyChanges       map[string]*SchemaChanges `json:"properties,omitempty" yaml:"properties,omitempty"`
	ExternalDocChanges          *ExternalDocChanges       `json:"externalDoc,omitempty" yaml:"externalDoc,omitempty"`
	XMLChanges                  *XMLChanges               `json:"xml,omitempty" yaml:"xml,omitempty"`
	ExtensionChanges            *ExtensionChanges         `json:"extensions,omitempty" yaml:"extensions,omitempty"`
	AdditionalPropertiesChanges *SchemaChanges            `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`

	// 3.1 specifics
	IfChanges                    *SchemaChanges            `json:"if,omitempty" yaml:"if,omitempty"`
	ElseChanges                  *SchemaChanges            `json:"else,omitempty" yaml:"else,omitempty"`
	ThenChanges                  *SchemaChanges            `json:"then,omitempty" yaml:"then,omitempty"`
	PropertyNamesChanges         *SchemaChanges            `json:"propertyNames,omitempty" yaml:"propertyNames,omitempty"`
	ContainsChanges              *SchemaChanges            `json:"contains,omitempty" yaml:"contains,omitempty"`
	UnevaluatedItemsChanges      *SchemaChanges            `json:"unevaluatedItems,omitempty" yaml:"unevaluatedItems,omitempty"`
	UnevaluatedPropertiesChanges *SchemaChanges            `json:"unevaluatedProperties,omitempty" yaml:"unevaluatedProperties,omitempty"`
	DependentSchemasChanges      map[string]*SchemaChanges `json:"dependentSchemas,omitempty" yaml:"dependentSchemas,omitempty"`
	DependentRequiredChanges     []*Change                 `json:"dependentRequired,omitempty" yaml:"dependentRequired,omitempty"`
	PatternPropertiesChanges     map[string]*SchemaChanges `json:"patternProperties,omitempty" yaml:"patternProperties,omitempty"`
	ContentSchemaChanges         *SchemaChanges            `json:"contentSchema,omitempty" yaml:"contentSchema,omitempty"`
	VocabularyChanges            []*Change                 `json:"$vocabulary,omitempty" yaml:"$vocabulary,omitempty"`
}

func (s *SchemaChanges) GetPropertyChanges() []*Change { _ = "STUB: not implemented"; return nil }

// GetAllChanges returns a slice of all changes made between Responses objects
func (s *SchemaChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns a count of the total number of changes made to this schema and all sub-schemas
func (s *SchemaChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes made to this schema and all sub-schemas.
func (s *SchemaChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// Count breaking changes in dependent required changes

// CompareSchemas accepts a left and right SchemaProxy and checks for changes. If anything is found, returns
// a pointer to SchemaChanges, otherwise returns nil
func CompareSchemas(l, r *base.SchemaProxy) *SchemaChanges { _ = "STUB: not implemented"; return nil }

// Added

// Removed

// if left proxy is a reference and right is a reference (we won't recurse into circular references here)

// points to the same schema

// check if this is a circular ref.

// if we have a circular reference, we can't do any more work here.

// local reference doesn't need following

// continue on because the external references are the same and we need to check things going forward.

// references are different, that's all we care to know.

// check if this is a circular ref.

// if we have a circular reference, we can't do any more work here.

// changed from inline to ref

// check if the referenced schema matches or not
// https://github.com/pb33f/libopenapi/issues/218

// check if this is a circular ref.

// if we have a circular reference, we can't do any more work here.

// changed from ref to inline

// check if the referenced schema matches or not
// https://github.com/pb33f/libopenapi/issues/218

// check if this is a circular ref.

// if we have a circular reference, we can't do any more work here.

// there is no point going on, we know nothing changed!

// check XML

// check examples

// check schema core properties for changes.

// now for the confusing part, there is also a schema's 'properties' property to parse.
// inception, eat your heart out.

// Check dependent required changes

// done

func checkSchemaXML(lSchema *base.Schema, rSchema *base.Schema, changes *[]*Change, sc *SchemaChanges) {
	_ = "STUB: not implemented"
	// XML removed
	return
}

// XML added

// compare XML

func checkMappedSchemaOfASchema(
	lSchema,
	rSchema *orderedmap.Map[low.KeyReference[string], low.ValueReference[*base.SchemaProxy]],
	changes *[]*Change,
) map[string]*SchemaChanges {
	_ = "STUB: not implemented"
	return nil
	// concurrent-safe map
}

// Convert the sync.Map into a regular map[string]*SchemaChanges.

func buildProperty(lProps, rProps []string, lEntities, rEntities map[string]*base.SchemaProxy,
	propChanges *sync.Map, changes *[]*Change, rKeyNodes, lKeyNodes map[string]*yaml.Node,
) {
	_ = "STUB: not implemented"
	return
}

// left and right equal.

// Handle keys that do not match.

// new property added.

// things removed

// stuff added

// Wait for all property comparisons to finish.

func checkSchemaPropertyChanges(
	lSchema *base.Schema,
	rSchema *base.Schema,
	lProxy *base.SchemaProxy,
	rProxy *base.SchemaProxy,
	changes *[]*Change, sc *SchemaChanges, skipSimpleScalarUnionDiff bool,
) {
	_ = "STUB: not implemented"
	return
}

// $schema (breaking change)

// ExclusiveMaximum

// ExclusiveMinimum

// Type

// Title

// MultipleOf

// Maximum

// Minimum

// MaxLength

// MinLength

// Pattern

// Format

// MaxItems

// MinItems

// MaxProperties

// MinProperties

// UniqueItems

// AdditionalProperties

// added AdditionalProperties

// removed AdditionalProperties

// Description

// ContentEncoding

// ContentMediaType

// Default

// Const

// Nullable

// ReadOnly

// WriteOnly

// Example

// Deprecated

// Required

// Enums

// Discriminator

// check if hash matches, if not then compare.

// added Discriminator

// removed Discriminator

// ExternalDocs

// check if hash matches, if not then compare.

// added ExternalDocs

// removed ExternalDocs

// 3.1 properties
// If

// added If

// removed If

// Else

// added Else

// removed Else

// Then

// added Then

// removed Then

// PropertyNames

// added PropertyNames

// removed PropertyNames

// Contains

// added Contains

// removed Contains

// UnevaluatedItems

// added UnevaluatedItems

// removed UnevaluatedItems

// UnevaluatedProperties

// added UnevaluatedProperties

// removed UnevaluatedProperties

// Not

// added Not

// removed not

// items

// added Items

// removed Items

// $dynamicAnchor (JSON Schema 2020-12)

// $dynamicRef (JSON Schema 2020-12)

// $id (JSON Schema 2020-12)

// $comment (JSON Schema 2020-12)

// contentSchema (JSON Schema 2020-12) - recursive schema comparison

// $vocabulary (JSON Schema 2020-12) - map comparison
// note: vocabulary changes are stored in VocabularyChanges and counted separately
// in TotalChanges(), so they should NOT be appended to the main changes slice

// check extensions

// check core properties

// Post-process: Update context line numbers for Type changes to use schema KeyNode for better context
// This provides line where "schema:" is defined, not "type: value"

// found the type change, no need to continue

func schemasUseEquivalentSimpleScalarUnion(l, r *base.SchemaProxy) bool {
	_ = "STUB: not implemented"
	return false
}

func schemaComparisonViewForSimpleAllOfObject(proxy *base.SchemaProxy, schema *base.Schema) *base.Schema {
	_ = "STUB: not implemented"
	return nil
}

func isSimpleAllOfObjectSchema(proxy *base.SchemaProxy, schema *base.Schema) bool {
	_ = "STUB: not implemented"
	return false
}

var simpleAllOfObjectBranchKeys = map[string]struct{}{
	v3.DescriptionLabel: {},
	v3.PropertiesLabel:  {},
	v3.RequiredLabel:    {},
	v3.TitleLabel:       {},
	v3.TypeLabel:        {},
}

func mergeSimpleAllOfObjectSchemaView(schema *base.Schema) (*base.Schema, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func mergeSimpleAllOfObjectType(schema *base.Schema) (low.NodeReference[base.SchemaDynamicValue[string, []low.ValueReference[string]]], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func mergeCompatibleStringNodeReference(
	baseRef low.NodeReference[string],
	allOf []low.ValueReference[*base.SchemaProxy],
	selector func(*base.Schema) low.NodeReference[string],
) (low.NodeReference[string], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func mergeSimpleAllOfObjectProperties(
	schema *base.Schema,
) (low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*base.SchemaProxy]]], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func mergeSimpleAllOfRequired(schema *base.Schema) low.NodeReference[[]low.ValueReference[string]] {
	_ = "STUB: not implemented"
	return nil
}

func schemaPairUsesEquivalentSimpleScalarUnion(typeProxy, anyOfProxy *base.SchemaProxy) bool {
	_ = "STUB: not implemented"
	return false
}

func isPureTypeArraySchema(proxy *base.SchemaProxy) bool { _ = "STUB: not implemented"; return false }

func isPureAnyOfUnionSchema(proxy *base.SchemaProxy) bool { _ = "STUB: not implemented"; return false }

func schemaNodeHasSingleKey(node *yaml.Node, key string) bool {
	_ = "STUB: not implemented"
	return false
}

func schemaNodeHasOnlyAllowedKeys(node *yaml.Node, allowed map[string]struct{}) bool {
	_ = "STUB: not implemented"
	return false
}

func extractTypeArraySet(proxy *base.SchemaProxy) (map[string]struct{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func extractSimpleAnyOfTypeSet(anyOf []low.ValueReference[*base.SchemaProxy]) (map[string]struct{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func extractScalarTypeName(node *yaml.Node) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func checkExamples(lSchema *base.Schema, rSchema *base.Schema, changes *[]*Change) {
	_ = "STUB: not implemented"
	return
}

// check examples (3.1+)

// create keys by hashing values

// if examples equal lengths, check for equality

// examples were removed.

// examples were added

func extractSchemaChanges(
	lSchema []low.ValueReference[*base.SchemaProxy],
	rSchema []low.ValueReference[*base.SchemaProxy],
	label string,
	sc *[]*SchemaChanges,
	changes *[]*Change,
) {
	_ = "STUB: not implemented"
	// if there is nothing here, there is nothing to do.
	return
}

// create hash key maps to check equality

// check for identical lengths

// keys are different, which means there are changes.

// things were removed

// things were added

type schemaCompositionEntry struct {
	identity  string
	stableKey string
	position  int
	proxy     *base.SchemaProxy
}

type schemaCompositionPair struct {
	left    schemaCompositionEntry
	right   schemaCompositionEntry
	changes *SchemaChanges
}

type schemaCompositionPairCandidate struct {
	left            schemaCompositionEntry
	right           schemaCompositionEntry
	changes         *SchemaChanges
	totalChanges    int
	breakingChanges int
	stableKeyMatch  bool
}

func isOrderInsensitiveSchemaCompositionLabel(label string) bool {
	_ = "STUB: not implemented"
	return false
}

func schemaCompositionEntryIdentity(proxy *base.SchemaProxy) string {
	_ = "STUB: not implemented"
	return ""
}

func buildSchemaCompositionEntries(schema []low.ValueReference[*base.SchemaProxy]) []schemaCompositionEntry {
	_ = "STUB: not implemented"
	return nil
}

func schemaCompositionEntryStableKey(proxy *base.SchemaProxy) string {
	_ = "STUB: not implemented"
	return ""
}

func pairExactSchemaCompositionEntries(
	leftEntries, rightEntries []schemaCompositionEntry,
) ([]schemaCompositionPair, []schemaCompositionEntry, []schemaCompositionEntry) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func buildSchemaCompositionPairCandidateMatrix(
	rowEntries, columnEntries []schemaCompositionEntry,
	swapPairDirection bool,
) [][]schemaCompositionPairCandidate {
	_ = "STUB: not implemented"
	return nil
}

func schemaCompositionPairingSignature(candidates []schemaCompositionPairCandidate) []int {
	_ = "STUB: not implemented"
	return nil
}

func schemaCompositionPairingIsBetter(
	candidates []schemaCompositionPairCandidate,
	best []schemaCompositionPairCandidate,
) bool {
	_ = "STUB: not implemented"
	return false
}

func selectBestSchemaCompositionPairCandidates(
	matrix [][]schemaCompositionPairCandidate,
) []schemaCompositionPairCandidate {
	_ = "STUB: not implemented"
	return nil
}

func pairUnmatchedSchemaCompositionEntries(
	leftEntries, rightEntries []schemaCompositionEntry,
) []schemaCompositionPair {
	_ = "STUB: not implemented"
	return nil
}

func schemaCompositionChangeBreaking(label string, changeType int) bool {
	_ = "STUB: not implemented"
	return false
}

func extractOrderInsensitiveSchemaChanges(
	lSchema []low.ValueReference[*base.SchemaProxy],
	rSchema []low.ValueReference[*base.SchemaProxy],
	label string,
	sc *[]*SchemaChanges,
	changes *[]*Change,
) {
	_ = "STUB: not implemented"
	return
}

// checkDependentRequiredChanges compares two DependentRequired maps and returns any changes found
func checkDependentRequiredChanges(
	left, right *orderedmap.Map[low.KeyReference[string], low.ValueReference[[]string]],
) []*Change {
	_ = "STUB: not implemented"
	// If both are nil, no changes
	return nil
}

// Build left map

// Build right map

// Check for property additions and modifications

// Property exists in both, check if requirements changed

// Property added

// Check for property removals

// slicesEqual compares two string slices for equality (order matters)
func slicesEqual(a, b []string) bool { _ = "STUB: not implemented"; return false }

// getNodeForProperty gets the YAML node for a specific property in a DependentRequired map
func getNodeForProperty(depMap *orderedmap.Map[low.KeyReference[string], low.ValueReference[[]string]], prop string) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// checkVocabularyChanges compares $vocabulary maps and returns a list of changes.
// the caller is responsible for appending the returned changes to their main changes slice.
func checkVocabularyChanges(lVocab, rVocab *orderedmap.Map[low.KeyReference[string], low.ValueReference[bool]]) []*Change {
	_ = "STUB: not implemented"
	return nil
}

// pre-allocate maps with size hints for better memory efficiency

// pre-allocate result slice with reasonable capacity

// check for removed or modified vocabularies

// vocabulary exists in both - check if value changed

// vocabulary was removed

// check for added vocabularies

// vocabulary was added
