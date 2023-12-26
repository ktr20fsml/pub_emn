package handler

import (
	domainUser "api/domain/model/user"
	"api/domain/service"
	"api/status"
	"api/usecase"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type homeHandler struct {
	homeUsecase usecase.HomeUsecase
	utility     service.UtilityService
}

type HomeHandler interface {
	Home(ctx *gin.Context)
	Signup(ctx *gin.Context)
	Signin(ctx *gin.Context)
	Signout(ctx *gin.Context)
}

func NewHomeHandler(hs usecase.HomeUsecase, util service.UtilityService) HomeHandler {
	return &homeHandler{
		homeUsecase: hs,
		utility:     util,
	}
}

func (hh *homeHandler) Home(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	id, ok := sess.Get("UserID").(string)
	if !ok {
		id = "Unknown ID"
	}
	name, ok := sess.Get("UserName").(string)
	if !ok {
		name = "Unknown Name"
	}
	admin, ok := sess.Get("Administrative").(bool)
	if !ok {
		admin = false
	}

	user := &domainUser.User{
		ID:             domainUser.UserID(id),
		Name:           name,
		Administrative: admin,
	}

	ctx.JSON(http.StatusOK, user)
}

func (hh *homeHandler) Signup(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	id, ok := sess.Get("UserID").(string)
	if !ok {
		id = "Unknown ID"
	}

	user := &domainUser.User{}
	errBind := ctx.BindJSON(user)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
		return
	}
	errValid := domainUser.Validate(user)
	if errValid.Message != nil {
		ctx.JSON(http.StatusBadRequest, status.Statuses{Message: errValid.Message})
		return
	}

	user.CreatedBy = domainUser.UserID(id)

	hashedPassword, errEncrypt := hh.utility.Encrypt(user.Password)
	if errEncrypt != nil {
		fmt.Println(errEncrypt)
	}
	user.Password = hashedPassword

	err := hh.homeUsecase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, status.Status{Message: "SUCCEEDED SIGN UP."})
}

func (hh *homeHandler) Signin(ctx *gin.Context) {
	user := &domainUser.User{}

	errBind := ctx.BindJSON(user)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
		return
	}

	plainPassword := user.Password
	_, errEncrypt := hh.utility.Encrypt(plainPassword)
	if errEncrypt != nil {
		fmt.Println(errEncrypt)
	}

	resUser, err := hh.homeUsecase.FindUserByID(user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	auth, errVerify := hh.utility.Verify(resUser.Password, plainPassword)
	if errVerify != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errVerify.Error()})
		return
	}
	if !auth {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: "AUTHENTICATION FAILED ERROR"})
		return
	}

	sess := sessions.Default(ctx)
	sess.Set("UserID", string(resUser.ID))
	sess.Set("UserName", resUser.Name)
	sess.Set("Administrative", resUser.Administrative)
	errSession := sess.Save()
	if errSession != nil {
		fmt.Println("SAVING SESSION ERROR: ", fmt.Errorf(errSession.Error()))
	}

	ctx.JSON(http.StatusOK, domainUser.ResponseUser(resUser))
}

func (hh *homeHandler) Signout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()

	ctx.JSON(http.StatusOK, status.Status{Message: "SUCCEEDED SIGN OUT."})
}
