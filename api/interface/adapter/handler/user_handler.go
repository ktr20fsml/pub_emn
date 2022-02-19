package handler

import (
	domainUser "api/domain/model/user"
	"api/status"
	"api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

type UserHandler interface {
	GetUsers(ctx *gin.Context)
}

func NewUserHandler(us usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: us,
	}
}

func (uh *userHandler) GetUsers(ctx *gin.Context) {
	users, err := uh.userUsecase.FindAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domainUser.ResponseUsers(users))
}
