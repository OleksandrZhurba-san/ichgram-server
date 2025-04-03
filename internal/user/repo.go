package user

import (
	"context"
	"time"

	"github.com/OleksandrZhurba-san/ichgram-server/common/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(user *User) error {
	collection := db.GetCollection("ichgram", "users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	user.CreatedAt = time.Now()

	return nil

}
