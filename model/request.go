package model

//TODO:Changing Request Field
type AirportCreateRequest struct {
	AirportName string `json:"airport_name"`
	Data        any    `json:"data"`
}
