package entity

type Token struct {
	Plaintext string `json:"token"`
	UserID    int64  `json:"-" db:"user_id"`
	Hash      []byte `json:"-" db:"token"`
}
