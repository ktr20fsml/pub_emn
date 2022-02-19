package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserID(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	id, ok := session.Get("UserID").(string)
	if ok != true {
		return ""
	}
	return id
}

func GetUserName(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	name, ok := session.Get("UserName").(string)
	if ok != true {
		return ""
	}
	return name
}

func GetAdministrative(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	_, ok := session.Get("Administrative").(bool)
	if ok != true {
		return false
	}
	return true
}
