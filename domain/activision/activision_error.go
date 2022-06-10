package activision

type ActivisionErrorResponse struct {
	StatusCode int               `json:"status_code"`
	Message    string            `json:"message"`
	Errors     []ActivisionError `json:"errors"`
}

type ActivisionError struct {
	Status string                      `json:"status"`
	Data   ActivisionErrorResponseData `json:"data"`
}

type ActivisionErrorResponseData struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e ActivisionErrorResponse) Error() string {
	return e.Message
}
