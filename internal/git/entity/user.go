package entity

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

type User struct {
	Username    string             `json:"username"`
	Email       string             `json:"email"`
	Name        string             `json:"name"`
	Permissions []AccessRepository `json:"permissions"`
}

type AccessRepository struct {
	Repository string `json:"repository,omitempty"`
	Permission string `json:"permission,omitempty"`
}

func NewUser(userName, email, name string, permission []AccessRepository) (*User, error) {
	user := &User{
		Username:    userName,
		Email:       email,
		Name:        name,
		Permissions: permission,
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewAccessRepository(repositoryName, permission string) (*AccessRepository, error) {
	access := &AccessRepository{
		Repository: repositoryName,
		Permission: permission,
	}
	err := access.validate()
	if err != nil {
		return nil, err
	}
	return access, nil
}

func (r *AccessRepository) validate() error {
	switch strings.ToUpper(r.Permission) {
	case Owner, Collaborator, Maintainer, Reader:
		return nil
	default:
		return fmt.Errorf("invalid user category permission: %s", r.Permission)
	}

}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return err
	}
	for _, access := range u.Permissions {
		err = access.validate()
		if err != nil {
			return err
		}
	}
	return nil
}
