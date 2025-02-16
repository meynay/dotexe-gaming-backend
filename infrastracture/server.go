package infrastracture

import (
	"context"
	"fmt"
	"os"
	"store/internal/delivery/middlewares"
	"store/internal/delivery/user_delivery"
	"store/internal/repositories/user_rep"
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
	userRep := user_rep.NewUserRepository(userCollection)
	userUsercase := user_usecase.NewUserUsecase(userRep, cacher)
	j := jwt.NewJWTTokenHandler(jwtsecret)
	userDelivery := user_delivery.NewUserDelivary(userUsercase, j)

	public := router.Group("/public-api")
	public.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hi, public works"})
	})
	public.POST("/signin", userDelivery.FirstStep)
	public.POST("/loginphone", userDelivery.LoginWithPhone)
	public.POST("/loginemail", userDelivery.LoginWithEmail)
	public.POST("/signupphone", userDelivery.SignupWithPhone)
	public.POST("/signupemail", userDelivery.SignupWithEmail)
	public.POST("/refreshtoken", userDelivery.RefreshToken)

	private := router.Group("/private-api")
	a := middlewares.NewAuth(j)
	private.Use(a.AuthMiddleware())
	//private.Use(middlewares.CSRFMiddleware())
	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hi, private works"})
	})
	router.Run(":8080")
}
