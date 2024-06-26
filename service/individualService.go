package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IndividualService struct {
	IndiviCollection *mongo.Collection
	// ProjectCollection *mongo.Collection
}

// 載入相關Collection
func NewIndividualService(db *mongo.Database) *IndividualService {
	return &IndividualService{
		IndiviCollection: db.Collection("individual"),
		// ProjectCollection: db.Collection("project"),
	}
}

func (is *IndividualService) CreateIndividual(ctx context.Context, indivi model.Individual, username, uuid string) (mongo.InsertOneResult, error) {
	indi := NewIndividualService(global.EP_MongoDB)

	// 檢查是否已存在user資料
	filter := bson.M{"username": username}
	var result bson.M
	err := indi.IndiviCollection.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {

		// 建立individual ID & User UUID
		indivi.ID = primitive.NewObjectID()
		indivi.UserUUID = uuid
		indivi.UserName = username

		// 建立SocialMedia UserId
		indivi.SocialMedia.UserID = indivi.ID

		// 建立Projects UserId
		if len(indivi.Projects) > 0 {
			for i, _ := range indivi.Projects {
				// 判斷ID如果為0則給新的ID
				if indivi.Projects[i].ID == "0" || indivi.Projects[i].ID == "" {
					indivi.Projects[i].ID = strconv.Itoa(i + 1)
					indivi.Projects[i].UserID = indivi.ID
				}
			}
		}
		// 建立時間
		indivi.CreatedAt = time.Now()
		indivi.UpdatedAt = time.Now()

		insert, err := indi.IndiviCollection.InsertOne(ctx, indivi)
		if err != nil {
			return mongo.InsertOneResult{}, err
		}
		// fmt.Println(insert)
		return *insert, nil
	} else if err != nil {
		return mongo.InsertOneResult{}, err
	} else {
		// 如果使用者資料已存在，返回錯誤
		return mongo.InsertOneResult{}, errors.New("user with UUID already exists")
	}
}

func (is *IndividualService) UpdateIndividual(ctx context.Context, indivi model.Individual, uuid string) (mongo.UpdateResult, error) {
	indi := NewIndividualService(global.EP_MongoDB)

	// 將user的individual資料找出來
	filter := bson.M{"user_uuid": uuid}
	var indiModel model.Individual
	err := indi.IndiviCollection.FindOne(ctx, filter).Decode(&indiModel)
	if err != nil {
		return mongo.UpdateResult{}, err
	}

	// 如果Struct的欄位是Struct則無法用omitempty解決，欄位為空值則忽略
	if (indivi.SocialMedia == model.SocialMedia{}) {
		indivi.SocialMedia = indiModel.SocialMedia
	}

	// 建立Projects UserId
	if len(indivi.Projects) > 0 {
		for i, _ := range indivi.Projects {
			// 判斷ID如果為0則給新的ID
			if indivi.Projects[i].ID == "0" {
				indivi.Projects[i].ID = strconv.Itoa(i + 1)
				indivi.Projects[i].UserID = indivi.ID
			}
		}
	}

	// 建立時間
	indivi.UpdatedAt = time.Now()

	update := bson.M{"$set": indivi}

	// 执行更新操作
	updateResult, err := indi.IndiviCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return *updateResult, err
	}

	return *updateResult, nil

}

func (is *IndividualService) DeleteIndividual(ctx context.Context, uuid string) (mongo.DeleteResult, error) {
	indi := NewIndividualService(global.EP_MongoDB)

	filter := bson.M{"user_uuid": uuid}

	deleteResult, err := indi.IndiviCollection.DeleteOne(ctx, filter)
	if err != nil {
		return *deleteResult, err
	}
	if deleteResult.DeletedCount == 0 {
		return  *deleteResult, errors.New("刪除資料失敗")
	}
	return *deleteResult, nil
}
