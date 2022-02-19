package middleware

import (
	"api/infrastructure/library/session"
	"api/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionInfo struct {
	UserID         string
	UserName       string
	Administrative bool
}

var sessionInfo SessionInfo

var (
	errInvalidSession              string = "Invalid Session. Please sign in."
	errInvalidSessionAdministrator string = "Invalid Session. Please sign in as a user with administrative permission."
)

func SessionCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionInfo.UserID = session.GetUserID(ctx)
		if sessionInfo.UserID == "" {
			ctx.JSON(http.StatusBadRequest, status.Status{Message: errInvalidSession})
			ctx.Abort()
		} else {
			ctx.Set("UserName", sessionInfo.UserName)
			ctx.Set("UserID", session.GetUserName(ctx))
			ctx.Set("Administrative", session.GetAdministrative(ctx))
			ctx.Next()
		}
	}
}

func SessionAdministratorCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionInfo.Administrative = session.GetAdministrative(ctx)
		if sessionInfo.Administrative == false {
			ctx.JSON(http.StatusBadRequest, status.Status{Message: errInvalidSession})
			ctx.Abort()
		} else {
			if sessionInfo.Administrative == false {
				ctx.JSON(http.StatusBadRequest, status.Status{Message: errInvalidSessionAdministrator})
				ctx.Abort()
			} else {
				ctx.Set("UserName", session.GetUserID(ctx))
				ctx.Set("UserID", session.GetUserName(ctx))
				ctx.Set("Administrative", session.GetAdministrative(ctx))
				ctx.Next()
			}
		}
	}
}
