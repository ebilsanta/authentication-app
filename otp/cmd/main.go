package main

import (
	// "context"
	otp "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/api/proto"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	// "net/smtp"
	// "os"
)

func GenerateOTP(n int) string {
	var numberRunes = []rune("0123456789")
    b := make([]rune, n)
    for i := range b {
        b[i] = numberRunes[rand.Intn(len(numberRunes))]
    }
    return string(b)
}

type myOTPServer struct {
	otp.UnimplementedOTPServer
}


func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	otpService := &myOTPServer{}
	otp.RegisterOTPServer(serverRegistrar, otpService)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}