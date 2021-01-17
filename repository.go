package user

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetAll() (Users, error)
}

func NewUserRepository(resource *Resource) UserRepository {
	collection := resource.DB.Collection("user")
	repository := &MongoUserRepository{resource: resource, collection: collection}
	return repository
}

type MongoUserRepository struct {
	resource   *Resource
	collection *mongo.Collection
}

func (repo *MongoUserRepository) GetAll() (Users, error) {
	users := Users{}
	ctx, cancel := initContext()
	defer cancel()

	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return Users{}, err
	}

	for cursor.Next(ctx) {
		var user User
		err = cursor.Decode(&user)
		if err != nil {
			logrus.Print(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}
