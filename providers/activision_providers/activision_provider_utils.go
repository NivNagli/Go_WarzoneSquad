// This file contains shared functions that will be used in activision_providers.go file in
// different methods that not have common request object as argument.

package activision_providers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/NivNagli/WarzoneSquad_Go/config"
	"github.com/NivNagli/WarzoneSquad_Go/domain/activision"
)

const (
	urlGetLastGameStats       = "https://my.callofduty.com/api/papi-client/crm/cod/v2/title/mw/platform/%s/gamer/%s/matches/wz/start/0/end/0/details"  // wildcard for the last game request url endpoint in the official API
	urlGetLastGameStatsByDate = "https://my.callofduty.com/api/papi-client/crm/cod/v2/title/mw/platform/%s/gamer/%s/matches/wz/start/0/end/%s/details" // wildcard for the last game request url endpoint in the official API
	urlGetLifetimeAndWeekly   = "https://my.callofduty.com/api/papi-client/stats/cod/v1/title/mw/platform/%s/gamer/%s/profile/type/wz"                 // wildcard for the lifetime and weekly request url endpoint in the official API
	headerAuthorizationFormat = "ACT_SSO_COOKIE=%s; ACT_SSO_COOKIE_EXPIRY=%d; atkn=%s;"                                                                // wildcard for the authorization cookie header
)

/***************************************** Help functions for setting headers *****************************************/
func AddHeadersForActivisionRequest() (*http.Header, error) {
	headers := http.Header{}
	headers.Add("user-agent", "golangNiv")
	cookieHeader, err := GetAuthorizationHeader()
	if err != nil {
		log.Printf("Error when try to get the tokens value from the environment variables")
		return nil, &activision.ActivisionErrorResponse{Message: "Error: Environment variables not set", StatusCode: 500}
	}

	headers.Add("Cookie", cookieHeader)
	return &headers, nil
}

// getAuthorizationHeader try to extract the tokens from the environment variable that needed for the authorization header Cookie.
// in case of an error that occurs when one pf the environment variable is not set we will return empty string and the error.
// a successful extract return err == nil and the authorization header string that contains the authorization tokens.
func GetAuthorizationHeader() (string, error) {
	tokens, err := config.GetActivisionAccessTokens()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(headerAuthorizationFormat, tokens.ACT_SSO_COOKIE, tokens.ACT_SSO_COOKIE_EXPIRY, tokens.ATKN), nil
}

/***************************************** Help functions for ActivisionRequest objects **********************************************/

// validatePlatform validates the platform which we received from the ActivisionRequest, in case of invalid platform name we will return an error, and in case of valid platform name we will return err==nil
func ValidatePlatform(r activision.ActivisionRequest) error {
	platforms := []string{"psn", "xbl", "battle", "uno"}
	for _, platform := range platforms {
		if platform == r.GetPlatform() {
			return nil
		}
	}
	return &activision.ActivisionErrorResponse{Message: "Error: invalid platform received for last games stats", StatusCode: 400}
}

// fixUsername function will made a url encoding for player from "battle" or "uno" which their username contain '#' that need to convert into %23 encoding,
// if player is not from those platform we just return his name as we received, in case of an error the function return empty string and the error.
// in successful case err == nil
func FixUsername(r activision.ActivisionRequest) (string, error) {
	// for example the username: "nivGolanigo#1234" will conveted into "nivGolanigo%231234" if his platform is "battle" or "uno".
	splittedUsername := strings.Split(r.GetUsername(), "#")

	if r.GetPlatform() == "battle" || r.GetPlatform() == "uno" {
		if len(splittedUsername) != 2 {
			return "", &activision.ActivisionErrorResponse{Message: "Error: invalid username received for last games stats request\n", StatusCode: 400}
		}
		return splittedUsername[0] + "%23" + splittedUsername[1], nil
	}

	// a user from the other platforms should not contain the '#' character in his name!
	if len(splittedUsername) != 1 {
		return "", &activision.ActivisionErrorResponse{Message: "Error: invalid username received for last games stats request\n", StatusCode: 400}
	}
	return r.GetUsername(), nil
}

/***************************************** Help functions for specific ActivisionRequest objects **********************************************/

// createLastGameStatsUrl will try to fill the urlGetLastGameStats url wildcard with the username and platform that received from the ActivisionRequest.
// in case of an error we will return empty string and the error, and in case of successful fill we return the fixed url and err==nil.
func CreateLastGamesStatsUrl(r activision.ActivisionRequest) (string, error) {
	if len(r.GetUsername()) == 0 || len(r.GetPlatform()) == 0 {
		return "", &activision.ActivisionErrorResponse{Message: "Error: missing argument for receive for the last games stats search\n", StatusCode: 400}
	}

	if err := ValidatePlatform(r); err != nil {
		return "", err
	}

	fixedUsername, err := FixUsername(r)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(urlGetLastGameStats, r.GetPlatform(), fixedUsername), nil
}

// createLastGameStatsUrl will try to fill the urlGetLastGameStats url wildcard with the username and platform that received from the ActivisionRequest.
// in case of an error we will return empty string and the error, and in case of successful fill we return the fixed url and err==nil.
func CreateLastGamesStatsByDateUrl(r activision.ActivisionRequest, d string) (string, error) {
	if len(r.GetUsername()) == 0 || len(r.GetPlatform()) == 0 || len(d) == 0 {
		return "", &activision.ActivisionErrorResponse{Message: "Error: missing argument for receive for the last games stats by date search\n", StatusCode: 400}
	}

	if err := ValidatePlatform(r); err != nil {
		return "", err
	}

	fixedUsername, err := FixUsername(r)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(urlGetLastGameStatsByDate, r.GetPlatform(), fixedUsername, d), nil
}

// CreateLifetimeAndWeeklyUrl will try to fill the urlGetLifetimeAndWeekly url wildcard with the username and platform that received from the ActivisionRequest.
// in case of an error we will return empty string and the error, and in case of successful fill we return the fixed url and err==nil.
func CreateLifetimeAndWeeklyUrl(r activision.ActivisionRequest) (string, error) {
	if len(r.GetUsername()) == 0 || len(r.GetPlatform()) == 0 {
		return "", &activision.ActivisionErrorResponse{Message: "Error: missing argument for receive for the lifetime and weekly stats search\n", StatusCode: 400}
	}

	if err := ValidatePlatform(r); err != nil {
		return "", err
	}

	fixedUsername, err := FixUsername(r)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(urlGetLifetimeAndWeekly, r.GetPlatform(), fixedUsername), nil
}
