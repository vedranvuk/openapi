// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Response Object
// Describes a single response from an API Operation, including design-time,
// static links to operations based on the response.
type Response struct {
	// REQUIRED. A short description of the response. CommonMark syntax MAY be
	// used for rich text representation.
	Description string `json:"description,omitempty"`
	// Holds Map[string, Header Object | Reference Object]
	//
	// Maps a header name to its definition. RFC7230 states header names are
	// case insensitive. If a response header is defined with the name
	// "Content-Type", it SHALL be ignored.
	Headers map[string]interface{} `json:"headers,omitempty"`
	// A map containing descriptions of potential response payloads. The key is
	// a media type or media type range and the value describes it. For
	// responses that match multiple keys, only the most specific key is
	// applicable. e.g. text/plain overrides text/*
	Content map[string]*MediaType `json:"content,omitempty"`
	// Holds Map[string, Link Object | Reference Object]
	//
	// A map of operations links that can be followed from the response. The
	// key of the map is a short name for the link, following the naming
	// constraints of the names for Component Objects.
	Links map[string]interface{} `json:"links,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// Response Object Examples
// Response of an array of a complex type:

// {
//   "description": "A complex object array response",
//   "content": {
//     "application/json": {
//       "schema": {
//         "type": "array",
//         "items": {
//           "$ref": "#/components/schemas/VeryComplexType"
//         }
//       }
//     }
//   }
// }
// description: A complex object array response
// content: 
//   application/json:
//     schema: 
//       type: array
//       items:
//         $ref: '#/components/schemas/VeryComplexType'
// Response with a string type:

// {
//   "description": "A simple string response",
//   "content": {
//     "text/plain": {
//       "schema": {
//         "type": "string"
//       }
//     }
//   }

// }
// description: A simple string response
// content:
//   text/plain:
//     schema:
//       type: string
// Plain text response with headers:

// {
//   "description": "A simple string response",
//   "content": {
//     "text/plain": {
//       "schema": {
//         "type": "string",
//         "example": "whoa!"
//       }
//     }
//   },
//   "headers": {
//     "X-Rate-Limit-Limit": {
//       "description": "The number of allowed requests in the current period",
//       "schema": {
//         "type": "integer"
//       }
//     },
//     "X-Rate-Limit-Remaining": {
//       "description": "The number of remaining requests in the current period",
//       "schema": {
//         "type": "integer"
//       }
//     },
//     "X-Rate-Limit-Reset": {
//       "description": "The number of seconds left in the current period",
//       "schema": {
//         "type": "integer"
//       }
//     }
//   }
// }
// description: A simple string response
// content:
//   text/plain:
//     schema:
//       type: string
//     example: 'whoa!'
// headers:
//   X-Rate-Limit-Limit:
//     description: The number of allowed requests in the current period
//     schema:
//       type: integer
//   X-Rate-Limit-Remaining:
//     description: The number of remaining requests in the current period
//     schema:
//       type: integer
//   X-Rate-Limit-Reset:
//     description: The number of seconds left in the current period
//     schema:
//       type: integer
// Response with no return value:

// {
//   "description": "object created"
// }
// description: object created
