package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/om13rajpal/dbgpt/internal/database"
	"github.com/om13rajpal/dbgpt/internal/models"
	"github.com/om13rajpal/dbgpt/utils"
)

func LoginHandler(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}

	var fetchedUser models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `SELECT username, email, password FROM users WHERE username = $1`
	err = database.Pool.QueryRow(ctx, query, user.Username).Scan(&fetchedUser.Username, &fetchedUser.Email, &fetchedUser.Password)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "User not found",
			})
			return

		}
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "database error",
			"error":   err.Error(),
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
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Login successful",
		"token":   token,
		"user":    fetchedUser})

}
