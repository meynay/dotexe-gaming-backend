package user_product_usecase

import (
	"store/internal/entities"
	"store/internal/repositories/user_product_rep"
)

type UserProductUsecase struct {
	rep *user_product_rep.UserProductRep
}

func NewUserProductUsecase(r *user_product_rep.UserProductRep) *UserProductUsecase {
	return &UserProductUsecase{rep: r}
}

func (up *UserProductUsecase) FaveProduct(userid, productid string) error {
	return up.rep.AddToFaves(productid, userid)
}

func (up *UserProductUsecase) UnfaveProduct(userid, productid string) error {
	return up.rep.DeleteFromFaves(productid, userid)
}

func (up *UserProductUsecase) CheckFave(userid, productid string) error {
	return up.rep.CheckFave(productid, userid)
}

func (up *UserProductUsecase) GetFaves(userid string) []entities.ProductLess {
	pr := []entities.ProductLess{}
	products, err := up.rep.GetFaves(userid)
	if err != nil {
		return pr
	}
	for _, product := range products {
		prdct := entities.ProductLess{
			ID:          product.ID,
			Image:       product.Image,
			Name:        product.Name,
			Price:       product.Price,
			Off:         product.Off,
			Rating:      product.Rating,
			RateCount:   product.RateCount,
			Category:    up.rep.GetCategoryName(product.CategoryID),
			Description: product.Description,
		}
		pr = append(pr, prdct)
	}
	return pr
}

func (up *UserProductUsecase) CommentOnProduct(c entities.Comment) error {
	return up.rep.AddComment(c)
}

func (up *UserProductUsecase) GetComments(productid string) []entities.CommentOut {
	cmnt, err := up.rep.GetComments(productid)
	comments := []entities.CommentOut{}
	if err != nil {
		return comments
	}
	for _, c := range cmnt {
		newcomment := entities.CommentOut{
			ID:        c.ID,
			Parent:    c.Parent,
			User:      up.rep.GetUsername(c.UserID),
			Comment:   c.Comment,
			CreatedAt: c.CreatedAt,
		}
		comments = append(comments, newcomment)
	}
	return comments
}

func (up *UserProductUsecase) RateProduct(r entities.Rating) error {
	return up.rep.AddRating(r)
}
func (up *UserProductUsecase) GetRates(productid string) []entities.RatingOut {
	rates, err := up.rep.GetRatings(productid)
	ratings := []entities.RatingOut{}
	if err != nil {
		return ratings
	}
	for _, r := range rates {
		newrating := entities.RatingOut{
			ID:        r.ID,
			Rate:      r.Rate,
			Username:  up.rep.GetUsername(r.UserID),
			Review:    r.Review,
			CreatedAt: r.CreatedAt,
			Likes:     r.Likes,
			Dislikes:  r.Dislikes,
		}
		ratings = append(ratings, newrating)
	}
	return ratings
}
