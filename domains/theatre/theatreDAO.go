package theatre

import (
	"context"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/utils/errors"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/datasources/mongodb/theatre"
)

func (t *Theatre) Get(ctx context.Context) error {
	if err := theatre.Client.Connect(ctx); err != nil {
		return errors.Internal("Unable to connect to DB")
	}
	defer theatre.Client.Disconnect(ctx)

	//TODO: Result to OBJ
	return nil
}

func (t *Theatre) Update(ctx *context.Context) error {
	return nil
}
