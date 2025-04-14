package post

import (
	"time"

	"github.com/OleksandrZhurba-san/ichgram-server/common/contextutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository struct {
	Collection *mongo.Collection
}

func NewPostRepository(db *mongo.Database) *PostRepository {
	return &PostRepository{
		Collection: db.Collection("posts"),
	}
}

func (r *PostRepository) InsertPost(post *Post) error {

	ctx, cancel := contextutil.WithTimeout()
	defer cancel()

	post.CreatedAt = time.Now()
	result, err := r.Collection.InsertOne(ctx, post)
	if err != nil {
		return err
	}

	post.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}
