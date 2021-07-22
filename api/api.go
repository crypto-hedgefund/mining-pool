package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Post() is a generic "POST" request, meant for a RPC server, returns http.response
func Post(path string, jsonRpc string, method string, params interface{}, id int) (*http.Response, error) {
	// `map[interface{}]interface{}` Allows you to assign any type key or value
	requestBody, err := json.Marshal(map[interface{}]interface{}{
		"jsonrpc": jsonRpc,
		"method":  method,
		"params":  params,
		"id":      id,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(path, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	// Check for success
	if resp.StatusCode != 200 {
		return nil, err
	}

	// Return as http.Resp to extract JSON easily
	return resp, nil
}
