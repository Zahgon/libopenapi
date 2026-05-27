// Copyright 2022-2025 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"log/slog"
)

// schemaIdRegistrationResult holds the result of a schema ID registration attempt.
type schemaIdRegistrationResult struct {
	registered bool   // true if successfully registered
	duplicate  bool   // true if a duplicate was found (first-wins policy applied)
	key        string // the key used for registration
}

// registerSchemaIdToRegistry is the common registration logic for both SpecIndex and Rolodex.
// Returns the registration result. Duplicates are logged but not treated as errors.
func registerSchemaIdToRegistry(
	registry map[string]*SchemaIdEntry,
	entry *SchemaIdEntry,
	logger *slog.Logger,
	registryName string,
) (*schemaIdRegistrationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// copySchemaIdRegistry creates a defensive copy of a schema ID registry.
func copySchemaIdRegistry(registry map[string]*SchemaIdEntry) map[string]*SchemaIdEntry {
	_ = "STUB: not implemented"
	return nil
}
