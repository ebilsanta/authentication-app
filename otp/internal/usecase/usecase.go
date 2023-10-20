package usecase

import (
	"io"
	"crypto/rand"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
	"time"
	// "strconv"
	// "log"
	// "fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"

	"github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/repository"
	"github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/utils"
)


var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateOTP(max int) string {
    b := make([]byte, max)
    n, err := io.ReadAtLeast(rand.Reader, b, max)
    if n != max {
        panic(err)
    }
    for i := 0; i < len(b); i++ {
        b[i] = table[int(b[i])%len(table)]
    }
	otp := string(b)
    return otp
}

func Encode(details map[string]string) (string, error) {
	crypt_password := os.Getenv("CRYPT_PASSWORD")
	secret_string := details["company"] + "\n" + details["expiration_date"] + "\n" + details["email"] + "\n" + details["otp"]
	block, err := aes.NewCipher([]byte(crypt_password))
	if err != nil {
		return "", err
	}
	bytes := []byte(os.Getenv("IV"))
	plainText := []byte(secret_string)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decode(verification_key string) (map[string]string, error) {
	crypt_password := os.Getenv("CRYPT_PASSWORD")
	block, err := aes.NewCipher([]byte(crypt_password))
	if err != nil {
		return nil, err
	}
	cipherText, err := base64.StdEncoding.DecodeString(verification_key)
	if err != nil {
		return nil, err
	}
	bytes := []byte(os.Getenv("IV"))
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	details := strings.Split(string(plainText), "\n")
	return map[string]string{"company": details[0], "expiration_date": details[1], "email": details[2], "otp": details[3]}, nil
}

func SendOTPEmail(email string, otp string) error {
	delete_otp_link := "https://google.com/"
	Sender := os.Getenv("AWS_SENDER_EMAIL")
	Subject := "Your OTP for Email Verification for Ascenda"
    HtmlBody :=  "<h2>Ascenda OTP Email</h2><p>Your OTP is: " + otp +
                "\n\nIf you did not request for an OTP please click <a href='" + delete_otp_link + "'>this link</a>.</p>"
    TextBody := "Your OTP is: " + otp + "\n\nIf you did not request for an OTP please click this link: " + delete_otp_link
    CharSet := "UTF-8"

	svc := utils.GetSES()

	input := &ses.SendEmailInput{
        Destination: &ses.Destination{
            CcAddresses: []*string{
            },
            ToAddresses: []*string{
                aws.String(email),
            },
        },
        Message: &ses.Message{
            Body: &ses.Body{
                Html: &ses.Content{
                    Charset: aws.String(CharSet),
                    Data:    aws.String(HtmlBody),
                },
                Text: &ses.Content{
                    Charset: aws.String(CharSet),
                    Data:    aws.String(TextBody),
                },
            },
            Subject: &ses.Content{
                Charset: aws.String(CharSet),
                Data:    aws.String(Subject),
            },
        },
        Source: aws.String(Sender),
    }

	_, err := svc.SendEmail(input)

	return err
}

type OTPUsecase interface {
	GetOTP(company string, email string) (string, string, error) 
	VerifyOTP(verificationKey string, otp string, email string) (string, string, string, string, error)
}

type otpUsecase struct {
	otpRepos repository.OTPRepository
}

func (o *otpUsecase) GetOTP(company string, email string) (string, string, error) {
	otp := GenerateOTP(6)
	expiration_delay_minutes := 2
	expiration_date := time.Now().Local().Add(time.Minute * time.Duration(expiration_delay_minutes))
	
	details := map[string]string{"company": company, "expiration_date": expiration_date.String(), "email": email, "otp": string(otp)}

	_, err := o.otpRepos.CreateOTP(otp, expiration_date.String())
	if err != nil {
		return "", "Error creating key", err
	}

	err = SendOTPEmail(email, otp)

	if err != nil {
		return "", "Error Sending OTP", err
	}

	verification_key, err := Encode(details)

	if err != nil {
		return "", "Failure", err
	}

	return verification_key, "Success", nil
}

func (o *otpUsecase) VerifyOTP(verification_key string, otp string, email string) (string, string, string, string, error) {
	details, err := Decode(verification_key)
	if err != nil {
		return "Failure", "Error processing verification key", "", email, err
	}
	now := time.Now()
	expiration_date, _ := time.Parse(now.String(), details["expiration_date"])
	if details["otp"] == otp && details["email"] == email && expiration_date.Before(now) {
		OTP, err := o.otpRepos.GetOTP(otp, details["expiration_date"])
		if err != nil || OTP == nil {
			return "Failure", "Invalid OTP", "", email, err
		}
		if OTP.Verified {
			return "Failure", "OTP has already been used", "", email, nil
		}
		_, err = o.otpRepos.UpdateOTP(otp, details["expiration_date"])
		if err != nil {
			return "Failure", "Something went wrong while verifying the OTP", "", email, err
		}
		return "Success", "OTP Verified", details["company"], email, nil
	}
	return "Failure", "Some details did not match", "", email, nil
}

func NewOTPUsecase(o repository.OTPRepository) OTPUsecase {
	return &otpUsecase{o}
}