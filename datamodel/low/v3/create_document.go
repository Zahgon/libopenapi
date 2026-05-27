// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"context"
	"net/url"

	"github.com/pb33f/libopenapi/datamodel"
	"github.com/pb33f/libopenapi/index"
	"go.yaml.in/yaml/v4"
)

type documentTopLevelNode struct {
	key   *yaml.Node
	value *yaml.Node
}

type documentTopLevelNodes struct {
	version           documentTopLevelNode
	jsonSchemaDialect documentTopLevelNode
	self              documentTopLevelNode
	info              documentTopLevelNode
	servers           documentTopLevelNode
	tags              documentTopLevelNode
	components        documentTopLevelNode
	security          documentTopLevelNode
	externalDocs      documentTopLevelNode
	paths             documentTopLevelNode
	webhooks          documentTopLevelNode
}

func selectDocumentNode(root *yaml.Node, preferred documentTopLevelNode, label string, topOnly bool) documentTopLevelNode {
	_ = "STUB: not implemented"
	return *new(documentTopLevelNode)
}

func collectDocumentTopLevelNodes(root *yaml.Node) documentTopLevelNodes {
	_ = "STUB: not implemented"
	return *new(documentTopLevelNodes)
}

// CreateDocument will create a new Document instance from the provided SpecInfo.
//
// Deprecated: Use CreateDocumentFromConfig instead. This function will be removed in a later version, it
// defaults to allowing file and remote references, and does not support relative file references.
func CreateDocument(info *datamodel.SpecInfo) (*Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateDocumentFromConfig Create a new document from the provided SpecInfo and DocumentConfiguration pointer.
func CreateDocumentFromConfig(info *datamodel.SpecInfo, config *datamodel.DocumentConfiguration) (*Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createDocument(info *datamodel.SpecInfo, config *datamodel.DocumentConfiguration) (*Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create an index config and shadow the document configuration.

// handle $self field for OpenAPI 3.2+ documents

// log error but continue with original config

// store error in spec info for later retrieval

// validate http/https URLs

// conflict detected

// use config BaseURL (programmatic control trumps document)

// use $self as BaseURL

// for non-http URLs (like file:// or custom schemes), use as-is if no conflict

// If basePath is provided, add a local filesystem to the rolodex.

// if a supplied local filesystem is provided, add it to the rolodex.

// wrap a plain fs.FS so it can be indexed.

// create a local filesystem

// add the filesystem to the rolodex

// Only create a remote filesystem when the caller explicitly allows remote references.

// create a remote filesystem

// add to the rolodex

// index the rolodex

// index all the things.

// check for circular references

// extract errors

// set root index.

// if set, extract jsonSchemaDialect (3.1)

// if set, extract $self (3.2)

func extractInfo(ctx context.Context, root *yaml.Node, nodes documentTopLevelNodes, doc *Document, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func extractSecurity(ctx context.Context, root *yaml.Node, nodes documentTopLevelNodes, doc *Document, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func extractExternalDocs(ctx context.Context, root *yaml.Node, nodes documentTopLevelNodes, doc *Document, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func extractComponents(ctx context.Context, root *yaml.Node, nodes documentTopLevelNodes, doc *Document, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func extractServers(ctx context.Context, root *yaml.Node, nodes documentTopLevelNodes, doc *Document, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func extractTags(ctx context.Context, root *yaml.Node, nodes documentTopLevelNodes, doc *Document, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func extractPaths(ctx context.Context, root *yaml.Node, nodes documentTopLevelNodes, doc *Document, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func extractWebhooks(ctx context.Context, root *yaml.Node, nodes documentTopLevelNodes, doc *Document, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func urlWithoutTrailingSlash(u *url.URL) *url.URL { _ = "STUB: not implemented"; return nil }
