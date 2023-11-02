package usecase

import (
	"context"
	// "crypto/rsa"
	"crypto/x509"
    "encoding/pem"
	"log"
	"time"
	// "fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/repository"
	"github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/utils"
	otp_server "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/external/proto"
)


func HashPassword(password string, cost int) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
    return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func GenerateIdToken(details map[string]interface{}) (string, error) {

	key_string := utils.GetPrivateKey()
	parsed_key_string := strings.ReplaceAll(key_string, "\\n", "\n")
	block, _ := pem.Decode([]byte(parsed_key_string))
    if block == nil {
        log.Println("failed to parse PEM block containing the public key")
		return "", nil
    }

    key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        log.Println("failed to parse DER encoded public key: " + err.Error())
		return "", err
    }
	
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, 
	jwt.MapClaims{ 
		"iss": details["iss"], 
		"sub": details["sub"], 
		"aud": details["aud"],
		"exp": details["exp"],
		"iat": details["iat"],
		"company": details["company"],
		"email": details["email"],
	})
	s, err := t.SignedString(key) 

	return s, err
}

type AuthenticationUsecase interface {
	Register(company string, email string, first_name string, last_name string, birthdate string, password string) (string, string, error)
	VerifyEmail(verification_key string, otp string, email string) (string, string, string, error)
	Login(company string, email string, password string) (string, string, error)
}

type authenticationUsecase struct {
	authenticationRepos repository.AuthenticationRepository
}

func (a *authenticationUsecase) Register(company string, email string, first_name string, last_name string, birthdate string, password string) (string, string, error) {
	user, err := a.authenticationRepos.GetUserByFullInfo(company, email, first_name, last_name, birthdate)
	if err != nil {
		return "Error", "", err
	}
	if user == nil {
		return "Sorry you are not enrolled in our database, please make sure that you have an account with one of our partners!", "", nil
	} 

	if user.Status != "pending" {
		return "You are already a registered user, please proceed to login instead!", "", nil
	}

	credential, err := a.authenticationRepos.GetCredentialByEmail(company, email)
	if err != nil {
		return "Error", "", err
	}
	if credential != nil {
		return "You are already a registered user, please proceed to verify your email!", "", nil
	}

	cost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))

	hashed_password, _ := HashPassword(password, cost)

	credential, err = a.authenticationRepos.RegisterUser(company, email, hashed_password)
	if err != nil {
		return "Error", "", err
	}

	conn, err := utils.ConnectOTPServer()
	defer conn.Close()
	if err != nil {
		return "Could not Generate OTP", "", err
	}

	client := otp_server.NewOTPClient(conn)

	response, err := client.GetOTP(context.Background(), &otp_server.OTPRequest{Company: company, Email: email})
	if err != nil {
		log.Fatalf("Error when calling GetOTP: %s", err)
	}

	return "Successfully Registered!", response.VerificationKey, nil
}

func (a *authenticationUsecase) VerifyEmail(verification_key string, otp string, email string) (string, string, string, error) {
	conn, err := utils.ConnectOTPServer()
	defer conn.Close()
	if err != nil {
		return "Could not Generate OTP", "", "", err
	}

	client := otp_server.NewOTPClient(conn)

	response, err := client.VerifyOTP(context.Background(), &otp_server.VerifyOTPRequest{VerificationKey: verification_key, Otp: otp, Email: email})
	if err != nil {
		log.Fatalf("Error when calling VerifyOTP: %s", err)
	}
	if response.Status == "Failure" {
		return response.Status, response.Details, email, err
	}
	_, err = a.authenticationRepos.UpdateUserByEmail(response.Company, response.Email)

	if err != nil {
		return "Failure", "Error updating verification status of user", email, err
	}
	return "Success", "Your email is verified, please proceed to login!", email, nil
}

func (a *authenticationUsecase) Login(company string, email string, password string) (string, string, error) {
	user, err := a.authenticationRepos.GetUserByEmail(company, email)
	if err != nil {
		return "Error", "", err
	}
	if user == nil {
		return "You are not an enrolled user!", "", nil
	}

	credential, err := a.authenticationRepos.GetCredentialByEmail(company, email)
	if err != nil {
		return "Error", "", err
	}
	if credential == nil {
		return "You are not a registered user, please proceed to register!", "", nil
	}

	password_match := CheckPasswordHash(password, credential.Password)

	if !password_match {
		return "Password did not match", "", nil
	}

	expiration_delay_minutes := 20

	key_details := map[string]interface{}{
		"iss": "authn.itsag2t1.com",
		"sub": "auth0|123456",
		"aud": "authz.itsag2t1.com",
		"exp": time.Now().Local().Add(time.Minute * time.Duration(expiration_delay_minutes)).Unix(),
		"iat": time.Now().Local().Unix(),
		"company": credential.Company,
		"email": credential.Email,
	}

	id_token, err := GenerateIdToken(key_details)
	if err != nil {
		return "Error generating ID Token", "", err
	}
	return "User verified", id_token, nil
}

func NewAuthenticationUsecase(a repository.AuthenticationRepository) AuthenticationUsecase {
	return &authenticationUsecase{a}
}