package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"tcms/m/internal/db/model"
)

type AutomationRepository interface {
	GetAll(ctx context.Context) ([]model.Automation, error)
}

type automationRepository struct {
	collection *mongo.Collection
}

func CreateAutomationRepository(db *mongo.Database) AutomationRepository {
	collection := db.Collection("automation")
	return automationRepository{collection: collection}
}

func (r automationRepository) GetAll(ctx context.Context) ([]model.Automation, error) {
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	list := make([]model.Automation, cur.RemainingBatchLength())
	for i := 0; cur.Next(ctx); i++ {
		var automation model.Automation
		err := cur.Decode(&automation)
		if err != nil {
			return nil, err
		}
		list[i] = automation
	}
	return list, nil
}
