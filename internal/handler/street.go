package handler

import (
	"fmt"
	"net/http"
	"ripProject/internal/entity"
	"ripProject/internal/handler/api"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createStreet(ctx *gin.Context) {
	var req api.CreateStreetRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	street := &entity.Street{
		Name: &req.Name,
	}

	err = h.Services.CreateStreet(ctx, street)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.Header("Locations", fmt.Sprintf("/streets/%d", street.ID))

	ctx.JSON(http.StatusCreated, &api.DefaultResponse{
		Code:    http.StatusCreated,
		Message: "street succesfully created",
	})
}
