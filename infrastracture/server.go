package infrastracture

import (
	"context"
	"fmt"
	"os"
	"store/internal/delivery/admin_delivery"
	"store/internal/delivery/comment_delivery"
	"store/internal/delivery/fave_delivery"
	"store/internal/delivery/middlewares"
	"store/internal/delivery/product_delivery"
	"store/internal/delivery/rating_delivery"
	"store/internal/delivery/user_delivery"
	"store/internal/repositories/admin_rep"
	"store/internal/repositories/category_rep"
	"store/internal/repositories/comment_rep"
	"store/internal/repositories/invoice_rep"
	"store/internal/repositories/product_rep"
	"store/internal/repositories/rating_rep"
	"store/internal/repositories/user_rep"
	"store/internal/usecases/admin_usecase"
	"store/internal/usecases/comment_usecase"
	"store/internal/usecases/fave_usecase"
	"store/internal/usecases/product_usecase"
	"store/internal/usecases/rating_usecase"
	"store/internal/usecases/user_usecase"
	"store/pkg/cacher"
	"store/pkg/jwt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartServer() {
	connectionString := os.Getenv("MONGO_CON_STRING")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	jwtsecret := os.Getenv("JWT_SECRET")
	//collections
	database := client.Database("store")
	userCollection := database.Collection("user")
	productCollection := database.Collection("product")
	categoryCollection := database.Collection("category")
	ratesCollection := database.Collection("rating")
	invoicesCollection := database.Collection("invoice")
	commentsCollection := database.Collection("comment")
	adminCollection := database.Collection(os.Getenv("ADMIN_COLLECTION"))

	//repositories
	userRep := user_rep.NewUserRepository(userCollection)
	productRep := product_rep.NewProductRep(productCollection)
	categoryRep := category_rep.NewCategoryRep(categoryCollection)
	ratingRep := rating_rep.NewRatingRep(ratesCollection)
	commentsRep := comment_rep.NewCommentRep(commentsCollection)
	invoiceRep := invoice_rep.NewInvoiceRep(invoicesCollection)
	adminRep := admin_rep.NewAdminRep(adminCollection)

	//usecases
	cacher := cacher.NewCacher(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASS"))
	userUsercase := user_usecase.NewUserUsecase(userRep, cacher)
	productUsecase := product_usecase.NewProductUseCase(productRep, categoryRep)
	faveUsecase := fave_usecase.NewFaveUsecase(userRep, productRep, categoryRep)
	commentUsecase := comment_usecase.NewCommentUsecase(commentsRep, userRep, adminRep)
	ratingUsecase := rating_usecase.NewRatingUsecase(ratingRep, productRep, userRep)
	adminUsecase := admin_usecase.NewAdminUsecase(productRep, categoryRep, invoiceRep, adminRep, userRep)

	//deliveries
	j := jwt.NewJWTTokenHandler(jwtsecret)
	userDelivery := user_delivery.NewUserDelivary(userUsercase, j)
	productDelivery := product_delivery.NewProductDelivery(productUsecase)
	adminDelivery := admin_delivery.NewAdminDelivery(adminUsecase, j)
	faveDelivery := fave_delivery.NewFaveDelivery(faveUsecase)
	commentDelivery := comment_delivery.NewCommentDelivery(commentUsecase)
	ratingDelivery := rating_delivery.NewRatingDelivery(ratingUsecase)

	router := gin.Default()
	//common middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.SecurityHeaders())
	//router.Use(middlewares.RateLimiter())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//router.Use(middlewares.APIKeyAuth())
	router.Use(gzip.Gzip(gzip.BestCompression))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Message": "app works"})
	})

	//public group
	public := router.Group("/public-api")
	public.POST("/signin", userDelivery.FirstStep)
	public.POST("/loginphone", userDelivery.LoginWithPhone)
	public.POST("/loginemail", userDelivery.LoginWithEmail)
	public.POST("/signupphone", userDelivery.SignupWithPhone)
	public.POST("/signupemail", userDelivery.SignupWithEmail)
	public.POST("/refreshtoken", userDelivery.RefreshToken)
	public.GET("/product/:id", productDelivery.GetProduct)
	public.GET("/products", productDelivery.GetProducts)
	public.GET("/categories", productDelivery.GetCategories)
	public.GET("/query", productDelivery.SearchQuery)
	public.GET("/comments/:id", commentDelivery.GetComments)
	public.GET("/rates/:id", ratingDelivery.GetRates)

	//private group (needs auth)
	private := router.Group("/private-api")
	a := middlewares.NewAuth(j)
	private.Use(a.AuthMiddleware())
	private.GET("/isinfaves/:id", faveDelivery.CheckFave)
	private.GET("/userfaves", faveDelivery.GetFaves)
	private.POST("/faveproduct/:id", faveDelivery.FaveProduct)
	private.DELETE("unfaveproduct/:id", faveDelivery.UnfaveProduct)

	//private.Use(middlewares.CSRFMiddleware())

	//admin group
	adminroute := os.Getenv("ADMIN_ROUTE")
	admin := router.Group(adminroute)
	//admin.Use()
	admin.POST("/addproduct", adminDelivery.AddProduct)
	admin.POST("/addcategory", adminDelivery.AddCategory)
	admin.DELETE("/deleteproduct/:id", adminDelivery.DeleteProduct)
	admin.DELETE("/deletecategory/:id", adminDelivery.DeleteCategory)
	admin.PUT("/editproduct/:id", adminDelivery.EditProduct)
	admin.PUT("/editcategory/:id", adminDelivery.EditCategory)
	router.Run(":8080")
}
