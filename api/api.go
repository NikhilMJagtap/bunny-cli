// Package api contains the functions to interact with the BunnyCDN API.
// Generic utility for interacting with the API should be added here.
package api

import (
	"encoding/json"
)

// Utility to convert a struct to a map[string]interface{} for use as query parameters.
// Generally used for GET requests.
func GetQueryParamsFromOptions(data interface{}) map[string]interface{} {
	content, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	var queryParams map[string]interface{}
	err = json.Unmarshal(content, &queryParams)
	if err != nil {
		return nil
	}
	return queryParams
}
