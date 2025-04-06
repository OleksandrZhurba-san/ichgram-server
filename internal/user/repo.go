package user

import (
	"context"
	"errors"
	"time"

	"github.com/OleksandrZhurba-san/ichgram-server/common/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(user *User) error {
	collection := db.GetCollection("ichgram", "users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user.CreatedAt = time.Now()

	var existing User

	err := collection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"email": user.Email},
			{"username": user.Username},
		},
	}).Decode(&existing)

	if err == nil {
		return errors.New("User Already Exists")
	}

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return nil

}
