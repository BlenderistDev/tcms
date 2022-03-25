package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection(ctx context.Context) (*mongo.Database, error) {
	url, err := getMongoHost()
	if err != nil {
		return nil, err
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	err = client.Database("tcms").CreateCollection(ctx, "automation")
	if err != nil {
		cmdErr, ok := err.(mongo.CommandError)
		if ok {
			if !cmdErr.HasErrorCode(48) {
				return nil, cmdErr
			}
		} else {
			if err != nil {
				return nil, cmdErr
			}

		}
	}
	return client.Database("tcms"), nil
}
