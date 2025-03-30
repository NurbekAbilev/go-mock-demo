package repo_test

import (
	"testing"

	"github.com/nurbekabilev/go-mock-demo/internal/app"
	mockrepo "github.com/nurbekabilev/go-mock-demo/mocks/repo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(mockrepo.MockUserRepository)

	// repo.User{Email: "asdf", Name: "alskdfjald"}
	// mockRepo.On("CreateUser").Return(nil)
	mockRepo.On("CreateUser", mock.Anything).Return(nil)

	err := app.CreateUser(mockRepo)
	require.NoError(t, err)

	mockRepo.AssertCalled(t, "CreateUser", mock.Anything)
}
