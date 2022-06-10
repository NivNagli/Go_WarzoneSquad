// Package restclient implements utility routines for sending http requests procedures
// the methods will receive the url, body, headers that needed for execute the send request and the restclient
// will preform them according them and will return the original http response and error.

// The restclient package in the future will also help us to implement unit test with a mocking mode option that will not preform a real send request but will
// return a mock response that we will set instead.

package restclient

import (
	"bytes"

	"encoding/json"
	"net/http"
)

func Get(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
