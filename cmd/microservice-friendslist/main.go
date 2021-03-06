package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ubivius/microservice-friendslist/pkg/database"
	"github.com/Ubivius/microservice-friendslist/pkg/handlers"
	"github.com/Ubivius/microservice-friendslist/pkg/resources"
	"github.com/Ubivius/microservice-friendslist/pkg/router"
	"go.opentelemetry.io/otel/exporters/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var log = logf.Log.WithName("friendslist-main")

func main() {
	// Starting k8s logger
	opts := zap.Options{}
	opts.BindFlags(flag.CommandLine)
	newLogger := zap.New(zap.UseFlagOptions(&opts), zap.WriteTo(os.Stdout))
	logf.SetLogger(newLogger.WithName("log"))

	// Initialising open telemetry
	// Creating console exporter
	exporter, err := stdout.NewExporter(
		stdout.WithPrettyPrint(),
	)
	if err != nil {
		log.Error(err, "Failed to initialize stdout export pipeline")
	}

	// Creating tracer provider
	ctx := context.Background()
	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(exporter)
	tracerProvider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(batchSpanProcessor))
	defer func() { _ = tracerProvider.Shutdown(ctx) }()

	// Resources init
	resources := resources.NewResources()

	// Database init
	db := database.NewMongoRelationships(resources)

	// Creating handlers
	relationshipHandler := handlers.NewRelationshipsHandler(db)

	// Router setup
	r := router.New(relationshipHandler)

	// Server setup
	server := &http.Server{
		Addr:        ":9090",
		Handler:     r,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	go func() {
		log.Info("Starting server", "port", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			log.Error(err, "Server error")
		}
	}()

	// Handle shutdown signals from operating system
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	receivedSignal := <-signalChannel

	log.Info("Received terminate, beginning graceful shutdown", "received_signal", receivedSignal.String())

	// DB connection shutdown
	db.CloseDB()
	
	// Server shutdown
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_ = server.Shutdown(timeoutContext)
}
