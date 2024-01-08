package entity

type City struct {
	ID       int64  `json:"id" db:"id"`
	CityName string `json:"city_name" db:"city_name"`
}
