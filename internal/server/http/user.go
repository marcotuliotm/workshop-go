package http

import (
	"net/http"

	"workshop-go/domain"

	"github.com/gin-gonic/gin"
)

func (h *handler) postUser(c *gin.Context) {
	user := &domain.User{}

	if err := c.BindJSON(&user); err != nil {
		return
	}

	if !h.userService.IsValid(user) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, user.ID)
}
