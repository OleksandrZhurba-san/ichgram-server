package post

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	UserID        primitive.ObjectID   `bson:"userId" json:"userId" binding:"required"`
	Images        []string             `bson:"images" json:"images" binding:"required"`
	Description   string               `bson:"description" json:"description" binding:"required"`
	LikesCount    int                  `bson:"likes_count,omitempty" json:"likes_count,omitempty"`
	CommentsCount int                  `bson:"comments_count,omitempty" json:"comments_count,omitempty"`
	Likes         []primitive.ObjectID `bson:"likes,omitempty" json:"likes,omitempty"`
	Comments      []primitive.ObjectID `bson:"comments,omitempty" json:"comments,omitempty"`
	CreatedAt     time.Time            `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

// NewPostFromInput returns a new Post with default values applied.
func NewPostFromInput(input *Post) *Post {
	return &Post{
		UserID:        input.UserID,
		Images:        input.Images,
		Description:   input.Description,
		LikesCount:    0,
		CommentsCount: 0,
		Likes:         []primitive.ObjectID{},
		Comments:      []primitive.ObjectID{},
		CreatedAt:     time.Now(),
	}
}
