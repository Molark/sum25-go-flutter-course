package user

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	// ErrInvalidEmail is returned when the email format is invalid
	ErrInvalidEmail = errors.New("invalid email format")
	// ErrInvalidAge is returned when the age is invalid
	ErrInvalidAge = errors.New("invalid age: must be between 0 and 150")
	// ErrEmptyName is returned when the name is empty
	ErrEmptyName = errors.New("name cannot be empty")
)

// User represents a user in the system
type User struct {
	Name  string
	Age   int
	Email string
}

// NewUser creates a new user with validation
func NewUser(name string, age int, email string) (*User, error) {
	// TODO: Implement user creation with validation
	if len(name) == 0 {
		return nil, ErrEmptyName
	}
	if age < 1 || age > 130 {
		return nil, ErrInvalidAge
	}
	check := regexp.MustCompile(`^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]+$`)
	if !check.Match([]byte(email)) {
		return nil, ErrInvalidEmail
	}
	user := User{
		Name:  name,
		Age:   age,
		Email: email,
	}
	return &user, nil
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	// TODO: Implement user validation
	if len(u.Name) == 0 {
		return ErrEmptyName
	}
	if u.Age < 1 || u.Age > 130 {
		return ErrInvalidAge
	}
	check := regexp.MustCompile(`^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]+$`)
	if !check.Match([]byte(u.Email)) {
		return ErrInvalidEmail
	}
	return nil
}

// String returns a string representation of the user
func (u *User) String() string {
	// TODO: Implement string representation
	return u.Name + " " + strconv.Itoa(u.Age) + " " + u.Email
}

// IsValidEmail checks if the email format is valid
func IsValidEmail(email string) bool {
	// TODO: Implement email validation
	check := regexp.MustCompile(`^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]+$`)
	if !check.Match([]byte(email)) {
		return false
	}
	return true
}
