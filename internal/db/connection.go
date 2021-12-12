package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"tcms/m/internal/dry"
)

func GetConnection(ctx context.Context) *mongo.Database {
	url, err := getMongoHost()
	dry.HandleErrorPanic(err)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	dry.HandleErrorPanic(err)
	err = client.Database("tcms").CreateCollection(ctx, "automation")
	if err != nil {
		cmdErr, ok := err.(mongo.CommandError)
		if ok {
			if !cmdErr.HasErrorCode(48) {
				dry.HandleErrorPanic(cmdErr)
			}
		} else {
			dry.HandleErrorPanic(cmdErr)
		}
	}
	return client.Database("tcms")
}
