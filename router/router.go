package router

import (
	"toko_belanja/controller"
	"toko_belanja/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/login", controller.UserLogin)
		userRouter.POST("/register", controller.UserRegister)
		userRouter.Use(middleware.Authentication())
		userRouter.PATCH("/topup", controller.UserTopUp)
	}

	categoryRouter := r.Group("/categories")
	{
		categoryRouter.Use(middleware.Authentication(), middleware.CategoryAuthorization())
		categoryRouter.POST("/", controller.CategoryCreate)
		categoryRouter.GET("/", controller.CategoryViewAll)
		categoryRouter.PATCH("/:categoryId", controller.CategoryUpdate)
		categoryRouter.DELETE("/:categgoryId", controller.CategoryDelete)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", middleware.ProductAuthorization(), controller.ProductRegister)
		productRouter.GET("/", controller.ProductViewAll)
		productRouter.PUT("/:productId", middleware.ProductAuthorization(), controller.ProductUpdate)
		productRouter.DELETE("/:productId", middleware.ProductAuthorization(), controller.ProductDelete)
	}

	transactionRouter := r.Group("/transactions")
	{
		transactionRouter.Use(middleware.Authentication())
		transactionRouter.POST("/", controller.TransactionCreate)
		transactionRouter.GET("/my-transactions", controller.TransactionViewMyTransaction)
		transactionRouter.GET("/user-transactions", middleware.TransactionAuthorization(), controller.TransactionViewAll)
	}

	return r
}
