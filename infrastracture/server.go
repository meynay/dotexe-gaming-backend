package infrastracture

import (
	"context"
	"fmt"
	"os"
	"store/internal/delivery/admin_delivery"
	"store/internal/delivery/cart_delivery"
	"store/internal/delivery/comment_delivery"
	"store/internal/delivery/fave_delivery"
	"store/internal/delivery/middlewares"
	"store/internal/delivery/product_delivery"
	"store/internal/delivery/rating_delivery"
	"store/internal/delivery/user_delivery"
	"store/internal/repositories/admin_rep"
	"store/internal/repositories/blogpost_rep"
	"store/internal/repositories/category_rep"
	"store/internal/repositories/comment_rep"
	"store/internal/repositories/invoice_rep"
	"store/internal/repositories/product_rep"
	"store/internal/repositories/rating_rep"
	"store/internal/repositories/user_rep"
	"store/internal/usecases/admin_usecase"
	"store/internal/usecases/cart_usecase"
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

	//collections
	database := client.Database("store")
	adminCollection := database.Collection(os.Getenv("ADMIN_COLLECTION"))
	blogPostCommentCollection := database.Collection("bpcomment")
	blogPostCollection := database.Collection("blogpost")
	categoryCollection := database.Collection("category")
	commentsCollection := database.Collection("comment")
	invoicesCollection := database.Collection("invoice")
	productCollection := database.Collection("product")
	ratesCollection := database.Collection("rating")
	userCollection := database.Collection("user")

	//repositories
	adminRep := admin_rep.NewAdminRep(adminCollection)
	blogPostRep := blogpost_rep.NewBlogPostRep(blogPostCollection, blogPostCommentCollection)
	categoryRep := category_rep.NewCategoryRep(categoryCollection)
	commentsRep := comment_rep.NewCommentRep(commentsCollection)
	productRep := product_rep.NewProductRep(productCollection)
	invoiceRep := invoice_rep.NewInvoiceRep(invoicesCollection)
	ratingRep := rating_rep.NewRatingRep(ratesCollection)
	userRep := user_rep.NewUserRepository(userCollection)

	//usecases
	cacher := cacher.NewCacher(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASS"))
	adminUsecase := admin_usecase.NewAdminUsecase(productRep, categoryRep, invoiceRep, adminRep, userRep, blogPostRep)
	cartUsecase := cart_usecase.NewCartUsecase(productRep, userRep, invoiceRep)
	commentUsecase := comment_usecase.NewCommentUsecase(commentsRep, userRep, adminRep)
	faveUsecase := fave_usecase.NewFaveUsecase(userRep, productRep, categoryRep)
	productUsecase := product_usecase.NewProductUseCase(productRep, categoryRep)
	ratingUsecase := rating_usecase.NewRatingUsecase(ratingRep, productRep, userRep)
	userUsercase := user_usecase.NewUserUsecase(userRep, cacher)

	//deliveries
	jwtsecret := os.Getenv("JWT_SECRET")
	j := jwt.NewJWTTokenHandler([]byte(jwtsecret))
	adminDelivery := admin_delivery.NewAdminDelivery(adminUsecase, j)
	cartDelivery := cart_delivery.NewCartDelivery(cartUsecase)
	commentDelivery := comment_delivery.NewCommentDelivery(commentUsecase)
	faveDelivery := fave_delivery.NewFaveDelivery(faveUsecase)
	productDelivery := product_delivery.NewProductDelivery(productUsecase)
	ratingDelivery := rating_delivery.NewRatingDelivery(ratingUsecase)
	userDelivery := user_delivery.NewUserDelivary(userUsercase, j)

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
	public.GET("/product/:productid", productDelivery.GetProduct)
	public.GET("/products", productDelivery.GetProducts)
	public.GET("/categories", productDelivery.GetCategories)
	public.GET("/query", productDelivery.SearchQuery)
	public.GET("/comments/:productid", commentDelivery.GetComments)
	public.GET("/rates/:productid", ratingDelivery.GetRates)

	//private group (needs auth)
	private := router.Group("/private-api")
	a := middlewares.NewAuth(j)
	private.Use(a.AuthMiddleware())
	private.GET("/info", userDelivery.GetInfo)
	private.PUT("/fillinfo", userDelivery.FillInfo)
	private.PUT("/changepass", userDelivery.ResetPassword)
	private.GET("/isincart/:id", cartDelivery.IsInCart)
	private.POST("/addtocart/:id", cartDelivery.AddToCart)
	private.PUT("/editincart/:id", cartDelivery.ChangeCountInCart)
	private.GET("/getcart", cartDelivery.GetCart)
	private.GET("/isinfaves/:id", faveDelivery.CheckFave)
	private.GET("/userfaves", faveDelivery.GetFaves)
	private.POST("/faveproduct/:id", faveDelivery.FaveProduct)
	private.DELETE("unfaveproduct/:id", faveDelivery.UnfaveProduct)
	private.POST("/comment/:productid", commentDelivery.CommentOnProduct)
	private.GET("/invoices", cartDelivery.GetInvoices)
	private.GET("/invoice/:invoiceid", cartDelivery.GetInvoice)
	private.GET("/getrate/:productid", ratingDelivery.GetRate)
	private.POST("/rate/:productid", ratingDelivery.RateProduct)
	//private.Use(middlewares.CSRFMiddleware())

	//admin group
	adminroute := os.Getenv("ADMIN_ROUTE")
	admin := router.Group(adminroute)
	//admin side
	admin.POST("/login", adminDelivery.Login)
	admin.Use(a.AuthMiddleware())
	admin.GET("/info", adminDelivery.GetInfo)
	admin.POST("/addadmin", adminDelivery.AddAdmin)
	admin.PUT("/fillfields", adminDelivery.FillFields)
	//product side
	admin.POST("/addproduct", adminDelivery.AddProduct)
	admin.PUT("/editproduct/:productid", adminDelivery.EditProduct)
	admin.DELETE("/deleteproduct/:productid", adminDelivery.DeleteProduct)
	admin.GET("/activeproductscount", adminDelivery.GetActiveProductsCount)
	//category side
	admin.POST("/addcategory", adminDelivery.AddCategory)
	admin.PUT("/editcategory/:categoryid", adminDelivery.EditCategory)
	admin.DELETE("/deletecategory/:categoryid", adminDelivery.DeleteCategory)
	//order side
	admin.GET("/invoices", adminDelivery.GetInvoices)
	admin.GET("/invoice/:invoiceid", adminDelivery.GetInvoice)
	admin.PUT("/changeorderstatus/:invoiceid", adminDelivery.ChangeInvoiceStatus)
	admin.GET("/ordercount", adminDelivery.GetNewInvoicesCount)
	admin.GET("/monthlysales", adminDelivery.GetMonthlySalesPrice)
	//user side
	admin.GET("/user/:userid", adminDelivery.GetUser)
	admin.GET("/activeusers", adminDelivery.GetActiveUsers)
	admin.GET("/activeuserscount", adminDelivery.GetActiveUsersCount)
	//chart
	admin.GET("/chart", adminDelivery.GetChart)

	router.Run(":8080")
}
