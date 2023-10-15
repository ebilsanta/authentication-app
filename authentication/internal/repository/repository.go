package repository

import models "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal"

type AuthenticationRepository interface {
	GetByFullInfo(company string, email string, first_name string, last_name string, birthdate string) (*models.User, error)
	UpdateByEmail(email string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	DeleteByEmail(email string) (string, error)
}