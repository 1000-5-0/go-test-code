package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

type A *mongo.Client

func Initialize(ip string, size uint64) *mongo.Client {
	var mongoLongpollConnectionPool *mongo.Client
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://" + ip).SetMaxPoolSize(size)
	println("MONGO DB HOST : " + ip + " ...")
	println("MONGO DB MAX POOL SIZE : " + strconv.FormatUint(size, 10) + " ...")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		println("MONGO DB CONNECTION FAIL ... ", err)
	}
	mongoLongpollConnectionPool = client
	err = mongoLongpollConnectionPool.Ping(ctx, nil)
	if err != nil {
		println("MONGO DB PING FAIL ... ", err)
	}
	return mongoLongpollConnectionPool
}

//func GetMongoCollection(db, collection string) *mongo.Collection {
//	return mongoLongpollConnectionPool.Database(db).Collection(collection)
//}
