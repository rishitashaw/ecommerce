package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishitashaw/ecommerce/controllers"
)

func UserRoutes(incomingRoute *gin.Engine){
	incomingRoute.POST("/users/signup", controllers.Signup())
	incomingRoute.POST("/users/login", controllers.Login())
	incomingRoute.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoute.GET("/users/productview", controllers.SearchProduct())
	incomingRoute.GET("/users/search", controllers.SearchProductByQuery())
}