package grpc

import (
	"bibit-test/src/constants"
	"bibit-test/src/controllers"
	"bibit-test/src/models/proto"
	"github.com/sarulabs/di"
	googleGrpc "google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type GRPC struct {
	ioc        di.Container
	controller *controllers.Controller
}

func New(ioc di.Container) *GRPC {
	return &GRPC{
		controller: ioc.Get(constants.CONTROLLER).(*controllers.Controller),
	}
}

func (r *GRPC) Run() {
	server := googleGrpc.NewServer()

	service := r.controller.MovieGRPC
	proto.RegisterMoviesServer(server, service)

	r.Serve(server)
}

func (r *GRPC) Serve(server *googleGrpc.Server) {
	port, found := os.LookupEnv("GRPC_PORT")
	if !found {
		port = "1333"
	}
	l, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("could not listen to %s: %v", port, err)
	}

	log.Fatal(server.Serve(l))
}
