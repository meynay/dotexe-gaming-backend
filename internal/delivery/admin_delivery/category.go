package admin_delivery

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"store/internal/entities"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ad *AdminDelivery) AddCategory(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse form"})
		return
	}
	category := entities.Category{}
	if err := c.BindJSON(&category); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	primaryImage := form.File["image"]
	dir := os.Getenv("CATEGORY_DIR")
	dst := dir + time.Now().String() + primaryImage[0].Filename
	if err = c.SaveUploadedFile(primaryImage[0], dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't save file"})
		return
	}
	host := os.Getenv("IMAGE_HOST")
	link := fmt.Sprintf("%s%s", host, dst)
	log.Printf("File %s uploaded successfully. Link: %s\n", primaryImage[0].Filename, link)
	category.Image = link
	err = ad.adminusecase.AddCategory(category)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
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
