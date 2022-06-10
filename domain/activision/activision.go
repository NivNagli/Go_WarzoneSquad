// In order to prevent multiple functions for the request validation and formatting
// ActivisionRequest interface will be implemented by all our request struct's

// This will give us the opportunity to use different request struct in the shared functions.

package activision

type ActivisionRequest interface {
	GetUsername() string
	GetPlatform() string
	GetGameID() string
	GetTarget() string
}
