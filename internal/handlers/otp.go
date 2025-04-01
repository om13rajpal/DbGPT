package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/om13rajpal/dbgpt/internal/database"
	"github.com/om13rajpal/dbgpt/internal/models"
	"github.com/om13rajpal/dbgpt/utils"
)

func VerifyOtpHandler(c *gin.Context) {
	var otp models.Otp

	err := c.BindJSON(&otp)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}

	query := `SELECT otp, username, expiresat FROM otp WHERE username = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var fetchedOtp models.Otp
	err = database.Pool.QueryRow(ctx, query, otp.Username).Scan(&fetchedOtp.Otp, &fetchedOtp.Username, &fetchedOtp.ExpiresAt)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error fetching OTP",
			"error":   err.Error(),
		})
		return
	}

	validOtp := utils.CheckOtpTime(fetchedOtp.ExpiresAt)

	if !validOtp {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "OTP expired",
		})
		return
	}

	if otp.Otp != fetchedOtp.Otp {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid OTP",
		})
		return
	}

	query = `UPDATE users SET isverified = true WHERE username = $1`
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = database.Pool.Exec(ctx, query, otp.Username)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error updating user verification status",
			"error":   err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "OTP verified successfully",
	})
}
