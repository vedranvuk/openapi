// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Discriminator Object
// When request bodies or response payloads may be one of a number of different 
// schemas, a discriminator object can be used to aid in serialization,
// deserialization, and validation. The discriminator is a specific object in a
// schema which is used to inform the consumer of the specification of an
// alternative schema based on the value associated with it.
//
// When using the discriminator, inline schemas will not be considered.
type Discriminator struct {
	// REQUIRED. The name of the property in the payload that will hold the
	// discriminator value.
	PropertyName string `json:"propertyName,omitempty"`
	// An object to hold mappings between payload values and schema names or
	// references.
	Mapping map[string]string `json:"mapping,omitempty"`
}