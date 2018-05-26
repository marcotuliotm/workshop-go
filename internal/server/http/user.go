package http

import (
	"fmt"
	"net/http"
	"workshop-go/domain"

	"github.com/gin-gonic/gin"
)

func (h *handler) postUser(c *gin.Context) {

	var user *domain.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	fmt.Printf("User: %s\n", user)

	if !h.userService.IsValid(user) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusCreated, user)
}
