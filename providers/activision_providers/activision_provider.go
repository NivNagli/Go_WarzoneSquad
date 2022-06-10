package activision_providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/NivNagli/WarzoneSquad_Go/clients/restclient"
	"github.com/NivNagli/WarzoneSquad_Go/config"
	"github.com/NivNagli/WarzoneSquad_Go/domain/activision"
)

const (
	urlGetLastGameStats       = "https://my.callofduty.com/api/papi-client/crm/cod/v2/title/mw/platform/%s/gamer/%s/matches/wz/start/0/end/0/details" // wildcard for the last game request url endpoint in the official API
	headerAuthorizationFormat = "ACT_SSO_COOKIE=%s; ACT_SSO_COOKIE_EXPIRY=%d; atkn=%s;"                                                               // wildcard for the authorization cookie header
)

// getAuthorizationHeader try to exctract the tokens from the environment variable that needed for the authorization header Cookie.
// in case of an error that occurs when one pf the environment variable is not set we will return empty string and the error.
// a successful exctract return err == nil and the authorization header string that contains the authorization tokens.
func getAuthorizationHeader() (string, error) {
	tokens, err := config.GetActivisionAccessTokens()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(headerAuthorizationFormat, tokens.ACT_SSO_COOKIE, tokens.ACT_SSO_COOKIE_EXPIRY, tokens.ATKN), nil
}

// createLastGameStatsUrl will try to fill the last game stats url wildcard with the username and platform that received from the LastGamesRequest.
// in case of an error we will return empty string and the error, and in case of successful fill we return the fixed url and err==nil.
func createLastGamesStatsUrl(r activision.LastGamesRequest) (string, error) {
	if len(r.Username) == 0 || len(r.Platform) == 0 {
		return "", &activision.ActivisionErrorResponse{Message: "Error: missing argument for receive for the search\n", StatusCode: 400}
	}

	if err := validatePlatform(r); err != nil {
		return "", err
	}

	fixedUsername, err := fixUsername(r)
	if err != nil {
		return "", err
	}

	r.Username = fixedUsername
	return fmt.Sprintf(urlGetLastGameStats, r.Platform, fixedUsername), nil
}

// validatePlatform validates the platform which we received from the LastGamesRequest, in case of invalid platform name we will return an error, and in case of valid platform name we will return err==nil
func validatePlatform(r activision.LastGamesRequest) error {
	platforms := []string{"psn", "xbl", "battle", "uno"}
	for _, platform := range platforms {
		if platform == r.Platform {
			return nil
		}
	}
	return &activision.ActivisionErrorResponse{Message: "Error: invalid platform received for last games stats", StatusCode: 400}
}

// fixUsername function will made a url encoding for player from "battle" or "uno" which their username contain '#' that need to convert into %23 encoding,
// if player is not from those platform we just return his name as we received, in case of an error the function return empty string and the error.
// in successful case err == nil
func fixUsername(r activision.LastGamesRequest) (string, error) {
	// for example the username: "nivGolanigo#1234" will conveted into "nivGolanigo%231234" if his platform is "battle" or "uno".
	splitedUsername := strings.Split(r.Username, "#")

	if r.Platform == "battle" || r.Platform == "uno" {
		if len(splitedUsername) != 2 {
			return "", &activision.ActivisionErrorResponse{Message: "Error: invalid username received for last games stats request\n", StatusCode: 400}
		}
		return splitedUsername[0] + "%23" + splitedUsername[1], nil
	}

	// a user from the other platforms should not contain the '#' character in his name!
	if len(splitedUsername) != 1 {
		return "", &activision.ActivisionErrorResponse{Message: "Error: invalid username received for last games stats request\n", StatusCode: 400}
	}
	return r.Username, nil
}

func addHeadersForLastGamesRequest() (*http.Header, error) {
	headers := http.Header{}
	headers.Add("user-agent", "golangNivos")
	cookieHeader, err := getAuthorizationHeader()
	if err != nil {
		log.Printf("Error when try to get the tokens value from the environment variables")
		return nil, &activision.ActivisionErrorResponse{Message: "Error: Enviroment variables not set", StatusCode: 500}
	}

	headers.Add("Cookie", cookieHeader)
	return &headers, nil
}

// GetLastGameStats return the response from the last game stats request that sent to the official API.
// In case of error we will return the ActivisionErrorResponse object that will contain the error status code that we set and a short message within him,
// the function argument is LastGamesRequest object that will contain the username and platform for the player that we want to get the last game stats for.

// In case in the future activision will change the struct of their response object or the authorization header we will get error from this method, i made a spereated tests
// of each case here and as of the date [10.6.2022] the response object and authorization headers are defined accordingly.
// We must make sure that the username and platform that we received here are fully
func GetLastGamesStats(r activision.LastGamesRequest) (*activision.LastGamesResponse, error) {
	// First we append the create for our request that will have the tokens from the environment variables and the user-agent header
	// in case we dont find the environment variables we will return error.
	headers, err := addHeadersForLastGamesRequest()
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, &activision.ActivisionErrorResponse{Message: "Error: Failed to set the headers request for last game stats request\n", StatusCode: 500}
	}
	// After we create the headers we creating now the url for the last games stats endpoint from the official API.
	// In case we received invalid username / platform we will return error.
	url, err := createLastGamesStatsUrl(r)
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
	// that the request failed because of an invalid player details, we need to check the response status string that comes from the response body.
	// So we have 2 cases which will return error from this unmarshal: first one is that our LastGamesResponse failed to receive the data from the response body and the second case is that
	// the data was not found for the givven username and platform...
	var result activision.LastGamesResponse
	if err := json.Unmarshal(body, &result); err != nil || result.Status == "error" {
		if err != nil {
			// in this case we need to check what are the diffrences between the activision.LastGamesResponse struct to the response that we receive from the official API for the last games endpoint.
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
