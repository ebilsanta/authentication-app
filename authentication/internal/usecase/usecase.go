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

func GenerateValidToken(details map[string]interface{}) (string, error) {

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
		"exp": details["exp"],
		"iat": details["iat"],
	})
	s, err := t.SignedString(key) 

	return s, err
}

func CheckValidToken(token_string string, email string) bool {
	key_string := utils.GetPublicKey()
	parsed_key_string := strings.ReplaceAll(key_string, "\\n", "\n")
	block, _ := pem.Decode([]byte(parsed_key_string))
    if block == nil {
        log.Println("failed to parse PEM block containing the public key")
		return false
    }

    key, err := x509.ParsePKCS1PublicKey(block.Bytes)
    if err != nil {
        log.Println("failed to parse DER encoded public key: " + err.Error())
		return false
    }

	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Println("unexpected signing method: %v", token.Header["alg"])
			return nil, err
		}
		return key, nil
	})

	if err != nil {
		log.Println("Error occured parsing token: ", err)
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Println("Invalid Token")
		return false
	}

	if claims["sub"] == nil || claims["exp"] == nil {
		log.Println("Invalid Token Content")
	}

	now := time.Now()

	parsed_time, err := time.Parse("1696969690", claims["exp"].(string))

	if err != nil {
		log.Println("Token Expiry is not a valid datetime")
		return false
	}

	if claims["sub"] != email || now.Before(parsed_time) {
		log.Println("Wrong email or Expired Token")
		return false
	}

	return true
}

type AuthenticationUsecase interface {
	Register(company string, email string, first_name string, last_name string, birthdate string, password string) (string, string, error)
	VerifyEmail(verification_key string, otp string, email string) (string, string, string, error)
	Login(company string, email string, password string) (string, string, error)
	NewOTP(company string, email string) (string, string, error)
	ValidToken(verification_key string, otp string, email string) (string, string, error)
	ChangePassword(valid_token string, company string, email string, password string) (string, string, string, error)
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
		return "Failure", "Could not Verify OTP", email, err
	}

	client := otp_server.NewOTPClient(conn)

	response, err := client.VerifyOTP(context.Background(), &otp_server.VerifyOTPRequest{VerificationKey: verification_key, Otp: otp, Email: email})
	if err != nil {
		log.Fatalf("Error when calling VerifyOTP: %s", err)
		return "Failure", "Error verifying OTP", email, err
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
		"sub": credential.Email,
		"aud": "authz.itsag2t1.com",
		"exp": time.Now().Local().Add(time.Minute * time.Duration(expiration_delay_minutes)).Unix(),
		"iat": time.Now().Local().Unix(),
		"company": credential.Company,
	}

	id_token, err := GenerateIdToken(key_details)
	if err != nil || len(id_token) == 0{
		return "Error generating ID Token", "", err
	}
	return "User verified", id_token, nil
}

func (a *authenticationUsecase) NewOTP(company string, email string) (string, string, error) {
	credential, err := a.authenticationRepos.GetCredentialByEmail(company, email)
	if err != nil {
		log.Println("Error getting user: ", err)
		return "Error getting user!", "", err
	} else if credential == nil {
		log.Println("User not registered")
		return "You are not registered with us!", "", nil
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
		return "Error sending OTP!", "", err
	}

	return "OTP Sent!", response.VerificationKey, nil
}

func (a *authenticationUsecase) ValidToken(verification_key string, otp string, email string) (string, string, error) {
	conn, err := utils.ConnectOTPServer()
	defer conn.Close()
	if err != nil {
		log.Println("Could not connect to OTP Server")
		return "Could not connect to OTP Server", "", nil
	}

	client := otp_server.NewOTPClient(conn)

	response, err := client.VerifyOTP(context.Background(), &otp_server.VerifyOTPRequest{VerificationKey: verification_key, Otp: otp, Email: email})
	if err != nil {
		log.Fatalf("Error when calling VerifyOTP: %s", err)
		return "Could not verify OTP", "", nil
	}
	if response.Status == "Failure" {
		return response.Details, "email", nil
	}
	_, err = a.authenticationRepos.UpdateUserByEmail(response.Company, response.Email)

	if err != nil {
		return "Error updating verification status of user", "", err
	}

	expiration_delay_minutes := 5

	key_details := map[string]interface{}{
		"iss": "authn.itsag2t1.com",
		"sub": email,
		"exp": time.Now().Local().Add(time.Minute * time.Duration(expiration_delay_minutes)).Unix(),
		"iat": time.Now().Local().Unix(),
	}

	valid_token, err := GenerateIdToken(key_details)
	if err != nil || len(valid_token) == 0 {
		log.Println("Error generating ID Token: ", err)
		return "Error generating ID Token", "", nil
	}

	return "Success", valid_token, nil
}

func (a *authenticationUsecase) ChangePassword(valid_token string, company string, email string, password string) (string, string, string, error) {
	credential, err := a.authenticationRepos.GetCredentialByEmail(company, email)
	if err != nil {
		log.Println("Error getting user: ", err)
		return "Failure", "Error getting user!", email, err
	} else if credential == nil {
		log.Println("User not registered")
		return "Failure", "You are not registered with us!", email, nil
	}

	isValid := CheckValidToken(valid_token, email)

	if !isValid {
		return "Failure", "Invalid Token", email, nil
	}

	cost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))

	hashed_password, _ := HashPassword(password, cost)

	_, err = a.authenticationRepos.UpdateUserPassword(company, email, hashed_password)

	if err != nil {
		return "Failure", "Error updating password", email, err
	}

	return "Success", "Password updated!", email, nil
}

func NewAuthenticationUsecase(a repository.AuthenticationRepository) AuthenticationUsecase {
	return &authenticationUsecase{a}
}