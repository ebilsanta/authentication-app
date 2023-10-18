package repository

import (
	"time"

	models "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal"
)

type OTPRepository interface {
	CreateOTP(otp string, expiration_date time.Time) (*models.OTP, error)
	GetOTP(otp string, expiration_date time.Time) (*models.OTP, error)
	UpdateOTP(otp string, expiration_date time.Time) (*models.OTP, error)
}