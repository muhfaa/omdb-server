package grpccontroller

import (
	"log"
	"net"

	"github.com/muhfaa/omdb-server/repository/grpcstub"
	"github.com/muhfaa/omdb-server/service"
	"google.golang.org/grpc"
)

func RunGRPCServer(port string, searchService service.SearchService, singleService service.SingleService) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	grpcstub.RegisterOmdbServer(grpcServer, NewGRPCService(searchService, singleService))
	grpcServer.Serve(lis)
}
