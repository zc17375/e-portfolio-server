package service

import (
	"context"

	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IndividualService struct {
	IndiviCollection  *mongo.Collection
	ProjectCollection *mongo.Collection
}

func NewIndividualService(db *mongo.Database) *IndividualService {
	return &IndividualService{
		IndiviCollection:  db.Collection("individual"),
		ProjectCollection: db.Collection("project"),
	}
}

func (is *IndividualService) CreateIndividual(ctx context.Context, indivi model.Individual, uuid string) (mongo.InsertOneResult, error) {
	indi := NewIndividualService(global.EP_MongoDB)

	// 建立individual ID & User UUID
	indivi.ID = primitive.NewObjectID()
	indivi.UserUUID = uuid

	// 建立SocialMedia UserId
	indivi.SocialMedia.UserID = indivi.ID

	insert, err := indi.IndiviCollection.InsertOne(ctx, indivi)
	if err != nil {
		return mongo.InsertOneResult{}, err
	}
	// fmt.Println(insert)
	return *insert, nil
}
