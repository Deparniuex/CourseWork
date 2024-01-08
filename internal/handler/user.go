package handler

import (
	"errors"
	"net/http"
	"ripProject/internal/entity"
	"ripProject/internal/handler/api"
	"ripProject/internal/repository"
	"ripProject/pkg/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var req api.CreateUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = h.Services.CreateUser(ctx, &entity.User{
		Username:  &req.Username,
		Email:     &req.Email,
		FirstName: &req.FirstName,
		LastName:  &req.LastName,
		Phone:     &req.Phone,
		Password: entity.Password{
			Plaintext: &req.Password,
		},
		Role: entity.USER,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.DefaultResponse{
		Code:    http.StatusCreated,
		Message: "user succesfully created",
	})
}

func (h *Handler) login(ctx *gin.Context) {
	var req api.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	token, err := h.Services.Login(ctx, req.Credentials, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, util.ErrMismatchedPassword) || errors.Is(err, repository.ErrRecordNotFound):
			ctx.JSON(http.StatusUnauthorized, &api.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "user does not exists or password does not match",
			})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusCreated, &api.LoginResponse{
		Code:    http.StatusCreated,
		Message: "token succesfully created",
		Body:    token,
	})
}
