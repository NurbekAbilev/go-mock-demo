package app

import (
	"fmt"

	"github.com/nurbekabilev/go-mock-demo/internal/repo"
)

func CreateUser(userRepository repo.UserRepository) error {
	user := repo.User{
		Email: "abilev.nurbek@gmail.com",
		Name:  "Nurbek",
	}

	err := userRepository.CreateUser(user)
	if err != nil {
		return fmt.Errorf("app error createUser %w", err)
	}

	return nil
}
