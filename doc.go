// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package openapi contains object model types for OpenAPI specification v3.1.0
//
// OpenAPI Specification
//
// Version 3.1.0
//
// The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD",
// "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this
// document are to be interpreted as described in BCP 14 RFC2119 RFC8174 when,
// and only when, they appear in all capitals, as shown here.
//
// This document is licensed under The Apache License, Version 2.0.
//
// Introduction
//
// The OpenAPI Specification (OAS) defines a standard, language-agnostic 
// interface to HTTP APIs which allows both humans and computers to discover 
// and understand the capabilities of the service without access to source code,
// documentation, or through network traffic inspection. When properly defined,
// a consumer can understand and interact with the remote service with a minimal
// amount of implementation logic.
//
// An OpenAPI definition can then be used by documentation generation tools to
// display the API, code generation tools to generate servers and clients in
// various programming languages, testing tools, and many other use cases.
//
// Definitions
//
// OpenAPI Document
//
// A self-contained or composite resource which defines or describes an API or
// elements of an API. The OpenAPI document MUST contain at least one paths 
// field, a components field or a webhooks field. An OpenAPI document uses and
// conforms to the OpenAPI Specification.
//
// Path Templating
//
// Path templating refers to the usage of template expressions, delimited by 
// curly braces ({}), to mark a section of a URL path as replaceable using path
// parameters.
//
// Each template expression in the path MUST correspond to a path parameter 
// that is included in the Path Item itself and/or in each of the Path Item's
// Operations. An exception is if the path item is empty, for example due to ACL
// constraints, matching path parameters are not required.
//
// The value for these path parameters MUST NOT contain any unescaped 
// "generic syntax" characters described by RFC3986: forward slashes (/), 
// question marks (?), or hashes (#).
//
// Media Types
//
// Media type definitions are spread across several resources. The media type
// definitions SHOULD be in compliance with RFC6838.
//
// Some examples of possible media type definitions:
//
//   text/plain; charset=utf-8
//   application/json
//   application/vnd.github+json
//   application/vnd.github.v3+json
//   application/vnd.github.v3.raw+json
//   application/vnd.github.v3.text+json
//   application/vnd.github.v3.html+json
//   application/vnd.github.v3.full+json
//   application/vnd.github.v3.diff
//   application/vnd.github.v3.patch
//
// HTTP Status Codes
//
// The HTTP Status Codes are used to indicate the status of the executed 
// operation. The available status codes are defined by RFC7231 and registered 
// status codes are listed in the IANA Status Code Registry.
//
// Specification
//
// Versions
//
// The OpenAPI Specification is versioned using a major.minor.patch versioning
// scheme. The major.minor portion of the version string (for example 3.1) SHALL
// designate the OAS feature set. .patch versions address errors in, or provide
// clarifications to, this document, not the feature set. Tooling which supports
// OAS 3.1 SHOULD be compatible with all OAS 3.1.* versions. The patch version 
// SHOULD NOT be considered by tooling, making no distinction between 3.1.0 and
// 3.1.1 for example.
//
// Occasionally, non-backwards compatible changes may be made in minor versions
// of the OAS where impact is believed to be low relative to the benefit 
// provided.
//
// An OpenAPI document compatible with OAS 3.*.* contains a required openapi 
// field which designates the version of the OAS that it uses.
//
// Format
//
// An OpenAPI document that conforms to the OpenAPI Specification is itself a 
// JSON object, which may be represented either in JSON or YAML format.
//
// For example, if a field has an array value, the JSON array representation 
// will be used:
//
//  {
//     "field": [ 1, 2, 3 ]
//  }
//
// All field names in the specification are case sensitive. This includes all 
// fields that are used as keys in a map, except where explicitly noted that
// keys are case insensitive.
//
// The schema exposes two types of fields: Fixed fields, which have a declared
// name, and Patterned fields, which declare a regex pattern for the field name.
//
// Patterned fields MUST have unique names within the containing object.
//
// In order to preserve the ability to round-trip between YAML and JSON formats,
// YAML version 1.2 is RECOMMENDED along with some additional constraints:
//
// Tags MUST be limited to those allowed by the JSON Schema ruleset.
// Keys used in YAML maps MUST be limited to a scalar string, as defined by the
// YAML Failsafe schema ruleset.
// Note: While APIs may be defined by OpenAPI documents in either YAML or JSON
// format, the API request and response bodies and other content are not 
// required to be JSON or YAML.
//
// Document Structure
//
// An OpenAPI document MAY be made up of a single document or be divided into
// multiple, connected parts at the discretion of the author. In the latter 
// case, Reference Objects and Schema Object $ref keywords are used.
//
// It is RECOMMENDED that the root OpenAPI document be named: openapi.json or
// openapi.yaml.
//
// Data Types
//
// Data types in the OAS are based on the types supported by the JSON Schema
// Specification Draft 2020-12. Note that integer as a type is also supported 
// and is defined as a JSON number without a fraction or exponent part. Models
// are defined using the Schema Object, which is a superset of JSON Schema 
// Specification Draft 2020-12.
//
// As defined by the JSON Schema Validation vocabulary, data types can have an
// optional modifier property: format. OAS defines additional formats to provide
// fine detail for primitive data types.
//
// The formats defined by the OAS are:
//
//  type	format	Comments
//  integer	int32	signed 32 bits
//  integer	int64	signed 64 bits (a.k.a long)
//  number	float	
//  number	double	
//  string	password	A hint to UIs to obscure input.
//
// Rich Text Formatting
//
// Throughout the specification description fields are noted as supporting 
// CommonMark markdown formatting. Where OpenAPI tooling renders rich text it 
// MUST support, at a minimum, markdown syntax as described by CommonMark 0.27.
// Tooling MAY choose to ignore some CommonMark features to address security 
// concerns.
//
// Relative References in URIs
//
// Unless specified otherwise, all properties that are URIs MAY be relative 
// references as defined by RFC3986.
//
// Relative references, including those in Reference Objects, PathItem Object 
// $ref fields, Link Object operationRef fields and Example Object externalValue
// fields, are resolved using the referring document as the Base URI according 
// to RFC3986.
//
// If a URI contains a fragment identifier, then the fragment should be resolved
// per the fragment resolution mechanism of the referenced document. If the
// representation of the referenced document is JSON or YAML, then the fragment
// identifier SHOULD be interpreted as a JSON-Pointer as per RFC6901.
//
// Relative references in Schema Objects, including any that appear as $id 
// values, use the nearest parent $id as a Base URI, as described by JSON Schema
// Specification Draft 2020-12. If no parent schema contains an $id, then the 
// Base URI MUST be determined according to RFC3986.
//
// Relative References in URLs
//
// Unless specified otherwise, all properties that are URLs MAY be relative 
// references as defined by RFC3986. Unless specified otherwise, relative 
// references are resolved using the URLs defined in the Server Object as a
// Base URL. Note that these themselves MAY be relative to the referring document.
//
// Schema
//
// In the following description, if a field is not explicitly REQUIRED or 
// described with a MUST or SHALL, it can be considered OPTIONAL.
package openapi

