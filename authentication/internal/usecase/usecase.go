package usecase

import (
	// "log"
	// "fmt"
	"os"
	"golang.org/x/crypto/bcrypt"
	"strconv"

	"github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/repository"
)


func HashPassword(password string, cost int) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
    return string(bytes), err
}

type AuthenticationUsecase interface {
	Register(company string, email string, first_name string, last_name string, birthdate string, password string) (string, error)
}

type authenticationUsecase struct {
	authenticationRepos repository.AuthenticationRepository
}

func (a *authenticationUsecase) Register(company string, email string, first_name string, last_name string, birthdate string, password string) (string, error) {
	user, err := a.authenticationRepos.GetUserByFullInfo(company, email, first_name, last_name, birthdate)
	if err != nil {
		return "Error", err
	}
	if user == nil {
		return "Sorry you are not enrolled in our database, please make sure that you have an account with one of our partners!", nil
	} 

	if user.Status != "pending" {
		return "You are already a registered user, please proceed to login instead!", nil
	}

	credential, err := a.authenticationRepos.GetCredentialByEmail(company, email)
	if err != nil {
		return "Error", err
	}
	if credential != nil {
		return "You are already a registered user, please proceed to verify your email!", nil
	}

	cost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))

	hashed_password, _ := HashPassword(password, cost)

	credential, err = a.authenticationRepos.RegisterUser(company, email, hashed_password)
	if err != nil {
		return "Error", err
	}

	return "Successfully Registered!", nil
}

func NewAuthenticationUsecase(a repository.AuthenticationRepository) AuthenticationUsecase {
	return &authenticationUsecase{a}
}