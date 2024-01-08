package handler

import (
	"fmt"
	"net/http"
	"ripProject/internal/entity"
	"ripProject/internal/handler/api"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createRestaurant(ctx *gin.Context) {
	var req api.CreateRestaurantRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	restaurant := &entity.Restaurant{
		Name:   &req.Name,
		City:   req.City,
		Street: req.Street,
	}
	err = h.Services.CreateRestaurant(ctx, restaurant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.Header("Locations", fmt.Sprintf("/restaurants/%d", restaurant.ID))

	ctx.JSON(http.StatusCreated, &api.DefaultResponse{
		Code:    http.StatusCreated,
		Message: "restaurant succesfully created",
	})
}
