// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Reference Object
// A simple object to allow referencing other components in the specification,
//  internally and externally.
//
// The Reference Object is defined by JSON Reference and follows the same
// structure, behavior and rules.
//
// For this specification, reference resolution is accomplished as defined by 
// the JSON Reference specification and not by the JSON Schema specification.
type Reference struct {
	// REQUIRED. The reference string.
	Ref string `json:"$ref,omitempty"`
	// This object cannot be extended with additional properties and any 
	// properties added SHALL be ignored.
}
