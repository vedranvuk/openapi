// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Components Object
// Holds a set of reusable objects for different aspects of the OAS. All objects
// defined within the components object will have no effect on the API unless
// they are explicitly referenced from properties outside the components object.
type Components struct {
	// Holds Map[string, Schema Object | Reference Object]
	//
	// An object to hold reusable Schema Objects.
	Schemas map[string]interface{} `json:"schemas,omitempty"`
	// Holds Map[string, Response Object | Reference Object]
	//
	// An object to hold reusable Response Objects.
	Responses map[string]interface{} `json:"responses,omitempty"`
	// Holds Map[string, Parameter Object | Reference Object]
	//
	// An object to hold reusable Parameter Objects.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// Holds Map[string, Example Object | Reference Object]
	//
	// An object to hold reusable Example Objects.
	Examples map[string]interface{} `json:"examples,omitempty"`
	// Holds Map[string, RequestBody Object | Reference Object]
	//
	// An object to hold reusable Request Body Objects.
	RequestBodies map[string]interface{} `json:"requestBodies,omitempty"`
	// Holds Map[string, Header Object | Reference Object]
	//
	// An object to hold reusable Header Objects.
	Headers map[string]interface{} `json:"headers,omitempty"`
	// Holds Map[string, SecurityScheme Object | Reference Object]
	//
	// An object to hold reusable Security Scheme Objects.
	SecuritySchemes map[string]interface{} `json:"securitySchemes,omitempty"`
	// Holds Map[string, Link Object | Reference Object]
	//
	// An object to hold reusable Link Objects.
	Links map[string]interface{} `json:"links,omitempty"`
	// Holds Map[string, Callback Object | Reference Object]
	//
	// An object to hold reusable Callback Objects.
	Callbacks map[string]interface{} `json:"callbacks,omitempty"`
	// Holds Map[string, Path Item Object | Reference Object]
	//
	// An object to hold reusable Path Item Object.
	PathItems map[string]interface{}

	// This object MAY be extended with Specification Extensions.
}

// All the fixed fields declared above are objects that MUST use keys that match the regular expression: ^[a-zA-Z0-9\.\-_]+$.

// Field Name Examples:

// User
// User_1
// User_Name
// user-name
// my.org.User
// Components Object Example
// "components": {
//   "schemas": {
//     "GeneralError": {
//       "type": "object",
//       "properties": {
//         "code": {
//           "type": "integer",
//           "format": "int32"
//         },
//         "message": {
//           "type": "string"
//         }
//       }
//     },
//     "Category": {
//       "type": "object",
//       "properties": {
//         "id": {
//           "type": "integer",
//           "format": "int64"
//         },
//         "name": {
//           "type": "string"
//         }
//       }
//     },
//     "Tag": {
//       "type": "object",
//       "properties": {
//         "id": {
//           "type": "integer",
//           "format": "int64"
//         },
//         "name": {
//           "type": "string"
//         }
//       }
//     }
//   },
//   "parameters": {
//     "skipParam": {
//       "name": "skip",
//       "in": "query",
//       "description": "number of items to skip",
//       "required": true,
//       "schema": {
//         "type": "integer",
//         "format": "int32"
//       }
//     },
//     "limitParam": {
//       "name": "limit",
//       "in": "query",
//       "description": "max records to return",
//       "required": true,
//       "schema" : {
//         "type": "integer",
//         "format": "int32"
//       }
//     }
//   },
//   "responses": {
//     "NotFound": {
//       "description": "Entity not found."
//     },
//     "IllegalInput": {
//       "description": "Illegal input for operation."
//     },
//     "GeneralError": {
//       "description": "General Error",
//       "content": {
//         "application/json": {
//           "schema": {
//             "$ref": "#/components/schemas/GeneralError"
//           }
//         }
//       }
//     }
//   },
//   "securitySchemes": {
//     "api_key": {
//       "type": "apiKey",
//       "name": "api_key",
//       "in": "header"
//     },
//     "petstore_auth": {
//       "type": "oauth2",
//       "flows": {
//         "implicit": {
//           "authorizationUrl": "http://example.org/api/oauth/dialog",
//           "scopes": {
//             "write:pets": "modify pets in your account",
//             "read:pets": "read your pets"
//           }
//         }
//       }
//     }
//   }
// }
// components:
//   schemas:
//     GeneralError:
//       type: object
//       properties:
//         code:
//           type: integer
//           format: int32
//         message:
//           type: string
//     Category:
//       type: object
//       properties:
//         id:
//           type: integer
//           format: int64
//         name:
//           type: string
//     Tag:
//       type: object
//       properties:
//         id:
//           type: integer
//           format: int64
//         name:
//           type: string
//   parameters:
//     skipParam:
//       name: skip
//       in: query
//       description: number of items to skip
//       required: true
//       schema:
//         type: integer
//         format: int32
//     limitParam:
//       name: limit
//       in: query
//       description: max records to return
//       required: true
//       schema:
//         type: integer
//         format: int32
//   responses:
//     NotFound:
//       description: Entity not found.
//     IllegalInput:
//       description: Illegal input for operation.
//     GeneralError:
//       description: General Error
//       content:
//         application/json:
//           schema:
//             $ref: '#/components/schemas/GeneralError'
//   securitySchemes:
//     api_key:
//       type: apiKey
//       name: api_key
//       in: header
//     petstore_auth:
//       type: oauth2
//       flows:
//         implicit:
//           authorizationUrl: http://example.org/api/oauth/dialog
//           scopes:
//             write:pets: modify pets in your account
//             read:pets: read your pets
