package user

import (
	"log"
	"net/http"

	"github.com/OleksandrZhurba-san/ichgram-server/internal/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Handler handles HTTP requests related to the user domain.
type Handler struct {
	Repo *UserRepository
}

// NewHanlder creates a new user Handler with the provided UserRepository.
func NewHanlder(repo *UserRepository) *Handler {
	return &Handler{Repo: repo}
}

// Register handles user registration.
// It validates the request body, creates a new user, hashes the password,
// checks for existing users, and inserts the new user into the database.
//
// @route POST /api/user/register
// @response 201 Created
func (h *Handler) Register(c *gin.Context) {
	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := NewUserFromInput(&input)

	_ = newUser.BeforeSave()

	if err := h.Repo.InsertUser(newUser); err != nil {
		if err.Error() == "User Already Exists" {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":      newUser.ID.Hex(),
		"message": "User created successfully",
	})
}

// LoginUser handles user login.
// It supports login by email or username, verifies credentials using bcrypt,
// and returns a signed JWT token if successful.
//
// @route POST /api/user/login
// @response 202 Accepted
func (h *Handler) LoginUser(c *gin.Context) {
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

	loggedInUser, err := h.Repo.FindByEmailOrUsername(input.Email, input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Bad credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(loggedInUser.Password),
		[]byte(input.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateToken(loggedInUser.ID.Hex())
	if err != nil {
		log.Println("Failed to generate token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"token": token})
}

