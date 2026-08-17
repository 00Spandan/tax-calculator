package main

import (
	"log"
	"net/http"
	"os"

	"github.com/00Spandan/financial-tools/services/bff/internal/server"
	"github.com/00Spandan/financial-tools/services/bff/internal/taxclient"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	taxCalculatorTarget := os.Getenv("TAX_CALCULATOR_GRPC_TARGET")
	if taxCalculatorTarget == "" {
		taxCalculatorTarget = "localhost:8080"
	}

	client, err := taxclient.New(taxCalculatorTarget)
	if err != nil {
		log.Fatalf("failed to initialize tax client: %v", err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("failed to close tax client: %v", err)
		}
	}()

	httpServer := server.New(client)
	log.Printf("financial-tools-bff listening on :%s", port)
	if err := http.ListenAndServe(":"+port, httpServer.Handler()); err != nil {
		log.Fatalf("failed to serve HTTP API: %v", err)
	}
}
