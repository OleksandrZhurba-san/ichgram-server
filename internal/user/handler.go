package user

import (
	"log"
	"net/http"

	"github.com/OleksandrZhurba-san/ichgram-server/internal/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := NewUserFromInput(&input)

	_ = newUser.BeforeSave()

	if err := InsertUser(newUser); err != nil {
		if err.Error() == "User Already Exists" {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":      newUser.ID.Hex(),
		"message": "User created successfully",
	})
}

func LoginUser(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Invalid inputs %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inputs"})
		return
	}

	if input.Email == "" && input.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Either email or username is required"},
		)
		return
	}

	loggedInUser, err := FindByEmailOrUsername(input.Email, input.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Bad credentials"})
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(loggedInUser.Password),
		[]byte(input.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}

	token, err := auth.GenerateToken(loggedInUser.ID.Hex())
	if err != nil {
		log.Println("Failed to generate token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"token": token})

}
