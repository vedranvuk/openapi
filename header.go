// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Header Object
// The Header Object follows the structure of the Parameter Object with the following changes:
//
// name MUST NOT be specified, it is given in the corresponding headers map.
// in MUST NOT be specified, it is implicitly in header.
// All traits that are affected by the location MUST be applicable to a location of header (for example, style).
// Header Object Example
// A simple header of type integer:
//
// {
//   "description": "The number of allowed requests in the current period",
//   "schema": {
//     "type": "integer"
//   }
// }
// description: The number of allowed requests in the current period
// schema:
//   type: integer
type Header struct {

}

// Header Object Example
// A simple header of type integer:

// {
//   "description": "The number of allowed requests in the current period",
//   "schema": {
//     "type": "integer"
//   }
// }
// description: The number of allowed requests in the current period
// schema:
//   type: integer