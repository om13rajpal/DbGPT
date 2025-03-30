package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/om13rajpal/dbgpt/internal/database"
	"github.com/om13rajpal/dbgpt/internal/models"
	"github.com/om13rajpal/dbgpt/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func LoginHandler(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var fetchedUser models.User
	result := database.UserCollection.FindOne(ctx, bson.M{
		"username": user.Username,
	}).Decode(&fetchedUser)

	if result != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "User not found",
		})
		return
	}

	isValidPassword := utils.ComparePassword(fetchedUser.Password, user.Password)
	if !isValidPassword {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid credentials",
		})
		return
	}

	token, err := utils.GenerateToken()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Internal server error",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Login successful",
		"token":   token,
		"user":    fetchedUser})
}
