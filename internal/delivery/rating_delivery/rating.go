package rating_delivery

import (
	"net/http"
	"store/internal/usecases/rating_usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingDelivery struct {
	ratingusecase *rating_usecase.RatingUsecase
}

func NewRatingDelivery(ru *rating_usecase.RatingUsecase) *RatingDelivery {
	return &RatingDelivery{ratingusecase: ru}
}

func (rd *RatingDelivery) RateProduct(c *gin.Context) {

}

func (rd *RatingDelivery) GetRates(c *gin.Context) {
	productid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	out := rd.ratingusecase.GetRates(productid)
	c.JSON(http.StatusOK, out)
}
