package entity

type Order struct {
	ID         int64      `json:"id" db:"id"`
	Restaurant Restaurant `json:"restaurant" db:"restaurant"`
	Dish       Dish       `json:"dish" db:"dish"`
}
