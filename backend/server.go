package main

import (
	"context"
	"log"
	"net/http"

	"egm.com/backend/proto"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type YourService struct {
	proto.UnimplementedHelloServiceServer
}

func (s *YourService) GetData(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	// Replace this with your actual data retrieval logic
	data := "Hello, " + req.Name
	return &proto.HelloResponse{Greeting: data}, nil
}

func main() {
	// Start gRPC server
	server := grpc.NewServer()
	proto.RegisterHelloServiceServer(server, &YourService{})

	// Enable reflection for tools like grpcwebproxy
	reflection.Register(server)

	// Wrap the gRPC server to make it compatible with gRPC-Web
	wrappedServer := grpcweb.WrapServer(server)

	// Create an HTTP server
	httpServer := &http.Server{
		Addr: ":9090", // Replace with your desired port
		Handler: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			wrappedServer.ServeHTTP(resp, req)
			// if wrappedServer.IsAcceptableGrpcCorsRequest(req) {
			// 	wrappedServer.ServeHTTP(resp, req)
			// } else {
			// 	http.Error(resp, "CORS not allowed", http.StatusForbidden)
			// }
		}),
	}

	log.Println("gRPC-Web server listening on :9090")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
