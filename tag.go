// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Tag Object
// Adds metadata to a single tag that is used by the Operation Object. It is 
// not mandatory to have a Tag Object per tag defined in the Operation Object 
// instances.
type Tag struct {
	// REQUIRED. The name of the tag.
	Name string `json:"name,omitempty"`
	// A short description for the tag. CommonMark syntax MAY be used for rich
	// text representation.
	Description string `json:"description,omitempty"`
	// Additional external documentation for this tag.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// Tag Object Example
// {
// 	"name": "pet",
// 	"description": "Pets operations"
// }
// name: pet
// description: Pets operations
