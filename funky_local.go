// +build local

package funky

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/itsfunky/funky/local"
)

// Handle sets up and starts a local RPC server.
func Handle(handler http.Handler) {
	fmt.Println("Built with Local type handler.")

	port := os.Getenv("FUNKY_SERVER_PORT")
	if port == "" {
		log.Fatal(errors.New("RPC port was not specified."))
	}

	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}

	svc := local.Service{Handler: handler}
	if err := rpc.Register(svc); err != nil {
		log.Fatal("failed to register handler function", err)
	}

	rpc.Accept(lis)
	log.Fatal("accept should not have returned")
}
