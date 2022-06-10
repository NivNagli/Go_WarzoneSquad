// Package config will be responsible for getting environment variables from the operating system
// and to return an error if the environment variables are not set.

package config

import (
	"errors"
	"os"
	"strconv"
)

type ActivisionTokens struct {
	ATKN                  string `json:"atkn"`
	ACT_SSO_COOKIE        string `json:"act_sso_cookie"`
	ACT_SSO_COOKIE_EXPIRY int    `json:"act_sso_cookie_expires"`
}

const (
	secret_ATKN                  = "ATKN"
	secret_ACT_SSO_COOKIE        = "ACT_SSO_COOKIE"
	secret_ACT_SSO_COOKIE_EXPIRY = "ACT_SSO_COOKIE_EXPIRY"
)

var (
	ATKN                  = os.Getenv(secret_ATKN)
	ACT_SSO_COOKIE        = os.Getenv(secret_ACT_SSO_COOKIE)
	ACT_SSO_COOKIE_EXPIRY = os.Getenv(secret_ACT_SSO_COOKIE_EXPIRY)
)

func GetActivisionAccessTokens() (*ActivisionTokens, error) {
	if len(ATKN) > 0 && len(ACT_SSO_COOKIE) > 0 && len(ACT_SSO_COOKIE_EXPIRY) > 0 {
		if stringToInt, err := strconv.Atoi(ACT_SSO_COOKIE_EXPIRY); err == nil {
			return &ActivisionTokens{ATKN, ACT_SSO_COOKIE, stringToInt}, nil
		}
	}
	return nil, errors.New("failed to read one of the tokens from the environment variable")
}
