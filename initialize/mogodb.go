package initialize

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/zc17375/e-portfolio-server/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var client *mongo.Client

func InitMongoDB() *mongo.Database {
	m := global.EP_CONFIG.Mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 設置連線選項
	clientOptions := options.Client().ApplyURI(m.GetURI())
	// 連線到 MongoDB
	var err error
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		global.EP_LOG.Fatal("Failed to connect to mongo database", zap.Error(err))
	}
	// 檢查連線
	err = client.Ping(context.Background(), nil)
	if err != nil {
		global.EP_LOG.Fatal("Failed to Ping to mongo database", zap.Error(err))
	}
	fmt.Println("Connected to MongoDB!")

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// 取得資料庫
	db := client.Database(m.Dbname)
	return db
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if client == nil {
		return
	}
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
