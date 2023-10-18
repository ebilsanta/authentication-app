package internal

import "time"

type OTP struct {
    OTP    			int         `dynamodbav:"otp"`
    ExpirationDate  time.Time   `dynamodbav:"expiration_date"`
	Verified		bool	    `dynamodbav:"verified"`
}