// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

import (
	"fmt"
	"testing"
)

func TestFromYAML(t *testing.T) {
	const data string = `openapi: 3.0.0
info:
	title: Sample API
	description: Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.
	version: 0.1.9

servers:
	- url: http://api.example.com/v1
	description: Optional server description, e.g. Main (production) server
	- url: http://staging-api.example.com
	description: Optional server description, e.g. Internal staging server for testing

paths:
	/users:
	get:
		summary: Returns a list of users.
		description: Optional extended description in CommonMark or HTML.
		responses:
		'200':    # status code
			description: A JSON array of user names
			content:
			application/json:
				schema: 
				type: array
				items: 
					type: string
`
	oa, err := FromYAML([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(oa)
}