// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// OpenAPI object.
// This is the root document object of the OpenAPI document.
//
// Specification Extensions
// While the OpenAPI Specification tries to accommodate most use cases, 
// additional data can be added to extend the specification at certain points.
//
// The extensions properties are implemented as patterned fields that are always 
// prefixed by "x-".
type OpenAPI struct {
	// REQUIRED. This string MUST be the semantic version number of the OpenAPI
	// Specification version that the OpenAPI document uses. The openapi field 
	// SHOULD be used by tooling specifications and clients to interpret the
	// OpenAPI document. This is not related to the API info.version string.
	OpenAPI      string        `json:"openapi,omitempty"`
	// REQUIRED. Provides metadata about the API. The metadata MAY be used by
	// tooling as required.
	Info         *Info         `json:"info,omitempty"`
	// An array of Server Objects, which provide connectivity information to a 
	// target server. If the servers property is not provided, or is an empty
	// array, the default value would be a Server Object with a url value of /.
	Servers      []*Server      `json:"servers,omitempty"`
	// REQUIRED. The available paths and operations for the API.
	Paths        *Paths        `json:"paths,omitempty"`
	// An element to hold various schemas for the specification.
	Components   *Components   `json:"components,omitempty"`
	// A declaration of which security mechanisms can be used across the API. 
	// The list of values includes alternative security requirement objects that
	// can be used. Only one of the security requirement objects need to be
	// satisfied to authorize a request. Individual operations can override this
	// definition. To make security optional, an empty security requirement ({})
	// can be included in the array.
	Security     []*SecurityRequirement     `json:"security,omitempty"`
	// A list of tags used by the specification with additional metadata. The 
	// order of the tags can be used to reflect on their order by the parsing 
	// tools. Not all tags that are used by the Operation Object must be
	// declared. The tags that are not declared MAY be organized randomly or 
	// based on the tools' logic. Each tag name in the list MUST be unique.
	Tags         []*Tag         `json:"tags,omitempty"`
	// Additional external documentation.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
	// This object MAY be extended with Specification Extensions.
}