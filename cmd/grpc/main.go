package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "golinks/internal/grpc"
	"golinks/internal/link"
)

type server struct {
	pb.UnimplementedGolinksServer
}

func (s *server) GetLinkImages(ctx context.Context, in *pb.LinkImagesRequest) (*pb.LinkImagesReply, error) {
	page, err := link.LoadHtml(in.GetUrl())
	if err != nil {
		return nil, err
	}

	links, err := link.GetImagesFromHtml(page)
	if err != nil {
		return nil, err
	}

	return &pb.LinkImagesReply{Message: links}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGolinksServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
