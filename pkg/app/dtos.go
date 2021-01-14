package app

type healthCheckResponse struct {
	Success bool
}

type errorResponse struct {
	Error string `json:",omitempty"`
}
