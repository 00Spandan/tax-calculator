package main

import (
	"log"
	"net"
	"os"

	taxv1 "github.com/00Spandan/financial-tools/gen/go/tax"
	"github.com/00Spandan/financial-tools/services/tax-calculator/internal/server"
	"github.com/00Spandan/financial-tools/services/tax-calculator/internal/tax"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	taxv1.RegisterTaxCalculatorServer(grpcServer, server.NewTaxServer(tax.NewCalculator()))

	log.Printf("tax-calculator gRPC server listening on :%s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
