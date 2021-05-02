// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Request Body Object
// Describes a single request body.
type RequestBody struct {
	// A brief description of the request body. This could contain examples of
	// use. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// REQUIRED. The content of the request body. The key is a media type or 
	// media type range and the value describes it. For requests that match 
	// multiple keys, only the most specific key is applicable. e.g.
	// text/plain overrides text/*
	Content map[string]*MediaType `json:"content,omitempty"`
	// Determines if the request body is required in the request. Defaults to
	// false.
	Required bool `json:"required,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// Request Body Examples
// A request body with a referenced model definition.

// {
//   "description": "user to add to the system",
//   "content": {
//     "application/json": {
//       "schema": {
//         "$ref": "#/components/schemas/User"
//       },
//       "examples": {
//           "user" : {
//             "summary": "User Example", 
//             "externalValue": "http://foo.bar/examples/user-example.json"
//           } 
//         }
//     },
//     "application/xml": {
//       "schema": {
//         "$ref": "#/components/schemas/User"
//       },
//       "examples": {
//           "user" : {
//             "summary": "User example in XML",
//             "externalValue": "http://foo.bar/examples/user-example.xml"
//           }
//         }
//     },
//     "text/plain": {
//       "examples": {
//         "user" : {
//             "summary": "User example in Plain text",
//             "externalValue": "http://foo.bar/examples/user-example.txt" 
//         }
//       } 
//     },
//     "*/*": {
//       "examples": {
//         "user" : {
//             "summary": "User example in other format",
//             "externalValue": "http://foo.bar/examples/user-example.whatever"
//         }
//       }
//     }
//   }
// }
// description: user to add to the system
// content: 
//   'application/json':
//     schema:
//       $ref: '#/components/schemas/User'
//     examples:
//       user:
//         summary: User Example
//         externalValue: 'http://foo.bar/examples/user-example.json'
//   'application/xml':
//     schema:
//       $ref: '#/components/schemas/User'
//     examples:
//       user:
//         summary: User Example in XML
//         externalValue: 'http://foo.bar/examples/user-example.xml'
//   'text/plain':
//     examples:
//       user:
//         summary: User example in text plain format
//         externalValue: 'http://foo.bar/examples/user-example.txt'
//   '*/*':
//     examples:
//       user: 
//         summary: User example in other format
//         externalValue: 'http://foo.bar/examples/user-example.whatever'
// A body parameter that is an array of string values:

// {
//   "description": "user to add to the system",
//   "content": {
//     "text/plain": {
//       "schema": {
//         "type": "array",
//         "items": {
//           "type": "string"
//         }
//       }
//     }
//   }
// }
// description: user to add to the system
// required: true
// content:
//   text/plain:
//     schema:
//       type: array
//       items:
//         type: string