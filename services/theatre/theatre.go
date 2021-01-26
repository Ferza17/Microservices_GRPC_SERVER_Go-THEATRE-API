package theatre

import (
	"context"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/domains/theatre"
)

var Services theatreInterface = &theatreStruct{}

type (
	theatreStruct struct {
	}

	theatreInterface interface {
		GetTheatreById(theatre theatre.Theatre, ctx context.Context) (*theatre.Theatre, error)
		GetSeatsByIdTheatre(theatre theatre.Theatre, ctx context.Context) ([]*theatre.Seats, error)
		UpdateSeat(theatre theatre.Theatre, ctx context.Context) (*theatre.Theatre, error)
	}
)

func (t *theatreStruct) GetTheatreById(theatre theatre.Theatre, ctx context.Context) (*theatre.Theatre, error) {
	if err := theatre.Get(ctx); err != nil {
		return nil, err
	}

	return &theatre, nil
}

func (t *theatreStruct) GetSeatsByIdTheatre(theatre theatre.Theatre, ctx context.Context) ([]*theatre.Seats, error) {
	panic("implement me")
}

func (t *theatreStruct) UpdateSeat(theatre theatre.Theatre, ctx context.Context) (*theatre.Theatre, error) {
	panic("implement me")
}
