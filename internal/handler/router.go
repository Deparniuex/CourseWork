package handler

import (
	"ripProject/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	restaurant := router.Group("/restaurants")
	city := router.Group("cities")
	user := router.Group("/users")
	street := router.Group("/streets")

	user.POST("login", h.login)
	user.POST("/register", h.createUser)
	restaurant.POST("/register", h.requireRole(entity.MODERATOR), h.createRestaurant)
	city.POST("/create", h.requireRole(entity.ADMIN), h.createCity)
	street.POST("/create", h.requireRole(entity.ADMIN), h.createStreet)
	return router
}
