package app

type healthCheckResponse struct {
	Success bool
}

type customResponse struct {
	Success bool        `json:"Success"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

type errorResponse struct {
	Error string `json:"Error,omitempty"`
}
