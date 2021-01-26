package applications

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/utils/logger"
	"log"
	"net"
)

func StartApplication() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	mapRPC(server)
	reflection.Register(server)

	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		logger.Error("unable to create server", err)
	}

	if err := server.Serve(lis); err != nil {
		logger.Error("Unable to serve ", err)
	}

	logger.Info("Server Running...")
}
