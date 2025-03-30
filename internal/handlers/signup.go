package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/om13rajpal/dbgpt/internal/database"
	"github.com/om13rajpal/dbgpt/internal/models"
	"github.com/om13rajpal/dbgpt/utils"
)

func SignupHandler(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": false, "error": "Invalid request"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": false, "error": "Error hashing password"})
		return
	}

	user.Password = string(hashedPassword)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": false, "error": "Error creating user"})
		return
	}

	err = utils.SendMail(user.Email, "Welcome to DBGPT", "Thank you for signing up!")

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": false, "error": "Error sending email"})
		fmt.Println("Error sending email:", err)
		return
	}

	token, err := utils.GenerateToken()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": false, "error": "Error generating token"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"status": true, "message": "user signed up successfully", "data": result, "token": token})
}
