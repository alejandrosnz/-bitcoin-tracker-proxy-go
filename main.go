package main

import "alejandrosnz/bitcoin-tracker-proxy-go/controllers"

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Import Controllers
	ticker := r.Group("/api/ticker")
	{
		ticker.GET("/current_price/:symbol", controllers.GetCurrentPriceBySymbol)
		ticker.GET("/closing_price/:symbol", controllers.GetClosingPriceBySymbol)
	}

	// Start server
	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
