// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

// Package libopenapi is a library containing tools for reading and in and manipulating Swagger (OpenAPI 2) and OpenAPI 3+
// specifications into strongly typed documents. These documents have two APIs, a high level (porcelain) and a
// low level (plumbing).
//
// Every single type has a 'GoLow()' method that drops down from the high API to the low API. Once in the low API,
// the entire original document data is available, including all comments, line and column numbers for keys and values.
//
// There are two steps to creating a using Document. First, create a new Document using the NewDocument() method
// and pass in a specification []byte array that contains the OpenAPI Specification. It doesn't matter if YAML or JSON
// are used.
package libopenapi

import (
	"github.com/pb33f/libopenapi/index"

	"github.com/pb33f/libopenapi/datamodel"
	v2high "github.com/pb33f/libopenapi/datamodel/high/v2"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/what-changed/model"
)

// Document Represents an OpenAPI specification that can then be rendered into a model or serialized back into
// a string document after being manipulated.
type Document interface {
	// GetVersion will return the exact version of the OpenAPI specification set for the document.
	GetVersion() string

	// GetRolodex will return the Rolodex instance that was used to load the document.
	GetRolodex() *index.Rolodex

	// GetSpecInfo will return the *datamodel.SpecInfo instance that contains all specification information.
	GetSpecInfo() *datamodel.SpecInfo

	// SetConfiguration will set the configuration for the document. This allows for finer grained control over
	// allowing remote or local references, as well as a BaseURL to allow for relative file references.
	SetConfiguration(configuration *datamodel.DocumentConfiguration)

	// GetConfiguration will return the configuration for the document. This allows for finer grained control over
	// allowing remote or local references, as well as a BaseURL to allow for relative file references.
	GetConfiguration() *datamodel.DocumentConfiguration

	// BuildV2Model will build out a Swagger (version 2) model from the specification used to create the document
	// If there are any issues, then no model will be returned, instead a slice of errors will explain all the
	// problems that occurred. This method will only support version 2 specifications and will throw an error for
	// any other types.
	BuildV2Model() (*DocumentModel[v2high.Swagger], error)

	// BuildV3Model will build out an OpenAPI (version 3+) model from the specification used to create the document
	// If there are any issues, then no model will be returned, instead a slice of errors will explain all the
	// problems that occurred. This method will only support version 3 specifications and will throw an error for
	// any other types.
	BuildV3Model() (*DocumentModel[v3high.Document], error)

	// RenderAndReload will render the high level model as it currently exists (including any mutations, additions
	// and removals to and from any object in the tree). It will then reload the low level model with the new bytes
	// extracted from the model that was re-rendered. This is useful if you want to make changes to the high level model
	// and then 'reload' the model into memory, so that line numbers and column numbers are correct and all update
	// according to the changes made.
	//
	// The method returns the raw YAML bytes that were rendered, and any errors that occurred during rebuilding of the model.
	// This is a destructive operation, and will re-build the entire model from scratch using the new bytes, so any
	// references to the old model will be lost. The second return is the new Document that was created, and the third
	// return is any errors hit trying to re-render.
	//
	// **IMPORTANT** This method only supports OpenAPI Documents. The Swagger model will not support mutations correctly
	// and will not update when called. This choice has been made because we don't want to continue supporting Swagger,
	// it's too old, so it should be motivation to upgrade to OpenAPI 3.
	RenderAndReload() ([]byte, Document, *DocumentModel[v3high.Document], error)

	// Render will render the high level model as it currently exists (including any mutations, additions
	// and removals to and from any object in the tree). Unlike RenderAndReload, Render will simply print the state
	// of the model as it currently exists, and will not re-load the model into memory. It means that the low-level and
	// the high-level models will be out of sync, and the index will only be useful for the original document.
	//
	// Why use this instead of RenderAndReload?
	//
	// The simple answer is that RenderAndReload is a destructive operation, and will re-build the entire model from
	// scratch using the new bytes, which is desirable if you want to make changes to the high level model and then
	// 'reload' the model into memory, so that line numbers and column numbers are correct and the index is accurate.
	// However, if you don't care about the low-level model, and you're not using the index, and you just want to
	// print the state of the model as it currently exists, then Render() is the method to use.
	// **IMPORTANT** This method only supports OpenAPI Documents.
	Render() ([]byte, error)

	// Serialize will re-render a Document back into a []byte slice. If any modifications have been made to the
	// underlying data model using low level APIs, then those changes will be reflected in the serialized output.
	//
	// It's important to know that this should not be used if the resolver has been used on a specification to
	// for anything other than checking for circular references. If the resolver is used to resolve the spec, then this
	// method may spin out forever if the specification backing the model has circular references.
	// Deprecated: This method is deprecated and will be removed in a future release. Use RenderAndReload() instead.
	// This method does not support mutations correctly.
	Serialize() ([]byte, error)

	// Release nils all internal state so that the YAML tree, SpecIndex, Rolodex,
	// and model objects can be garbage-collected even if something still holds
	// a reference to the Document interface value.
	Release()
}

type document struct {
	rolodex           *index.Rolodex
	version           string
	info              *datamodel.SpecInfo
	config            *datamodel.DocumentConfiguration
	highOpenAPI3Model *DocumentModel[v3high.Document]
	highSwaggerModel  *DocumentModel[v2high.Swagger]
}

// DocumentModel represents either a Swagger document (version 2) or an OpenAPI document (version 3) that is
// built from a parent Document.
type DocumentModel[T v2high.Swagger | v3high.Document] struct {
	Model T
	Index *index.SpecIndex // index created from the document.
}

// NewDocument will create a new OpenAPI instance from an OpenAPI specification []byte array. If anything goes
// wrong when parsing, reading or processing the OpenAPI specification, there will be no document returned, instead
// a slice of errors will be returned that explain everything that failed.
//
// After creating a Document, the option to build a model becomes available, in either V2 or V3 flavors. The models
// are about 70% different between Swagger and OpenAPI 3, which is why two different models are available.
//
// This function will NOT automatically follow (meaning load) any file or remote references that are found.
//
// If this isn't the behavior you want, then you can use the NewDocumentWithConfiguration() function instead, which allows you to set a configuration that
// will allow you to control if file or remote references are allowed. In particular the `AllowFileReferences` and `AllowRemoteReferences`
// properties.
func NewDocument(specByteArray []byte) (Document, error) {
	_ = "STUB: not implemented"
	return *new(Document), nil
}

func NewDocumentWithTypeCheck(specByteArray []byte, bypassCheck bool) (Document, error) {
	_ = "STUB: not implemented"
	return *new(Document), nil
}

// NewDocumentWithConfiguration is the same as NewDocument, except it's a convenience function that calls NewDocument
// under the hood and then calls SetConfiguration() on the returned Document.
func NewDocumentWithConfiguration(specByteArray []byte, configuration *datamodel.DocumentConfiguration) (Document, error) {
	_ = "STUB: not implemented"
	return *new(Document), nil
}

func (d *document) Release() { _ = "STUB: not implemented"; return }

// This method intentionally does not call SpecIndex.Release(). Low-level
// model objects (Schema, PathItem, etc.) retain their own references to the
// SpecIndex and require its config and root node for hashing and comparison
// operations that may run after a Document is released. Callers that own the
// full lifecycle should call SpecIndex.Release() separately once all model
// consumers are finished.

func (d *document) GetRolodex() *index.Rolodex { _ = "STUB: not implemented"; return nil }

func (d *document) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (d *document) GetSpecInfo() *datamodel.SpecInfo { _ = "STUB: not implemented"; return nil }

func (d *document) GetConfiguration() *datamodel.DocumentConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (d *document) SetConfiguration(configuration *datamodel.DocumentConfiguration) {
	_ = "STUB: not implemented"
	return
}

func (d *document) Serialize() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *document) RenderAndReload() ([]byte, Document, *DocumentModel[v3high.Document], error) {
	_ = "STUB: not implemented"
	return nil, *new(Document), nil, nil
}

// build the model.

// this document is now dead, long live the new document!

func (d *document) Render() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// check for Swagger model first, to give a more helpful error message.

func (d *document) BuildV2Model() (*DocumentModel[v2high.Swagger], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do not short-circuit on circular reference errors, so the client
// has the option of ignoring them.

func (d *document) BuildV3Model() (*DocumentModel[v3high.Document], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do not short-circuit on circular reference errors, so the client
// has the option of ignoring them.

// CompareDocuments will accept a left and right Document implementing struct, build a model for the correct
// version and then compare model documents for changes.
//
// If there are any errors when building the models, those errors are returned with a nil pointer for the
// model.DocumentChanges. If there are any changes found however between either Document, then a pointer to
// model.DocumentChanges is returned containing every single change, broken down, model by model.
func CompareDocuments(original, updated Document) (*model.DocumentChanges, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
