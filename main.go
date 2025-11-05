package main

import (
	"dzabrail/connect"
	"dzabrail/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	connect.ConnectBase()
	r.GET("/drinks", controllers.GetAllDrinks) // list grinks
	r.GET("/drinks/ins", controllers.DrinksInstock)// true drink
	r.GET("/drinks/:id", controllers.GetByID) // all drink info 
	r.POST("/drinks", controllers.AddDrink) // add drinks
	r.PATCH("/drinks/:id", controllers.UpdateDrink)// updated drink
	r.DELETE("/drinks/:id", controllers.DeleteDrink) // just delete

	r.Run("localhost:8080")

}