package grpc

import (
	"context"
	"google.golang.org/grpc"

	"github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/usecase"
	otp "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/api/proto"
)

type OTPServer struct {
	useCase usecase.OTPUsecase
	otp.UnimplementedOTPServer
}

func NewOTPServer(grpcServer *grpc.Server, usecase usecase.OTPUsecase){
	userGrpc := &OTPServer{useCase: usecase}
	otp.RegisterOTPServer(grpcServer, userGrpc)
}

func (srv *OTPServer) GetOTP(c context.Context, receipient *otp.OTPRequest) (*otp.OTPResponse, error) {
	verificationKey, message, error := srv.useCase.GetOTP(receipient.Email)
	return &otp.OTPResponse{VerificationKey: verificationKey, Message: message}, error
}

func (srv *OTPServer) VerifyOTP(c context.Context, receipient *otp.VerifyOTPRequest) (*otp.VerifyOTPResponse, error) {
	status, details, email, error := srv.useCase.VerifyOTP(receipient.VerificationKey, receipient.Otp, receipient.Email)
	return &otp.VerifyOTPResponse{Status: status, Details: details, Email: email}, error
}