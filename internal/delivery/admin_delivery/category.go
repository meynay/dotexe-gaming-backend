package admin_delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"store/internal/entities"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ad *AdminDelivery) AddCategory(c *gin.Context) {
	const maxMemory = 10 << 20 // 10MB
	categoryDir := os.Getenv("CATEGORY_DIR")
	if categoryDir == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "category upload directory not configured"})
		return
	}
	host := os.Getenv("IMAGE_HOST")
	if host == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "image host not configured"})
		return
	}
	if err := c.Request.ParseMultipartForm(maxMemory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse form"})
		return
	}
	categoryJSON := c.PostForm("category")
	if categoryJSON == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing category data"})
		return
	}
	var category entities.Category
	if err := json.Unmarshal([]byte(categoryJSON), &category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category data"})
		return
	}
	if err := os.MkdirAll(categoryDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create upload directory"})
		return
	}
	primaryImage, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category image is required"})
		return
	}
	uniqueID := uuid.New().String()
	fileExt := strings.ToLower(filepath.Ext(primaryImage.Filename))
	allowedExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	if !allowedExtensions[fileExt] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file type for category image"})
		return
	}
	filename := fmt.Sprintf("category_%s%s", uniqueID, fileExt)
	filePath := filepath.Join(categoryDir, filename)
	if err := c.SaveUploadedFile(primaryImage, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save category image"})
		return
	}
	category.Image = fmt.Sprintf("%s/%s", strings.TrimSuffix(host, "/"), filename)
	if err := ad.adminusecase.AddCategory(category); err != nil {
		os.Remove(filePath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category added seccessfully"})
}

func (ad *AdminDelivery) EditCategory(c *gin.Context) {
	category := entities.Category{}
	if c.BindJSON(&category) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	err := ad.adminusecase.EditCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category edited seccessfully"})
}

func (ad *AdminDelivery) DeleteCategory(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("categoryid"))
	err := ad.adminusecase.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted successfully"})
}
