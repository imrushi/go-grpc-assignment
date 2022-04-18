package main

import (
	"fmt"
	"net"
	"os"

	"github.com/imrushi/go-grpc-assignment/data"
	protos "github.com/imrushi/go-grpc-assignment/protos/api/proto/v1/users"
	"github.com/imrushi/go-grpc-assignment/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var log *logrus.Logger

func init() {
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	log = &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		},
		Level: logLevel,
	}

}

func main() {
	user := data.NewUserDetail(*log)

	// create a new gRPC server, use WithInsecure to allow http connections
	grpcServer := grpc.NewServer()

	// create an instance of the Users Server
	u := server.NewUsers(user, *log)

	// register the Users detail server
	protos.RegisterUserDetailServer(grpcServer, u)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	// listen for requests
	log.Infof("Starting gRPC server on port: 9092")
	grpcServer.Serve(l)
}
