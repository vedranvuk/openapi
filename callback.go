// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// Callback Object
// A map of possible out-of band callbacks related to the parent operation.
// Each value in the map is a Path Item Object that describes a set of requests
// that may be initiated by the API provider and the expected responses. The key
// value used to identify the path item object is an expression, evaluated at 
// runtime, that identifies a URL to use for the callback operation.
type Callback struct {
	// A Path Item Object used to define a callback request and expected 
	// responses. A complete example is available.
	// {expression} *PathItem

	// This object MAY be extended with Specification Extensions.
}

// Key Expression
// The key that identifies the Path Item Object is a runtime expression that can be evaluated in the context of a runtime HTTP request/response to identify the URL to be used for the callback request. A simple example might be $request.body#/url. However, using a runtime expression the complete HTTP message can be accessed. This includes accessing any part of a body that a JSON Pointer RFC6901 can reference.

// For example, given the following HTTP request:

// POST /subscribe/myevent?queryUrl=http://clientdomain.com/stillrunning HTTP/1.1
// Host: example.org
// Content-Type: application/json
// Content-Length: 187

// {
//   "failedUrl" : "http://clientdomain.com/failed",
//   "successUrls" : [
//     "http://clientdomain.com/fast",
//     "http://clientdomain.com/medium",
//     "http://clientdomain.com/slow"
//   ] 
// }

// 201 Created
// Location: http://example.org/subscription/1
// The following examples show how the various expressions evaluate, assuming the callback operation has a path parameter named eventType and a query parameter named queryUrl.

// Expression	Value
// $url	http://example.org/subscribe/myevent?queryUrl=http://clientdomain.com/stillrunning
// $method	POST
// $request.path.eventType	myevent
// $request.query.queryUrl	http://clientdomain.com/stillrunning
// $request.header.content-Type	application/json
// $request.body#/failedUrl	http://clientdomain.com/failed
// $request.body#/successUrls/2	http://clientdomain.com/medium
// $response.header.Location	http://example.org/subscription/1
// Callback Object Examples
// The following example uses the user provided queryUrl query string parameter to define the callback URL. This is an example of how to use a callback object to describe a WebHook callback that goes with the subscription operation to enable registering for the WebHook.

// myCallback:
//   '{$request.query.queryUrl}':
//     post:
//       requestBody:
//         description: Callback payload
//         content: 
//           'application/json':
//             schema:
//               $ref: '#/components/schemas/SomePayload'
//       responses:
//         '200':
//           description: callback successfully processed
// The following example shows a callback where the server is hard-coded, but the query string parameters are populated from the id and email property in the request body.

// transactionCallback:
//   'http://notificationServer.com?transactionId={$request.body#/id}&email={$request.body#/email}':
//     post:
//       requestBody:
//         description: Callback payload
//         content: 
//           'application/json':
//             schema:
//               $ref: '#/components/schemas/SomePayload'
//       responses:
//         '200':
//           description: callback successfully processed
