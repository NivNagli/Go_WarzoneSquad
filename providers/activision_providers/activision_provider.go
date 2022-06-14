package activision_providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/NivNagli/WarzoneSquad_Go/clients/restclient"
	"github.com/NivNagli/WarzoneSquad_Go/domain/activision"
)

/********************************* Functions for getting the last games stats by recent/date/cycles *************************************/

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
			log.Printf("error when trying to unmarshal last games successful response: %s\n", err.Error())
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

// GetLastGamesStatsByDate func responsible for return the player last games stats according to date string
// the represented in utc time format, the result will be according to the last 20 games from the
// given date string.
func GetLastGamesStatsByDate(r activision.LastGamesRequest, d string) (*activision.LastGamesResponse, error) {
	// First we append the create for our request that will have the tokens from the environment variables and the user-agent header
	// in case we dont find the environment variables we will return error.
	headers, err := AddHeadersForActivisionRequest()
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, &activision.ActivisionErrorResponse{Message: "Error: Failed to set the headers request for last game stats request\n", StatusCode: 500}
	}
	// After we create the headers we creating now the url for the last games stats endpoint from the official API.
	// In case we received invalid username / platform we will return error.
	url, err := CreateLastGamesStatsByDateUrl(r, d)
	if err != nil {
		return nil, err
	}
	// Now that we the url and the headers we can send the request, we are doing that with the "restclient" Get method which will handle the sending procedure for us,
	// i implemented the restclient because i want to reduce the repeated code and also to have the option to mock the response result for mocking that will serve us in the tests.
	// In case of successful request we will get the *http.Response object and err == nil, in the case of failure we will recive nil and the err.
	response, err := restclient.Get(url, nil, *headers)
	if err != nil {
		log.Printf("error when trying to get last games stats by date from activision API: %s\n", err.Error())
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
			log.Printf("error when trying to unmarshal last games by date successful response: %s\n", err.Error())
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

// GetLastGamesStatsByCycles will try to fill the player's last games array with more then 20 games
// because activision API return for us only the last 20 games for given request i create this
// function which will return the cycles * 20 games list.
func GetLastGamesStatsByCycles(r activision.LastGamesRequest, c int) (*activision.LastGamesResponse, error) {
	// First we are getting the recent 20 games with the 'GetLastGamesStats' method
	firstResult, err := GetLastGamesStats(r)
	if err != nil {
		return nil, err
	}
	// Now we create 2 slices which will help us to extract the games past the last 20
	responsesArray := make([]activision.LastGamesResponse, c) // This array will hold the responses that comes from 'GetLastGamesStatsByDate'
	responsesArray[0] = *firstResult                          // In this stage we already have the first response so we insert her
	matchesArray := make([][]activision.Match, c)             // This array will hold the 'Matches' array field from each response that comes from 'GetLastGamesStatsByDate'
	matchesArray[0] = firstResult.Data.Matches
	// In this loop we will use the last recorded game date that we have in order to find the
	// 20 games that occur after him, and then will save the results in order to use them if we didn't
	// finish all the cycles and for the return value.
	for i := 0; i < c-1; i++ {
		// Reading the last game date that we have for this cycle
		dateInUtc := fmt.Sprintf("%d", int(responsesArray[i].Data.Matches[len(responsesArray[i].Data.Matches)-1].UtcStartSeconds))
		// Getting the 20 games past this date.
		newResult, err := GetLastGamesStatsByDate(r, dateInUtc+"000") // Needed to add the '000' because the official API does not save the date in the correct format.
		if err != nil {
			return nil, err
		}
		// save them for the next cycle and for the result array.
		responsesArray[i+1] = *newResult
		matchesArray[i+1] = newResult.Data.Matches
	}
	// after we save the slices in the 'matchArray' slice we need to append each one of them
	// into the first 20 games that we extract and that what this loop does.
	for i, m := range matchesArray {
		// if i == 0 thats mean that we are on the first slice that we are concat so we don't need to concat him to himself..
		if i != 0 {
			firstResult.Data.Matches = append(firstResult.Data.Matches, m...)
		}
	}
	// Result that contain 20*cycles games array.
	return firstResult, nil
}

/***************************************** Function for getting the weekly and lifetime stats ********************************/

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
			log.Printf("error when trying to unmarshal life time and weekly successful response: %s\n", err.Error())
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
