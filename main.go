package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jjtimmons/sblast/api"
	pb "github.com/jjtimmons/sblast/gen/sblast"

	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.Default()
	logger.Info("starting server")

	mux := runtime.NewServeMux()

	// register gateway
	if err := pb.RegisterSemanticBlastServiceHandlerServer(context.Background(), mux, api.New()); err != nil {
		log.Fatal("failed to register gateway", "error", err)
	}

	// register ui
	fileServer := http.FileServer(http.Dir("./ui/out"))
	mux.HandlePath("GET", "/{dir_path=**}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		fileServer.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
