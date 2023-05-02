package mongodb

import (
	"context"
	"github.com/Ryeom/soft-formula/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func New(platform, target string) *mongo.Client {
	ip := ""
	// 여기서 통신 테스트 하기
	return newMongoClient(ip)
}

func newMongoClient(key string) *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	clientOptions := options.Client().ApplyURI(key).SetMaxPoolSize(3)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Logger.Error("Client Connection %s", err)

	}
	client.Ping(ctx, nil)
	if err != nil {
		log.Logger.Error("Client Ping %s", err)

	}
	return client
}
v