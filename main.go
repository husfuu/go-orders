package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/husfuu/go-order/handler"
	"github.com/husfuu/go-order/repository"
	"github.com/husfuu/go-order/service"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("ngak konek")
	}

	orderRepository := repository.NewOrderRepository(db)

	orderService := service.NewOrderService(orderRepository)

	userHandler := handler.NewOrderHandler(orderService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.GET("/orders", userHandler.GetOrders)
	api.GET("/orders/:id", userHandler.GetOrder)
	api.POST("/orders", userHandler.CreateOrder)
	api.PUT("/orders/:id", userHandler.UpdateOrder)
	api.DELETE("/orders/:id", userHandler.DeleteOrder)

	router.Run()
}
