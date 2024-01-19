package git

type Repository struct {
	Name        string
	Description string
	Private     bool
	Users       []User
}
