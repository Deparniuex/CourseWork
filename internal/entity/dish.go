package entity

type Dish struct {
	ID     int64   `json:"id" db:"id"`
	Name   *string `json:"name" db:"name"`
	Weight int64   `json:"weight" db:"weight"`
	Price  int64   `json:"price" db:"price"`
	Menu   Menu    `json:"menu" db:"menu"`
}
