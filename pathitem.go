// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Path Item Object
// Describes the operations available on a single path. A Path Item MAY be
// empty, due to ACL constraints. The path itself is still exposed to the
// documentation viewer but they will not know which operations and parameters
// are available.
type PathItem struct {
	// Allows for an external definition of this path item. The referenced 
	// structure MUST be in the format of a Path Item Object. In case a Path
	// Item Object field appears both in the defined object and the referenced
	// object, the behavior is undefined.
	Ref string `json:"$ref,omitempty"`
	// An optional, string summary, intended to apply to all operations in this 
	// path.
	Summary string `json:"summary,omitempty"`
	// An optional, string description, intended to apply to all operations in 
	// this path. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// A definition of a GET operation on this path.
	Get *Operation `json:"get,omitempty"`
	// A definition of a PUT operation on this path.
	Put *Operation `json:"put,omitempty"`
	// A definition of a POST operation on this path.
	Post *Operation `json:"post,omitempty"`
	// A definition of a DELETE operation on this path.
	Delete *Operation `json:"delete,omitempty"`
	// A definition of a OPTIONS operation on this path.
	Options *Operation `json:"options,omitempty"`
	// A definition of a HEAD operation on this path.
	Head *Operation `json:"head,omitempty"`
	// A definition of a PATCH operation on this path.
	Patch *Operation `json:"patch,omitempty"`
	// A definition of a TRACE operation on this path.
	Trace *Operation `json:"trace,omitempty"`
	// An alternative server array to service all operations in this path.
	Servers []*Server `json:"servers,omitempty"`
	// A list of parameters that are applicable for all the operations described
	// under this path. These parameters can be overridden at the operation 
	// level, but cannot be removed there. The list MUST NOT include duplicated 
	// parameters. A unique parameter is defined by a combination of a name and
	// location. The list can use the Reference Object to link to parameters 
	// that are defined at the OpenAPI Object's components/parameters.
	Parameters []interface{} `json:"parameters,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// Path Item Object Example
// {
//   "get": {
//     "description": "Returns pets based on ID",
//     "summary": "Find pets by ID",
//     "operationId": "getPetsById",
//     "responses": {
//       "200": {
//         "description": "pet response",
//         "content": {
//           "*/*": {
//             "schema": {
//               "type": "array",
//               "items": {
//                 "$ref": "#/components/schemas/Pet"
//               }
//             }
//           }
//         }
//       },
//       "default": {
//         "description": "error payload",
//         "content": {
//           "text/html": {
//             "schema": {
//               "$ref": "#/components/schemas/ErrorModel"
//             }
//           }
//         }
//       }
//     }
//   },
//   "parameters": [
//     {
//       "name": "id",
//       "in": "path",
//       "description": "ID of pet to use",
//       "required": true,
//       "schema": {
//         "type": "array",
//         "items": {
//           "type": "string"
//         }
//       },
//       "style": "simple"
//     }
//   ]
// }
// get:
//   description: Returns pets based on ID
//   summary: Find pets by ID
//   operationId: getPetsById
//   responses:
//     '200':
//       description: pet response
//       content:
//         '*/*' :
//           schema:
//             type: array
//             items:
//               $ref: '#/components/schemas/Pet'
//     default:
//       description: error payload
//       content:
//         'text/html':
//           schema:
//             $ref: '#/components/schemas/ErrorModel'
// parameters:
// - name: id
//   in: path
//   description: ID of pet to use
//   required: true
//   schema:
//     type: array
//     items:
//       type: string
//   style: simple
