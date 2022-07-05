package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rishitashaw/ecommerce/controllers"
	"github.com/rishitashaw/ecommerce/database"
	"github.com/rishitashaw/ecommerce/middleware"
	"github.com/rishitashaw/ecommerce/routes"
)

func main() {
    port:=os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    app:=controllers.NewApplication(database.ProductData(database.Client,"Products"), database.UserData(database.Client,"Users"))

    router := gin.New()
    router.Use(middleware.Logger())

    routes.UserRoutes(router)
    router.Use(middleware.Auth())

    router.GET("/addtocart",app.Addtocart())
    router.GET("/removeitem",app.RemoveItem())
    router.GET("/cartcheckout",app.BuyFromCart())
    router.GET("/instantbuy",app.InstantBuy())

    log.Fatal((router.Run(":"+port)))
}