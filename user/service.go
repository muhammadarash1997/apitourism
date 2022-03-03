package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(userInput RegisterUserInput) (User, error)
	DeleteUserByUUID(id string) error
	Login(loginInput LoginInput) (User, error)
	GetUser(id string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (this *service) RegisterUser(userInput RegisterUserInput) (User, error) {
	user := User{}

	user.Name = userInput.Name
	user.Email = userInput.Email
	user.Role = userInput.Role
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	userRegistered, err := this.repository.Save(user)
	if err != nil {
		return userRegistered, err
	}

	return userRegistered, nil
}

func (this *service) DeleteUserByUUID(id string) error {
	err := this.repository.DeleteByUUID(id)
	if err != nil {
		return err
	}

	return nil
}

func (this *service) Login(loginInput LoginInput) (User, error) {
	email := loginInput.Email
	password := loginInput.Password

	user, err := this.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, errors.New("wrong password")
	}

	return user, nil
}

func (this *service) GetUser(id string) (User, error) {
	user, err := this.repository.GetByUUID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}