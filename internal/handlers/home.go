package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Backend for DB GPT is up and running",
	})
}