package api

type CreateCityRequest struct {
	CityName string `json:"city_name" binding:"required"`
}

type DeleteCityRequest struct {
	ID
	CityName string `json:"city_name" binding:"required"`
}
