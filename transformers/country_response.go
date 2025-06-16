package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformCountry(country models.Country) dto.CountryResponse {

	return dto.CountryResponse{
		ID: country.Id,
		Name: country.Name,
	}
}