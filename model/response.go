package model

type ResponseWeb struct {
	Data       any    `json:"data"`
	Message    string `json:"message"`
	StatusCode int16  `json:"status_code"`
	Status     string `json:"status"`
}

type ResponseFailWeb struct {
	Error      any    `json:"error"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}
