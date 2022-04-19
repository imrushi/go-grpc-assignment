package main

import (
	"flag"
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

func main() {
	log_level := flag.String("log-level", "info", "To set log level like debug,info...")
	serverPort := flag.Int("server-port", 9092, "Server serve apis on given port")
	flag.Parse()

	logLevel, err := logrus.ParseLevel(*log_level)
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

	user := data.NewUserDetail(*log)

	// create a new gRPC server, use WithInsecure to allow http connections
	grpcServer := grpc.NewServer()
	log.Debugf("gRPC server is initalized")
	// create an instance of the Users Server
	u := server.NewUsers(user, *log)
	log.Debugf("Users is initalized")
	// register the Users detail server
	protos.RegisterUserDetailServer(grpcServer, u)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *serverPort))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	// listen for requests
	log.Infof("Starting gRPC server on port: %d", *serverPort)
	grpcServer.Serve(l)
}
