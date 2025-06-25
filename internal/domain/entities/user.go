package entities

import gonanoid "github.com/matoous/go-nanoid/v2"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name, email string) *User {
	id, err := gonanoid.New(21)

	if err != nil {
		panic(err)
	}

	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
