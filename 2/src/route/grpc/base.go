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
	Controller *controllers.Controller
}

func New(ioc di.Container) *GRPC {
	return &GRPC{
		Controller: ioc.Get(constants.CONTROLLER).(*controllers.Controller),
	}
}

func (r *GRPC) Run() {
	server := googleGrpc.NewServer()

	service := r.Controller.MovieGRPC
	proto.RegisterMoviesServer(server, service)

	r.Serve(server)
}

func (r *GRPC) Serve(server *googleGrpc.Server) {
	port, found := os.LookupEnv("GRPC_PORT")
	if !found {
		port = "8000"
	}
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("could not listen to %s: %v", port, err)
	}
	log.Printf("Listenting in %s", port)

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
