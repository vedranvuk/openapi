// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// XML Object
// A metadata object that allows for more fine-tuned XML model definitions.
// 
// When using arrays, XML element names are not inferred
// (for singular/plural forms) and the name property SHOULD be used to add that
// information. See examples for expected behavior.
type XML struct {
	// Replaces the name of the element/attribute used for the described schema
	// property. When defined within items, it will affect the name of the
	// individual XML elements within the list. When defined alongside type
	// being array (outside the items), it will affect the wrapping element and
	// only if wrapped is true. If wrapped is false, it will be ignored.
	Name string `json:"name,omitempty"`
	// The URI of the namespace definition. Value MUST be in the form of an
	// absolute URI.
	Namespace string `json:"namespace,omitempty"`
	// The prefix to be used for the name.
	Prefix string `json:"prefix,omitempty"`
	// Declares whether the property definition translates to an attribute
	// instead of an element. Default value is false.
	Attribute string `json:"attribute,omitempty"`
	// MAY be used only for an array definition. Signifies whether the array is
	// wrapped (for example, <books><book/><book/></books>) or unwrapped
	// (<book/><book/>). Default value is false. The definition takes effect
	// only when defined alongside type being array (outside the items).
	Wrapped string `json:"wrapped,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// XML Object Examples
// The examples of the XML object definitions are included inside a property definition of a Schema Object with a sample of the XML representation of it.

// No XML Element
// Basic string property:

// {
//     "animals": {
//         "type": "string"
//     }
// }
// animals:
//   type: string
// <animals>...</animals>
// Basic string array property (wrapped is false by default):

// {
//     "animals": {
//         "type": "array",
//         "items": {
//             "type": "string"
//         }
//     }
// }
// animals:
//   type: array
//   items:
//     type: string
// <animals>...</animals>
// <animals>...</animals>
// <animals>...</animals>
// XML Name Replacement
// {
//   "animals": {
//     "type": "string",
//     "xml": {
//       "name": "animal"
//     }
//   }
// }
// animals:
//   type: string
//   xml:
//     name: animal
// <animal>...</animal>
// XML Attribute, Prefix and Namespace
// In this example, a full model definition is shown.

// {
//   "Person": {
//     "type": "object",
//     "properties": {
//       "id": {
//         "type": "integer",
//         "format": "int32",
//         "xml": {
//           "attribute": true
//         }
//       },
//       "name": {
//         "type": "string",
//         "xml": {
//           "namespace": "http://example.com/schema/sample",
//           "prefix": "sample"
//         }
//       }
//     }
//   }
// }
// Person:
//   type: object
//   properties:
//     id:
//       type: integer
//       format: int32
//       xml:
//         attribute: true
//     name:
//       type: string
//       xml:
//         namespace: http://example.com/schema/sample
//         prefix: sample
// <Person id="123">
//     <sample:name xmlns:sample="http://example.com/schema/sample">example</sample:name>
// </Person>
// XML Arrays
// Changing the element names:

// {
//   "animals": {
//     "type": "array",
//     "items": {
//       "type": "string",
//       "xml": {
//         "name": "animal"
//       }
//     }
//   }
// }
// animals:
//   type: array
//   items:
//     type: string
//     xml:
//       name: animal
// <animal>value</animal>
// <animal>value</animal>
// The external name property has no effect on the XML:

// {
//   "animals": {
//     "type": "array",
//     "items": {
//       "type": "string",
//       "xml": {
//         "name": "animal"
//       }
//     },
//     "xml": {
//       "name": "aliens"
//     }
//   }
// }
// animals:
//   type: array
//   items:
//     type: string
//     xml:
//       name: animal
//   xml:
//     name: aliens
// <animal>value</animal>
// <animal>value</animal>
// Even when the array is wrapped, if a name is not explicitly defined, the same name will be used both internally and externally:

// {
//   "animals": {
//     "type": "array",
//     "items": {
//       "type": "string"
//     },
//     "xml": {
//       "wrapped": true
//     }
//   }
// }
// animals:
//   type: array
//   items:
//     type: string
//   xml:
//     wrapped: true
// <animals>
//   <animals>value</animals>
//   <animals>value</animals>
// </animals>
// To overcome the naming problem in the example above, the following definition can be used:

// {
//   "animals": {
//     "type": "array",
//     "items": {
//       "type": "string",
//       "xml": {
//         "name": "animal"
//       }
//     },
//     "xml": {
//       "wrapped": true
//     }
//   }
// }
// animals:
//   type: array
//   items:
//     type: string
//     xml:
//       name: animal
//   xml:
//     wrapped: true
// <animals>
//   <animal>value</animal>
//   <animal>value</animal>
// </animals>
// Affecting both internal and external names:

// {
//   "animals": {
//     "type": "array",
//     "items": {
//       "type": "string",
//       "xml": {
//         "name": "animal"
//       }
//     },
//     "xml": {
//       "name": "aliens",
//       "wrapped": true
//     }
//   }
// }
// animals:
//   type: array
//   items:
//     type: string
//     xml:
//       name: animal
//   xml:
//     name: aliens
//     wrapped: true
// <aliens>
//   <animal>value</animal>
//   <animal>value</animal>
// </aliens>
// If we change the external element but not the internal ones:

// {
//   "animals": {
//     "type": "array",
//     "items": {
//       "type": "string"
//     },
//     "xml": {
//       "name": "aliens",
//       "wrapped": true
//     }
//   }
// }
// animals:
//   type: array
//   items:
//     type: string
//   xml:
//     name: aliens
//     wrapped: true
// <aliens>
//   <aliens>value</aliens>
//   <aliens>value</aliens>
// </aliens>
