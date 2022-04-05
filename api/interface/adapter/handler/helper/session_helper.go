package session_helper

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserID(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	id, ok := session.Get("UserID").(string)
	if !ok {
		return ""
	}
	return id
}

func GetUserName(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	name, ok := session.Get("UserName").(string)
	if !ok {
		return ""
	}
	return name
}

func GetAdministrative(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	admin, ok := session.Get("Administrative").(bool)
	if !ok {
		return false
	}
	return admin
}
