package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jjtimmons/sblast/api"
	pb "github.com/jjtimmons/sblast/pb/server"
	"github.com/jjtimmons/sblast/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.Default()
	logger.Info("starting server")
	ctx := context.Background()

	// register metrics
	if err := metrics.Register(); err != nil {
		log.Fatal(err)
	}

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

	// register ui and metrics handlers
	metricsHandler := promhttp.Handler()
	fileServer := http.FileServer(http.Dir("./ui/out"))
	err = mux.HandlePath("GET", "/{dir=**}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		if pathParams["dir"] == "metrics" {
			metricsHandler.ServeHTTP(w, r)
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", mux))
}
