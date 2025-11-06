package models

import "go.mongodb.org/mongo-driver/mongo"

type MongodbRepo struct {
	Client *mongo.Client
}

// func NewMongodbRepo() MongodbRepo {
// 	return &MongodbRepo{
// 		Client: mongo.Client,
// 	}
// }
