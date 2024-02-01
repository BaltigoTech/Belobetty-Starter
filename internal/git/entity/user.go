package entity

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

type User struct {
	UserName    string             `json:"username"`
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
		UserName:    strings.ToLower(userName),
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
		Repository: strings.ToLower(repositoryName),
		Permission: strings.ToLower(permission),
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
	err := validUserName(u.UserName)
	if err != nil {
		return err
	}
	err = validEmail(u.Email)
	if err != nil {
		return err
	}
	if u.Name == "" {
		return errors.New("name is required")
	}
	for _, access := range u.Permissions {
		err = access.validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func validUserName(userName string) error {
	if userName == "" {
		return errors.New("username is required")
	}
	if len(userName) > 39 {
		return errors.New("username is too long (maximum is 39 characters)")
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$`)
	if !regex.MatchString(userName) {
		return errors.New("invalid username, may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen")
	}
	return nil
}

func validEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid e-mail")
	}

	return nil
}
