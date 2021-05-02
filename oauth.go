// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// OAuth Flow Object
// Configuration details for a supported OAuth Flow
type OAuthFlow struct {
	// Applies to: oauth2 ("implicit", "authorizationCode")
	// REQUIRED. The authorization URL to be used for this flow. This MUST be
	// in the form of a URL.
	AuthorizationURL string `json:"authorizationUrl,omitempty"`
	// Applies to: oauth2 ("password", "clientCredentials", "authorizationCode")
	// REQUIRED. The token URL to be used for this flow. This MUST be in the
	// form of a URL.
	TokenURL string `json:"tokenUrl,omitempty"`
	// Applies to: oauth2
	// The URL to be used for obtaining refresh tokens. This MUST be in the form
	// of a URL.
	RefreshURL string `json:"refreshUrl,omitempty"`
	// Applies to: oauth2
	// REQUIRED. The available scopes for the OAuth2 security scheme. A map
	// between the scope name and a short description for it. The map MAY be
	// empty.
	Scopes map[string]string `json:"scopes,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// OAuth Flows Object
// Allows configuration of the supported OAuth Flows.
type OAuthFlows struct {
	// Configuration for the OAuth Implicit flow.
	Implicit *OAuthFlow `json:"implicit,omitempty"`
	// Configuration for the OAuth Resource Owner Password flow.
	Password *OAuthFlow `json:"password,omitempty"`
	// Configuration for the OAuth Client Credentials flow. Previously called
	// application in OpenAPI 2.0.
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty"`
	// Configuration for the OAuth Authorization Code flow. Previously called
	// accessCode in OpenAPI 2.0.
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// OAuth Flow Object Examples
// {
//   "type": "oauth2",
//   "flows": {
//     "implicit": {
//       "authorizationUrl": "https://example.com/api/oauth/dialog",
//       "scopes": {
//         "write:pets": "modify pets in your account",
//         "read:pets": "read your pets"
//       }
//     },
//     "authorizationCode": {
//       "authorizationUrl": "https://example.com/api/oauth/dialog",
//       "tokenUrl": "https://example.com/api/oauth/token",
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
//   authorizationCode:
//     authorizationUrl: https://example.com/api/oauth/dialog
//     tokenUrl: https://example.com/api/oauth/token
//     scopes:
//       write:pets: modify pets in your account
//       read:pets: read your pets
