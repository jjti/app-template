package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jjtimmons/sblast/api"
	pb "github.com/jjtimmons/sblast/gen/server"

	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.Default()
	logger.Info("starting server")
	ctx := context.Background()

	// create the router
	mux := runtime.NewServeMux()

	// create the api
	grpcServer, err := api.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// register gateway
	if err = pb.RegisterEchoServiceHandlerServer(context.Background(), mux, grpcServer); err != nil {
		log.Fatal(err)
	}

	// register ui
	fileServer := http.FileServer(http.Dir("./ui/static"))
	err = mux.HandlePath("GET", "/{dir_path=**}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		fileServer.ServeHTTP(w, r)
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", mux))
}
