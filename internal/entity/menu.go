package entity

type Menu struct {
	ID         int64      `json:"id" db:"id"`
	Name       *string    `json:"name" db:"name"`
	Restaurant Restaurant `json:"restaurant" db:"restaurant"`
}
