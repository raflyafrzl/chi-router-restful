package model

type ResponseWeb struct {
	Data       any    `json:"data"`
	Message    string `json:"message"`
	StatusCode int16  `json:"status_code"`
	Status     string `json:"status"`
}

type ResponseFailWeb struct {
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}
