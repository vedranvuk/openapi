// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Example object
type Example struct {
	// Short description for the example.
	Summary string `json:"summary,omitempty"`
	// Long description for the example. CommonMark syntax MAY be used for rich
	// text representation.
	Description string `json:"description,omitempty"`
	// Embedded literal example. The value field and externalValue field are
	// mutually exclusive. To represent examples of media types that cannot
	// naturally represented in JSON or YAML, use a string value to contain the
	// example, escaping where necessary.
	Value interface{} `json:"value,omitempty"`
	// A URL that points to the literal example. This provides the capability
	// to reference examples that cannot easily be included in JSON or YAML
	// documents. The value field and externalValue field are mutually exclusive.
	ExternalValue string `json:"externalValue,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// In all cases, the example value is expected to be compatible with the type schema of its associated value. Tooling implementations MAY choose to validate compatibility automatically, and reject the example value(s) if incompatible.

// Example Object Examples
// In a request body:

// requestBody:
//   content:
//     'application/json':
//       schema:
//         $ref: '#/components/schemas/Address'
//       examples: 
//         foo:
//           summary: A foo example
//           value: {"foo": "bar"}
//         bar:
//           summary: A bar example
//           value: {"bar": "baz"}
//     'application/xml':
//       examples: 
//         xmlExample:
//           summary: This is an example in XML
//           externalValue: 'http://example.org/examples/address-example.xml'
//     'text/plain':
//       examples:
//         textExample: 
//           summary: This is a text example
//           externalValue: 'http://foo.bar/examples/address-example.txt'
// In a parameter:

// parameters:
//   - name: 'zipCode'
//     in: 'query'
//     schema:
//       type: 'string'
//       format: 'zip-code'
//     examples:
//       zip-example: 
//         $ref: '#/components/examples/zip-example'
// In a response:

// responses:
//   '200':
//     description: your car appointment has been booked
//     content: 
//       application/json:
//         schema:
//           $ref: '#/components/schemas/SuccessResponse'
//         examples:
//           confirmation-success:
//             $ref: '#/components/examples/confirmation-success'