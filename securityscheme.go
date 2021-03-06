// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Security Scheme Object
// Defines a security scheme that can be used by the operations. Supported
// schemes are HTTP authentication, an API key (either as a header, a cookie
// parameter or as a query parameter), OAuth2's common flows
// (implicit, password, client credentials and authorization code) as defined
// in RFC6749, and OpenID Connect Discovery.
type SecurityScheme struct {
	// Applies to Any.
	// REQUIRED. The type of the security scheme. Valid values are "apiKey",
	// "http", "oauth2", "openIdConnect".
	Type string `json:"type,omitempty"`
	// Applies to Any.
	// A short description for security scheme. CommonMark syntax MAY be used
	// for rich text representation.
	Description string `json:"description,omitempty"`
	// Applies to: apiKey.
	// REQUIRED. The name of the header, query or cookie parameter to be used.
	Name string `json:"name,omitempty"`
	// Applies to: apiKey.
	// REQUIRED. The location of the API key. Valid values are "query", "header"
	// or "cookie".
	In string `json:"in,omitempty"`
	// Applies to: http
	// REQUIRED. The name of the HTTP Authorization scheme to be used in the
	// Authorization header as defined in RFC7235. The values used SHOULD be
	// registered in the IANA Authentication Scheme registry.
	Scheme string `json:"scheme,omitempty"`
	// Applies to: http ("bearer")
	// A hint to the client to identify how the bearer token is formatted.
	// Bearer tokens are usually generated by an authorization server, so this
	// information is primarily for documentation purposes.
	BearerFormat string `json:"bearerFormat,omitempty"`
	// Applies to: oauth2
	// REQUIRED. An object containing configuration information for the flow
	// types supported.
	Flows *OAuthFlows `json:"flows,omitempty"`
	// Applies to: openIdConnect
	// REQUIRED. OpenId Connect URL to discover OAuth2 configuration values.
	// This MUST be in the form of a URL.
	OpenIDConnectURL string `json:"openIdConnectUrl,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// Security Scheme Object Example
// Basic Authentication Sample
// {
//   "type": "http",
//   "scheme": "basic"
// }
// type: http
// scheme: basic
// API Key Sample
// {
//   "type": "apiKey",
//   "name": "api_key",
//   "in": "header"
// }
// type: apiKey
// name: api_key
// in: header
// JWT Bearer Sample
// {
//   "type": "http",
//   "scheme": "bearer",
//   "bearerFormat": "JWT",
// }
// type: http
// scheme: bearer
// bearerFormat: JWT
// Implicit OAuth2 Sample
// {
//   "type": "oauth2",
//   "flows": {
//     "implicit": {
//       "authorizationUrl": "https://example.com/api/oauth/dialog",
//       "scopes": {
//         "write:pets": "modify pets in your account",
//         "read:pets": "read your pets"
//       }
//     }
//   }
// }
// type: oauth2
// flows:
//   implicit:
//     authorizationUrl: https://example.com/api/oauth/dialog
//     scopes:
//       write:pets: modify pets in your account
//       read:pets: read your pets
