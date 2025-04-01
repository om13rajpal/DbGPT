package utils

import (
	"fmt"
	"time"
)

func CheckOtpTime(expiresAt time.Time) bool {

	fmt.Println("time difference", time.Since(expiresAt))

	return time.Now().Before(expiresAt)
}
