package domain

type User struct {
	Id uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	email string `json:"email"`
}
