// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// License Object
// License information for the exposed API.
type License struct {
	// REQUIRED. The license name used for the API.
	Name string `json:"name,omitempty"`
	// A URL to the license used for the API. MUST be in the format of a URL.
	URL  string `json:"url,omitempty"`
	// This object MAY be extended with Specification Extensions.
}