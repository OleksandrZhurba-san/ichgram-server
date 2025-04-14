package user

import (
	"errors"
	"time"

	"github.com/OleksandrZhurba-san/ichgram-server/common/contextutil"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: db.Collection("users"),
	}
}

func (r *UserRepository) InsertUser(user *User) error {

	ctx, cancel := contextutil.WithTimeout()
	defer cancel()

	user.CreatedAt = time.Now()

	var existing User

	err := r.Collection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"email": user.Email},
			{"username": user.Username},
		},
	}).Decode(&existing)

	if err == nil {
		return errors.New("User Already Exists")
	}

	result, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return nil

}

func (r *UserRepository) FindByEmailOrUsername(email, username string) (*User, error) {

	ctx, cancel := contextutil.WithTimeout()
	defer cancel()

	var user User

	err := r.Collection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"email": email},
			{"username": username},
		},
	}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *UserRepository) FindByUserID(userID primitive.ObjectID) (*User, error) {

	ctx, cancel := contextutil.WithTimeout()
	defer cancel()

	var user User

	filter := bson.M{"_id": userID}
	err := r.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
