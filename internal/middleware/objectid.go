package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateObjectID(paramKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param(paramKey)
		objectID, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID format",
			})
			return
		}

		//Store the parsed ObjectID in context
		c.Set("objectID", objectID)
		c.Next()
	}
}
