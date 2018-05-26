package user

import (
	"testing"
	"workshop-go/domain"

	"github.com/stretchr/testify/assert"
)

func TestIsAValidUser(t *testing.T) {
	user := &domain.User{Name: "Tulio", Age: 33}

	userService := NewService()

	assert.True(t, userService.IsValid(user))

}

func TestIsntAValidUser(t *testing.T) {
	user := &domain.User{Name: "Tulio", Age: 17}

	userService := NewService()

	assert.False(t, userService.IsValid(user))

}
