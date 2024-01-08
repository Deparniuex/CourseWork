package entity

type OrderHistory struct {
	ID    int64 `json:"id" db:"id"`
	Order Order `json:"order" db:"ord"`
}
