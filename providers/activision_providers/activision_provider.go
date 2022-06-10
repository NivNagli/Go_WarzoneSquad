package activision_providers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/NivNagli/WarzoneSquad_Go/clients/restclient"
	"github.com/NivNagli/WarzoneSquad_Go/domain/activision"
)

// GetLastGameStats return the response from the last game stats request that sent to the official API.
// In case of error we will return the ActivisionErrorResponse object that will contain the error status code that we set and a short message within him,
// the function argument is LastGamesRequest object that will contain the username and platform for the player that we want to get the last game stats for.

// In case in the future activision will change the struct of their response object or the authorization header we will get error from this method, i made a spereated tests
// of each case here and as of the date [10.6.2022] the response object and authorization headers are defined accordingly.
// We must make sure that the username and platform that we received here are fully
func GetLastGamesStats(r activision.LastGamesRequest) (*activision.LastGamesResponse, error) {
	// First we append the create for our request that will have the tokens from the environment variables and the user-agent header
	// in case we dont find the environment variables we will return error.
	headers, err := AddHeadersForActivisionRequest()
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, &activision.ActivisionErrorResponse{Message: "Error: Failed to set the headers request for last game stats request\n", StatusCode: 500}
	}
	// After we create the headers we creating now the url for the last games stats endpoint from the official API.
	// In case we received invalid username / platform we will return error.
	url, err := CreateLastGamesStatsUrl(r)
	if err != nil {
		return nil, err
	}
	// Now that we the url and the headers we can send the request, we are doing that with the "restclient" Get method which will handle the sending procedure for us,
	// i implemented the restclient because i want to reduce the repeated code and also to have the option to mock the response result for mocking that will serve us in the tests.
	// In case of successful request we will get the *http.Response object and err == nil, in the case of failure we will recive nil and the err.
	response, err := restclient.Get(url, nil, *headers)
	if err != nil {
		log.Printf("error when trying to get last games stats from activision API: %s\n", err.Error())
		return nil, &activision.ActivisionErrorResponse{Message: "Error: Failed to get data from activision API\n", StatusCode: 500}
	}
	// We must to close the response.Body object when we finish to work on him, so i set defer that will execute when this method scope will close.
	defer response.Body.Close()
	// Attempt to read the response body into []byte object.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to parse response body: %s\n", err.Error())
		return nil, &activision.ActivisionErrorResponse{Message: "Error: Invalid response body\n", StatusCode: 500}
	}
	// Now that we have the response body in []byte format, we can try to unmarshal the response into our LastGamesResponse object, because activision API doesn't use Status code to indicate
	// if the request failed because of an invalid player details, we need to check the response status string that comes from the response body.
	// So we have 2 cases which we will return error from this unmarshal: first one is that our LastGamesResponse failed to receive the data from the response body and the second case is that
	// the data was not found for the given username and platform...
	var result activision.LastGamesResponse
	if err := json.Unmarshal(body, &result); err != nil || result.Status == "error" {
		if err != nil {
			// in this case we need to check what are the differences between the activision.LastGamesResponse struct to the response that we receive from the official API for the last games endpoint.
			log.Printf("error when trying to unmarshal create repo successful response: %s\n", err.Error())
			return nil, &activision.ActivisionErrorResponse{Message: "Error: Invalid response body\n", StatusCode: 500}
		} else {
			// The case the request failed that can happen if the tokens expired or if the user gave us invalid username or platform.
			log.Printf("Error: failed to access to activision API due invalid arguments, ask the user to verify his info and check your activision tokens!")
			return nil, &activision.ActivisionErrorResponse{Message: "Error: invalid player details make sure you have public profile\nif you do have public profile contact us with error code NN97\n", StatusCode: 500}
		}

	}
	// Finally that is the good case, i added manually the username and platform for the LastGamesResponse result object.
	result.Username = r.Username
	result.Platform = r.Platform
	return &result, nil
}

// Pretty much just like the GetLastGamesStats method, except this time we are pointing to the lifetime and weekly endpoint
// thus we need to work with different response object and url, except that the logic almost the same...
func GetLifetimeAndWeeklyStats(r activision.LifetimeAndWeeklyRequest) (*activision.LifetimeAndWeeklyResponse, error) {
	// First we append the create for our request that will have the tokens from the environment variables and the user-agent header
	// in case we don't find the environment variables we will return error.
	headers, err := AddHeadersForActivisionRequest()
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, &activision.ActivisionErrorResponse{Message: "Error: Failed to set the headers request for last game stats request\n", StatusCode: 500}
	}
	// After we create the headers we creating now the url for the last games stats endpoint from the official API.
	// In case we received invalid username / platform we will return error.
	url, err := CreateLifetimeAndWeeklyUrl(r)
	if err != nil {
		return nil, err
	}
	// Now that we the url and the headers we can send the request, we are doing that with the "restclient" Get method which will handle the sending procedure for us,
	// i implemented the restclient because i want to reduce the repeated code and also to have the option to mock the response result for mocking that will serve us in the tests.
	// In case of successful request we will get the *http.Response object and err == nil, in the case of failure we will recive nil and the err.
	response, err := restclient.Get(url, nil, *headers)
	if err != nil {
		log.Printf("error when trying to get last games stats from activision API: %s\n", err.Error())
		return nil, &activision.ActivisionErrorResponse{Message: "Error: Failed to get data from activision API\n", StatusCode: 500}
	}
	// We must to close the response.Body object when we finish to work on him, so i set defer that will execute when this method scope will close.
	defer response.Body.Close()
	// Attempt to read the response body into []byte object.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to parse response body: %s\n", err.Error())
		return nil, &activision.ActivisionErrorResponse{Message: "Error: Invalid response body\n", StatusCode: 500}
	}
	// Now that we have the response body in []byte format, we can try to unmarshal the response into our LifetimeAndWeeklyResponse object, because activision API doesn't use Status code to indicate
	// if the request failed because of an invalid player details, we need to check the response status string that comes from the response body.
	// So we have 2 cases which we will return error from this unmarshal: first one is that our LifetimeAndWeeklyResponse failed to receive the data from the response body and the second case is that
	// the data was not found for the given username and platform...
	var result activision.LifetimeAndWeeklyResponse
	if err := json.Unmarshal(body, &result); err != nil || result.Status == "error" {
		if err != nil {
			// in this case we need to check what are the differences between the activision.LifetimeAndWeeklyResponse struct to the response that we receive from the official API for the Lifetime&Weekly endpoint.
			log.Printf("error when trying to unmarshal create repo successful response: %s\n", err.Error())
			return nil, &activision.ActivisionErrorResponse{Message: "Error: Invalid response body\n", StatusCode: 500}
		} else {
			// The case the request failed that can happen if the tokens expired or if the user gave us invalid username or platform.
			log.Printf("Error: failed to access to activision API due invalid arguments, ask the user to verify his info and check your activision tokens!")
			return nil, &activision.ActivisionErrorResponse{Message: "Error: invalid player details make sure you have public profile\nif you do have public profile contact us with error code NN97\n", StatusCode: 500}
		}

	}
	// Finally that is the good case, i added manually the username and platform for the LifetimeAndWeeklyResponse result object.
	return &result, nil
}
