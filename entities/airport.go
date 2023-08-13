package entities

type AirportEntity struct {
	Id              string `gorm:"type:uuid;column:id" json:"id"`
	AirportName     string `gorm:"type:varchar(40);column:airport_name" json:"airport_name"`
	AirportCode     string `gorm:"type:varchar(3);column:airport_code" json:"airport_code"`
	Location        string `gorm:"type:text;column:location" json:"location"`
	LocationAcronym string `gorm:"type:varchar(3);column:location_acronym" json:"location_acronym"`
}
