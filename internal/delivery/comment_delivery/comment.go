package comment_delivery

import (
	"net/http"
	"store/internal/entities"
	"store/internal/usecases/comment_usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentDelivery struct {
	commentusecase *comment_usecase.CommentUsecase
}

func NewCommentDelivery(cu *comment_usecase.CommentUsecase) *CommentDelivery {
	return &CommentDelivery{commentusecase: cu}
}

func (cd *CommentDelivery) CommentOnProduct(c *gin.Context) {
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
		Comment string `json:"comment"`
		Parent  string `json:"parent"`
		IsAdmin bool   `json:"is_admin"`
	}{}
	if c.BindJSON(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	if input.Parent == "" {
		input.Parent = "0"
	}
	par, _ := strconv.Atoi(input.Parent)
	prnt := uint(par)
	cmnt := entities.Comment{
		Comment:   input.Comment,
		IsAdmin:   input.IsAdmin,
		UserID:    userID,
		ProductID: uint(productid),
		ParentID:  &prnt,
	}
	err := cd.commentusecase.CommentOnProduct(cmnt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "comment added"})
}

func (cd *CommentDelivery) GetComments(c *gin.Context) {
	productid, _ := strconv.Atoi(c.Param("productid"))
	out := cd.commentusecase.GetComments(uint(productid))
	c.JSON(http.StatusOK, out)
}
