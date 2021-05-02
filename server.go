// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Server Object
// An object representing a Server.
type Server struct {
	// REQUIRED. A URL to the target host. This URL supports Server Variables
	// and MAY be relative, to indicate that the host location is relative to
	// the location where the OpenAPI document is being served. Variable
	// substitutions will be made when a variable is named in {brackets}.
	URL string `json:"url,omitempty"`
	// An optional string describing the host designated by the URL. CommonMark
	// syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// A map between a variable name and its value. The value is used for
	// substitution in the server's URL template.
	Variables map[string]*ServerVariable `json:"variables,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// Server Object Example
// A single server would be described as:

// {
//   "url": "https://development.gigantic-server.com/v1",
//   "description": "Development server"
// }
// url: https://development.gigantic-server.com/v1
// description: Development server
// The following shows how multiple servers can be described, for example, at the OpenAPI Object's servers:

// {
//   "servers": [
//     {
//       "url": "https://development.gigantic-server.com/v1",
//       "description": "Development server"
//     },
//     {
//       "url": "https://staging.gigantic-server.com/v1",
//       "description": "Staging server"
//     },
//     {
//       "url": "https://api.gigantic-server.com/v1",
//       "description": "Production server"
//     }
//   ]
// }
// servers:
// - url: https://development.gigantic-server.com/v1
//   description: Development server
// - url: https://staging.gigantic-server.com/v1
//   description: Staging server
// - url: https://api.gigantic-server.com/v1
//   description: Production server
// The following shows how variables can be used for a server configuration:

// {
//   "servers": [
//     {
//       "url": "https://{username}.gigantic-server.com:{port}/{basePath}",
//       "description": "The production API server",
//       "variables": {
//         "username": {
//           "default": "demo",
//           "description": "this value is assigned by the service provider, in this example `gigantic-server.com`"
//         },
//         "port": {
//           "enum": [
//             "8443",
//             "443"
//           ],
//           "default": "8443"
//         },
//         "basePath": {
//           "default": "v2"
//         }
//       }
//     }
//   ]
// }
// servers:
// - url: https://{username}.gigantic-server.com:{port}/{basePath}
//   description: The production API server
//   variables:
//     username:
//       # note! no enum here means it is an open value
//       default: demo
//       description: this value is assigned by the service provider, in this example `gigantic-server.com`
//     port:
//       enum:
//         - '8443'
//         - '443'
//       default: '8443'
//     basePath:
//       # open meaning there is the opportunity to use special base paths as assigned by the provider, default is `v2`
//       default: v2

