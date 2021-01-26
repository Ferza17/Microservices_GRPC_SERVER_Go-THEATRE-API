package theatre

import (
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/utils/env"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-THEATRE-API/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

var (
	mongodbUrl = env.GetEnvironmentVariable("MONGODB_URL")
	Client     *mongo.Client
)


func init()  {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbUrl))
	if err != nil {
		logger.Error("Unable to connect to DB ", err)
	}
	Client = client
}