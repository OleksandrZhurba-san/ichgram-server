package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginInput struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Username string `json:"username" binding:"omitempty"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	FullName string `bson:"full_name" json:"full_name" binding:"required"`
	Username string `bson:"username" json:"username" binding:"required"`
	Email    string `bson:"email" json:"email" binding:"required,email"`
	Password string `bson:"password" json:"password" binding:"required,min=6"`
}

// TODO: unfinished user with *populated* posts
type UserWithPosts struct {
	ID primitive.ObjectID `json:"id"`
	Email string `json:"email"`
}

