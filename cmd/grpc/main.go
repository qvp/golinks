package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"net"

	"google.golang.org/grpc"

	pb "golinks/internal/grpc"
	"golinks/internal/parser"
)

type server struct {
	pb.UnimplementedGolinksServer
}

func (s *server) GetLinkImages(ctx context.Context, in *pb.LinkImagesRequest) (*pb.LinkImagesReply, error) {
	page, err := parser.LoadHtml(in.GetUrl())
	if err != nil {
		return nil, err
	}

	links, err := parser.GetImagesFromHtml(page)
	if err != nil {
		return nil, err
	}

	return &pb.LinkImagesReply{Message: links}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGolinksServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}
