// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Contact Object
// Contact information for the exposed API.
type Contact struct {
	// The identifying name of the contact person/organization.
	Name string `json:"name,omitempty"`
	// The URL pointing to the contact information. MUST be in the format of a URL.
	URL string `json:"url,omitempty"`
	// The email address of the contact person/organization. MUST be in the
	// format of an email address.
	EMail string `json:"email,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// Contact Object Example
// {
//   "name": "API Support",
//   "url": "http://www.example.com/support",
//   "email": "support@example.com"
// }
// name: API Support
// url: http://www.example.com/support
// email: support@example.com
