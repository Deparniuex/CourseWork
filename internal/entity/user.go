package entity

type User struct {
	ID        int64    `json:"id" db:"id"`
	FirstName *string  `json:"first_name" db:"first_name"`
	LastName  *string  `json:"last_name" db:"last_name"`
	Phone     *string  `json:"phone" db:"phone"`
	Email     *string  `json:"email" db:"email"`
	Role      Role     `json:"role" db:"role"`
	Username  *string  `json:"username" db:"username"`
	Password  Password `json:"-" db:"password_hash"`
}

type Password struct {
	Plaintext *string `json:"-"`
	Hash      *[]byte `json:"-" db:"password_hash"`
}
