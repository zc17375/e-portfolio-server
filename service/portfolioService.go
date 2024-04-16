package service

import (
	"context"

	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model"
	"github.com/zc17375/e-portfolio-server/model/common"
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
	if pageRequest.KeyWords != "" {
		// 使用正则表达式进行模糊匹配
		keywordFilter := bson.M{"$regex": pageRequest.KeyWords, "$options": "i"}

		// 为每个字段添加搜索条件
		filter["$or"] = []bson.M{
			{"job_title": keywordFilter},
			bson.M{"skills": bson.M{"$elemMatch": keywordFilter}},
			// 添加更多字段
		}
	}

	indi := NewIndividualService(global.EP_MongoDB)

	totalRecords, err := indi.IndiviCollection.CountDocuments(ctx, filter)
	if err != nil {
		global.EP_LOG.Error("取得Portfolio List失敗!", zap.Error(err))
		return totalRecords, err
	}

	// 计算总页数
	totalPages := totalRecords / int64(pageRequest.PageSize)
	if totalRecords%int64(pageRequest.PageSize) > 0 {
		totalPages++
	}

	// 执行分页查询
	skip := int64((pageRequest.CurrentPage - 1) * pageRequest.PageSize)
	limit := int64(pageRequest.PageSize)
	options := options.Find().SetSkip(skip).SetLimit(limit)
	cursor, err := indi.IndiviCollection.Find(ctx, filter, options)
	if err != nil {
		global.EP_LOG.Error("取得分頁列表失敗!", zap.Error(err))
		return cursor, err
	}
	defer cursor.Close(ctx)

	// 解码查询结果
	var results []model.Individual
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
