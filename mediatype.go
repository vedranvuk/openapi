// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Media Type Object
// Each Media Type Object provides schema and examples for the media type 
// identified by its key.
type MediaType struct {
	// Contains Schema or Reference object.
	// The schema defining the content of the request, response, or parameter.
	Schema interface{} `json:"schema,omitempty"`
	// Contains Any type.
	// Example of the media type. The example object SHOULD be in the correct
	// format as specified by the media type. The example field is mutually
	// exclusive of the examples field. Furthermore, if referencing a schema
	// which contains an example, the example value SHALL override the example
	// provided by the schema.
	Example interface{} `json:"example,omitempty"`
	// Holds Map[ string, Example Object | Reference Object]
	// Examples of the media type. Each example object SHOULD match the media
	// type and specified schema if present. The examples field is mutually
	// exclusive of the example field. Furthermore, if referencing a schema
	// which contains an example, the examples value SHALL override the example
	// provided by the schema.
	Examples map[string]interface{} `json:"examples,omitempty"`
	// Holds Map[string, Encoding Object]
	// A map between a property name and its encoding information. The key,
	// being the property name, MUST exist in the schema as a property. The
	// encoding object SHALL only apply to requestBody objects when the media
	// type is multipart or application/x-www-form-urlencoded.
	Encoding map[string]*Encoding `json:"encoding,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// Media Type Examples
// {
//   "application/json": {
//     "schema": {
//          "$ref": "#/components/schemas/Pet"
//     },
//     "examples": {
//       "cat" : {
//         "summary": "An example of a cat",
//         "value": 
//           {
//             "name": "Fluffy",
//             "petType": "Cat",
//             "color": "White",
//             "gender": "male",
//             "breed": "Persian"
//           }
//       },
//       "dog": {
//         "summary": "An example of a dog with a cat's name",
//         "value" :  { 
//           "name": "Puma",
//           "petType": "Dog",
//           "color": "Black",
//           "gender": "Female",
//           "breed": "Mixed"
//         },
//       "frog": {
//           "$ref": "#/components/examples/frog-example"
//         }
//       }
//     }
//   }
// }
// application/json: 
//   schema:
//     $ref: "#/components/schemas/Pet"
//   examples:
//     cat:
//       summary: An example of a cat
//       value:
//         name: Fluffy
//         petType: Cat
//         color: White
//         gender: male
//         breed: Persian
//     dog:
//       summary: An example of a dog with a cat's name
//       value:
//         name: Puma
//         petType: Dog
//         color: Black
//         gender: Female
//         breed: Mixed
//     frog:
//       $ref: "#/components/examples/frog-example"
// Considerations for File Uploads
// In contrast with the 2.0 specification, file input/output content in OpenAPI is described with the same semantics as any other schema type. Specifically:

// # content transferred with base64 encoding
// schema:
//   type: string
//   format: base64
// # content transferred in binary (octet-stream):
// schema:
//   type: string
//   format: binary
// These examples apply to either input payloads of file uploads or response payloads.

// A requestBody for submitting a file in a POST operation may look like the following example:

// requestBody:
//   content:
//     application/octet-stream:
//       schema:
//         # a binary file of any type
//         type: string
//         format: binary
// In addition, specific media types MAY be specified:

// # multiple, specific media types may be specified:
// requestBody:
//   content:
//       # a binary file of type png or jpeg
//     'image/jpeg':
//       schema:
//         type: string
//         format: binary
//     'image/png':
//       schema:
//         type: string
//         format: binary        
// To upload multiple files, a multipart media type MUST be used:

// requestBody:
//   content:
//     multipart/form-data:
//       schema:
//         properties:
//           # The property name 'file' will be used for all files.
//           file:
//             type: array
//             items:
//               type: string
//               format: binary

// Support for x-www-form-urlencoded Request Bodies
// To submit content using form url encoding via RFC1866, the following definition may be used:

// requestBody:
//   content:
//     application/x-www-form-urlencoded:
//       schema:
//         type: object
//         properties:
//           id:
//             type: string
//             format: uuid
//           address:
//             # complex types are stringified to support RFC 1866
//             type: object
//             properties: {}
// In this example, the contents in the requestBody MUST be stringified per RFC1866 when passed to the server. In addition, the address field complex object will be stringified.

// When passing complex objects in the application/x-www-form-urlencoded content type, the default serialization strategy of such properties is described in the Encoding Object's style property as form.

// Special Considerations for multipart Content
// It is common to use multipart/form-data as a Content-Type when transferring request bodies to operations. In contrast to 2.0, a schema is REQUIRED to define the input parameters to the operation when using multipart content. This supports complex structures as well as supporting mechanisms for multiple file uploads.

// When passing in multipart types, boundaries MAY be used to separate sections of the content being transferred — thus, the following default Content-Types are defined for multipart:

// If the property is a primitive, or an array of primitive values, the default Content-Type is text/plain
// If the property is complex, or an array of complex values, the default Content-Type is application/json
// If the property is a type: string with format: binary or format: base64 (aka a file object), the default Content-Type is application/octet-stream
// Examples:

// requestBody:
//   content:
//     multipart/form-data:
//       schema:
//         type: object
//         properties:
//           id:
//             type: string
//             format: uuid
//           address:
//             # default Content-Type for objects is `application/json`
//             type: object
//             properties: {}
//           profileImage:
//             # default Content-Type for string/binary is `application/octet-stream`
//             type: string
//             format: binary
//           children:
//             # default Content-Type for arrays is based on the `inner` type (text/plain here)
//             type: array
//             items:
//               type: string
//           addresses:
//             # default Content-Type for arrays is based on the `inner` type (object shown, so `application/json` in this example)
//             type: array
//             items:
//               type: '#/components/schemas/Address'
// An encoding attribute is introduced to give you control over the serialization of parts of multipart request bodies. This attribute is only applicable to multipart and application/x-www-form-urlencoded request bodies.
