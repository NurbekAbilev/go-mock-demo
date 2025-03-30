package sqlite

import (
	"database/sql"
	"fmt"

	rep "github.com/nurbekabilev/go-mock-demo/internal/repo"
)

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) *repo {
	return &repo{db: db}
}

func (r *repo) CreateUser(user rep.User) error {
	_, err := r.db.Exec("insert into users(email, name) values ($1, $2)", user.Email, user.Name)
	if err != nil {
		return fmt.Errorf("error create user %w", err)
	}

	return nil
}
