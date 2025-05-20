package admin_delivery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"store/internal/entities"
	"store/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ad *AdminDelivery) AddProduct(c *gin.Context) {
	const maxMemory = 10 << 20
	if err := c.Request.ParseMultipartForm(maxMemory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse form"})
		return
	}
	productJSON := c.PostForm("product")
	if productJSON == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing product data"})
		return
	}
	product := entities.Product{}
	if err := json.Unmarshal([]byte(productJSON), &product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product data"})
		return
	}
	primaryImage, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "primary image is required"})
		return
	}
	uniqueID := uuid.New().String()
	fileExt := filepath.Ext(primaryImage.Filename)
	allowedExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true, ".svg": true}
	if !allowedExtensions[fileExt] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file type for primary image"})
		return
	}
	primaryFilename := fmt.Sprintf("product_%s_primary%s", uniqueID, fileExt)
	primaryDst := filepath.Join(os.Getenv("PRODUCT_DIR"), primaryFilename)
	if err := c.SaveUploadedFile(primaryImage, primaryDst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save primary image"})
		return
	}
	host := os.Getenv("IMAGE_HOST")
	link := fmt.Sprintf("%sproduct/%s", host, primaryFilename)
	log.Printf("File %s uploaded successfully. Link: %s\n", primaryFilename, link)
	product.Image = link
	secondaryImages := c.Request.MultipartForm.File["images"]
	var imageLinks []string
	var savedFiles []string
	for i, fileHeader := range secondaryImages {
		if fileHeader.Size > maxMemory {
			pkg.CleanupFiles(savedFiles)
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("file %s is too large", fileHeader.Filename)})
			return
		}
		fileExt := filepath.Ext(fileHeader.Filename)
		if !allowedExtensions[fileExt] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file type for secondary image"})
			return
		}
		uniqueID = uuid.New().String()
		secondaryFilename := fmt.Sprintf("product_%s_secondary_%d%s", uniqueID, i, fileExt)
		secondaryDst := filepath.Join(os.Getenv("PRODUCT_DIR"), secondaryFilename)
		if err := c.SaveUploadedFile(fileHeader, secondaryDst); err != nil {
			pkg.CleanupFiles(savedFiles)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save secondary images"})
			return
		}
		savedFiles = append(savedFiles, secondaryDst)
		link := fmt.Sprintf("%sproduct/%s", host, secondaryFilename)
		imageLinks = append(imageLinks, link)
	}
	product.Images = entities.JSONB{"images": imageLinks}
	err = ad.adminusecase.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added successfully"})
}

func (ad *AdminDelivery) EditProduct(c *gin.Context) {
	const maxMemory = 10 << 20
	if err := c.Request.ParseMultipartForm(maxMemory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse form"})
		return
	}
	productJSON := c.PostForm("product")
	if productJSON == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing product data"})
		return
	}
	product := entities.Product{}
	if err := json.Unmarshal([]byte(productJSON), &product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product data"})
		return
	}
	primaryImage, err := c.FormFile("image")
	if err == nil {
		uniqueID := uuid.New().String()
		fileExt := filepath.Ext(primaryImage.Filename)
		allowedExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true, ".svg": true}
		if !allowedExtensions[fileExt] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file type for primary image"})
			return
		}
		primaryFilename := fmt.Sprintf("product_%s_primary%s", uniqueID, fileExt)
		primaryDst := filepath.Join(os.Getenv("PRODUCT_DIR"), primaryFilename)
		if err := c.SaveUploadedFile(primaryImage, primaryDst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save primary image"})
			return
		}
		host := os.Getenv("IMAGE_HOST")
		link := fmt.Sprintf("%sproduct/%s", host, primaryFilename)
		log.Printf("File %s uploaded successfully. Link: %s\n", primaryFilename, link)
		product.Image = link
	}
	secondaryImages := c.Request.MultipartForm.File["images"]
	var imageLinks []string
	var savedFiles []string
	for i, fileHeader := range secondaryImages {
		if fileHeader.Size > maxMemory {
			pkg.CleanupFiles(savedFiles)
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("file %s is too large", fileHeader.Filename)})
			return
		}
		fileExt := filepath.Ext(fileHeader.Filename)
		allowedExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true, ".svg": true}
		if !allowedExtensions[fileExt] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file type for secondary image"})
			return
		}
		uniqueID := uuid.New().String()
		secondaryFilename := fmt.Sprintf("product_%s_secondary_%d%s", uniqueID, i, fileExt)
		secondaryDst := filepath.Join(os.Getenv("PRODUCT_DIR"), secondaryFilename)
		if err := c.SaveUploadedFile(fileHeader, secondaryDst); err != nil {
			pkg.CleanupFiles(savedFiles)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save secondary images"})
			return
		}
		host := os.Getenv("IMAGE_HOST")
		savedFiles = append(savedFiles, secondaryDst)
		link := fmt.Sprintf("%sproduct/%s", host, secondaryFilename)
		imageLinks = append(imageLinks, link)
	}
	product.Images = entities.JSONB{"images": imageLinks}
	err = ad.adminusecase.EditProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "product edited successfully"})
}

func (ad *AdminDelivery) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("productid"))
	err := ad.adminusecase.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}

func (ad *AdminDelivery) GetActiveProductsCount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"count": ad.adminusecase.GetActiveProductsCount()})
}
