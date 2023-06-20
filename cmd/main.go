package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/hashicorp/go-hclog"

	pb "github.com/jjtimmons/echo/pb/server"
	"github.com/jjtimmons/echo/pkg/api"
	"github.com/jjtimmons/echo/pkg/metrics"
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

	// make the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// make a done channel
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-done
		logger.Error("shutdown signal received", "signal", sig)
		if err := server.Shutdown(ctx); err != nil {
			logger.Error("failed to shut down server", "error", err)
		}
	}()

	// start server and wait
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error("failure seen in server", "error", err)
	}
}
