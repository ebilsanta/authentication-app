package usecase

import (
	"log"

	"github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/repository"
)

type AuthenticationUsecase interface {
	CheckUserEnrolled(company string, email string, first_name string, last_name string, birthdate string) (string, error)
	Register(email string, password string) (string, error)
}

type authenticationUsecase struct {
	authenticationRepos repository.AuthenticationRepository
}

func (a *authenticationUsecase) CheckUserEnrolled(company string, email string, first_name string, last_name string, birthdate string) (string, error) {
	user, err := a.authenticationRepos.GetByFullInfo(company, email, first_name, last_name, birthdate)
	if err != nil {
		log.Fatal(err)
	}
	if user == nil {
		return "Sorry you are not enrolled in our database, please make sure that you have an account with one of our partners!", nil
	} 

	if user.Status != "pending" {
		return "You are already a registered user, please proceed to login instead!", nil
	}

	return "You are enrolled in in our database, please proceed to register an account with us!", nil
}

func (a *authenticationUsecase) Register(email string, password string) (string, error) {
	return "Successfully Registered!", nil
}

func NewAuthenticationUsecase(a repository.AuthenticationRepository) AuthenticationUsecase {
	return &authenticationUsecase{a}
}