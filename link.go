// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Link Object
// The Link object represents a possible design-time link for a response. The
// presence of a link does not guarantee the caller's ability to successfully 
// invoke it, rather it provides a known relationship and traversal mechanism 
// between responses and other operations.
// 
// Unlike dynamic links (i.e. links provided in the response payload), the OAS
// linking mechanism does not require link information in the runtime response.
// 
// For computing links, and providing instructions to execute them, a runtime 
// expression is used for accessing values in an operation and using them as
// parameters while invoking the linked operation.
type Link struct {
	// A relative or absolute URI reference to an OAS operation. This field is
	// mutually exclusive of the operationId field, and MUST point to an
	// Operation Object. Relative operationRef values MAY be used to locate an
	// existing Operation Object in the OpenAPI definition.
	OperationRef string `json:"operationRef,omitempty"`
	// The name of an existing, resolvable OAS operation, as defined with a
	// unique operationId. This field is mutually exclusive of the operationRef
	// field.
	OperationId string `json:"operationId,omitempty"`
	// Holds Map[string, Any | {expression}]
	// A map representing parameters to pass to an operation as specified with
	// operationId or identified via operationRef. The key is the parameter name
	// to be used, whereas the value can be a constant or an expression to be
	// evaluated and passed to the linked operation. The parameter name can be
	// qualified using the parameter location [{in}.]{name} for operations that
	// use the same parameter name in different locations (e.g. path.id).
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// Any | {expression}
	// A literal value or {expression} to use as a request body when calling the
	// target operation.
	RequestBody interface{} `json:"requestBody,omitempty"`
	// A description of the link. CommonMark syntax MAY be used for rich text
	// representation.
	Description string `json:"description,omitempty"`
	// A server object to be used by the target operation.
	Server *Server `json:"server,omitempty"`
	// This object MAY be extended with Specification Extensions.
}

// A linked operation MUST be identified using either an operationRef or operationId. In the case of an operationId, it MUST be unique and resolved in the scope of the OAS document. Because of the potential for name clashes, the operationRef syntax is preferred for specifications with external references.

// Examples
// Computing a link from a request operation where the $request.path.id is used to pass a request parameter to the linked operation.

// paths:
//   /users/{id}:
//     parameters:
//     - name: id
//       in: path
//       required: true
//       description: the user identifier, as userId 
//       schema:
//         type: string
//     get:
//       responses:
//         '200':
//           description: the user being returned
//           content:
//             application/json:
//               schema:
//                 type: object
//                 properties:
//                   uuid: # the unique user id
//                     type: string
//                     format: uuid
//           links:
//             address:
//               # the target link operationId
//               operationId: getUserAddress
//               parameters:
//                 # get the `id` field from the request path parameter named `id`
//                 userId: $request.path.id
//   # the path item of the linked operation
//   /users/{userid}/address:
//     parameters:
//     - name: userid
//       in: path
//       required: true
//       description: the user identifier, as userId 
//       schema:
//         type: string
//     # linked operation
//     get:
//       operationId: getUserAddress
//       responses:
//         '200':
//           description: the user's address
// When a runtime expression fails to evaluate, no parameter value is passed to the target operation.

// Values from the response body can be used to drive a linked operation.

// links:
//   address:
//     operationId: getUserAddressByUUID
//     parameters:
//       # get the `uuid` field from the `uuid` field in the response body
//       userUuid: $response.body#/uuid
// Clients follow all links at their discretion. Neither permissions, nor the capability to make a successful call to that link, is guaranteed solely by the existence of a relationship.

// OperationRef Examples
// As references to operationId MAY NOT be possible (the operationId is an optional field in an Operation Object), references MAY also be made through a relative operationRef:

// links:
//   UserRepositories:
//     # returns array of '#/components/schemas/repository'
//     operationRef: '#/paths/~12.0~1repositories~1{username}/get'
//     parameters:
//       username: $response.body#/username
// or an absolute operationRef:

// links:
//   UserRepositories:
//     # returns array of '#/components/schemas/repository'
//     operationRef: 'https://na2.gigantic-server.com/#/paths/~12.0~1repositories~1{username}/get'
//     parameters:
//       username: $response.body#/username
// Note that in the use of operationRef, the escaped forward-slash is necessary when using JSON references.

// Runtime Expressions
// Runtime expressions allow defining values based on information that will only be available within the HTTP message in an actual API call. This mechanism is used by Link Objects and Callback Objects.

// The runtime expression is defined by the following ABNF syntax

//       expression = ( "$url" / "$method" / "$statusCode" / "$request." source / "$response." source )
//       source = ( header-reference / query-reference / path-reference / body-reference )
//       header-reference = "header." token
//       query-reference = "query." name  
//       path-reference = "path." name
//       body-reference = "body" ["#" json-pointer ]
//       json-pointer    = *( "/" reference-token )
//       reference-token = *( unescaped / escaped )
//       unescaped       = %x00-2E / %x30-7D / %x7F-10FFFF
//          ; %x2F ('/') and %x7E ('~') are excluded from 'unescaped'
//       escaped         = "~" ( "0" / "1" )
//         ; representing '~' and '/', respectively
//       name = *( CHAR )
//       token = 1*tchar
//       tchar = "!" / "#" / "$" / "%" / "&" / "'" / "*" / "+" / "-" / "." /
//         "^" / "_" / "`" / "|" / "~" / DIGIT / ALPHA
// Here, json-pointer is taken from RFC 6901, char from RFC 7159 and token from RFC 7230.

// The name identifier is case-sensitive, whereas token is not.

// The table below provides examples of runtime expressions and examples of their use in a value:

// Examples
// Source Location	example expression	notes
// HTTP Method	$method	The allowable values for the $method will be those for the HTTP operation.
// Requested media type	$request.header.accept	
// Request parameter	$request.path.id	Request parameters MUST be declared in the parameters section of the parent operation or they cannot be evaluated. This includes request headers.
// Request body property	$request.body#/user/uuid	In operations which accept payloads, references may be made to portions of the requestBody or the entire body.
// Request URL	$url	
// Response value	$response.body#/status	In operations which return payloads, references may be made to portions of the response body or the entire body.
// Response header	$response.header.Server	Single header values only are available
// Runtime expressions preserve the type of the referenced value. Expressions can be embedded into string values by surrounding the expression with {} curly braces.
