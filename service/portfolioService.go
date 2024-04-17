package service

import (
	"context"

	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model"
	"github.com/zc17375/e-portfolio-server/model/common"
	"github.com/zc17375/e-portfolio-server/model/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type PortfolioService struct {
	IndiviCollection *mongo.Collection
	// ProjectCollection *mongo.Collection
}

// 載入相關Collection
func NewPortfolioService(db *mongo.Database) *IndividualService {
	return &IndividualService{
		IndiviCollection: db.Collection("individual"),
		// ProjectCollection: db.Collection("project"),
	}
}

func (ps *PortfolioService) GetUserPortfolio(ctx context.Context, userName string) (model.Individual, error) {
	indi := NewIndividualService(global.EP_MongoDB)

	// 將user的individual資料找出來
	filter := bson.M{"username": userName}
	var indiModel model.Individual
	err := indi.IndiviCollection.FindOne(ctx, filter).Decode(&indiModel)
	if err != nil {
		global.EP_LOG.Error("取得User Portfolio失敗!", zap.Error(err))
		return indiModel, err
	}
	return indiModel, err
}

func (ps *PortfolioService) GetAllPortfolio(ctx context.Context, pageRequest common.Pagination) (interface{}, error) {
	filter := bson.M{}

	if len(pageRequest.KeyWords) > 0 {
	// 使用正則表達式進行模糊匹配
		keywordFilter := bson.M{"$regex": pageRequest.KeyWords, "$options": "i"}
		filter["$or"] = []bson.M{
			{"job_title": keywordFilter},
			{"skills": bson.M{"$elemMatch": keywordFilter}},
			// 增加更多搜尋欄位
		}
	}
	indi := NewIndividualService(global.EP_MongoDB)

	totalRecords, err := indi.IndiviCollection.CountDocuments(ctx, filter)
	if err != nil {
		global.EP_LOG.Error("取得Portfolio List失敗!", zap.Error(err))
		return totalRecords, err
	}

	// 計算總筆數
	totalPages := totalRecords / int64(pageRequest.PageSize)
	if totalRecords%int64(pageRequest.PageSize) > 0 {
		totalPages++
	}

	// 分頁條件
	skip := int64((pageRequest.CurrentPage - 1) * pageRequest.PageSize)
	limit := int64(pageRequest.PageSize)

	// 如果要篩選返回的欄位options.SetProjection(projection)
	// projection := bson.M{
	// 	"name": 1,
	// 	"head_img_path":     1,
	// 	// "_id":      0, // 不返回 _id 字段
	// }
	
	//排序
	sortRule := bson.D{{Key: "name", Value: 1}} 
	options := options.Find().SetSkip(skip).SetLimit(limit).SetSort(sortRule)

	cursor, err := indi.IndiviCollection.Find(ctx, filter, options)
	if err != nil {
		global.EP_LOG.Error("取得分頁列表失敗!", zap.Error(err))
		return cursor, err
	}
	defer cursor.Close(ctx)

	// 解码查询结果
	var results []response.PortfolioListResponse
	if err := cursor.All(ctx, &results); err != nil {
		global.EP_LOG.Error("查詢結果失敗!", zap.Error(err))
		return results, err
	}

	// 返回结果
	response := map[string]interface{}{
		"records":      results,
		"currentPage":  pageRequest.CurrentPage,
		"totalPages":   totalPages,
		"totalRecords": totalRecords,
	}

	return response, nil
}
