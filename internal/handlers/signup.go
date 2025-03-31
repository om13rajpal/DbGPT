package handlers

import (
	"context"
	"math/rand"
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error hashing password",
		})
		return
	}

	user.Password = string(hashedPassword)

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`

	var id int
	err = database.Pool.QueryRow(ctx, query, user.Username, user.Email, user.Password).Scan(&id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error creating user",
			"error":   err.Error(),
		})
		return
	}

	go utils.SendMail(user.Email, "Welcome to DBGPT", "Thank you for signing up!")

	query = `INSERT INTO otp (username, otp) VALUES ($1, $2) RETURNING id`
	var otpId int

	otp := rand.Intn(900000) + 100000

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = database.Pool.QueryRow(ctx, query, user.Username, otp).Scan(&otpId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error creating OTP",
			"error":   err.Error(),
		})
		return
	}

	go utils.SendOTP(user.Email, otp)

	token, err := utils.GenerateToken()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error generating token",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User created successfully",
		"token":   token,
		"id":      id,
	})

}
