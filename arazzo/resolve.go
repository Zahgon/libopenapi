// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	high "github.com/pb33f/libopenapi/datamodel/high/arazzo"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

var resolveFilepathAbs = filepath.Abs

// OpenAPIDocumentFactory creates a parsed OpenAPI document from raw bytes.
// The sourceURL provides location context for relative reference resolution.
type OpenAPIDocumentFactory func(sourceURL string, bytes []byte) (*v3high.Document, error)

// ArazzoDocumentFactory creates a parsed Arazzo document from raw bytes.
// The sourceURL provides location context for relative reference resolution.
type ArazzoDocumentFactory func(sourceURL string, bytes []byte) (*high.Arazzo, error)

// ResolveConfig configures how source descriptions are resolved.
type ResolveConfig struct {
	OpenAPIFactory OpenAPIDocumentFactory // Creates *v3high.Document from bytes
	ArazzoFactory  ArazzoDocumentFactory  // Creates *high.Arazzo from bytes
	BaseURL        string
	HTTPHandler    func(url string) ([]byte, error)
	HTTPClient     *http.Client
	FSRoots        []string

	Timeout        time.Duration // Per-source fetch timeout (default: 30s)
	MaxBodySize    int64         // Max response body in bytes (default: 10MB)
	AllowedSchemes []string      // URL scheme allowlist (default: ["https", "http", "file"])
	AllowedHosts   []string      // Host allowlist (nil = allow all)
	MaxSources     int           // Max source descriptions to resolve (default: 50)
}

// ResolvedSource represents a successfully resolved source description.
type ResolvedSource struct {
	Name            string           // SourceDescription name
	URL             string           // Resolved URL
	Type            string           // "openapi" or "arazzo"
	OpenAPIDocument *v3high.Document // Non-nil when Type == "openapi"
	ArazzoDocument  *high.Arazzo     // Non-nil when Type == "arazzo"
}

// ResolveSources resolves all source descriptions in an Arazzo document.
func ResolveSources(doc *high.Arazzo, config *ResolveConfig) ([]*ResolvedSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply defaults

// 10MB

// Default per spec

// Auto-attach OpenAPI source documents to the Arazzo model so that
// validation and the engine can resolve operation references without
// the caller needing to wire this up manually.

func parseAndResolveSourceURL(rawURL, base string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Detect Windows absolute paths (e.g. "C:\Users\..." or "D:/foo/bar").
// url.Parse misinterprets the drive letter as a URL scheme ("c:", "d:").
// A single-letter scheme is always a Windows drive letter; real URL schemes
// are at least two characters. Use strings.ReplaceAll instead of
// filepath.ToSlash so backslashes are normalized on all platforms.

// Resolve relative URLs against BaseURL when provided.

func validateSourceURL(sourceURL *url.URL, config *ResolveConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func fetchSourceBytes(sourceURL *url.URL, config *ResolveConfig) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// On Windows, file URLs without a leading slash (e.g. "file://C:/path")
// cause url.Parse to place the drive letter in Host ("C:") and strip it
// from Path ("/path"). Reconstruct the full path.

func fetchHTTPSourceBytes(sourceURL string, config *ResolveConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getResolveHTTPClient(config *ResolveConfig) *http.Client {
	_ = "STUB: not implemented"
	return nil
}

func readFileWithLimit(path string, maxBytes int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveFilePath(path string, roots []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If no roots are configured, resolve relative paths from the current working directory.

// Absolute paths must be inside one of the configured roots.
// Canonicalize the cleaned path for comparison only (resolves Windows 8.3
// short names and macOS /var -> /private/var symlinks) so that the path
// matches canonicalRoots. The original cleaned path is returned to callers.

// Relative paths are resolved against each root in order.
// Use absRoots for building candidates (preserves original paths) but
// canonicalRoots for security checks.

// isPathWithinRoots checks whether path falls inside at least one of the given roots.
// Both path and roots must be absolute paths; no filepath.Abs calls are made here
// since callers already guarantee absolute inputs.
func isPathWithinRoots(path string, roots []string) bool { _ = "STUB: not implemented"; return false }

func canonicalizeRoots(roots []string) []string { _ = "STUB: not implemented"; return nil }

func ensureResolvedPathWithinRoots(path string, roots []string) error {
	_ = "STUB: not implemented"
	return nil
}

func containsFold(values []string, value string) bool { _ = "STUB: not implemented"; return false }
