// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Responses Object
// A container for the expected responses of an operation. The container maps 
// a HTTP response code to the expected response.
//
// The documentation is not necessarily expected to cover all possible HTTP 
// response codes because they may not be known in advance. However, 
// documentation is expected to cover a successful operation response and any 
// known errors.
//
// The default MAY be used as a default response object for all HTTP codes that 
// are not covered individually by the specification.
//
// The Responses Object MUST contain at least one response code, and it SHOULD
// be the response for a successful operation call.
type Responses struct {
	// Contains Response Object | Reference Object
	//
	// The documentation of responses other than the ones declared for specific 
	// HTTP response codes. Use this field to cover undeclared responses. A 
	// Reference Object can link to a response that the OpenAPI Object's 
	// components/responses section defines.
	Default interface{}

	// Contains Response Object | Reference Object
	//
	// Any HTTP status code can be used as the property name, but only one 
	// property per code, to describe the expected response for that HTTP 
	// status code. A Reference Object can link to a response that is defined 
	// in the OpenAPI Object's components/responses section. This field MUST be
	// enclosed in quotation marks (for example, "200") for compatibility 
	// between JSON and YAML. To define a range of response codes, this field 
	// MAY contain the uppercase wildcard character X. For example, 2XX 
	// represents all response codes between [200-299]. Only the following range
	// definitions are allowed: 1XX, 2XX, 3XX, 4XX, and 5XX. If a response is
	// defined using an explicit code, the explicit code definition takes
	// precedence over the range definition for that code.
	// HTTP_STATUS_CODE interface{}
	
	// This object MAY be extended with Specification Extensions.
}

// Responses Object Example
// A 200 response for a successful operation and a default response for others (implying an error):

// {
//   "200": {
//     "description": "a pet to be returned",
//     "content": {
//       "application/json": {
//         "schema": {
//           "$ref": "#/components/schemas/Pet"
//         }
//       }
//     }
//   },
//   "default": {
//     "description": "Unexpected error",
//     "content": {
//       "application/json": {
//         "schema": {
//           "$ref": "#/components/schemas/ErrorModel"
//         }
//       }
//     }
//   }
// }
// '200':
//   description: a pet to be returned
//   content:
//     application/json:
//       schema:
//         $ref: '#/components/schemas/Pet'
// default:
//   description: Unexpected error
//   content:
//     application/json:
//       schema:
//         $ref: '#/components/schemas/ErrorModel'
