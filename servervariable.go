// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi


// Server Variable Object
// An object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	// An enumeration of string values to be used if the substitution options
	// are from a limited set. The array SHOULD NOT be empty.
	Enum []string `json:"enum,omitempty"`
	// REQUIRED. The default value to use for substitution, which SHALL be sent
	// if an alternate value is not supplied. Note this behavior is different
	// than the Schema Object's treatment of default values, because in those
	// cases parameter values are optional. If the enum is defined, the value
	// SHOULD exist in the enum's values.
	Default string `json:"default,omitempty"`
	// An optional description for the server variable. CommonMark syntax MAY
	// be used for rich text representation.
	Description string `json:"description,omitempty"`
	// This object MAY be extended with Specification Extensions.
}
