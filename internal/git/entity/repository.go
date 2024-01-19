package entity

import (
	"errors"
	"fmt"
	"regexp"
)

type Repository struct {
	Name        string
	Description string
	Private     bool
	Users       UserRepository
}

type UserRepository map[string]string

const (
	Owner        = "Owner"
	Collaborator = "Collaborator"
	Maintainer   = "Maintainer"
	Reader       = "Reader"
)

func NewRepository(name, description string, private bool, users UserRepository) (*Repository, error) {
	repository := &Repository{
		Name:        name,
		Description: description,
		Private:     private,
		Users:       users,
	}

	err := repository.validate()
	if err != nil {
		return nil, err
	}

	return repository, nil
}

func (r *Repository) validate() error {
	err := validRepositoryName(r.Name)
	if err != nil {
		return err
	}
	if len(r.Description) > 350 {
		return errors.New("description cannot be more than 350 characters")
	}
	err = validPermissionCategory(r.Users)
	if err != nil {
		return err
	}
	return nil
}

func validPermissionCategory(users UserRepository) error {
	if users == nil {
		return nil
	}
	result := "invalid user category permission:"
	for user, permission := range users {
		switch permission {
		case Owner, Collaborator, Maintainer, Reader:
			continue
		default:
			result += fmt.Sprintf("\nUser: %s with permission: %s", user, permission)
		}
	}
	if result != "invalid user category permission:" {
		return errors.New(result)
	}
	return nil
}

func validRepositoryName(n string) error {
	if n == "" {
		return errors.New("n repository is required")
	}
	regexPattern := "^[a-zA-Z0-9._-]+$"
	ok, err := regexp.MatchString(regexPattern, n)
	if err != nil || !ok {
		return errors.New("the repository name can only contain ASCII letters, digits, and the characters ., -, and _")
	}
	return nil
}
