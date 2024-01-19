package entity

type User struct {
	Username    string
	Email       string
	Name        string
	Permissions []AccessRepository
}

type AccessRepository struct {
	Repository string
	Permission string
}
