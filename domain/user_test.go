package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	user := &User{Name: "Tulio", Age: 33}
	assert.Equal(t, "Tulio", user.String(), "they should be equal")

}
