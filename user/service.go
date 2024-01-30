package user

// create a contract (it's like a method in a controller that accesible)

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// function for register user
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "USER"

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, err
}

// function for validate login
func (s *service) Login(input LoginInput) (User, error) {
	// get user input
	email := input.Email
	password := input.Password

	// check email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	// check if user not exist
	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	// check password (using bcrypt library to compare hash with user input)
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	// check if user not exist
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}