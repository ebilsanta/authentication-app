package repository

import models "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal"

type AuthenticationRepository interface {
	GetUserByFullInfo(company string, email string, first_name string, last_name string, birthdate string) (*models.User, error)
	UpdateUserByEmail(company string, email string) (*models.User, error)
	GetCredentialByEmail(company string, email string) (*models.Credential, error)
	GetUserByEmail(company string, email string) (*models.User, error)
	DeleteCredentialByEmail(company string, email string) (string, error)
	RegisterUser(company string, email string, password string) (*models.Credential, error)
	UpdateUserPassword(company string, email string, password string) (*models.Credential, error)
}