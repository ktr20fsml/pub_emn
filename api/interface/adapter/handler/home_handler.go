package handler

import (
	domainUser "api/domain/model/user"
	"api/status"
	"api/usecase"
	"fmt"
	"net/http"

	"api/infrastructure/library/crypt"
	"api/infrastructure/library/session"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type homeHandler struct {
	homeUsecase usecase.HomeUsecase
}

type HomeHandler interface {
	Home(ctx *gin.Context)
	Signup(ctx *gin.Context)
	Signin(ctx *gin.Context)
	Signout(ctx *gin.Context)
}

func NewHomeHandler(hs usecase.HomeUsecase) HomeHandler {
	return &homeHandler{
		homeUsecase: hs,
	}
}

func (hh *homeHandler) Home(ctx *gin.Context) {
	user := &domainUser.User{
		ID:             domainUser.UserID(session.GetUserID(ctx)),
		Name:           session.GetUserName(ctx),
		Administrative: session.GetAdministrative(ctx),
	}

	ctx.JSON(http.StatusOK, user)
}

func (hh *homeHandler) Signup(ctx *gin.Context) {
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

	user.CreatedBy = domainUser.UserID(session.GetUserID(ctx))

	hashedPassword, errEncrypt := crypt.Encrypt(user.Password)
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
	_, errEncrypt := crypt.Encrypt(plainPassword)
	if errEncrypt != nil {
		fmt.Println(errEncrypt)
	}

	resUser, err := hh.homeUsecase.FindUserByID(&user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	auth, errVerify := crypt.Verify(resUser.Password, plainPassword)
	if errVerify != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errVerify.Error()})
		return
	}
	if auth != true {
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
