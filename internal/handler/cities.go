package handler

import (
	"fmt"
	"net/http"
	"ripProject/internal/entity"
	"ripProject/internal/handler/api"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createCity(ctx *gin.Context) {
	var req api.CreateCityRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	city := &entity.City{
		CityName: req.CityName,
	}

	err = h.Services.CreateCity(ctx, city)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.Header("Locations", fmt.Sprintf("/cities/%d", city.ID))

	ctx.JSON(http.StatusCreated, &api.DefaultResponse{
		Code:    http.StatusCreated,
		Message: "city succesfully created",
	})
}

func (h *Handler) deleteCity(ctx *gin.Context) {

	var id api.ID

	err := ctx.ShouldBindUri(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	err = h.Services.DeleteCity(ctx, id.Value)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.DefaultResponse{
		Code:    http.StatusOK,
		Message: "city succesfully deleted",
	})
}

func (h *Handler) getCities(ctx *gin.Context) {
	cities, err := h.Services.GetCities(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, &api.GetCityResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Body:    cities,
	})
}
