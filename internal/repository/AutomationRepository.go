package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"tcms/internal/model"
)

type AutomationRepository interface {
	GetAll(ctx context.Context) ([]model.Automation, error)
	Save(ctx context.Context, automation model.NewAutomation) error
	Update(ctx context.Context, id string, automation model.NewAutomation) error
	Remove(ctx context.Context, id string) error
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

func (r automationRepository) Save(ctx context.Context, automation model.NewAutomation) error {
	_, err := r.collection.InsertOne(ctx, automation)
	return err
}

func (r automationRepository) Update(ctx context.Context, id string, automation model.NewAutomation) error {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.ReplaceOne(ctx, bson.M{"_id": idPrimitive}, automation)

	return err
}

func (r automationRepository) Remove(ctx context.Context, id string) error {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": idPrimitive})

	return err
}
