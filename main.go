package main

import (
	"log"

	"gonewaje/final/config"
	"gonewaje/final/controllers"
	"gonewaje/final/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	database := db.Open(cfg.DatabaseURL)
	defer database.Close()

	auth := controllers.AuthController{DB: database}
	rest := controllers.RestaurantsController{DB: database}
	menu := controllers.MenusController{DB: database}
	order := controllers.OrdersController{DB: database}

	r := gin.Default()

	// Auth
	r.POST("/api/users/register", auth.Register)
	r.POST("/api/users/login", auth.Login)

	// Public restaurant browsing
	r.GET("/api/restaurants", rest.List)
	r.GET("/api/restaurants/:id", rest.Detail)
	r.GET("/api/restaurants/:id/menu", menu.ListByRestaurant)

	// Protected
	api := r.Group("/api")
	api.Use(controllers.JWTAuth())

	// Admin-only for managing restaurants & menus
	admin := api.Group("/admin")
	admin.Use(controllers.RequireAdmin())
	admin.POST("/restaurants", rest.Create)
	admin.PUT("/restaurants/:id", rest.Update)
	admin.DELETE("/restaurants/:id", rest.Delete)
	admin.POST("/menu", menu.Create)
	admin.PUT("/menu/:id", menu.Update)
	admin.DELETE("/menu/:id", menu.Delete)
	admin.PUT("/orders/:id/status", order.UpdateStatus)

	// Orders (customer)
	api.POST("/orders", order.Create)
	api.GET("/orders/mine", order.MyOrders)

	log.Printf("🚀 Server running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
