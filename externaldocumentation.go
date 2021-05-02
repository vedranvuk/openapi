// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// External Documentation Object
// Allows referencing an external resource for extended documentation.
type ExternalDocumentation struct {
	// A short description of the target documentation. CommonMark syntax MAY be
	// used for rich text representation.
	Description string `json:"description,omitempty"`
	// REQUIRED. The URL for the target documentation. Value MUST be in the
	// format of a URL.
	URL string `json:"url,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// External Documentation Object Example
// {
//   "description": "Find more info here",
//   "url": "https://example.com"
// }
// description: Find more info here
// url: https://example.com
