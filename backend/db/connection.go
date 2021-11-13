package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"tcms/m/dry"
)

func GetConnection(ctx context.Context) *mongo.Database {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		dry.HandleErrorPanic(err)
	}
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
