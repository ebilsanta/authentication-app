package grpc

import (
	"context"
	"google.golang.org/grpc"

	"github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/usecase"
	authentication "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/api/proto"
)

type AuthenticationServer struct {
	useCase usecase.AuthenticationUsecase
	authentication.UnimplementedAuthenticationServer
}

func NewAuthenticationServer(grpcServer *grpc.Server, usecase usecase.AuthenticationUsecase){
	userGrpc := &AuthenticationServer{useCase: usecase}
	authentication.RegisterAuthenticationServer(grpcServer, userGrpc)
}

func (srv *AuthenticationServer) CheckUserEnrolled(c context.Context, user *authentication.UserEnrolledRequest) (*authentication.UserEnrolledResponse, error) {
	message, error := srv.useCase.CheckUserEnrolled(user.Company, user.Email, user.FirstName, user.LastName, user.Birthdate)
	return &authentication.UserEnrolledResponse{Message: message}, error
}

func (srv *AuthenticationServer) Register(c context.Context, userCredentials *authentication.RegistrationRequest) (*authentication.RegistrationResponse, error) {
	message, error := srv.useCase.Register(userCredentials.Email, userCredentials.Password)
	return &authentication.RegistrationResponse{Message: message}, error
}