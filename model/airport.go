package model

type CreateAirportModel struct {
	Location        string `json:"location" validate:"required,min=4"`
	AirportName     string `json:"airport_name" validate:"required,min=3"`
	LocationAcronym string `json:"location_acronym" validate:"required,max=3"`
}

type UpdateAirportModel struct {
	Location        string `json:"location" validate:"min=4"`
	AirportName     string `json:"airport_name" validate:"min=3"`
	LocationAcronym string `json:"location_acronym" validate:"alpha,max=3"`
}
