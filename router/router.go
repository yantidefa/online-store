package router

import (
	"fmt"
	"log"
	authhandler "online-store/handler/auth_handler"
	carthandler "online-store/handler/cart_handler"
	categoryproducthandler "online-store/handler/category_product_handler"
	paymenthandler "online-store/handler/payment_handler"
	producthandler "online-store/handler/product_handler"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Routes() *gin.Engine {

	err_host := godotenv.Load(".env")
	if err_host != nil {
		fmt.Println(err_host)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v3noauth := r.Group("/v1")

	auth := v3noauth.Group("/auth")
	{
		auth.POST("/register-customer", authhandler.RegisterCustomer)
		auth.POST("/register-admin", authhandler.RegisterAdmin)
		auth.POST("/login", authhandler.Login)
		auth.POST("/logout", authhandler.Logout)
	}

	categoryProduct := v3noauth.Group("/category")
	{
		categoryProduct.GET("/get", categoryproducthandler.GetCategoryProductAll)
		categoryProduct.GET("/get-by-id", categoryproducthandler.GetCategoryProductById)
		categoryProduct.DELETE("/delete-by-id", categoryproducthandler.DeleteCategoryProductById)
		categoryProduct.PUT("/update", categoryproducthandler.UpdateCategoryProductById)
		categoryProduct.POST("/create", categoryproducthandler.CreateCategoryProduct)
	}

	product := v3noauth.Group("/product")
	{
		product.GET("/get", producthandler.GetProduct)
		product.DELETE("/delete-by-id", producthandler.DeleteProductById)
		product.PUT("/update", producthandler.UpdateProductById)
		product.POST("/create", producthandler.CreateProduct)
	}

	cart := v3noauth.Group("/cart")
	{
		cart.GET("/get", carthandler.GetCart)
		cart.DELETE("/delete-by-id", carthandler.DeleteCartById)
		cart.PUT("/update", carthandler.UpdateCartById)
		cart.POST("/create", carthandler.CreateCart)
	}

	payment := v3noauth.Group("/payment")
	{
		payment.GET("/get-by-user-id", paymenthandler.GetPaymentByUserId)
		payment.DELETE("/delete-by-id", paymenthandler.DeletePaymentById)
		payment.POST("/create", paymenthandler.CreatePayment)
	}

	return r
}

func InitialRouter() {
	port := ":" + os.Getenv("ACTIVE_PORT")
	if err := Routes().Run(port); err != nil {
		log.Fatalln(err)
	}
}
