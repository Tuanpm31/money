package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"money/internal/api"
	"money/internal/db"
)

func main() {
	connectionUrl := os.Getenv("DATABASE_URL")
	fmt.Println("connectionUrl: ", connectionUrl)
	if connectionUrl == "" {
		log.Fatal("DATABASE_URL must be set")
	}
	dbConn := db.InitDB(connectionUrl)
	db.CreateMoneyTable(dbConn)

	handler := api.NewHandler(dbConn)
	router := gin.Default()

	router.GET("/money/:id", handler.GetMoney)
	router.PUT("/money/:id", handler.SetMoney)

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
