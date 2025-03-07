package infrastracture

import (
	"context"
	"fmt"
	"os"
	"store/internal/delivery/middlewares"
	"store/internal/delivery/product_delivery"
	"store/internal/delivery/user_delivery"
	"store/internal/delivery/user_product_delivery"
	"store/internal/repositories/product_rep"
	"store/internal/repositories/user_product_rep"
	"store/internal/repositories/user_rep"
	"store/internal/usecases/product_usecase"
	"store/internal/usecases/user_product_usecase"
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
	database := client.Database("store")
	jwtsecret := os.Getenv("JWT_SECRET")

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.SecurityHeaders())
	//router.Use(middlewares.RateLimiter())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorizaion"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//router.Use(middlewares.APIKeyAuth())
	router.Use(gzip.Gzip(gzip.BestCompression))

	cacher := cacher.NewCacher(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASS"))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Message": "app works"})
	})
	userCollection := database.Collection("user")
	productCollection := database.Collection("product")
	categoryCollection := database.Collection("category")
	ratesCollection := database.Collection("rate")
	invoicesCollection := database.Collection("invoice")
	commentsCollection := database.Collection("comment")

	userRep := user_rep.NewUserRepository(userCollection)
	userUsercase := user_usecase.NewUserUsecase(userRep, cacher)
	j := jwt.NewJWTTokenHandler(jwtsecret)
	userDelivery := user_delivery.NewUserDelivary(userUsercase, j)

	productRep := product_rep.NewProductRep(productCollection, categoryCollection)
	productUsecase := product_usecase.NewProductUseCase(productRep)
	productDelivery := product_delivery.NewProductDelivery(productUsecase)

	userproductRep := user_product_rep.NewUserProductRep(ratesCollection, userCollection, productCollection, invoicesCollection, commentsCollection, categoryCollection)
	userproductUsecase := user_product_usecase.NewUserProductUsecase(userproductRep)
	userproductDelivery := user_product_delivery.NewUserProductDelivery(userproductUsecase)

	public := router.Group("/public-api")
	public.POST("/signin", userDelivery.FirstStep)
	public.POST("/loginphone", userDelivery.LoginWithPhone)
	public.POST("/loginemail", userDelivery.LoginWithEmail)
	public.POST("/signupphone", userDelivery.SignupWithPhone)
	public.POST("/signupemail", userDelivery.SignupWithEmail)
	public.POST("/refreshtoken", userDelivery.RefreshToken)
	public.GET("/getproduct/:id", productDelivery.GetProduct)
	public.GET("/getproducts", productDelivery.GetProducts)
	public.GET("/query", productDelivery.SearchQuery)
	public.GET("/getcomments/:id", userproductDelivery.GetComments)
	public.GET("/getrates/:id", userproductDelivery.GetRates)
	private := router.Group("/private-api")
	a := middlewares.NewAuth(j)
	private.Use(a.AuthMiddleware())
	private.GET("/isinfaves/:id", userproductDelivery.CheckFave)
	private.GET("/userfaves", userproductDelivery.GetFaves)
	//private.Use(middlewares.CSRFMiddleware())
	admin := router.Group(os.Getenv("ADMIN_ROUTE"))
	admin.Use()
	admin.POST("/addproduct", productDelivery.AddProduct)
	admin.POST("/addcategory", productDelivery.AddCategory)
	admin.DELETE("/deleteproduct/:id", productDelivery.DeleteProduct)
	admin.PUT("/editproduct/:id", productDelivery.EditProduct)
	router.Run(":8080")
}
