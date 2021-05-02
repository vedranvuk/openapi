// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Paths Object
// Holds the relative paths to the individual endpoints and their operations.
// The path is appended to the URL from the Server Object in order to construct
// the full URL. The Paths MAY be empty, due to ACL constraints.
type Paths struct {
	// A relative path to an individual endpoint. The field name MUST begin with
	// a forward slash (/). The path is appended (no relative URL resolution) 
	// to the expanded URL from the Server Object's url field in order to 
	// construct the full URL. Path templating is allowed. When matching URLs,
	// concrete (non-templated) paths would be matched before their templated 
	// counterparts. Templated paths with the same hierarchy but different 
	// templated names MUST NOT exist as they are identical. In case of
	// ambiguous matching, it's up to the tooling to decide which one to use.
	// /{path} *PathItem

	// This object MAY be extended with Specification Extensions.
}

// Path Templating Matching
// Assuming the following paths, the concrete definition, /pets/mine, will be matched first if used:

//   /pets/{petId}
//   /pets/mine
// The following paths are considered identical and invalid:

//   /pets/{petId}
//   /pets/{name}
// The following may lead to ambiguous resolution:

//   /{entity}/me
//   /books/{id}
// Paths Object Example
// {
//   "/pets": {
//     "get": {
//       "description": "Returns all pets from the system that the user has access to",
//       "responses": {
//         "200": {          
//           "description": "A list of pets.",
//           "content": {
//             "application/json": {
//               "schema": {
//                 "type": "array",
//                 "items": {
//                   "$ref": "#/components/schemas/pet"
//                 }
//               }
//             }
//           }
//         }
//       }
//     }
//   }
// }
// /pets:
//   get:
//     description: Returns all pets from the system that the user has access to
//     responses:
//       '200':
//         description: A list of pets.
//         content:
//           application/json:
//             schema:
//               type: array
//               items:
//                 $ref: '#/components/schemas/pet'