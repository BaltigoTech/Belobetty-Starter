package git

type User struct {
	Username     string
	Email        string
	Name         string
	Repositories []AccessRepository
}

type AccessRepository struct {
	Repository string
	Permission string
}
