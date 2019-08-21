package main

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "weslango/proto"
)

func TestWeslango_DetectLanguage(t *testing.T) {
	ctx := context.Background()
	s, lis := startGRPCServer()
	// it is here to properly stop the server
	//defer func() {time.Sleep(10 * time.Millisecond)}()
	defer s.Stop()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(getBufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	c := pb.NewLanguageClient(conn)

	examples := []struct {
		text string
		lang string
	}{
		{"short", ""},
		{"this is a sentende in english", "eng"},
		{"esta es una frase en español", "spa"},
		{"das ist ein Satz aug Deutsch", "deu"},
	}
	for _, ex := range examples {
		r, err := c.DetectLanguage(ctx, &pb.DetectRequest{Text: ex.text})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if r.Language != ex.lang {
			t.Fatalf("unexpected lang: got '%s' expected '%s'", r.Language, ex.lang)
		}
	}
}

func startGRPCServer() (*grpc.Server, *bufconn.Listener) {
	bufSize := 1024 * 1024
	lis := bufconn.Listen(bufSize)
	s := newGrpcServer()
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()
	return s, lis
}

func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}
