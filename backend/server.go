package main

import (
	"fmt"
	"log"
	"net/http"

	"egm.com/backend/pb"
	"egm.com/backend/services"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	Port          = 9090
	AllowedOrigin = "http://localhost:4200"
)

func main() {
	// Start gRPC server
	server := grpc.NewServer()
	pb.RegisterGridServiceServer(server, &services.GridService{})

	// Enable reflection for tools like grpcwebproxy
	reflection.Register(server)

	// Wrap the gRPC server to make it compatible with gRPC-Web
	wrappedServer := grpcweb.WrapServer(
		server,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return origin == AllowedOrigin
		}))

	// Create an HTTP server
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", Port),
		Handler: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			wrappedServer.ServeHTTP(resp, req)
		}),
	}

	log.Printf("gRPC-Web server listening on :%d\n", Port)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
