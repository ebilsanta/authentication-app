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

func (srv *AuthenticationServer) Register(c context.Context, user *authentication.RegisterRequest) (*authentication.RegisterResponse, error) {
	message, error := srv.useCase.Register(user.Company, user.Email, user.FirstName, user.LastName, user.Birthdate, user.Password)
	return &authentication.RegisterResponse{Email: user.Email, Message: message}, error
}