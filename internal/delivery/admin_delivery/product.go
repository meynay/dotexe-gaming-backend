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

func (ad *AdminDelivery) AddProduct(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse form"})
		return
	}
	product := entities.Product{}
	if c.BindJSON(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	primaryImage := form.File["image"]
	secondaryImages := form.File["images"]
	dir := os.Getenv("PRODUCT_DIR")
	dst := dir + time.Now().String() + primaryImage[0].Filename
	if err = c.SaveUploadedFile(primaryImage[0], dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't save file"})
		return
	}
	host := os.Getenv("IMAGE_HOST")
	link := fmt.Sprintf("%s%s", host, dst)
	log.Printf("File %s uploaded successfully. Link: %s\n", primaryImage[0].Filename, link)
	product.Image = link
	images := []string{}
	for i, file := range secondaryImages {
		dst = dir + time.Now().String() + string(i) + file.Filename
		if err = c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't save file"})
			return
		}
		link := fmt.Sprintf("%s%s", host, dst)
		log.Printf("File %s uploaded successfully. Link: %s\n", file.Filename, link)
		images = append(images, link)
	}
	product.Images = images
	err = ad.adminusecase.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added successfully"})
}

func (ad *AdminDelivery) EditProduct(c *gin.Context) {
	product := entities.Product{}
	if c.BindJSON(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	err := ad.adminusecase.EditProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "product edited successfully"})
}

func (ad *AdminDelivery) DeleteProduct(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	err := ad.adminusecase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "product deleted successfully"})
}
