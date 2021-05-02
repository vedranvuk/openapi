// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Schema Object
// The Schema Object allows the definition of input and output data types.
// These types can be objects, but also primitives and arrays. This object is
// an extended subset of the JSON Schema Specification Wright Draft 00.
//
// For more information about the properties, see JSON Schema Core and JSON
// Schema Validation. Unless stated otherwise, the property definitions follow
// the JSON Schema.
//
// Properties
// The following properties are taken directly from the JSON Schema definition
// and follow the same specifications:
//
// title
// multipleOf
// maximum
// exclusiveMaximum
// minimum
// exclusiveMinimum
// maxLength
// minLength
// pattern (This string SHOULD be a valid regular expression, according to the Ecma-262 Edition 5.1 regular expression dialect)
// maxItems
// minItems
// uniqueItems
// maxProperties
// minProperties
// required
// enum
//
// The following properties are taken from the JSON Schema definition but their
// definitions were adjusted to the OpenAPI Specification.
//
// type - Value MUST be a string. Multiple types via an array are not supported.
// allOf - Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.
// oneOf - Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.
// anyOf - Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.
// not - Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.
// items - Value MUST be an object and not an array. Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema. items MUST be present if the type is array.
// properties - Property definitions MUST be a Schema Object and not a standard JSON Schema (inline or referenced).
// additionalProperties - Value can be boolean or object. Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema. Consistent with JSON Schema, additionalProperties defaults to true.
// description - CommonMark syntax MAY be used for rich text representation.
// format - See Data Type Formats for further details. While relying on JSON Schema's defined formats, the OAS offers a few additional predefined formats.
// default - The default value represents what would be assumed by the consumer of the input as the value of the schema if one is not provided. Unlike JSON Schema, the value MUST conform to the defined type for the Schema Object defined at the same level. For example, if type is string, then default can be "foo" but cannot be 1.
// Alternatively, any time a Schema Object can be used, a Reference Object can be used in its place. This allows referencing definitions instead of defining them inline.
//
// Additional properties defined by the JSON Schema specification that are not
// mentioned here are strictly unsupported.
//
// Other than the JSON Schema subset fields, the following fields MAY be used
// for further schema documentation:
type Schema struct {
	// A true value adds "null" to the allowed type specified by the type
	// keyword, only if type is explicitly defined within the same Schema
	// Object. Other Schema Object constraints retain their defined behavior,
	// and therefore may disallow the use of null as a value. A false value
	// leaves the specified or default type unmodified. The default value is
	// false.
	Nullable bool `json:"nullable,omitempty"`
	// Adds support for polymorphism. The discriminator is an object name that
	// is used to differentiate between other schemas which may satisfy the
	// payload description. See Composition and Inheritance for more details.
	Discriminator *Discriminator `json:"discriminator,omitempty"`
	// Relevant only for Schema "properties" definitions. Declares the property
	// as "read only". This means that it MAY be sent as part of a response but
	// SHOULD NOT be sent as part of the request. If the property is marked as
	// readOnly being true and is in the required list, the required will take
	// effect on the response only. A property MUST NOT be marked as both
	// readOnly and writeOnly being true. Default value is false.
	ReadOnly bool `json:"readOnly,omitempty"`
	// Relevant only for Schema "properties" definitions. Declares the property
	// as "write only". Therefore, it MAY be sent as part of a request but
	// SHOULD NOT be sent as part of the response. If the property is marked as
	// writeOnly being true and is in the required list, the required will take
	// effect on the request only. A property MUST NOT be marked as both
	// readOnly and writeOnly being true. Default value is false.
	WriteOnly bool `json:"writeOnly,omitempty"`
	// This MAY be used only on properties schemas. It has no effect on root
	// schemas. Adds additional metadata to describe the XML representation of
	// this property.
	XML *XML `json:"xml,omitempty"`
	// Additional external documentation for this schema.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
	// A free-form property to include an example of an instance for this schema.
	// To represent examples that cannot be naturally represented in JSON or
	// YAML, a string value can be used to contain the example with escaping
	// where necessary.
	Example interface{} `json:"example,omitempty"`
	// Specifies that a schema is deprecated and SHOULD be transitioned out of
	// usage. Default value is false.
	Deprecated bool `json:"deprecated,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// Composition and Inheritance (Polymorphism)
// The OpenAPI Specification allows combining and extending model definitions using the allOf property of JSON Schema, in effect offering model composition. allOf takes an array of object definitions that are validated independently but together compose a single object.

// While composition offers model extensibility, it does not imply a hierarchy between the models. To support polymorphism, the OpenAPI Specification adds the discriminator field. When used, the discriminator will be the name of the property that decides which schema definition validates the structure of the model. As such, the discriminator field MUST be a required field. There are two ways to define the value of a discriminator for an inheriting instance.

// Use the schema name.
// Override the schema name by overriding the property with a new value. If a new value exists, this takes precedence over the schema name. As such, inline schema definitions, which do not have a given id, cannot be used in polymorphism.
// XML Modeling
// The xml property allows extra definitions when translating the JSON definition to XML. The XML Object contains additional information about the available options.

// Schema Object Examples
// Primitive Sample
// {
//   "type": "string",
//   "format": "email"
// }
// type: string
// format: email
// Simple Model
// {
//   "type": "object",
//   "required": [
//     "name"
//   ],
//   "properties": {
//     "name": {
//       "type": "string"
//     },
//     "address": {
//       "$ref": "#/components/schemas/Address"
//     },
//     "age": {
//       "type": "integer",
//       "format": "int32",
//       "minimum": 0
//     }
//   }
// }
// type: object
// required:
// - name
// properties:
//   name:
//     type: string
//   address:
//     $ref: '#/components/schemas/Address'
//   age:
//     type: integer
//     format: int32
//     minimum: 0
// Model with Map/Dictionary Properties
// For a simple string to string mapping:

// {
//   "type": "object",
//   "additionalProperties": {
//     "type": "string"
//   }
// }
// type: object
// additionalProperties:
//   type: string
// For a string to model mapping:

// {
//   "type": "object",
//   "additionalProperties": {
//     "$ref": "#/components/schemas/ComplexModel"
//   }
// }
// type: object
// additionalProperties:
//   $ref: '#/components/schemas/ComplexModel'
// Model with Example
// {
//   "type": "object",
//   "properties": {
//     "id": {
//       "type": "integer",
//       "format": "int64"
//     },
//     "name": {
//       "type": "string"
//     }
//   },
//   "required": [
//     "name"
//   ],
//   "example": {
//     "name": "Puma",
//     "id": 1
//   }
// }
// type: object
// properties:
//   id:
//     type: integer
//     format: int64
//   name:
//     type: string
// required:
// - name
// example:
//   name: Puma
//   id: 1
// Models with Composition
// {
//   "components": {
//     "schemas": {
//       "ErrorModel": {
//         "type": "object",
//         "required": [
//           "message",
//           "code"
//         ],
//         "properties": {
//           "message": {
//             "type": "string"
//           },
//           "code": {
//             "type": "integer",
//             "minimum": 100,
//             "maximum": 600
//           }
//         }
//       },
//       "ExtendedErrorModel": {
//         "allOf": [
//           {
//             "$ref": "#/components/schemas/ErrorModel"
//           },
//           {
//             "type": "object",
//             "required": [
//               "rootCause"
//             ],
//             "properties": {
//               "rootCause": {
//                 "type": "string"
//               }
//             }
//           }
//         ]
//       }
//     }
//   }
// }
// components:
//   schemas:
//     ErrorModel:
//       type: object
//       required:
//       - message
//       - code
//       properties:
//         message:
//           type: string
//         code:
//           type: integer
//           minimum: 100
//           maximum: 600
//     ExtendedErrorModel:
//       allOf:
//       - $ref: '#/components/schemas/ErrorModel'
//       - type: object
//         required:
//         - rootCause
//         properties:
//           rootCause:
//             type: string
// Models with Polymorphism Support
// {
//   "components": {
//     "schemas": {
//       "Pet": {
//         "type": "object",
//         "discriminator": {
//           "propertyName": "petType"
//         },
//         "properties": {
//           "name": {
//             "type": "string"
//           },
//           "petType": {
//             "type": "string"
//           }
//         },
//         "required": [
//           "name",
//           "petType"
//         ]
//       },
//       "Cat": {
//         "description": "A representation of a cat. Note that `Cat` will be used as the discriminator value.",
//         "allOf": [
//           {
//             "$ref": "#/components/schemas/Pet"
//           },
//           {
//             "type": "object",
//             "properties": {
//               "huntingSkill": {
//                 "type": "string",
//                 "description": "The measured skill for hunting",
//                 "default": "lazy",
//                 "enum": [
//                   "clueless",
//                   "lazy",
//                   "adventurous",
//                   "aggressive"
//                 ]
//               }
//             },
//             "required": [
//               "huntingSkill"
//             ]
//           }
//         ]
//       },
//       "Dog": {
//         "description": "A representation of a dog. Note that `Dog` will be used as the discriminator value.",
//         "allOf": [
//           {
//             "$ref": "#/components/schemas/Pet"
//           },
//           {
//             "type": "object",
//             "properties": {
//               "packSize": {
//                 "type": "integer",
//                 "format": "int32",
//                 "description": "the size of the pack the dog is from",
//                 "default": 0,
//                 "minimum": 0
//               }
//             },
//             "required": [
//               "packSize"
//             ]
//           }
//         ]
//       }
//     }
//   }
// }
// components:
//   schemas:
//     Pet:
//       type: object
//       discriminator:
//         propertyName: petType
//       properties:
//         name:
//           type: string
//         petType:
//           type: string
//       required:
//       - name
//       - petType
//     Cat:  ## "Cat" will be used as the discriminator value
//       description: A representation of a cat
//       allOf:
//       - $ref: '#/components/schemas/Pet'
//       - type: object
//         properties:
//           huntingSkill:
//             type: string
//             description: The measured skill for hunting
//             enum:
//             - clueless
//             - lazy
//             - adventurous
//             - aggressive
//         required:
//         - huntingSkill
//     Dog:  ## "Dog" will be used as the discriminator value
//       description: A representation of a dog
//       allOf:
//       - $ref: '#/components/schemas/Pet'
//       - type: object
//         properties:
//           packSize:
//             type: integer
//             format: int32
//             description: the size of the pack the dog is from
//             default: 0
//             minimum: 0
//         required:
//         - packSize
