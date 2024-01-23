package entity_test

import (
	"Belobetty-Starter/internal/git/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepository(t *testing.T) {
	users := make(entity.UserRepository)
	users["Gabriel"] = "Owner"
	users["Dani"] = "collaborator"
	repo, err := entity.NewRepository("Repo_name", "Description Test", true, users)
	assert.Nil(t, err)
	assert.NotNil(t, repo)
	assert.Equal(t, repo.Name, "Repo_name")
}

func TestCreateRepositoryWithInvalidName(t *testing.T) {
	repo, err := entity.NewRepository("test@", "Description test", true, nil)
	assert.Nil(t, repo)
	assert.Error(t, err, "the repository n can only contain ASCII letters, digits, and the characters ., -, and _")
}

func TestCreateRepositoryWithInvalidDescription(t *testing.T) {
	bigString := "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz" +
		"zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz" +
		"zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz" +
		"zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"

	repo, err := entity.NewRepository("Test", bigString, true, nil)
	assert.True(t, len(bigString) > 350)
	assert.Nil(t, repo)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "description cannot be more than 350 characters")
}

func TestCreateRepositoryWithInvalidUserPermission(t *testing.T) {
	users := map[string]string{
		"UserOk":    "Maintainer",
		"UserFail1": "Master",
		"UserOk2":   "Collaborator",
	}

	repo, err := entity.NewRepository("Repo_name", "Description Test", true, users)
	assert.Nil(t, repo)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid user category permission:\nUser: UserFail1 with permission: Master")
}
