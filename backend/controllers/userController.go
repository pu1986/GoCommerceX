package controllers

import (
	"context"
	"fmt"
	"gocommerce/config"
	"gocommerce/models"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("supersecretkey")

// 🏠 Home route
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GoCommerce Backend Running 🚀",
	})
}

// 🧾 Signup Handler
func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ✅ Check if user already exists
	var existing models.User
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existing)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// ✅ Hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)

	// ✅ Default: new users are not admins
	user.IsAdmin = false
	user.ID = primitive.NewObjectID()

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":      user.ID.Hex(),
			"name":    user.Name,
			"email":   user.Email,
			"isAdmin": user.IsAdmin,
		},
	})
}

// 🔐 Login Handler
func Login(c *gin.Context) {
	var creds models.LoginInput
	var user models.User

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ Clean inputs
	creds.Email = strings.TrimSpace(creds.Email)
	creds.Password = strings.TrimSpace(creds.Password)

	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"email": creds.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// ✅ Compare hashed passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// ✅ Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  user.ID.Hex(),
		"email":   user.Email,
		"isAdmin": user.IsAdmin,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// ✅ Send user info + token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
		"user": gin.H{
			"name":    user.Name,
			"email":   user.Email,
			"isAdmin": user.IsAdmin,
		},
	})
}

// 🛠️ Make a user admin manually
func MakeAdmin(c *gin.Context) {
	email := c.Param("email")
	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{"isAdmin": true}}
	res, err := collection.UpdateOne(ctx, bson.M{"email": email}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to promote user"})
		return
	}

	if res.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("✅ %s promoted to admin successfully!", email),
	})
}
