package utils

func ReplaceWindowsDriveWithLinuxPath(path string) string { _ = "STUB: not implemented"; return "" }

// CheckPathOverlap joins pathA and pathB while avoiding duplicated overlapping segments.
// It tolerates mixed separators in the inputs and uses OS-specific path comparison rules.
// It also handles leading ".." segments in pathB by first resolving them against pathA
// before checking for overlap, preventing path doubling issues.
func CheckPathOverlap(pathA, pathB, sep string) string { _ = "STUB: not implemented"; return "" }

// Split on both separators so mixed-path inputs are handled safely.

// preserve path prefix (absolute path marker or Windows drive letter)

// If pathA has a Windows drive letter, drop it from aParts
// to avoid duplicating it when rebuilding with prefix.

// Handle leading ".." segments in pathB by going up pathA.
// This prevents path doubling when resolving refs like "../components/schemas/X.yaml"
// from a base path like "/path/to/components/schemas".

// Go up one directory
// Remove the ".."

// rebuild pathA if we resolved any ".." segments

// Use the OS separator to join parts, then add a prefix

// find the longest suffix of aParts that matches the prefix of bParts.

// sep is used to build the tail; filepath.Join normalizes separators for the OS.
