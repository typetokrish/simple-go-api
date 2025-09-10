package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/typetokrish/simple-go-api/api"
)

func main() {

	log.Print("Main method is started")

	r := gin.Default()

	productsHandler := api.NewProductsHandler()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		log.Print("Serving HTTP/GET health api")
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// List products endpoint
	r.GET("/products", func(c *gin.Context) {
		log.Print("Serving HTTP/GET products api")

		products := productsHandler.GetProducts()
		c.JSON(http.StatusOK, products)
	})

	// Cloud Run uses PORT env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on PORT:%s", port)

	err := r.Run(":" + port) // Start server
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
