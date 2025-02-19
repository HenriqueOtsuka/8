package entity

import (
	"github.com/HenriqueOtsuka/8/pkg/entity"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
}

func NewUser(username, password, email string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Username: username,
		Password: string(hash),
		Email:    email,
	}, nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password),
		[]byte(password))
}
