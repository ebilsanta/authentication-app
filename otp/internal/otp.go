package internal

type OTP struct {
    OTP    			string      `dynamodbav:"otp"`
    ExpirationDate  string      `dynamodbav:"expiration_date"`
	Verified		bool	    `dynamodbav:"verified"`
}