package index

// ResolveReferenceValue resolves a reference string to a decoded value.
//
// Resolution order:
//  1. Resolve using SpecIndex when available.
//  2. Fallback to local JSON pointer resolution (e.g. "#/components/schemas/Foo")
//     using getDocData when provided.
//
// Returns nil when the reference cannot be resolved.
func ResolveReferenceValue(ref string, specIndex *SpecIndex, getDocData func() map[string]interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Fallback parser only supports local JSON pointers ("#" root or "#/...").

func resolveLocalJSONPointer(docData map[string]interface{}, ref string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func decodeJSONPointerToken(token string) string { _ = "STUB: not implemented"; return "" }
