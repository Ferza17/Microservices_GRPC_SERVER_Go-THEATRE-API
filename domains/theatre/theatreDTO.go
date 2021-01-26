package theatre

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Seats struct {
		Seat string `bson:"seat"`
		Cost float64 `bson:"cost"`
		Status bool `bson:"status"`
	}
	Auditorium struct {
		Type string `bson:"type"`
		IdMovies string `bson:"idMovies"`
		Schedule string `bson:"date"`
		Seats []*Seats `bson:"seats"`
	}
	Theatre struct {
		Id primitive.ObjectID `bson:"_id"`
		Name string `bson:"name"`
		Address string `bson:"address"`
		Auditorium Auditorium `bson:"auditorium"`
	}
)
