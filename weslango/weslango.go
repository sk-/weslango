package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"cld3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "weslango/proto"
)

var port = flag.Int("port", 50051, "The server port")

type languageServer struct {
	det *cld3.LanguageDetector
}

func newServer() (*languageServer, error) {
	det, err := cld3.New(8, 1024)
	if err != nil {
		return &languageServer{}, err
	}
	return &languageServer{det}, nil
}

func (s *languageServer) DetectLanguage(ctx context.Context, req *pb.DetectRequest) (*pb.DetectResponse, error) {
	r := s.det.FindLanguage(req.Text)
	log.Printf("Lang: %s\n", r.Language)
	res := pb.DetectResponse{Language: r.Language, Probability: r.Probability, IsReliable: r.IsReliable, Latin: r.Latin}
	return &res, nil
}

func newGrpcServer() (s *grpc.Server) {
	// Set the maximum size to 512 kiB
	s = grpc.NewServer(grpc.MaxRecvMsgSize(512 * 1024 * 1024))
	langServer, err := newServer()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	pb.RegisterLanguageServer(s, langServer)
	reflection.Register(s)
	return
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := newGrpcServer()
	log.Printf("Listening on localhost:%d", *port)
	grpcServer.Serve(lis)
}
