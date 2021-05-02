// Copyright 2021 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

// Returns an OpenAPI instance from JSON data or an error.
func FromJSON(data []byte) (result *OpenAPI, err error) {
	result = &OpenAPI{}
	if err = json.Unmarshal(data, result); err != nil {
		result = nil
	}
	return
}

// Returns an OpenAPI instance from YAML data or an error.
func FromYAML(data []byte) (result *OpenAPI, err error) {
	result = &OpenAPI{}
	if err = yaml.Unmarshal(data, result); err != nil {
		result = nil
	}
	return
}