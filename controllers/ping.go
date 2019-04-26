package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.String(http.StatusOK, "This is an API developed in Golang. For more info, check the language manual: * https://golang.org *")
}

func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
