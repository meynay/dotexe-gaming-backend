package rating_delivery

import (
	"net/http"
	"store/internal/entities"
	"store/internal/usecases/rating_usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RatingDelivery struct {
	ratingusecase *rating_usecase.RatingUsecase
}

func NewRatingDelivery(ru *rating_usecase.RatingUsecase) *RatingDelivery {
	return &RatingDelivery{ratingusecase: ru}
}

func (rd *RatingDelivery) RateProduct(c *gin.Context) {
	productid, _ := strconv.Atoi(c.Param("productid"))
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	input := struct {
		Review string  `json:"review"`
		Rate   float64 `json:"rate"`
	}{}
	if c.BindJSON(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	rate := entities.Rating{
		UserID:    userID,
		ProductID: uint(productid),
		Rate:      input.Rate,
		Review:    input.Review,
		Likes:     0,
		Dislikes:  0,
	}
	err := rd.ratingusecase.RateProduct(rate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "rate added"})
}

func (rd *RatingDelivery) GetRates(c *gin.Context) {
	productid, _ := strconv.Atoi(c.Param("productid"))
	out := rd.ratingusecase.GetRates(uint(productid))
	c.JSON(http.StatusOK, out)
}

func (rd *RatingDelivery) GetRate(c *gin.Context) {
	productid, _ := strconv.Atoi((c.Param("productid")))
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	rate, err := rd.ratingusecase.GetRating(uint(productid), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product is not rated by user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rate": rate})
}
