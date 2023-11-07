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
	message, verification_key, error := srv.useCase.Register(user.Company, user.Email, user.FirstName, user.LastName, user.Birthdate, user.Password)
	return &authentication.RegisterResponse{Email: user.Email, Message: message, VerificationKey: verification_key}, error
}

func (srv *AuthenticationServer) VerifyEmail(c context.Context, verification *authentication.VerifyEmailRequest) (*authentication.VerifyEmailResponse, error) {
	status, details, _, error := srv.useCase.VerifyEmail(verification.VerificationKey, verification.Otp, verification.Email)
	return &authentication.VerifyEmailResponse{Status: status, Details: details, Email: verification.Email}, error
}

func (srv *AuthenticationServer) Login(c context.Context, credentials *authentication.LoginRequest) (*authentication.LoginResponse, error) {
	status, idToken, error := srv.useCase.Login(credentials.Company, credentials.Email, credentials.Password)
	return &authentication.LoginResponse{Status: status, IdToken: idToken}, error
}

func (srv *AuthenticationServer) NewOTP(c context.Context, userInfo *authentication.NewOTPRequest) (*authentication.NewOTPResponse, error) {
	message, verification_key, error := srv.useCase.NewOTP(userInfo.Company, userInfo.Email)
	return &authentication.NewOTPResponse{Message: message, VerificationKey: verification_key}, error
}

func (srv *AuthenticationServer) ChangePassword(c context.Context, verification *authentication.ChangePasswordRequest) (*authentication.ChangePasswordResponse, error) {
	status, details, email, error := srv.useCase.ChangePassword(verification.VerificationKey, verification.Otp, verification.Company, verification.Email, verification.Password)
	return &authentication.ChangePasswordResponse{Status: status, Details: details, Email: email}, error
}

