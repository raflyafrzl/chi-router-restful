package model

type ResponseWeb struct {
	Data       any    `json:"data"`
	Message    string `json:"message"`
	StatusCode int16  `json:"status_code"`
	Status     string `json:"status"`
}
