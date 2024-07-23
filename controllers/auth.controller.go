package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAccount(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}
