package repo

type User struct {
	Email string
	Name  string
}

type UserRepository interface {
	CreateUser(user User) error
}
