package entity

type Restaurant struct {
	ID     int64   `json:"id" db:"id"`
	Name   *string `json:"name" db:"name"`
	City   int64   `json:"city" db:"city"`
	Street int64   `json:"street" db:"street"`
}
