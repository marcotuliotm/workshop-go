package http

import (
	"net/http"
	"workshop-go/domain/user"

	"github.com/gin-gonic/gin"
)

type handler struct {
	userService *user.Service
}

func NewHandler(userService *user.Service) http.Handler {
	handler := &handler{
		userService: userService,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(handler.recovery())

	v1 := router.Group("/v1")
	v1.POST("/users", handler.postUser)

	return router
}

// defer é um finally

func (h *handler) recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
