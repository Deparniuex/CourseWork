package api

type CreateCityRequest struct {
	CityName string `json:"city_name" binding:"required"`
}

type GetCitiesRequest struct {
	CityName string `json:"city_name" binding:"required"`
}
