package comment_delivery

import (
	"net/http"
	"store/internal/usecases/comment_usecase"

	"github.com/gin-gonic/gin"
)

type CommentDelivery struct {
	commentusecase *comment_usecase.CommentUsecase
}

func NewCommentDelivery(cu *comment_usecase.CommentUsecase) *CommentDelivery {
	return &CommentDelivery{commentusecase: cu}
}

func (cd *CommentDelivery) CommentOnProduct(c *gin.Context) {

}

func (cd *CommentDelivery) GetComments(c *gin.Context) {
	productid := c.Param("id")
	out := cd.commentusecase.GetComments(productid)
	c.JSON(http.StatusOK, out)
}
