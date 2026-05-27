// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package reports

import (
	"github.com/pb33f/libopenapi/what-changed/model"
)

// Changed provides a simple wrapper for changed counts
type Changed struct {
	Total    int `json:"totalChanges"`
	Breaking int `json:"breakingChanges"`
}

// OverallReport provides a Document level overview of all changes to an OpenAPI doc.
type OverallReport struct {
	ChangeReport map[string]*Changed `json:"overallSummaryReport"`
}

// CreateOverallReport will create a high level report for all top level changes (but with deep counts)
func CreateOverallReport(changes *model.DocumentChanges) *OverallReport {
	_ = "STUB: not implemented"
	return nil
}

func mergeRootPropertyChanges(changedReport map[string]*Changed, propertyChanges *model.PropertyChanges) {
	_ = "STUB: not implemented"
	return
}

func mergeChangedModel(changedReport map[string]*Changed, label string, next *Changed) {
	_ = "STUB: not implemented"
	return
}

func getOrCreateChanged(changedReport map[string]*Changed, label string) *Changed {
	_ = "STUB: not implemented"
	return nil
}

func createChangedModel(ch HasChanges) *Changed { _ = "STUB: not implemented"; return nil }

func createChangedModelFromSlice(ch []HasChanges) *Changed { _ = "STUB: not implemented"; return nil }
