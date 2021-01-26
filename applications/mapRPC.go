package applications

import (
	theatreProto "github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/protos/server/theatre"
	theatreController "github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/controllers/theatre"
	"google.golang.org/grpc"
)

func mapRPC(s *grpc.Server) {
	// Register RPC
	theatreProto.RegisterTheatreServiceServer(s, &theatreController.Server{})
}
