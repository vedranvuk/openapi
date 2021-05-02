// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Encoding Object
// A single encoding definition applied to a single schema property.
type Encoding struct {
	// The Content-Type for encoding a specific property. Default value depends
	// on the property type: for string with format being
	// binary – application/octet-stream;
	// for other primitive types – text/plain;
	// for object - application/json;
	// for array – the default is defined based on the inner type. The value can
	// be a specific media type (e.g. application/json), a wildcard media type
	// (e.g. image/*), or a comma-separated list of the two types.
	ContentType string `json:"contentType,omitempty"`
	// Holds Map[string, Header Object | Reference Object]
	// A map allowing additional information to be provided as headers, for
	// example Content-Disposition. Content-Type is described separately and
	// SHALL be ignored in this section. This property SHALL be ignored if the
	// request body media type is not a multipart.
	Headers map[string]interface{} `json:"headers,omitempty"`
	// Describes how a specific property value will be serialized depending on
	// its type. See Parameter Object for details on the style property. The
	// behavior follows the same values as query parameters, including default
	// values. This property SHALL be ignored if the request body media type is
	// not application/x-www-form-urlencoded.
	Style string `json:"style,omitempty"`
	// When this is true, property values of type array or object generate
	// separate parameters for each value of the array, or key-value-pair of the
	// map. For other types of properties this property has no effect. When
	// style is form, the default value is true. For all other styles, the
	// default value is false. This property SHALL be ignored if the request
	// body media type is not application/x-www-form-urlencoded.
	Explode bool `json:"explode,omitempty"`
	// Determines whether the parameter value SHOULD allow reserved characters,
	// as defined by RFC3986 :/?#[]@!$&'()*+,;= to be included without
	// percent-encoding. The default value is false. This property SHALL be
	// ignored if the request body media type is not
	// application/x-www-form-urlencoded.
	AllowReserved bool `json:"allowReserved,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// Encoding Object Example
// requestBody:
//   content:
//     multipart/mixed:
//       schema:
//         type: object
//         properties:
//           id:
//             # default is text/plain
//             type: string
//             format: uuid
//           address:
//             # default is application/json
//             type: object
//             properties: {}
//           historyMetadata:
//             # need to declare XML format!
//             description: metadata in XML format
//             type: object
//             properties: {}
//           profileImage:
//             # default is application/octet-stream, need to declare an image type only!
//             type: string
//             format: binary
//       encoding:
//         historyMetadata:
//           # require XML Content-Type in utf-8 encoding
//           contentType: application/xml; charset=utf-8
//         profileImage:
//           # only accept png/jpeg
//           contentType: image/png, image/jpeg
//           headers:
//             X-Rate-Limit-Limit:
//               description: The number of allowed requests in the current period
//               schema:
//                 type: integer
