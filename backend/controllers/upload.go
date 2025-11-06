package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println("❌ Upload error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image uploaded"})
		return
	}

	// ✅ Ensure folder exists
	os.MkdirAll("uploads", os.ModePerm)

	// ✅ Save file
	filename := filepath.Base(file.Filename)
	savePath := filepath.Join("uploads", filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		fmt.Println("❌ Save error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	imageURL := fmt.Sprintf("http://localhost:8080/uploads/%s", filename)
	fmt.Println("✅ Uploaded:", imageURL)

	c.JSON(http.StatusOK, gin.H{
		"message": "Image uploaded successfully",
		"url":     imageURL,
	})
}
