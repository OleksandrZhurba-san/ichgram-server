package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	FullName       string               `bson:"full_name" json:"full_name" binding:"required"`
	Username       string               `bson:"username" json:"username" binding:"required"`
	Email          string               `bson:"email" json:"email" binding:"required,email"`
	Password       string               `bson:"password" json:"password" binding:"required,min=6"`
	Bio            string               `bson:"bio,omitempty" json:"about,omitempty"`
	Website        string               `bson:"website,omitempty" json:"website,omitempty"`
	Image          string               `bson:"image,omitempty" json:"image,omitempty"`
	PostsCount     int                  `bson:"posts_count,omitempty" json:"posts_count,omitempty"`
	FollowersCount int                  `bson:"followers_count,omitempty" json:"followers_count,omitempty"`
	FollowingCount int                  `bson:"following_count,omitempty" json:"following_count,omitempty"`
	Posts          []primitive.ObjectID `bson:"posts,omitempty" json:"posts,omitempty"`
	Followers      []primitive.ObjectID `bson:"followers,omitempty" json:"followers,omitempty"`
	Following      []primitive.ObjectID `bson:"following,omitempty" json:"following,omitempty"`
	CreatedAt      time.Time            `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

// NewUserFromInput returns a new User with default values applied.
func NewUserFromInput(input *User) *User {
	return &User{
		FullName:       input.FullName,
		Username:       input.Username,
		Email:          input.Email,
		Password:       input.Password,
		Bio:            "",
		Website:        "",
		Image:          "",
		PostsCount:     0,
		FollowersCount: 0,
		FollowingCount: 0,
		Posts:          []primitive.ObjectID{},
		Followers:      []primitive.ObjectID{},
		Following:      []primitive.ObjectID{},
		CreatedAt:      time.Now(),
	}

}
