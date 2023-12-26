package middleware

import (
	session_helper "api/interface/adapter/handler/helper"
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
		sessionInfo.UserID = session_helper.GetUserID(ctx)
		if sessionInfo.UserID == "" {
			ctx.JSON(http.StatusBadRequest, status.Status{Message: errInvalidSession})
			ctx.Abort()
		} else {
			ctx.Set("UserName", sessionInfo.UserName)
			ctx.Set("UserID", session_helper.GetUserName(ctx))
			ctx.Set("Administrative", session_helper.GetAdministrative(ctx))
			ctx.Next()
		}
	}
}

func SessionAdministratorCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionInfo.UserID = session_helper.GetUserID(ctx)
		sessionInfo.Administrative = session_helper.GetAdministrative(ctx)
		if sessionInfo.UserID == "" {
			ctx.JSON(http.StatusBadRequest, status.Status{Message: errInvalidSession})
			ctx.Abort()
		} else {
			if !sessionInfo.Administrative {
				ctx.JSON(http.StatusBadRequest, status.Status{Message: errInvalidSessionAdministrator})
				ctx.Abort()
			} else {
				ctx.Set("UserName", session_helper.GetUserID(ctx))
				ctx.Set("UserID", session_helper.GetUserName(ctx))
				ctx.Set("Administrative", session_helper.GetAdministrative(ctx))
				ctx.Next()
			}
		}
	}
}
