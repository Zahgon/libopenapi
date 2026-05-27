// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package libopenapi

// ClearAllCaches resets every global in-process cache in libopenapi.
// Call this between document lifecycles in long-running processes
// (servers, CLI tools that process many specs) to release memory that
// would otherwise accumulate and never be garbage-collected.
func ClearAllCaches() { _ = "STUB: not implemented"; return }

// hashCache + indexCollectionCache
// SchemaQuickHashMap
// nodeHashCache

// Drain sync.Pool instances that hold *yaml.Node pointers.
// Pooled slices/maps keep the entire YAML parse tree alive.
