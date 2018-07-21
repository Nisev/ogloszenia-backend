package users

import "github.com/gin-gonic/gin"

func RouterRegister(router *gin.RouterGroup) {
	router.GET("/login", login)
}

func login(c *gin.Context) {
	c.String(200, "Hello")
}
