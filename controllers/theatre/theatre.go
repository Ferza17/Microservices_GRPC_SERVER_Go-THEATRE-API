package theatre

import (
	"context"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/domains/theatre"
	theatreService "github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/services/theatre"
	theatreProto "github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/protos/server/theatre"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/utils/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
}

func (s *Server) GetTheatreById(ctx context.Context, request *theatreProto.GetTheatreByIdRequest) (*theatreProto.GetTheatreByIdResponse, error) {
	oid, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, errors.InvalidArgument("Invalid type of id")
	}

	res, err := theatreService.Services.GetTheatreById(theatre.Theatre{
		Id: oid,
	}, ctx)

	if err != nil {
		return nil, err
	}

	//TODO: Create utils for generate from object to proto response3
	_ = res
	//return &theatreProto.GetTheatreByIdResponse{
	//	Theatre:
	//}

	return nil, nil
}
