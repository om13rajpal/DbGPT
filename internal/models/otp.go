package models

import "time"

type Otp struct {
	Otp       string `json:"otp"`
	Username  string `json:"username"`
	ExpiresAt time.Time `json:"expires_at"`
}