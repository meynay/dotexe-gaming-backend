package product_delivery

import (
	"net/http"
	"store/internal/entities"
	"store/internal/usecases/product_usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductDelivery struct {
	pu *product_usecase.ProductUseCase
}

func NewProductDelivery(u *product_usecase.ProductUseCase) *ProductDelivery {
	return &ProductDelivery{pu: u}
}

func (pd *ProductDelivery) AddProduct(c *gin.Context) {
	product := entities.Product{}
	if c.BindJSON(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad json format"})
		return
	}
	if pd.pu.AddProduct(product) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't add product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added successfully"})
}

func (pd *ProductDelivery) GetProduct(c *gin.Context) {
	productid := c.Param("id")
	product, err := pd.pu.GetProduct(productid)
	if err.Error() == "couldn't find product" {
		c.JSON(http.StatusNotFound, gin.H{"message": "no products found with given id"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error occured during decoding product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pd *ProductDelivery) GetProducts(c *gin.Context) {
	query := c.Query("query")
	categoryid := c.Query("categoryid")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	ipp, err := strconv.Atoi(c.Query("ipp"))
	if err != nil {
		ipp = 12
	}
	order, err := strconv.Atoi(c.Query("order"))
	if err != nil {
		order = entities.MostRelevant
	}
	products, pages, err := pd.pu.FilterProducts(entities.Filter{
		Query:         query,
		CategoryID:    categoryid,
		Page:          page,
		NumberOfItems: ipp,
		Order:         order,
	})
	if err.Error() == "couldn't get products" {
		c.JSON(http.StatusNotFound, gin.H{"message": "no products found with given filters"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error occured during decoding products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pages":    pages,
		"products": products,
	})
}

func (pd *ProductDelivery) EditProduct(c *gin.Context) {
	product := entities.Product{}
	if c.BindJSON(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad json format"})
		return
	}
	if pd.pu.EditProduct(product) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't update product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product updated successfully"})
}

func (pd *ProductDelivery) DeleteProduct(c *gin.Context) {
	id := c.Query("id")
	if pd.pu.DeleteProduct(id) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
